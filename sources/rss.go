package sources

import (
	"fmt"
	"io"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"github.com/mmcdole/gofeed"
)

// RSSFeed implements the Source interface for RSS feeds
type RSSFeed struct {
	rssFeedURL string
	fields     []string
	parser     *gofeed.Parser
}

const (
	fieldTitle       = "title"
	fieldDate        = "date"
	fieldAuthor      = "author"
	fieldDescription = "description"
	fieldLink        = "link"
)

func (r *RSSFeed) Print(writer io.Writer) {
	feed, err := r.parser.ParseURL(r.rssFeedURL)
	if err != nil {
		fmt.Fprintf(writer, "Failed to retrieve rss feed: %v", err)
		return
	}

	fmt.Fprintf(writer, "Headlines from: %s\n", feed.Title)
	fmt.Fprintf(writer, "Updated on %s\n", feed.Published)
	for _, item := range feed.Items {
		for _, field := range r.fields {
			fmt.Fprintf(writer, "%s:\t %s\n", field, getArticleField(item, field))
		}

	}

}

func getArticleField(item *gofeed.Item, field string) string {
	switch field {
	case fieldTitle:
		return item.Title
	case fieldDate:
		return item.Published
	case fieldAuthor:
		authors := ""
		for idx, author := range item.Authors {
			authors += author.Name
			if idx < len(item.Authors)-1 {
				authors += ", "
			}
		}
		return authors
	case fieldDescription:

		rval := item.Description

		if rval == "" {
			rval = item.Content
		}

		rval = bluemonday.StrictPolicy().Sanitize(rval)
		return rval
	default:
		return ""
	}
}

func (r *RSSFeed) Configure(config map[string]string) error {

	r.parser = gofeed.NewParser()
	r.parser.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"
	url, ok := config["url"]
	if !ok {
		return fmt.Errorf("url is required")
	}
	fields, ok := config["fields"]
	if !ok {
		r.fields = []string{"title", "date", "author"}
	} else {
		aFields := strings.Split(fields, ",")
		for i := range aFields {
			aFields[i] = strings.TrimSpace(aFields[i])
		}
		r.fields = aFields
	}

	r.rssFeedURL = url
	return nil
}
