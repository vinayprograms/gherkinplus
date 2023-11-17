package util

import (
	"testing"

	"github.com/cucumber/gherkin-go/v19"
	"github.com/stretchr/testify/assert"
)

func TestWithNullDialectAndEmptyContent(t *testing.T) {
	dp := NewDialect("adm", nil)
	docs, err := Load(dp, []string{})

	assert.Equal(t, nil, err)
	assert.Equal(t, 0, len(docs))
}

func TestWithEmptyContent(t *testing.T) {
	dialect := gherkin.GherkinDialect{
		Language: "adm",
		Name:     "attack-defense modeling",
		Native:   "attack-defense modeling",
		Keywords: map[string][]string{
			// Customized for ADM
			"feature":    {"Model"},
			"scenario":   {"Attack", "Defense"},
			"rule":       {"Policy"},
			"background": {"Assumption"},

			// Default (from Gherkin)
			"examples": {"Examples"},
			"given":    {"Given"},
			"when":     {"When"},
			"then":     {"Then"},
			"and":      {"And"},
			"but":      {"But"},
		},
	}
	dp := NewDialect("adm", &dialect)
	docs, err := Load(dp, []string{})

	assert.Equal(t, nil, err)
	assert.Equal(t, 0, len(docs))
}

func TestWithSimpleADMContent(t *testing.T) {
	dialect := gherkin.GherkinDialect{
		Language: "adm",
		Name:     "attack-defense modeling",
		Native:   "attack-defense modeling",
		Keywords: map[string][]string{
			// Customized for ADM
			"feature":    {"Model"},
			"scenario":   {"Attack", "Defense"},
			"rule":       {"Policy"},
			"background": {"Assumption"},

			// Default (from Gherkin)
			"examples": {"Examples"},
			"given":    {"Given"},
			"when":     {"When"},
			"then":     {"Then"},
			"and":      {"And"},
			"but":      {"But"},
		},
	}
	adm := `
		Model: Sample Model
			Attack: Sample Attack
				When the attack step happens
				Then the attack step is successful
			Defense: Sample Defense
				When the attack step is successful
				Then monitoring folks are notified
	`
	dp := NewDialect("adm", &dialect)
	docs, err := Load(dp, []string{adm})

	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(docs))
	assert.Equal(t, "Model", docs[0].Feature.Keyword)
	assert.Equal(t, "Sample Model", docs[0].Feature.Name)
	assert.Equal(t, "Attack", docs[0].Feature.Children[0].Scenario.Keyword)
	assert.Equal(t, "Sample Attack", docs[0].Feature.Children[0].Scenario.Name)
	assert.Equal(t, "Defense", docs[0].Feature.Children[1].Scenario.Keyword)
	assert.Equal(t, "Sample Defense", docs[0].Feature.Children[1].Scenario.Name)
	assert.Equal(t, docs[0].Feature.Children[0].Scenario.Steps[1].Text, docs[0].Feature.Children[1].Scenario.Steps[0].Text)
}

func TestWithWrongDialect(t *testing.T) {
	dialect := gherkin.GherkinDialect{
		Language: "none",
		Name:     "None",
		Native:   "None",
		Keywords: map[string][]string{
			// Uncustomized
			"feature":    {"Feature"},
			"scenario":   {"Scenario"},
			"rule":       {"Rule"},
			"background": {"Background"},

			// Default (from Gherkin)
			"examples": {"Examples"},
			"given":    {"Given"},
			"when":     {"When"},
			"then":     {"Then"},
			"and":      {"And"},
			"but":      {"But"},
		},
	}
	adm := `
		Model: Sample Model
			Attack: Sample Attack
				When the attack step happens
				Then the attack step is successful
			Defense: Sample Defense
				When the attack step is successful
				Then monitoring folks are notified
	`
	dp := NewDialect("none", &dialect)
	docs, err := Load(dp, []string{adm})

	assert.NotNil(t, err)
	assert.Nil(t, docs)
}

// TestWithSimpleADMContentWithNoDialect(...) not implemented since gherkin-go
// triggers a panic, crashin the program

func TestWithComplexADMContent(t *testing.T) {
	dialect := gherkin.GherkinDialect{
		Language: "adm",
		Name:     "attack-defense modeling",
		Native:   "attack-defense modeling",
		Keywords: map[string][]string{
			// Customized for ADM
			"feature":    {"Model"},
			"scenario":   {"Attack", "Defense"},
			"rule":       {"Policy"},
			"background": {"Assumption"},

			// Default (from Gherkin)
			"examples": {"Examples"},
			"given":    {"Given"},
			"when":     {"When"},
			"then":     {"Then"},
			"and":      {"And"},
			"but":      {"But"},
		},
	}
	adm := `
		Model: Sample Model
			Attack: Sample Attack
				When the attack step happens
				Then the attack step is successful
			Defense: Sample Defense
				When the attack step is successful
				Then monitoring folks are notified
			Policy: Sample Policy
				Assumption: Sample Assumption
					Given the assumption is true
				Defense: Sample Defense 1
					When monitoring folks are notified
					Then Company is fully defended
	`
	dp := NewDialect("adm", &dialect)
	docs, err := Load(dp, []string{adm})

	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(docs))
	assert.Equal(t, "Model", docs[0].Feature.Keyword)
	assert.Equal(t, "Sample Model", docs[0].Feature.Name)
	assert.Equal(t, "Attack", docs[0].Feature.Children[0].Scenario.Keyword)
	assert.Equal(t, "Sample Attack", docs[0].Feature.Children[0].Scenario.Name)
	assert.Equal(t, "Defense", docs[0].Feature.Children[1].Scenario.Keyword)
	assert.Equal(t, "Sample Defense", docs[0].Feature.Children[1].Scenario.Name)
	assert.Equal(t, docs[0].Feature.Children[0].Scenario.Steps[1].Text, docs[0].Feature.Children[1].Scenario.Steps[0].Text)
}
