package main

import (
	"fmt"
	"gpsd"
	"time"
)

func main() {
	disconnectTest()
}

func connectTest() {
	var gps *gpsd.Session
	var err error

	if gps, err = gpsd.Dial(gpsd.DefaultAddress); err != nil {
		panic(fmt.Sprintf("Failed to connect to GPSD: %s", err))
	}

	gps.AddFilter("TPV", func(r interface{}) {
		tpv := r.(*gpsd.TPVReport)
		fmt.Println("TPV", tpv.Mode, tpv.Time)
	})

	skyfilter := func(r interface{}) {
		sky := r.(*gpsd.SKYReport)

		fmt.Println("SKY", len(sky.Satellites), "satellites")
	}

	gps.AddFilter("SKY", skyfilter)

	done := gps.Watch()
	<-done

}

func disconnectTest() {
	gps, err := gpsd.Dial(gpsd.DefaultAddress)
	if err != nil {
		panic("uh oh.")
	}

	gps.AddFilter("TPV", func(r interface{}) {
		tpv := r.(*gpsd.TPVReport)
		fmt.Println("TPV", tpv.Mode, tpv.Time)
	})

	gps.AddFilter("SKY", func(r interface{}) {
		sky := r.(*gpsd.SKYReport)

		fmt.Println("SKY", len(sky.Satellites), "satellites")
	})

	done := gps.Watch()

	fmt.Println("Sleeping...")
	time.Sleep(3 * time.Second)

	fmt.Println("Done sleeping")
	done <- true
	fmt.Println("Shut down gpsd watcher")

	fmt.Println("Sleeping...")
	time.Sleep(5 * time.Second)
	fmt.Println("Done sleeping")
}
