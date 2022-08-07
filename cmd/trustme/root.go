package trustme

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "trustme",
	Short: "TrustMe - a simple CLI to manage Trust Relationship of AWS IAM Role",
	Long: `TrustMe - a simple CLI to manage Trust Relationship of AWS IAM Role.

If you want to know about Trust Relationship, then recommend read this article: 
  https://aws.amazon.com/blogs/security/how-to-use-trust-policies-with-iam-roles/`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error!, '%s'", err)
		os.Exit(1)
	}
}
