package network1
 
func Network(){
 
	go SendImAlive()
	go RecieveImAlive()
	go PrintRecievedMessages()
 }
