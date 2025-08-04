package cmd

import (
	"elotus/cmd/dsa"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "elotus-interview",
	Short: "elotus-interview include dsa and authentication , upload file",
	Long:  `elotus-interview include dsa and authentication , upload file`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(dsaCmd)
}

var dsaCmd = &cobra.Command{
	Use:   "dsa [algorithm]",
	Short: "run dsa test",
	Long:  `run dsa test with specified algorithm (grayCode, sumOfDistancesInTree, findLength)`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		algorithm := args[0]
		fmt.Printf("Running DSA test for algorithm: %s\n", algorithm)
		dsa.RunDSATest(algorithm)
	},
}

var authenticationServer = &cobra.Command{
	Use:   "authentication",
	Short: "authentication server",
	Long:  `Authentication server,....`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run authentication server")
	},
}

var uploadServer = &cobra.Command{
	Use:   "upload",
	Short: "Upload server",
	Long:  `Upload file server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run upload server")
	},
}
