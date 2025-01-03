module github.com/arikkfir/go-jsonschema/tests

go 1.22.0

replace (
	github.com/arikkfir/go-jsonschema => ../
	github.com/arikkfir/go-jsonschema/tests/helpers/other => ./helpers/other
)

require (
	github.com/arikkfir/go-jsonschema v0.16.0
	github.com/arikkfir/go-jsonschema/tests/helpers/other v0.0.0-20240909221408-bcba1cdc5eb2
	github.com/go-viper/mapstructure/v2 v2.1.0
	github.com/google/go-cmp v0.6.0
	github.com/stretchr/testify v1.10.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/goccy/go-yaml v1.15.13 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/sanity-io/litter v1.5.5 // indirect
	golang.org/x/exp v0.0.0-20241217172543-b2144cdd0a67 // indirect
)
