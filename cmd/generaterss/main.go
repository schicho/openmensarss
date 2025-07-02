package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

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
			defer wg.Done()
			feed, err := openmensarss.FeedForCanteenID(canteenId, time.Now())
			if err != nil {
				fmt.Printf("err: %v\n", err)
				return
			}

			file, err := os.Create("rss/" + strconv.Itoa(canteenId) + ".xml")
			if err != nil {
				fmt.Printf("err: %v\n", err)
				return
			}

			err = feed.WriteRss(file)
			if err != nil {
				fmt.Printf("err: %v\n", err)
				return
			}
		}()
	}
	wg.Wait()
}
