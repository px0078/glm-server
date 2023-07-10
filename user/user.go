package user


import (
	"net/http"
	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)


/**
通过用户id获取用户信息
**/ 
func GetUser(c *gin.Context) {
	user := c.Params.ByName("name")
	value, ok := db[user]
	if ok {
		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	}
}