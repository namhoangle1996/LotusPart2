package main

import (
	"context"
	"finan/mvt-adapter/conf"
	"finan/mvt-adapter/pkg/route"
	"finan/mvt-adapter/pkg/utils"
	"os"

	"gitlab.com/goxp/cloud0/logger"

	_ "finan/mvt-adapter/docs"
)

const (
	APPNAME = "DEMO"
)

// @title DEMO API
// @version 1.0
// @description This is DEMO api docs.
// @termsOfService http://swagger.io/terms/

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost
// @BasePath  /
func main() {
	conf.SetEnv()
	logger.Init(APPNAME)
	utils.LoadMessageError()

	app := route.NewService()
	ctx := context.Background()
	err := app.Start(ctx)
	if err != nil {
		logger.Tag("main").Error(err)
	}
	os.Clearenv()
}
