package main

import (
	"github.com/111/111/Plugins"
	"github.com/111/111/common"
)

func main() {
	var Info common.HostInfo
	common.Flag(&Info)
	common.Parse(&Info)
	Plugins.Scan(Info)
	print("scan end\n")
}
