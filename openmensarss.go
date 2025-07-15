package openmensarss

import (
	"strings"
	"time"

	"github.com/gorilla/feeds"
	"github.com/j0hax/go-openmensa"
)

const OpenMensaRSSGenerator string = "OpenMensa RSS Generator"

var RSSMetadata = struct {
	Description   string
	GeneratorName string
	Author        *feeds.Author
	Link          *feeds.Link
	Image         *feeds.Image
}{
	Description:   "Automated RSS feed using OpenMensa",
	GeneratorName: OpenMensaRSSGenerator,
	Author:        &feeds.Author{Name: OpenMensaRSSGenerator, Email: "johann.schicho@tuwien.ac.at"},
	Link:          &feeds.Link{Href: "https://schicho.github.io/openmensarss/"},
	Image:         &feeds.Image{Url: "https://schicho.github.io/openmensarss/omrss.gif", Title: OpenMensaRSSGenerator, Link: "https://schicho.github.io/openmensarss/"},
}

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

	t := time.Now()

	b := strings.Builder{}
	b.WriteString(t.Format("2006-01-02"))
	b.WriteString(", ")
	b.WriteString(canteen.Name)

	feed := &feeds.Feed{
		Title:       b.String(),
		Link:        RSSMetadata.Link,
		Description: RSSMetadata.Description,
		Author:      RSSMetadata.Author,
		Image:       RSSMetadata.Image,
		Created:     t,
	}

	feed.Items = make([]*feeds.Item, 0, 10)

	for _, meal := range menu.Meals {
		feed.Add(createFeedItem(meal))
	}

	return feed, nil
}

func createFeedItem(meal openmensa.Meal) *feeds.Item {
	desc := strings.Join(append(meal.Notes, meal.Category), ", ")

	return &feeds.Item{
		Title:       meal.Name,
		Description: desc,
		Link:        RSSMetadata.Link,
	}
}
