package api

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

type Proxy struct {
}

func NewProxy() *Proxy {
	return new(Proxy)
}

func (p *Proxy) Proxy(ctx *gin.Context) {
	serviceProvider := ctx.Param("serviceProvider")
	deployment := ctx.Param("deployment")
	target, err := url.Parse("https://api.openai.com")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"serviceProvider": serviceProvider,
			"deployment":      deployment,
		})
		return
	}
	reverseProxy := httputil.NewSingleHostReverseProxy(target)
	defaultDirector := reverseProxy.Director
	reverseProxy.Director = func(req *http.Request) {
		defaultDirector(req)
	}
	reverseProxy.ModifyResponse = p.ModifyResponse
	reverseProxy.ServeHTTP(ctx.Writer, ctx.Request)
}

func (p *Proxy) ModifyResponse(resp *http.Response) error {
	return nil
}

func ProxyNotImplemented(ctx *gin.Context) {
	serviceProvider := ctx.Param("serviceProvider")
	deployment := ctx.Param("deployment")
	ctx.JSON(200, gin.H{
		"serviceProvider": serviceProvider,
		"deployment":      deployment,
		"uri":             ctx.FullPath(),
	})
}
