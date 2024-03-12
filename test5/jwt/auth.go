package middleware

import (
	"2108-a-homework/test5/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

func AuthMiddleware(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	userID, err := utils.GetToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "invalid token")
		c.Abort()
	}
	c.Set("user_id", userID)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(9000) + 1000
	fmt.Println("生成的4位随机数为:", randomNumber)
}
