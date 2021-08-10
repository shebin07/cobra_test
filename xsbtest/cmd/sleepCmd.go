package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

var(
	duration int
	startTime time.Time

	sleepCmd = &cobra.Command{
		Use:"sleep",
		Short: "used when sleep",
		Run: WhenSleep,
	}
)
func init(){
	sleepCmd.PersistentFlags().IntVarP(&duration,"duration","d",0,"how many hours the person sleep")
	//sleepCmd.PersistentFlags().StringVarP(&startTime,"startTime","s","not sleep yet","time hte person start to sleep")

}

func WhenSleep(cmd *cobra.Command,args[]string){
	startTime=time.Now()
	fmt.Println(name," slept ",duration," hours from ",startTime)
}