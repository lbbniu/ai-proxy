
<div align="center">

# AI Proxy

_✨ 通过对应厂商大模型标准的 API 格式进行访问，开箱即用 ✨_

</div>

## 功能
1. 支持多种大模型：
   + [ ] [OpenAI ChatGPT 系列模型](https://platform.openai.com/docs/guides/gpt/chat-completions-api)（支持 [Azure OpenAI API](https://learn.microsoft.com/en-us/azure/ai-services/openai/reference)）
   + [ ] [Anthropic Claude 系列模型](https://anthropic.com)
   + [ ] [Google PaLM2/Gemini 系列模型](https://developers.generativeai.google)
   + [ ] [百度文心一言系列模型](https://cloud.baidu.com/doc/WENXINWORKSHOP/index.html)
   + [ ] [阿里通义千问系列模型](https://help.aliyun.com/document_detail/2400395.html)
   + [ ] [讯飞星火认知大模型](https://www.xfyun.cn/doc/spark/Web.html)
   + [ ] [智谱 ChatGLM 系列模型](https://bigmodel.cn)
   + [ ] [360 智脑](https://ai.360.cn)
   + [ ] [腾讯混元大模型](https://cloud.tencent.com/document/product/1729)
   + [ ] [Moonshot AI](https://platform.moonshot.cn/)
   + [ ] [字节云雀大模型](https://www.volcengine.com/product/ark)
   + [ ] [MINIMAX](https://api.minimax.chat/)
2. 支持配置镜像以及众多[第三方代理服务](https://iamazing.cn/page/openai-api-third-party-services)。
3. 支持通过**负载均衡**的方式访问多个渠道。
4. 支持 **stream 模式**，可以通过流式传输实现打字机效果。
5. 支持**多机部署**，[详见此处](#多机部署)。
6. 支持**令牌管理**，设置令牌的过期时间和配额。
7. 支持失败自动重试。
8. 支持绘图接口。
9. 支持用户管理，支持**多种用户登录注册方式**：
    + 邮箱登录注册（支持注册邮箱白名单）以及通过邮箱进行密码重置。
    + [GitHub 开放授权](https://github.com/settings/applications/new)。
    + 微信公众号授权（需要额外部署 [WeChat Server](https://github.com/songquanpeng/wechat-server)）。
## 部署
### 基于 Docker 进行部署
```shell
# 使用 SQLite 的部署命令
# 使用 MySQL 的部署命令
```

### 基于 Docker Compose 进行部署

> 仅启动方式不同，参数设置不变，请参考基于 Docker 部署部分

```shell
# 目前支持 MySQL 启动，数据存储在 ./data/mysql 文件夹内
docker-compose up -d

# 查看部署状态
docker-compose ps
```

## 配置
系统本身开箱即用。

你可以通过设置环境变量或者命令行参数进行配置。

等到系统启动后，使用 `lbbniu` 用户登录系统并做进一步的配置。

**Note**：如果你不知道某个配置项的含义，可以临时删掉值以看到进一步的提示文字。

## 相关项目
* [FastGPT](https://github.com/labring/FastGPT): 基于 LLM 大语言模型的知识库问答系统
* [ChatGPT Next Web](https://github.com/Yidadaa/ChatGPT-Next-Web):  一键拥有你自己的跨平台 ChatGPT 应用
* [One API](https://github.com/songquanpeng/one-api):  通过标准的 OpenAI API 格式访问所有的大模型，开箱即用


## 注意

本项目使用 GPL 协议进行开源，**在此基础上**，必须在页面底部保留署名以及指向本项目的链接。如果不想保留署名，必须首先获得授权。

同样适用于基于本项目的二开项目。

依据 GPL 协议，使用者需自行承担使用本项目的风险与责任，本开源项目开发者与此无关。
