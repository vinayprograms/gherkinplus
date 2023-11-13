# Gherkin+

A wrapper around cucumber's official gherkin [package](https://github.com/cucumber/gherkin/tree/main/go) to enable connections between multiple parts of a feature or across features. One such example - `ScenarioChainRule` is available under [graph](/graph/sample_rule.go) package.

## Purpose

Certain tools use relationships between features and scenarios as a way to build decision graphs and for other graph based analysis.

## Design

The `loader` package includes a `Load(...)` function. It accepts a `DialectProvider` structure to allow for custom dialects. If you to use a dialect for a specific language, use one of gherkin's builtin dialects to build your custom `DialectProvider`. For example, `gherkin.GherkinDialectsBuildin().GetDialect("en")` for English.

The graph package contains a single interface - `Rule`. Each of your rules must implement the `Index(...)` function to build a map of statements to specific gherkin entities. For example `ScenarioChainRule` creates a map of a Step statement to `*messages.Scenario` so that the `Apply(...)` can use it to build a scenario-to-scenario map.
