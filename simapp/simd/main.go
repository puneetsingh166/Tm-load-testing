package main

import (
	"os"

	"github.com/puneetsingh166/tm-load-test/server"
	svrcmd "github.com/puneetsingh166/tm-load-test/server/cmd"
	"github.com/puneetsingh166/tm-load-test/simapp"
	"github.com/puneetsingh166/tm-load-test/simapp/simd/cmd"
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
