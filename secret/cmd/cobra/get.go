package cobra

import (
	"fmt"

	"github.com/santosh/gophercises/secret"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a secret from your secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		v := secret.File(encodingKey, secretsPath())
		fmt.Println(args)
		key := args[0]
		value, err := v.Get(key)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s=%s\n", key, value)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
