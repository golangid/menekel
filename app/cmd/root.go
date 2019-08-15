package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	// imported for Mysql driver
	_ "github.com/go-sql-driver/mysql"

	"github.com/golangid/menekel"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "menekel",
		Short: "Article Management CLI",
	}
	articleUsecase    menekel.ArticleUsecase
	articleRepository menekel.ArticleRepository
	dbConn            *sql.DB
)

// Execute will run the CLI app of menekel
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig, initHTTPServiceDependecies)
}

func initConfig() {
	viper.SetConfigType("toml")
	viper.SetConfigFile("config.toml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	logrus.Info("Using Config file: ", viper.ConfigFileUsed())

	if viper.GetBool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Warn("Comment service is Running in Debug Mode")
		return
	}
	logrus.SetLevel(logrus.InfoLevel)
	logrus.Warn("Comment service is Running in Production Mode")
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func initHTTPServiceDependecies() {
	// DATABASE
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	var err error
	dbConn, err = sql.Open(`mysql`, dsn)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
