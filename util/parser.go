package util

import (
	"strings"

	"github.com/cucumber/gherkin-go/v19"
	"github.com/cucumber/messages-go/v16"
)

// Parse a set of Feature strings into a set of GherkinDocuments.
//
// NOTE: Each string must contain a complete feature specification.
func Load(dialect DialectProvider, content []string) ([]*messages.GherkinDocument, error) {
	var documents []*messages.GherkinDocument
	for _, c := range content {
		builder := gherkin.NewAstBuilder((&messages.Incrementing{}).NewId)
		parser := gherkin.NewParser(builder)
		parser.StopAtFirstError(false)
		matcher := gherkin.NewLanguageMatcher(dialect, dialect.Name)
		err := parser.Parse(gherkin.NewScanner(strings.NewReader(c)), matcher)
		if err != nil {
			return nil, err
		}
		newDocument := builder.GetGherkinDocument()
		fixDocument(newDocument)
		documents = append(documents, newDocument)
	}
	return documents, nil
}

////////////////////////////////////////
// Internal functions

// Remove spaces around keywords and statements;
// Replace 'And" statements with the keyword of parent statement (Given, When or Then)
func fixDocument(document *messages.GherkinDocument) {
	// Remove spaces around keywords and statements
	for _, child := range document.Feature.Children {
		if child.Scenario != nil {
			fixSScenario(child.Scenario)
		} else if child.Background != nil {
			fixBackground(child.Background)
		} else if child.Rule != nil {
			fixRule(child.Rule)
		}
	}
}

func fixSScenario(scenario *messages.Scenario) {
	for i, step := range scenario.Steps {
		step.Keyword = strings.TrimSpace(step.Keyword)
		step.Text = strings.TrimSpace(step.Text)
		if step.Keyword == "And" {
			scenario.Steps[i].Keyword = scenario.Steps[i-1].Keyword
		}
	}
}

func fixBackground(background *messages.Background) {
	for i, step := range background.Steps {
		step.Keyword = strings.TrimSpace(step.Keyword)
		step.Text = strings.TrimSpace(step.Text)
		if step.Keyword == "And" {
			background.Steps[i].Keyword = "Given"
		}
	}
}

func fixRule(rule *messages.Rule) {
	for _, ruleChild := range rule.Children {
		if ruleChild.Background != nil {
			fixBackground(ruleChild.Background)
		} else if ruleChild.Scenario != nil {
			fixSScenario(ruleChild.Scenario)
		}
	}
}
