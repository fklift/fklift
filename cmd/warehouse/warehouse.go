package warehouse

import (
	"github.com/fklift/fklift/cmd"

	"github.com/spf13/cobra"
)

var warehouseCmd = &cobra.Command{
	Use:   "warehouse",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func init() {
	cmd.RootCmd.AddCommand(warehouseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// warehouseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// warehouseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
