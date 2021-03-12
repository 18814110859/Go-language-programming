package search

import "log"

// contains the result of search
// 包搜索结果
type Result struct {
	Field   string
	Content string
}

// defines the behavior required by types that want
// to implement a new search type
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// TODO >>
// a goroutine for each individual feed to run
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	// search against the specified matcher
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	// write the results to channel
	for _,result := range searchResults {
		results <- result
	}
}

// writes results to the console window as they
// are received by the individual goroutines
func Display(results chan *Result) {

	// a result is written to the channel
	// once the channel is closed the for loop terminates
	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}


