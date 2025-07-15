package main

import (
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

var Canteens []int = []int{TU_WIEN, UNI_PASSAU}

func main() {
	wg := &sync.WaitGroup{}
	for _, canteenId := range Canteens {
		wg.Add(1)
		go func() {
			fmt.Printf("Generating RSS for canteen ID %v\n", canteenId)
			defer wg.Done()
			feed, err := openmensarss.FeedForCanteenID(canteenId, time.Now())
			if err != nil {
				fmt.Printf("err: canteen ID: %v: %v\n", canteenId, err)
				return
			}

			// overwrite image with custom 88x31 image per feed
			img := &feeds.Image{Url: fmt.Sprintf("https://schicho.github.io/openmensarss/%v.gif", canteenId), Title: openmensarss.OpenMensaRSSGenerator, Link: "https://schicho.github.io/openmensarss/"}
			feed.Image = img

			file, err := os.Create("rss/" + strconv.Itoa(canteenId) + ".xml")
			if err != nil {
				fmt.Printf("err: canteen ID: %v: %v\n", canteenId, err)
				return
			}

			// manually convert to RssFeed struct, so we can set the Generator field.
			rss := (&feeds.Rss{Feed: feed}).RssFeed()
			rss.Generator = openmensarss.OpenMensaRSSGenerator

			err = feeds.WriteXML(rss, file)
			if err != nil {
				fmt.Printf("err: canteen ID: %v: %v\n", canteenId, err)
				return
			}
		}()
	}
	wg.Wait()
}
