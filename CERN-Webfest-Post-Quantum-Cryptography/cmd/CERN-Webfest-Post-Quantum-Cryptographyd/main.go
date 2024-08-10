package main

import (
	"fmt"
	"os"

	clienthelpers "cosmossdk.io/client/v2/helpers"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/EishaMazhar/CERN-Webfest-Post-Quantum-Cryptography/app"
	"github.com/EishaMazhar/CERN-Webfest-Post-Quantum-Cryptography/cmd/CERN-Webfest-Post-Quantum-Cryptographyd/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, clienthelpers.EnvPrefix, app.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}
}
