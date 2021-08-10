package cmd

import (
	"github.com/spf13/cobra"
)

var(
	age int
	sex string
	country string
	name string

	rootCmd = &cobra.Command{
		Use:"person",
		Long:"a command write to practice cobra",

	}
)

func Execute() error {
	return rootCmd.Execute()
}
func init(){

	rootCmd.PersistentFlags().IntVarP(&age,"age","a",20,"age of this person")
	rootCmd.PersistentFlags().StringVarP(&sex,"sex","s","woman","sex of this person")
	rootCmd.PersistentFlags().StringVarP(&country,"country","c","China","Country of this person")
	rootCmd.PersistentFlags().StringVarP(&name,"name","n","Name","name of this person")

	rootCmd.AddCommand(eatCmd)
	rootCmd.AddCommand(sleepCmd)

}