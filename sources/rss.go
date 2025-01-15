package sources

import (
	"fmt"
	"io"

	"github.com/mmcdole/gofeed"
)

// RSSFeed implements the Source interface for RSS feeds
type RSSFeed struct {
	rssFeedURL string
	parser     *gofeed.Parser
}

func (r *RSSFeed) Print(writer io.Writer) {
	feed, err := r.parser.ParseURL(r.rssFeedURL)
	if err != nil {
		fmt.Fprintf(writer, "Failed to retrieve rss feed: %v", err)
		return
	}

	fmt.Fprintf(writer, "Headlines from: %s\n", feed.Title)
	fmt.Fprintf(writer, "Updated on %s\n", feed.Published)
	for _, item := range feed.Items {
		fmt.Fprintf(writer, "%v: %s\n", item.Published, item.Title)
	}

}

func (r *RSSFeed) Configure(config map[string]string) error {

	r.parser = gofeed.NewParser()
	r.parser.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"
	url, ok := config["url"]
	if !ok {
		return fmt.Errorf("url is required")
	}
	r.rssFeedURL = url
	return nil
}
