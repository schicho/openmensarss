package openmensarss

import (
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
		Title:  canteen.Name,
		Link:   &feeds.Link{Href: "https://github.com/schicho/openmensarss"},
		Author: &feeds.Author{Name: "OpenMensa RSS Generator"},
	}

	for meal := range menu.Meals {
		print(meal)
	}

	return feed, nil
}
