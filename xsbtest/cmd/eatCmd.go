package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var(
	food string
	drink string
	eatCmd = &cobra.Command{
		Use:"eat",
		Short:"used to eat food or drink",
		Run: WhenEat,
	}

)

func init(){
	eatCmd.Flags().StringVarP(&food,"food","f","nothing","give the person food to eat")
	eatCmd.Flags().StringVarP(&drink,"drink","d","nothing","give the person something to drink")

}
func WhenEat(cmd *cobra.Command,args []string){
	if food!="nothing" {
		fmt.Println(name," eat",food)
	}
	if drink!="nothing"{
		fmt.Println(name," drink",drink)
	}

}