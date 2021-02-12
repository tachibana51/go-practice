package web
import (
	"2fa/config"
	"github.com/labstack/echo"
)


func Init(config config.Config){
	e := echo.New()
	e.GET("/", func(c echo.Context) error {return c.String(200,"OK")})
	e.Logger.Fatal(e.Start("0.0.0.0:8888"))
}