package model

import (
	"net/http"
	"net/url"

	"github.com/lbbniu/ai-proxy/proxy/api"
)

type Model interface {
	Target(api.ProxyApi) (*url.URL, error)
	Header() http.Header
}
