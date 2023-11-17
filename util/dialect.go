package util

import (
	"github.com/cucumber/gherkin-go/v19"
)

type DialectProvider struct {
	Name    string
	Dialect *gherkin.GherkinDialect
}

func NewDialect(name string, dialect *gherkin.GherkinDialect) DialectProvider {
	return DialectProvider{
		Name:    name,
		Dialect: dialect,
	}
}

func (dp DialectProvider) GetDialect(language string) *gherkin.GherkinDialect {
	if dp.Name == language {
		return dp.Dialect
	} else {
		// This line should never be reached since parser is going to send the name of the dialect.
		return nil
	}
}
