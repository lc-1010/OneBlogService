package ping

import (
	"blog-server/global"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Ping struct{}

func NewPing() Ping {
	return Ping{}
}
func (p *Ping) Pong(c *gin.Context) {
	global.Logger.SetTraceInfo(c).Infof(c, "%s for test ping,path:%s", c.HandlerName(), c.Request.URL.Path)

	c.JSON(http.StatusOK, map[string]string{
		"msg": "pong",
	})
}
