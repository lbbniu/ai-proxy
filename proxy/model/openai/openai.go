package openai

import (
	"net/http"
	"net/url"

	"github.com/lbbniu/ai-proxy/proxy/api"
)

const (
	baseURL = "https://api.openai.com"
)

type OpenAi struct {
}

func New() *OpenAi {
	return &OpenAi{}
}

func (o *OpenAi) Target(proxyApi api.ProxyApi) (*url.URL, error) {
	var path string
	switch proxyApi {
	case api.ProxyApiChatCompletions:
		path = "/v1/chat/completions"
	}
	return url.Parse(baseURL + path)
}
func (o *OpenAi) Header() http.Header {
	header := http.Header{}
	// todo: 增加认证请求头
	return header
}
