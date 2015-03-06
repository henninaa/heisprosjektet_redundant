
package main
import(
	"driver_elev"
	)

func main(){
	
	init()

	go network()

	deadChan := make(chan int)
	<-deadChan
}
