package openmensarss

import (
	"fmt"
	"strings"
	"time"

	"github.com/gorilla/feeds"
	"github.com/j0hax/go-openmensa"
)

const OpenMensaRSSGenerator string = "OpenMensa RSS Generator"

// RSSMetadata allows setting of general metadata for the feeds generated using this library's functions.
// All fields can be modified here before generation, or of course after generation on a per feed basis.
var RSSMetadata = struct {
	Description string
	Author      *feeds.Author
	Link        *feeds.Link
	Image       *feeds.Image
}{
	Description: "Automated RSS feed using OpenMensa",
	Author:      &feeds.Author{Name: OpenMensaRSSGenerator, Email: "johann.schicho@tuwien.ac.at"},
	Link:        &feeds.Link{Href: "https://schicho.github.io/openmensarss/"},
	Image:       &feeds.Image{Url: "https://schicho.github.io/openmensarss/omrss.gif", Title: OpenMensaRSSGenerator, Link: "https://schicho.github.io/openmensarss/"},
}

// FeedForCanteenID creates a feed of the canteen menu on a certain day.
// The id specifies the canteen from OpenMensa. The day is selected by passing a time stamp.
//
// Throws an error if the OpenMensa API does not provide data for the specified input.
func FeedForCanteenID(id int, date time.Time) (*feeds.Feed, error) {
	canteen, err := openmensa.GetCanteen(id)
	if err != nil {
		return nil, err
	}

	return generateFeed(canteen, date)
}

// FeedForCanteen creates a feed of the canteen menu on a certain day.
// The Canteen struct specifies the canteen from OpenMensa. The day is selected by passing a time stamp.
//
// Throws an error if the OpenMensa API does not provide data for the specified input.
func FeedForCanteen(canteen *openmensa.Canteen, date time.Time) (*feeds.Feed, error) {
	return generateFeed(canteen, date)
}

// generatedFeed builds the gorilla/feeds Feed.
// It uses the values of the RSSMetadata struct and converts each meal of OpenMensa to a feed item.
func generateFeed(canteen *openmensa.Canteen, date time.Time) (*feeds.Feed, error) {
	menu, err := canteen.MenuOn(date)
	if err != nil {
		return nil, err
	}

	t := time.Now()

	feed := &feeds.Feed{
		Title:       canteen.Name + " " + t.Format("(Mon, 02 Jan 2006)"),
		Link:        RSSMetadata.Link,
		Description: RSSMetadata.Description,
		Author:      RSSMetadata.Author,
		Image:       RSSMetadata.Image,
		Created:     t,
	}

	feed.Items = make([]*feeds.Item, 0, len(menu.Meals))

	for _, meal := range menu.Meals {
		feed.Add(createFeedItem(meal))
	}

	return feed, nil
}

func createFeedItem(meal openmensa.Meal) *feeds.Item {
	prices := []string{}

	for k, v := range meal.Prices {
		// null values of the OpenMensa API are unmarshalled into 0.0
		if v == 0.0 {
			continue
		}
		prices = append(prices, fmt.Sprintf("<i>%v: %.2f</i>", strings.Title(k), v))
	}

	priceInfo := strings.Join(prices, ", ")
	allergenInfo := fmt.Sprintf("(%v)", strings.Join(meal.Notes, ", "))

	return &feeds.Item{
		Title:       meal.Name,
		Description: fmt.Sprintf("%v <p>%v %v</p>", meal.Category, priceInfo, allergenInfo),
		Link:        RSSMetadata.Link,
	}
}
