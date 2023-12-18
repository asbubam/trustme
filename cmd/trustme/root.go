package trustme

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"os"

	"github.com/asbubam/trustme/pkg/trustme"
)

var whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Check your current IAM identity.",
	Run: func(cmd *cobra.Command, args []string) {
		arn, err := trustme.GetCurrentIAMArn(context.TODO())
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(arn)
	},
}

var AssumeCmd = &cobra.Command{
	Use:   "assume",
	Short: "Assume IAM Role",
	Run: func(cmd *cobra.Command, args []string) {
		arn, err := trustme.AssumeRole(context.TODO(), args[0], args[1])
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(arn)
	},
}

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

func init() {
	rootCmd.AddCommand(whoamiCmd)
	rootCmd.AddCommand(AssumeCmd)
}
