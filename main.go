package main

///////////////////////////////////////////////////////////////////////////////////////////
//    _    _   ___   _      _       _____  _    _   ___  ______ ______  _____ ______     //
//   | |  | | / _ \ | |    | |     /  ___|| |  | | / _ \ | ___ \| ___ \|  ___|| ___ \    //
//   | |  | |/ /_\ \| |    | |     \ `--. | |  | |/ /_\ \| |_/ /| |_/ /| |__  | |_/ /    //
//   | |/\| ||  _  || |    | |      `--. \| |/\| ||  _  ||  __/ |  __/ |  __| |    /     //
//   \  /\  /| | | || |____| |____ /\__/ /\  /\  /| | | || |    | |    | |___ | |\ \     //
//    \/  \/ \_| |_/\_____/\_____/ \____/  \/  \/ \_| |_/\_|    \_|    \____/ \_| \_|    //
//                                                                                       //
///////////////////////////////////////////////////////////////////////////////////////////

import (
	"time"

	_ "github.com/pbnjay/strptime"
	"github.com/reujab/wallpaper"
)

func main() {
	//// ~ DEFINE VARIABLES ~ ////

	morningTime := "00:00"
	eveningTime := "12:00"
	nightTime := "19:00"

	//date format
	layout := "15:04"
	morningTimef, err := time.Parse(layout, morningTime)
	eveningTimef, err := time.Parse(layout, eveningTime)
	nightTimef, err := time.Parse(layout, nightTime)
	print(err)

	daypart := "default"
	setimg := "defaultimg"
	presentImg := "default"

	morningImg := "C:\\Users\\User\\go\\GoPaper\\media\\morning.jpg"
	eveningImg := "C:\\Users\\User\\go\\GoPaper\\media\\evening.jpg"
	nightImg := "C:\\Users\\User\\go\\GoPaper\\media\\night.jpg"

	running := true

	//// ~ PROGRAM ~ ////

	for running == true {

		presentTime := time.Now()
		presentTimestr := presentTime.Format(layout)
		print(presentTimestr)
		presentTimef, err := time.Parse(layout, presentTimestr)
		print(err)

		//// ~ define day part ~ ////
		//morning
		if presentTimef.After(morningTimef) && presentTimef.Before(nightTimef) && presentTimef.Before(eveningTimef) {
			daypart = "morning"
			setimg = morningImg

			//night
		} else if presentTimef.After(eveningTimef) && presentTimef.After(morningTimef) && presentTimef.After(nightTimef) {
			daypart = "night"
			setimg = nightImg

			//evening
		} else if presentTimef.After(eveningTimef) && presentTimef.After(morningTimef) && presentTimef.Before(nightTimef) {
			daypart = "evening"
			setimg = eveningImg

		} else {
			print("\nerror\n")
		}
		print("\n" + setimg + " is for " + daypart + "\n")

		//// ~ change picture ~ ////
		//if good picture is not setted
		if setimg != presentImg {
			//change wallpaper
			if daypart == "morning" {
				//change for morning
				wallpaper.SetFromFile(setimg)
				presentImg := setimg
				print("Wallpaper changed by " + presentImg)

			} else if daypart == "evening" {
				//change for evening
				wallpaper.SetFromFile(setimg)
				presentImg := setimg
				print("Wallpaper changed by " + presentImg)

			} else {
				//change for night
				wallpaper.SetFromFile(setimg)
				presentImg := setimg
				print("Wallpaper changed by " + presentImg)
			}

		} else {
			print("Good picture already here")
		}

		time.Sleep(60000 * time.Millisecond)
	}
}
