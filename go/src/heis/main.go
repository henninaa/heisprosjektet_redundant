
package elevator

import(
	"driver"
	)

func main(){
	
	init()

	go fsm()
	go sensor()
	go network()

	deadChan := make(chan int)
	<-deadChan
}
