package graph

import "github.com/cucumber/messages-go/v16"

type Indexer interface {
	// Generate indices that can then be used for applying the rule.
	Index(features []*messages.Feature) error
}

type Rule interface {
	// Apply this rule to a graph. The rule may modify the graphs in place.
	// The rule *MUST* create edges between nodes within th graph *AS WELL AS* subgraphs.
	Apply() (map[*interface{}][]*interface{}, error)
}
