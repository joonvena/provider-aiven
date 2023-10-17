/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	database "github.com/joonvena/provider-aiven/internal/controller/mysql/database"
	service "github.com/joonvena/provider-aiven/internal/controller/mysql/service"
	user "github.com/joonvena/provider-aiven/internal/controller/mysql/user"
	databasepostgres "github.com/joonvena/provider-aiven/internal/controller/postgres/database"
	servicepostgres "github.com/joonvena/provider-aiven/internal/controller/postgres/service"
	userpostgres "github.com/joonvena/provider-aiven/internal/controller/postgres/user"
	providerconfig "github.com/joonvena/provider-aiven/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		database.Setup,
		service.Setup,
		user.Setup,
		databasepostgres.Setup,
		servicepostgres.Setup,
		userpostgres.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
