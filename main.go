package main

import (
	"LotusPart2/conf"
	"LotusPart2/pkg/route"
	"context"
	"os"

	_ "LotusPart2/docs"
	"gitlab.com/goxp/cloud0/logger"
)

const (
	APPNAME = "eLotus SVC"
)

// @title eLotus SVC API
// @version 1.0
// @description This is eLotus SVC api docs.
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

	app := route.StartNewService()
	ctx := context.Background()
	err := app.Start(ctx)
	if err != nil {
		logger.Tag("main").Error(err)
	}
	os.Clearenv()
}
