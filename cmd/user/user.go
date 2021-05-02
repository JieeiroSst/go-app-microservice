package main

import (
	"github.com/JIeeiroSst/go-app/config"
	"github.com/JIeeiroSst/go-app/internal/models/user"
	server "github.com/JIeeiroSst/go-app/internal/server/user"
	"github.com/JIeeiroSst/go-app/pkg/jwt"
	"github.com/JIeeiroSst/go-app/pkg/logger"
	"github.com/JIeeiroSst/go-app/pkg/postgres"
	"github.com/JIeeiroSst/go-app/utils"
	"github.com/labstack/gommon/log"
	"os"
)

func main(){
	log.Info("Starting server")
	configPath := config.GetConfigPath(os.Getenv("docker"))
	cfg, err := config.GetConfig(configPath)
	if err != nil {
		log.Fatalf("Loading config: %v", err)
	}
	appLogger := logger.NewApiLogger(cfg)
	appLogger.InitLogger()
	appLogger.Infof(
		"AppVersion: %s, LogLevel: %s, Mode: %s, SSL: %v",
		cfg.Server.AppVersion,
		cfg.Logger.Level,
		cfg.Server.Mode,
		cfg.Server.SSL,
	)
	appLogger.Infof("Success parsed config: %#v", cfg.Server.AppVersion)
	plsqlDB, err := postgres.NewPsqlDB(cfg)
	if err != nil {
		appLogger.Fatalf("PostgresSQL init: %s", err)
	}

	appLogger.Infof("PostgresSQL connected: %#v", plsqlDB.Logger)

	tokenUser :=jwt.NewTokenUser(user.Users{},cfg)
	hash:=utils.NewHash()

	s:=server.NewServer(plsqlDB,cfg,tokenUser,hash)

	appLogger.Fatal(s.Run())
}