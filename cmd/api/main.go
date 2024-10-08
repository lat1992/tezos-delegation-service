package main

import (
	"log/slog"

	"github.com/lat1992/tezos-delegation-service/database"
	"github.com/lat1992/tezos-delegation-service/external"
	"github.com/lat1992/tezos-delegation-service/handlers"
	"github.com/lat1992/tezos-delegation-service/services"
	"github.com/spf13/viper"
)

func main() {
	slog.Info("starting service")

	viper.SetDefault("port", "8080")
	viper.SetDefault("database_host", "localhost")
	viper.SetDefault("database_port", "5432")
	viper.SetDefault("database_database", "app")
	viper.SetDefault("database_user", "postgres")
	viper.SetDefault("tezos_api", "https://api.tzkt.io/v1")
	viper.AutomaticEnv()

	db, err := database.NewStore(viper.GetString("database_host"), viper.GetString("database_port"), viper.GetString("database_database"), viper.GetString("database_user"), viper.GetString("database_password"))
	if err != nil {
		slog.Error("error when create new storage", "error", err)
		return
	}
	defer db.Close()

	tc := external.NewTezosClient(viper.GetString("tezos_api"))

	slog.Info("initializing...")
	tds := services.NewTezosDelegation(db, tc)
	if err := tds.Index(true); err != nil {
		slog.Error("error when initializing", "error", err)
	}
	slog.Info("initialize end")

	go func() {
		if err = tds.Start(); err != nil {
			slog.Error("error in tezos delegation service", "error", err)
		}
	}()

	router := handlers.GetRouter(tds)

	slog.Info("service started")
	if err := router.Run(":" + viper.GetString("port")); err != nil {
		slog.Error("error when service start", "error", err)
	}
}
