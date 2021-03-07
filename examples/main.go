package main

import (
	"fmt"
	"gpsd"
	"time"
)

func main() {
	removeFilterTest()
}

func setup() *gpsd.Session {
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

	return gps
}

func connectTest() {
	gps := setup()
	done := gps.Watch()
	<-done

}

func disconnectTest() {
	gps := setup()
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

func removeFilterTest() {
	gps := setup()
	done := gps.Watch()

	fmt.Println("Started watching")
	time.Sleep(3 * time.Second)

	fmt.Println("Removing TPV Filter")
	gps.RemoveFilter("TPV")

	time.Sleep(8 * time.Second)
	done <- true
	fmt.Println("Finished.")
}
