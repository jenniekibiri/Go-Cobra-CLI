
package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "flourish",
    Short: "Flourish Coding Challenge CLI",
    Long:  "A CLI application for managing customers and orders.",
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
	
}
