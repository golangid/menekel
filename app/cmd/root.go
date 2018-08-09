package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golangid/menekel"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	RootCMD = &cobra.Command{
		Use:   "menekel",
		Short: "Article Management CLI",
	}
	articleUsecase    menekel.ArticleUsecase
	articleRepository menekel.ArticleRepository
	autorRepository   menekel.AuthorRepository
	configFile        string
	dbConn            *sql.DB
)

func Execute() {
	if err := RootCMD.Execute(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig, initHTTPServiceDependecies)
	RootCMD.PersistentFlags().StringVar(&configFile, `config`, `./config.json`, `JSON file consists all configurations`)
}
func initConfig() {
	viper.SetConfigType("json")
	if configFile != "" {
		viper.SetConfigFile(configFile)
	}
	viper.AddConfigPath(".")
	logrus.SetFormatter(&logrus.JSONFormatter{})
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		logrus.Errorf("Cannot use config: %v. Got error:%v ", viper.ConfigFileUsed(), err)
		os.Exit(1)
	}
	logrus.Info("Using Config file: ", viper.ConfigFileUsed())
	if viper.GetBool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Warn("Menekel is Running in Debug Mode")
		return
	}
	logrus.SetLevel(logrus.InfoLevel)
	logrus.Warn("Menekel is Running in Production Mode")
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
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer dbConn.Close()

}
