package network1

import (
	"fmt"
	"net"
	"time"
	"encoding/json"
	)

var threadMsg = make(chan int,1)

const(
	port = ":20003"
	IP = "129.241.187.141"
	primary = 1
	secondary = 0
)
type Jasonstruct struct {
		Msg int
	}

func transmit(message int){

	var tmp Jasonstruct

	tmp.Msg = message

	addr, err := net.ResolveUDPAddr("udp", IP+port)
	handleError(err)

	sock, err := net.DialUDP("udp",nil,addr)
	handleError(err)

	//message := ""
	//fmt.Scanf("%s", &message)

	fmt.Println("I'm Alive!!!! WOHOHOHOHOHOOOOO!!")
	bmsg, err := json.Marshal(tmp)
	//fmt.Println(bmsg)
	sock.Write(bmsg)
	handleError(err)
}



func recieve() {

	buffer := make([]byte, 1024)
	var integer Jasonstruct

	addr, _ := net.ResolveUDPAddr("udp",port)
	sock, _ := net.ListenUDP("udp", addr)

	for{
		n, _,error := sock.ReadFromUDP(buffer)
		if error != nil{
			handleError(error)
		}

		json.Unmarshal(buffer[:n], &integer)

		threadMsg <- integer.Msg
	}
}


func handleError(err error){
	if err != nil{
		fmt.Println(err)
	}

}

func SendImAlive(){

	for{
		transmit(-1)
		time.Sleep(100*time.Millisecond)
	}

}

/*
func main(){

	var isPrimary bool
	timer := (time.Now().Add(3 * time.Second))
	counter := 0

	isPrimary = false

	go recieve()
	

	for{

		if(!isPrimary){


			
			select{

			case msg := <- threadMsg:

				if(msg == -1){
					timer = (time.Now().Add(1 * time.Second))
				}else{
					counter = msg
				}
				fmt.Println("Recieved: ", counter)

			default:
				if(time.Now().After(timer)){
					isPrimary = true
				}
				fmt.Println(counter)
			}
			

		} else {

				transmit(-1)
				transmit(counter)

				fmt.Println(counter)

				counter += 1

		}
		time.Sleep(500*time.Millisecond)
	}


}
*/