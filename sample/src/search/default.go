package search

// implements the default matcher
type defaultMatcher struct{}

// registers the default matcher with the program
func init() {
	var matcher defaultMatcher
	Register("default", matcher);
}

// implements the behavior for the default matcher
func (m defaultMatcher) Search(feeds *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}








