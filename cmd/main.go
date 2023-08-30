package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Main(args []string) {
	var cmd = &cobra.Command{
		Use:     "ytools [command]",
		Short:   "This is a tools for check middleware status",
		Version: "0.1",
	}

	var kafkaCmd = &cobra.Command{
		Use:   "kafka",
		Short: "Check kafka status",
		Run:   kafkaCommand,
	}

	kafkaCmd.Flags().StringP("brokers", "b", "localhost:9092", "Set the broker addr for kafka")

	cmd.AddCommand(kafkaCmd)

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
