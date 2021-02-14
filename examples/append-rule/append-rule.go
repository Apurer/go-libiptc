package main

import (
	"fmt"
	"os"

	common "github.com/apurer/go-libiptc"
	"github.com/apurer/go-libiptc/libip4tc"
	"github.com/apurer/go-libiptc/libip6tc"
)


func main() {
	h, err := TableInit("filter")
	if err != nil {
		fmt.Fprintf(os.Stderr, "append-rule: %s\n", err)
		os.Exit(3)
	}

	var a C.uint
	

	entry := new(IptEntry)
	entry.ip.src.s_addr = 

	var c common.XtChainLabel = "INPUT"

	h.AppendEntry(c,)
}
