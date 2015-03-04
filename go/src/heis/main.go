
package main
import(
	"driver_elev"
	"yeye"
	)

func main(){
	
	init()

	go fsm()
	go sensor()
	go network()

	deadChan := make(chan int)
	<-deadChan
}
