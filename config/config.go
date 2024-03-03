package config

type ServiceProvider string


const (
	Ai360 ServiceProvider = "ai360"
	Aliyun ServiceProvider = "aliyun"
	Anthropic ServiceProvider = "anthropic"
	BaiChuan ServiceProvider = "baichuan"
	Baidu ServiceProvider = "baidu"
	Gemini ServiceProvider = "gemini"
	MiniMax ServiceProvider = "minimax"
	MoonShot ServiceProvider = "moonshot"
	OpenAi ServiceProvider = "openai"
	Palm ServiceProvider = "palm"
	Tencent ServiceProvider = "tencent"
	XunFei ServiceProvider = "xunfei"
	ZhiPu ServiceProvider = "zhipu"
)

type Deployment struct {
	Name string
	ServiceProvider ServiceProvider
	Upstreams []Upstream
}

type Upstream struct {
	ApiKey string
	BaseURL string
}

type App struct {
	ID int64
	Deployment string
	ApiKey string
	Tpm int64
	Rpm int64
	Qps int64
}

type Config struct {
	Deployments []Deployment
	Apps []App
}