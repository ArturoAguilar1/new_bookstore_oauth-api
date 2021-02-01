package app
import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func Pong() {
	fmt.Println("Ponggggg")
}

func StartApplication() {
	//router.GET("/oauth/access_token/:access_token_id",Pong())
	router.Run(":8080")
	fmt.Println("funcionando")
}
