package main

import (
	_ "rssplus/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"rssplus/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
