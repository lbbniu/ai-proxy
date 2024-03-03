package api

import (
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"

	"github.com/lbbniu/ai-proxy/proxy/api"
	"github.com/lbbniu/ai-proxy/proxy/model"
	"github.com/lbbniu/ai-proxy/proxy/model/openai"
)

type Proxy struct {
}

func NewProxy() *Proxy {
	return new(Proxy)
}

func (p *Proxy) Proxy(ctx *gin.Context) {
	serviceProvider := ctx.Param("serviceProvider")
	deployment := ctx.Param("deployment")
	var m model.Model = openai.New()
	target, err := m.Target(api.ProxyApiChatCompletions)
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
		req.URL.Path, req.URL.RawPath = target.Path, target.RawPath
		for key, value := range m.Header() {
			req.Header[key] = value
		}
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
