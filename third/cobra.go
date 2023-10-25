package third

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go_study",
	Short: "go_study is a project for recording go learning process",
	Long: `go_study is a project for recording go learning process, 
				  Complete url is available at https://github.com/heshanu3d/go_study`,
	Run: func(cmd *cobra.Command, args []string) {
	  fmt.Println("run hugo...")
	  fmt.Println("args: ", args)
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of go_study",
	Long: "Print the version number of go_study",
	Run: func(cmd *cobra.Command, args []string) {
	  fmt.Println("v0.1")
	},
}
  
func Execute() {
	rootCmd.AddCommand(versionCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}