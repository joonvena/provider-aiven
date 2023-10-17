package mysql

import "github.com/upbound/upjet/pkg/config"

const (
	ConfigPackagePath = "github.com/joonvena/provider-aiven/config"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aiven_mysql_database", func(r *config.Resource) {
		r.References["service_name"] = config.Reference{
			Type: "Service",
		}
	})

	p.AddResourceConfigurator("aiven_mysql_user", func(r *config.Resource) {
		r.References["service_name"] = config.Reference{
			Type: "Service",
		}
	})
}
