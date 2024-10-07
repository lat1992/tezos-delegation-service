package main

import (
	"log/slog"

	"github.com/lat1992/tezos-delegation-service/handlers"
	"github.com/lat1992/tezos-delegation-service/services"
	"github.com/lat1992/tezos-delegation-service/storages"
	"github.com/spf13/viper"
)

func main() {
	slog.Info("starting service")

	viper.SetDefault("port", "8080")
	viper.SetDefault("storage_host", "localhost")
	viper.SetDefault("storage_port", "5432")
	viper.SetDefault("storage_database", "app")
	viper.SetDefault("storage_user", "postgres")
	viper.SetDefault("storage_password", "postgres")
	viper.AutomaticEnv()

	db, err := storages.NewStorage(viper.GetString("storage_host"), viper.GetString("storage_port"), viper.GetString("storage_database"), viper.GetString("storage_user"), viper.GetString("storage_password"))
	if err != nil {
		slog.Error("error when create new storage", "error", err)
		return
	}
	defer db.Close()

	tds := services.NewTezosDelegation()

	router := handlers.GetRouter(tds)

	slog.Info("service started")
	if err := router.Run(":" + viper.GetString("port")); err != nil {
		slog.Error("error when service start", "error", err)
	}
}
