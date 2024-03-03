package api

import "strings"

type ProxyApi int

const (
	ProxyApiUnknown ProxyApi = iota
	ProxyApiAudioTranscription
	ProxyApiChatCompletions
	ProxyApiEmbeddings
	ProxyApiImagesGenerations
)

func Path2ProxyApi(path string) ProxyApi {
	if strings.HasSuffix(path, "audio/transcriptions") {
		return ProxyApiAudioTranscription
	} else if strings.HasSuffix(path, "chat/completions") {
		return ProxyApiChatCompletions
	} else if strings.HasSuffix(path, "embeddings") {
		return ProxyApiEmbeddings
	} else if strings.HasSuffix(path, "images/generations") {
		return ProxyApiImagesGenerations
	} else {
		return ProxyApiUnknown
	}
}
