package nux

import (
	"fmt"
	"testing"
)

func TestMetrics(t *testing.T) {

	fmt.Println("======kernel======")
	fmt.Print("KernelMaxFiles:")
	fmt.Println(KernelMaxFiles())

	fmt.Print("KernelAllocateFiles:")
	fmt.Println(KernelAllocateFiles())

	fmt.Print("KernelMaxProc:")
	fmt.Println(KernelMaxProc())

	fmt.Print("KernelHostname:")
	fmt.Println(KernelHostname())

	fmt.Println("======loadavg======")
	fmt.Print("LoadAvg:")
	fmt.Println(LoadAvg())

	fmt.Println("======cpuinfo======")
	fmt.Print("NumCpu:")
	fmt.Println(NumCpu())

	fmt.Print("CpuMHz:")
	fmt.Println(CpuMHz())

	fmt.Println("======cpustat======")
	if ps, err := CurrentProcStat(); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Print("CPU :")
		fmt.Println(ps.Cpu)
		for i, o := range ps.Cpus {
			fmt.Printf("CPU%d:", i)
			fmt.Println(o)
		}
	}

	fmt.Println("======dfstat======")
	if L, err := ListMountPoint(); err != nil {
		fmt.Println("error:", err)
	} else {
		for _, arr := range L {
			fmt.Println(BuildDeviceUsage(arr[0], arr[1], arr[2]))
		}
	}

	fmt.Println("======NetIfs======")
	if L, err := NetIfs([]string{}); err != nil {
		fmt.Println("error:", err)
	} else {
		for _, i := range L {
			fmt.Println(i)
		}
	}

	fmt.Println("======ListDiskStats======")
	if L, err := ListDiskStats(); err != nil {
		fmt.Println("error:", err)
	} else {
		for _, i := range L {
			fmt.Println(i)
		}
	}

	fmt.Println("======MemInfo:======")
	fmt.Println(MemInfo())

}
