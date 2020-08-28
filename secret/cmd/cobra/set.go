package cobra

import (
	"fmt"

	"github.com/santosh/gophercises/secret"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets a secret in your secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		v := secret.File(encodingKey, secretsPath())
		key, value := args[0], args[1]
		err := v.Set(key, value)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Value set successfully!")
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}
