package main

import (
	"fmt"
	"os"

	common "github.com/apurer/go-libiptc"
	"github.com/apurer/go-libiptc/libip4tc"
	"github.com/apurer/go-libiptc/libip6tc"
)

func showSyntax() {
	fmt.Fprintf(os.Stderr, "Usage: dump-table-rules [-4|-6] <table-name>\n")
	os.Exit(1)
}

func main() {
	if len(os.Args) != 3 {
		showSyntax()
	}

	var useIpv6 bool
	switch os.Args[1] {
	case `-4`:
	case `-6`:
		useIpv6 = true
	default:
		showSyntax()
	}

	acquired, err := common.XtablesLock(false, 0)
	if err != nil {
		panic(err)
	}
	if !acquired {
		fmt.Fprintf(os.Stderr, "dump-table-rules: could not acquire xtables lock!\n")
		os.Exit(1)
	}
	defer func() {
		_, err := common.XtablesUnlock()
		if err != nil {
			panic(err)
		}
	}()

	tableName := os.Args[2]

	if !useIpv6 {
		table, err := libip4tc.TableInit(tableName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dump-table-rules: %s\n", err)
			os.Exit(3)
		}
		defer table.Free()

		// traverse trough chains
		chain, err := table.FirstChain()
		if err != nil {
			fmt.Fprintf(os.Stderr, "dump-table-rules: %s\n", err)
			os.Exit(3)
		}
		for chain != "" {
			// use Go-native rules conversion
			e, err := table.FirstRule(chain)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dump-table-rules: %s\n", err)
				os.Exit(3)
			}
			for !e.IsEmpty() {
				fmt.Println(chain+":", table.IptEntry2Rule(&e).String())
				e, err = table.NextRule(e)
				if err != nil {
					fmt.Fprintf(os.Stderr, "dump-table-rules: %s\n", err)
					os.Exit(3)
				}
			}

			chain, err = table.NextChain()
			if err != nil {
				fmt.Fprintf(os.Stderr, "dump-table-rules: %s\n", err)
				os.Exit(3)
			}
		}
	} else {
		table, err := libip6tc.TableInit(tableName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dump-table-rules: %s\n", err)
			os.Exit(3)
		}
		defer table.Free()

		// traverse trough chains
		chain, err := table.FirstChain()
		if err != nil {
			fmt.Fprintf(os.Stderr, "dump-table-rules: %s\n", err)
			os.Exit(3)
		}
		for chain != "" {
			// use Go-native rules conversion
			e, err := table.FirstRule(chain)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dump-table-rules: %s\n", err)
				os.Exit(3)
			}
			for !e.IsEmpty() {
				fmt.Println(chain+":", table.IptEntry2Rule(&e).String())
				e, err = table.NextRule(e)
				if err != nil {
					fmt.Fprintf(os.Stderr, "dump-table-rules: %s\n", err)
					os.Exit(3)
				}
			}

			chain, err = table.NextChain()
			if err != nil {
				fmt.Fprintf(os.Stderr, "dump-table-rules: %s\n", err)
				os.Exit(3)
			}
		}
	}
}
