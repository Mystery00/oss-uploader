package main

import (
	"github.com/spf13/cobra"
	"go-cli-template/cmd"
	"go-cli-template/logger"
)

func main() {
	cobra.OnInitialize(logger.InitLog)
	cmd.Execute()
}
