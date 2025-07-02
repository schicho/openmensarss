package openmensarss

import (
	"strings"
	"time"

	"github.com/gorilla/feeds"
	"github.com/j0hax/go-openmensa"
)

func FeedForCanteenID(id int, date time.Time) (*feeds.Feed, error) {
	canteen, err := openmensa.GetCanteen(id)
	if err != nil {
		return nil, err
	}

	return generateFeed(canteen, date)
}

func FeedForCanteen(canteen *openmensa.Canteen, date time.Time) (*feeds.Feed, error) {
	return generateFeed(canteen, date)
}

func generateFeed(canteen *openmensa.Canteen, date time.Time) (*feeds.Feed, error) {
	menu, err := canteen.MenuOn(date)
	if err != nil {
		return nil, err
	}

	feed := &feeds.Feed{
		Title:       canteen.Name,
		Link:        &feeds.Link{Href: "https://github.com/schicho/openmensarss"},
		Description: "Automated RSS feed using OpenMensa",
		Author:      &feeds.Author{Name: "OpenMensa RSS Generator"},
		Created:     time.Now(),
	}

	feed.Items = make([]*feeds.Item, 0, 10)

	for _, meal := range menu.Meals {
		feed.Items = append(feed.Items, createFeedItem(meal))
	}

	return feed, nil
}

func createFeedItem(meal openmensa.Meal) *feeds.Item {
	desc := strings.Join(append(meal.Notes, meal.Category), ", ")

	return &feeds.Item{
		Title:       meal.Name,
		Description: desc,
	}
}
