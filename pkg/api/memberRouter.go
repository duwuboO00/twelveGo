package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func MemberHeartbeat(c *gin.Context) {
	c.String(http.StatusOK, "member ok")
}
