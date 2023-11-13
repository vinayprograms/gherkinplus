package graph

import (
	"errors"
	"strings"

	"github.com/cucumber/messages-go/v16"
)

// //////////////////////////////////////
// Sample chain rule for use in tests

type entry struct {
	keyword  string
	scenario *messages.Scenario
}

// The 'Then' statement of a scenario matches the 'When' statement of a previous scenario
type ScenarioChainRule struct {
	statementsIndex map[string][]*entry
}

func NewScenarioChainRule() (r *ScenarioChainRule) {
	r = &ScenarioChainRule{
		statementsIndex: make(map[string][]*entry),
	}
	return
}

func (r *ScenarioChainRule) Index(features []*messages.Feature) error {
	if r.statementsIndex == nil {
		return errors.New("create new ScenarioChainRule object first by calling `NewScenarioChainRule(...)`")
	}

	for _, feature := range features {
		for _, child := range feature.Children {
			if child.Scenario != nil {
				r.indexScenario(child.Scenario)
			} else if child.Rule != nil {
				r.indexRule(child.Rule)
			}
		}
	}

	return nil
}

func (r *ScenarioChainRule) indexScenario(scenario *messages.Scenario) error {
	if r.statementsIndex == nil {
		return errors.New("create new ScenarioChainRule object first by calling `NewScenarioChainRule(...)`")
	}

	for _, s := range scenario.Steps {
		r.statementsIndex[strings.TrimSpace(s.Text)] = append(r.statementsIndex[strings.TrimSpace(s.Text)],
			&entry{
				keyword:  strings.TrimSpace(s.Keyword),
				scenario: scenario,
			})
	}

	return nil
}

func (r *ScenarioChainRule) indexRule(rule *messages.Rule) error {
	if r.statementsIndex == nil {
		return errors.New("create new ScenarioChainRule object first by calling `NewScenarioChainRule(...)`")
	}

	for _, child := range rule.Children {
		if child.Scenario != nil {
			r.indexScenario(child.Scenario)
		}
	}

	return nil
}

func (r *ScenarioChainRule) Apply() (map[interface{}][]interface{}, error) {
	edges := make(map[interface{}][]interface{})

	if r.statementsIndex == nil {
		return nil, errors.New("create new object by calling `NewScenarioChainRule(...)` and indexing features by calling `Index(...)` before calling `Apply(...)`")
	}

	for _, scenarios := range r.statementsIndex {
		if len(scenarios) > 1 {
			// Confirm that there is only one "Then" statement
			thenCount := 0
			for _, scenario := range scenarios {
				if scenario.keyword == "Then" {
					thenCount++
				}
			}
			if thenCount != 1 {
				return nil, errors.New("multiple 'Then' statements found")
			}

			// Add edges from the scenario with 'When' statement to associated scenarios with same 'Then' statement
			for _, scenario := range scenarios {
				if scenario.keyword == "Then" {
					for _, other := range scenarios {
						if other.keyword == "When" {
							edges[scenario.scenario] = append(edges[scenario.scenario], other.scenario)
						}
					}
				}
			}
		}
	}

	return edges, nil
}
