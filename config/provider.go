/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	mysql "github.com/joonvena/provider-aiven/config/mysql"
	postgres "github.com/joonvena/provider-aiven/config/postgres"
)

const (
	resourcePrefix = "aiven"
	modulePath     = "github.com/joonvena/provider-aiven"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithShortName("aiven"),
		ujconfig.WithRootGroup("aiven.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			GroupKindOverrides(),
			KindOverrides(),
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		postgres.Configure,
		mysql.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
