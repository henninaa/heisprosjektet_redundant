package main

import( 
	"driver_elev"
	"network1"
	. "fmt"
	"time"
	)

func main(){
	driver_elev.Elev_init();
	driver_elev.Elev_stop_engine()
	time.Sleep(1000 * time.Millisecond)
	driver_elev.Elev_start_engine(true)

	state := 0;
	var gotofloor int;
	cfloor := 0;

	network1.Network()

	for(true){

		switch (state){
		case 0:

			
			if driver_elev.Elev_get_button_signal(2, 0) {
				gotofloor = 0;
				state = 1;

			}else if driver_elev.Elev_get_button_signal(2, 1) {
				gotofloor = 1;
				state = 1;

			}else if driver_elev.Elev_get_button_signal(2, 2) {
				gotofloor = 2;
				state = 1;

			}else if driver_elev.Elev_get_button_signal(2, 3) {
				gotofloor = 3;
				state = 1;
				

			}
		
		case 1:
			Println("aaa")
			if(gotofloor == cfloor){
				state = 0;
			}else if(gotofloor < cfloor){
				driver_elev.Elev_start_engine(false);
			}else if(gotofloor > cfloor){
				driver_elev.Elev_start_engine(true);
			}
			state = 2;
			Print("uuuuu")

		case 2:
			
			if(driver_elev.Elev_get_floor_sensor_signal() == gotofloor){
				driver_elev.Elev_stop_engine();
				state = 0;
				cfloor = gotofloor;
				Println("ee");
			}
		}



	}


}
