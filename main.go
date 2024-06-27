package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"

	"github.com/RoshanShrestha123/sconfig/utils"
	"github.com/common-nighthawk/go-figure"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

type Info struct {
	Os               string
	Ram              int
	Core             int
	BootTime         time.Time
	Version          string
	UsedRamInPercent float64
}

func main() {
	fig := figure.NewFigure("SCONFIG", "", true)
	totalCpuCount, _ := cpu.Counts(false)
	bootTime, _ := host.BootTime()
	version, _ := host.KernelVersion()
	fig.Print()
	v, _ := mem.VirtualMemory()
	_, _, totalGb := utils.MemUnitConversion(int(v.Total))

	info := Info{
		Os:               runtime.GOOS,
		Ram:              totalGb,
		Core:             totalCpuCount,
		BootTime:         time.Unix(int64(bootTime), 0),
		Version:          version,
		UsedRamInPercent: v.UsedPercent,
	}

	values := reflect.ValueOf(info)
	typeOfField := values.Type()
	fmt.Println()

	for i := 0; i < values.NumField(); i++ {
		fmt.Printf("%v : %v\n", typeOfField.Field(i).Name, values.Field(i).Interface())
	}

}
