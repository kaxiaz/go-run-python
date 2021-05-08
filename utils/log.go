package util

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UnprocessableLog(c *gin.Context, err error) {
	c.JSON(http.StatusUnprocessableEntity, "")
	log.Println(err.Error())
}
