package types

type MatchResults struct {
	File     string
	Matchers []MatcherResult
}

type MatcherResult struct {
	MatcherName string
	Result      map[string]Match
}

type Match struct {
	Match      string
	LineNumber string
	Message    string
}
