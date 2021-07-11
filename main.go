package main

import (
	"fmt"

	"github.com/mcholismalik/solid-go/dip"
	"github.com/mcholismalik/solid-go/isp"
	"github.com/mcholismalik/solid-go/lsp"
	"github.com/mcholismalik/solid-go/ocp"
	"github.com/mcholismalik/solid-go/srp"
)

func main() {
	fmt.Println("Learn solid principle in go")

	srp.Run()
	ocp.Run()
	lsp.Run()
	isp.Run()
	dip.Run()
}
