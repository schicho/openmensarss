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
const AKBILD_WIEN int = 1957

var Canteens []int = []int{TU_WIEN, UNI_PASSAU, AKBILD_WIEN}

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
