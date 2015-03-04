package main

import( 
	"driver_elev"
	. "fmt"
	)

func main(){
	driver_elev.Elev_init();
	driver_elev.Elev_set_speed(0)



	state := 0;
	var gotofloor int;
	cfloor := 0;

	for(true){

		switch (state){
		case 0:

			Println("ee");
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
				driver_elev.Elev_start_engine(true);
			}else if(gotofloor > cfloor){
				driver_elev.Elev_start_engine(false);
			}
			state = 2;

		case 2:
			Print("uuuuu")
			if(driver_elev.Elev_get_floor_sensor_signal() == gotofloor){
				driver_elev.Elev_stop_engine();
				state = 0;
				cfloor = gotofloor;
			}
		}




	}


}