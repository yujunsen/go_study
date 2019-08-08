package matchers

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"

	"github.com/goinaction/mycode/demo2/search"
)

type (
	// item defines the fields associated with the item tag
	// in the rss document.
	item struct {
		XMLNAME     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        string   `xml:"guid"`
		GeoRssPoint string   `xml:"georss:point"`
	}

	// image defines the fields associated with the image tag
	// in the rss document.
	image struct {
		XMLNAME xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}

	// channel defines the fields associated with the channel tag
	// in the rss document.
	channel struct {
		XMLNAME        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubDate        string   `xml:"pubDate"`
		LastBuildDate  string   `xml:"lastBuildDate"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          image    `xml:"image"`
		Item           []item   `xml:"item"`
	}

	// rssDocument defines the fields associated with the rss document.
	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
)

// rssMatcher implements the Matcher interface.
type rssMatcher struct{}

// init registers the matcher with the program.
func init() {
	var matcher rssMatcher
	search.Registers("rss", matcher)
}

// Search looks at the document for the specified search term.
func (m rssMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, err) {
	var results []*search.Result

	log.Printf("Search Feed Type[%s] Site[%s] For URI[%s]\n", feed.Type, feed.Name, feed.URI)

	// Retrieve the data to search.
	document, err := m.retrieve(feed)
	if err != nil{
		return nil, err
	}

}

// retrieve performs a HTTP Get request for the rss feed and decodes the results.
func (m rssMatcher) retrieve(feed *search.Feed) (*rssDocument, error) {
	if feed.URI == "" {
		return nil, error.New("No rss feed uri provided")
	}

	// Retrieve the rss feed document from the web.
	resp, err := http.Get(feed.URI)
	if err != nil {
		return nil, err
	}
	// Retrieve the rss feed document from the web.
	defer resp.Body.Close()

	// Check the status code for a 200 so we know we have received a
	// proper response.
	if resp.StatusCode != 200 {
		//return nil, fmt.Errorf("HTTP Response Error %d\n", resp.StatusCode)
		return nil, fmt.Errorf("HTTP Response Error %d\n ", resp.StatusCode)
	}

	// Decode the rss feed document into our struct type.
	// We don't need to check for errors, the caller can do this.
	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)

	return &document, err
}