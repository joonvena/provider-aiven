package config

import (
	"strings"

	"github.com/upbound/upjet/pkg/config"
	"github.com/upbound/upjet/pkg/types/name"
)

// GroupKindOverrides overrides the group and kind of the resource if it matches
// any entry in the GroupMap.
func GroupKindOverrides() config.ResourceOption {
	return func(r *config.Resource) {
		if f, ok := GroupMap[r.Name]; ok {
			r.ShortGroup, r.Kind = f(r.Name)
		}
	}
}

// KindOverrides overrides the kind of the resources given in KindMap.
func KindOverrides() config.ResourceOption {
	return func(r *config.Resource) {
		if k, ok := KindMap[r.Name]; ok {
			r.Kind = k
		}
	}
}

// GroupKindCalculator returns the correct group and kind name for given TF
// resource.
type GroupKindCalculator func(resource string) (string, string)

// ReplaceGroupWords uses given group as the group of the resource and removes
// a number of words in resource name before calculating the kind of the resource.
func ReplaceGroupWords(group string, count int) GroupKindCalculator {
	return func(resource string) (string, string) {
		words := strings.Split(strings.TrimPrefix(resource, "aiven_"), "_")
		snakeKind := strings.Join(words[count:], "_")
		return group, name.NewFromSnake(snakeKind).Camel
	}
}

// GroupMap contains all overrides we'd like to make to the default group search.
// Keep the same structure as in the Terraform docs: https://registry.terraform.io/providers/aiven/aiven/latest/docs
var GroupMap = map[string]GroupKindCalculator{
	// PG
	"aiven_pg":          ReplaceGroupWords("postgres", 0),
	"aiven_pg_database": ReplaceGroupWords("postgres", 0),
	"aiven_pg_user":     ReplaceGroupWords("postgres", 0),

	// MySQL
	"aiven_mysql":          ReplaceGroupWords("mysql", 0),
	"aiven_mysql_database": ReplaceGroupWords("mysql", 0),
	"aiven_mysql_user":     ReplaceGroupWords("mysql", 0),
}

// KindMap contains kind string overrides.
var KindMap = map[string]string{
	// PG
	"aiven_pg":          "Service",
	"aiven_pg_database": "Database",
	"aiven_pg_user":     "User",

	// MySQL
	"aiven_mysql":          "Service",
	"aiven_mysql_database": "Database",
	"aiven_mysql_user":     "User",
}
