package main

import (
	"embed"

	"github.com/Fuabioo/altalune/cmd"
)

//go:embed frontend/dist
var staticFiles embed.FS

func main() {
	cmd.Execute(staticFiles)
}
