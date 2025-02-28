package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitMemberRouter(r *gin.Engine) {
	member := r.Group("/member")
	{
		member.GET("/heartbeat", MemberHeartbeat)
	}
}

func MemberHeartbeat(c *gin.Context) {
	c.String(http.StatusOK, "member ok")
}
