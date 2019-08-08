package search

import (
	"log"
	"sync"
)

// Result contains the result of a search.
type Result struct {
	Filed   string
	Content string
}

func Run(searchTerm string) {
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}
	// Create an unbuffered channel to receive match results to display.
	results := make(chan *Result)
	// Setup a wait group so we can process all the feeds.
	var waitGroup sync.WaitGroup

	// Set the number of goroutines we need to wait for while
	// they process the individual feeds.
	waitGroup.Add(len(feeds))

	// Launch a goroutine for each feed to find the results.
	// for _, exists := range feeds {

	// }
}
