package main

import (
	"github.com/lbbniu/ai-proxy/router"
)

func main() {
	g := router.New()

	g.Run()
}
