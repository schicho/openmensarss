package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/feeds"
	"github.com/schicho/openmensarss"
)

const TU_WIEN int = 1098
const UNI_PASSAU int = 196
const AKBILD_WIEN int = 1957

var Canteens []int = []int{TU_WIEN, UNI_PASSAU, AKBILD_WIEN}

// writeRSSWithStylesheet writes the RSS feed to a file with an XSL stylesheet processing instruction
func writeRSSWithStylesheet(rss *feeds.RssFeed, file *os.File) error {
	// Wrap the RssFeed in RssFeedXml to get the <rss> root element
	rssFeedXml := &feeds.RssFeedXml{
		Version:          "2.0",
		ContentNamespace: "http://purl.org/rss/1.0/modules/content/",
		Channel:          rss,
	}
	
	// Write XML declaration
	_, err := file.WriteString(xml.Header)
	if err != nil {
		return err
	}
	
	// Write XSL stylesheet processing instruction
	_, err = file.WriteString(`<?xml-stylesheet type="text/xsl" href="feed.xsl"?>` + "\n")
	if err != nil {
		return err
	}
	
	// Marshal the RSS feed to XML
	var buf bytes.Buffer
	encoder := xml.NewEncoder(&buf)
	encoder.Indent("", "  ")
	err = encoder.Encode(rssFeedXml)
	if err != nil {
		return err
	}
	
	// Write the RSS content
	_, err = file.Write(buf.Bytes())
	return err
}

func main() {
	wg := &sync.WaitGroup{}
	for _, canteenId := range Canteens {
		wg.Add(1)
		go func() {
			fmt.Printf("Generating RSS for canteen ID %v\n", canteenId)
			defer wg.Done()
			feed, err := openmensarss.FeedForCanteenID(canteenId, time.Now())
			if err != nil {
				fmt.Printf("error: %v\n", err)
				return
			}

			file, err := os.Create("rss/" + strconv.Itoa(canteenId) + ".xml")
			if err != nil {
				fmt.Printf("error: %v\n", err)
				return
			}
			defer file.Close()

			// manually convert to RssFeed struct, so we can set the Generator field.
			rss := (&feeds.Rss{Feed: feed}).RssFeed()
			rss.Generator = openmensarss.OpenMensaRSSGenerator

			// Write XML with XSL stylesheet processing instruction
			err = writeRSSWithStylesheet(rss, file)
			if err != nil {
				fmt.Printf("error: %v\n", err)
				return
			}
		}()
	}
	wg.Wait()
}
