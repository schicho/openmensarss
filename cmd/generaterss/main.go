package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/feeds"
	"github.com/schicho/openmensarss/v2"
)

const TU_WIEN int = 1098
const UNI_PASSAU int = 196
const AKBILD_WIEN int = 1957

var Canteens []int = []int{TU_WIEN, UNI_PASSAU, AKBILD_WIEN}

func writeRSSWithStylesheet(feed feeds.XmlFeed, w io.Writer) error {
	x := feed.FeedXml()

	if _, err := w.Write([]byte(xml.Header)); err != nil {
		return err
	}

	if _, err := w.Write([]byte(`<?xml-stylesheet type="text/xsl" href="feed.xsl"?>` + "\n")); err != nil {
		return err
	}

	e := xml.NewEncoder(w)
	e.Indent("", "  ")
	return e.Encode(x)
}

// main starts the process to generate all requested canteen RSS feeds for tomorrow.
// tomorrow is calculated based on the current local time + 24 hours.
// Unless it's before monday noon, then monday's menu is fetched.
func main() {
	now := time.Now()
	fetchTime := now

	// if it's before monday noon, we want to fetch monday's menu.
	// that's because monday menus are often not available on sunday evening.
	if !(now.Weekday() == time.Monday && now.Hour() < 12) {
		fetchTime = now.Add(24 * time.Hour)
	}

	wg := &sync.WaitGroup{}
	for _, canteenId := range Canteens {
		wg.Add(1)
		go func() {
			fmt.Printf("Generating RSS for canteen ID %v\n", canteenId)
			defer wg.Done()
			feed, err := openmensarss.FeedForCanteenID(canteenId, fetchTime)
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

			err = writeRSSWithStylesheet(rss, file)
			if err != nil {
				fmt.Printf("error: %v\n", err)
				return
			}
		}()
	}
	wg.Wait()
}
