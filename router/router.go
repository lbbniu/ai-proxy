package router

import (
	"github.com/gin-gonic/gin"

	"github.com/lbbniu/ai-proxy/api"
)

func New() *gin.Engine {
	g := gin.New()
	// [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
	g.SetTrustedProxies([]string{"*"})
	g.Use(gin.Recovery())

	// 注册路由
	addRouter(g)

	return g
}

func addRouter(g *gin.Engine) {
	// 参考微软 和 OpenAi 接口路径规则
	r := g.Group("/:serviceProvider/deployments/:deployment")
	{
		r.POST("audio/transcriptions", api.ProxyNotImplemented)
		r.POST("chat/completions", api.ProxyNotImplemented)
		r.POST("embeddings", api.ProxyNotImplemented)
		r.POST("images/generations", api.ProxyNotImplemented)
	}
}
