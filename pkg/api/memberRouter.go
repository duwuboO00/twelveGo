package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MemberHeartbeat(c *gin.Context) {
	c.String(http.StatusOK, "member ok")
}
