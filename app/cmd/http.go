package cmd

import (
	"time"

	"github.com/spf13/cobra"

	"github.com/golangid/menekel/article"
	"github.com/golangid/menekel/internal/database/mysql"
	delivery "github.com/golangid/menekel/internal/http"
	"github.com/golangid/menekel/internal/http/middleware"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

var (
	httpCmd = &cobra.Command{
		Use:   "http",
		Short: "Start HTTP REST API",
		Run:   initHTTP,
	}
)

func initHTTP(cmd *cobra.Command, args []string) {
	e := echo.New()
	middL := middleware.InitMiddleware()
	e.Use(middL.CORS)
	articleRepository = mysql.NewArticleRepository(dbConn)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	au := article.NewArticleUsecase(articleRepository, timeoutContext)
	delivery.InitArticleHandler(e, au)

	e.Start(viper.GetString("server.address"))
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
