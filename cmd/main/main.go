package main

import (
	"github.com/macox/buildpack-template/helloworld"
	"os"

	"github.com/paketo-buildpacks/libpak"
	"github.com/paketo-buildpacks/libpak/bard"
)

func main() {
	libpak.Main(
		helloworld.Detect{},
		helloworld.Build{Logger: bard.NewLogger(os.Stdout)},
	)
}
