package main
import(
	"awesomeProject1/xsbtest/cmd"
	"os"
)
func main(){
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}

}
