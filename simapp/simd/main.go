package main

import (
	"os"

	"github.com/mydexchain/chain-sdk/server"
	svrcmd "github.com/mydexchain/chain-sdk/server/cmd"
	"github.com/mydexchain/chain-sdk/simapp"
	"github.com/mydexchain/chain-sdk/simapp/simd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, simapp.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
