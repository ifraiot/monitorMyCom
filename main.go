package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/distatus/battery"
	"github.com/ifraiot/monitorMyCom/ifrasdk"
	"github.com/mackerelio/go-osstat/memory"
	"github.com/shirou/gopsutil/cpu"
)

func main() {

	host := flag.String("host", "", "MQTT host")
	topic := flag.String("topic", "", "Organization's topic")
	username := flag.String("username", "", "MQTT username")
	password := flag.String("password", "", "MQTT password")

	flag.Parse()

	ifraCon := ifrasdk.NewIFRA(
		*host,
		*topic,
		*username,
		*password)

	for {
		memory, err := memory.Get()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			return
		}

		// fmt.Printf("memory total: %d bytes\n", memory.Total)
		// fmt.Printf("memory used: %d bytes\n", memory.Used)
		// fmt.Printf("memory cached: %d bytes\n", memory.Cached)
		// fmt.Printf("memory free: %d bytes\n", memory.Free)

		percent, _ := cpu.Percent(time.Second, true)
		// fmt.Printf("  User: %.2f\n", percent[cpu.CPUser])
		// fmt.Printf("  Nice: %.2f\n", percent[cpu.CPNice])
		// fmt.Printf("   Sys: %.2f\n", percent[cpu.CPSys])
		// fmt.Printf("  Intr: %.2f\n", percent[cpu.CPIntr])
		// fmt.Printf("  Idle: %.2f\n", percent[cpu.CPIdle])
		// fmt.Printf("States: %.2f\n", percent[cpu.CPUStates])

		//Memory
		ifraCon.AddMeasurement("memory_total", float64(memory.Total)/1024/1024)
		ifraCon.AddMeasurement("memory_used", float64(memory.Used)/1024/1024)
		ifraCon.AddMeasurement("memory_cached", float64(memory.Cached)/1024/1024)
		ifraCon.AddMeasurement("memory_free", float64(memory.Free)/1024/1024)

		//CPU
		ifraCon.AddMeasurement("cpu_sys", percent[cpu.CPSys])
		ifraCon.AddMeasurement("cpu_user", percent[cpu.CPUser])
		ifraCon.AddMeasurement("cpu_usage", percent[cpu.CPUser]+percent[cpu.CPSys])
		ifraCon.AddMeasurement("cpu_idle", percent[cpu.CPIdle])
		ifraCon.Send()

		//battery
		batteries, err := battery.GetAll()
		if err != nil {
			fmt.Println("Could not get battery info!")
			return
		}
		for _, battery := range batteries {

			if battery.State.String() == "Full" {
				ifraCon.AddMeasurement("battery_state", 1)
			} else {
				ifraCon.AddMeasurement("battery_state", 0)
			}

			ifraCon.AddMeasurement("battery_capacity", battery.Current)
			ifraCon.AddMeasurement("battery_last_capacity", battery.Full)
			ifraCon.AddMeasurement("battery_charge_rate", battery.Current/battery.Full*100)

			// fmt.Printf("current capacity: %f mWh, ", battery.Current)
			// fmt.Printf("last full capacity: %f mWh, ", battery.Full)
			// fmt.Printf("design capacity: %f mWh, ", battery.Design)
			// fmt.Printf("charge rate: %f mW, ", battery.ChargeRate)
			// fmt.Printf("voltage: %f V, ", battery.Voltage)
			// fmt.Printf("design voltage: %f V\n", battery.DesignVoltage)
		}

		time.Sleep(5 * time.Second)
	}
}
