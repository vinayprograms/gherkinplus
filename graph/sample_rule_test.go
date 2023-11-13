package graph

import (
	"testing"

	"github.com/cucumber/gherkin-go/v19"
	"github.com/cucumber/messages-go/v16"
	"github.com/stretchr/testify/assert"
	"github.com/vinayprograms/gherkinplus/loader"
)

// //////////////////////////////////////
// Tests
func TestScenarioChaining(t *testing.T) {
	docs, err := loader.Load(
		loader.NewDialectProvider("en", gherkin.GherkinDialectsBuildin().GetDialect("en")),
		[]string{
			`
			Feature: Sample feature
				Scenario: Sample scenario 1
					When First step is executed
					Then First result is returned
				Scenario: Sample scenario 2
					When First result is returned
					Then Result 2 is returned
				Scenario: Sample scenario 3
					When First result is returned
					Then Result 3 is returned
				Scenario: Sample scenario 4
					When Result 2 is returned
					And Result 3 is returned
					Then Result 4 is returned
				Rule: Sample Rule
					Background: Sample background
						Given Background step 1
						And Background step 2
					Scenario: Sample scenario 5
						When Result 2 is returned
						Then Result 5 is returned
			`,
		})
	if err != nil {
		t.Error(err)
	}

	r := NewScenarioChainRule()
	r.Index([]*messages.Feature{docs[0].Feature})
	connections, err := r.Apply()
	if err != nil {
		t.Error(err)
	}

	assert.Len(t, connections, 3) // Three 1..N connections
	assert.Len(t, connections[docs[0].Feature.Children[0].Scenario], 2)
	assert.Len(t, connections[docs[0].Feature.Children[1].Scenario], 2)
	assert.Len(t, connections[docs[0].Feature.Children[2].Scenario], 1)
}
