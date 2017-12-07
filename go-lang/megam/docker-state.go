// Copyright 2013 go-dockerclient authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// https://github.com/docker/docker/blob/0d445685b8d628a938790e50517f3fb949b300e0/api/client/stats.go#L199

package main

import (
  "fmt"
  "github.com/fsouza/go-dockerclient"
)

type Stats struct{
	ContainerId  string
	Image        string
	MemoryUsage  uint64   //in bytes
	SystemMemory uint64
	CPUStats     CPUStats  //in percentage of total cpu used
	PreCPUStats  CPUStats
	NetworkIn    uint64
	NetworkOut   uint64
 	AccountId    string
	AssemblyId   string
  	AssembliesId string
	Status       string
}

type CPUStats struct {
	PercpuUsage       []uint64
	UsageInUsermode   uint64
	TotalUsage        uint64
	UsageInKernelmode uint64
  	SystemCPUUsage    uint64
}


func TestStats(id string) {
	//id := "7adc361d7d7f"

	client, _ := docker.NewClient("tcp://185.88.29.253:2375")
	client.SkipServerVersionCheck = true
	errC := make(chan error, 1)
	statsC := make(chan *docker.Stats)
	done := make(chan bool)
	defer close(done)
	go func() {
		errC <- client.Stats(docker.StatsOptions{ID: id, Stats: statsC, Stream: false, Done: done})
		close(errC)
	}()
	var resultStats []*docker.Stats
	for {
		stats, ok := <-statsC
		if !ok {
			break
		}
    
    cpuDelta := float64(stats.CPUStats.CPUUsage.TotalUsage) -  float64(stats.PreCPUStats.CPUUsage.TotalUsage)
    systemDelta := float64(stats.CPUStats.SystemCPUUsage) - float64(stats.PreCPUStats.SystemCPUUsage)
    if systemDelta > 0.0 && cpuDelta > 0.0 {
       RESULT_CPU_USAGE := (cpuDelta / systemDelta) * float64(len(stats.CPUStats.CPUUsage.PercpuUsage)) * 100.0
       fmt.Println("CPUStats   % ",RESULT_CPU_USAGE)
    }
//   RESULT_CPU_USAGE := cpuDelta / systemDelta * 100;
    fmt.Println("Memory_states :\n","Usage  :",stats.MemoryStats.Usage,"\n Max Useage",stats.MemoryStats.MaxUsage,"\nLimit :  ",stats.MemoryStats.Limit)

    //fmt.Println("Networks ",stats.Networks)
		resultStats = append(resultStats,  stats)
	}
	err := <-errC
	if err != nil {
    fmt.Println("******Error: 3 ",err)
	}
	if len(resultStats) != 2 {
		fmt.Println("Stats: Expected 2 results. Got %d.", len(resultStats))
	}

}

func TestListContainersParams() string {
    input := docker.ListContainersOptions{
      Filters: map[string][]string{"status": {"running"}},
      }
	  client, _ := docker.NewClient("tcp://185.88.29.253:2375")
		 ls, err := client.ListContainers(input)
    if err != nil {
			fmt.Println("\n\n\n*****************************************",err)
      return ""
    }
    fmt.Println(ls)
    fmt.Println("\nID :",ls[0].ID,"\tLables :",ls[0].Labels )
    return ls[0].ID
}

func TestStatsAllContainers() ([]interface{},error) {
  var resultStats []interface{}
  input := docker.ListContainersOptions{All: true}
   client, _ := docker.NewClient("tcp://185.88.29.253:2375")
   ps, err := client.ListContainers(input)
  	for _, v := range ps {
  		id := v.ID
  		errC := make(chan error, 1)
  		statsC := make(chan *docker.Stats)
  		done := make(chan bool)
  		defer close(done)
  		go func() {
  			errC <- client.Stats(docker.StatsOptions{ID: id, Stats: statsC, Stream: false, Done: done})
  			close(errC)
  		}()

  		for {
  			stats, ok := <-statsC
  			if !ok {
  				break
  			}
  			fmt.Println("Success ")
  			resultStats = append(resultStats, parseContainerStats(v,stats))
                  cpuDelta := float64(stats.CPUStats.CPUUsage.TotalUsage) -  float64(stats.PreCPUStats.CPUUsage.TotalUsage)
                  systemDelta := float64(stats.CPUStats.SystemCPUUsage) - float64(stats.PreCPUStats.SystemCPUUsage)
                  if systemDelta > 0.0 && cpuDelta > 0.0 {
                    RESULT_CPU_USAGE := (cpuDelta / systemDelta) * float64(len(stats.CPUStats.CPUUsage.PercpuUsage)) * 100.0
                    fmt.Println("CPUStats ",float64(cpuDelta/1000/1000) , "/",float64(systemDelta/1000/1000), "2 *100   = %",RESULT_CPU_USAGE)
                  }
  		}
  		err := <-errC
  		if err != nil {
  			return nil, err
  		}

  	}
  	if err != nil {
  		return nil, err
  	}

  	return resultStats, nil
}



func main() {
  //id :=  TestListContainersParams()
  //TestStats(id)
   res,err := TestStatsAllContainers()
   if err != nil {
     fmt.Println("Error:  ",err)
   }

   fmt.Println(res)
   for _,v := range res {
     f, ok := v.(*Stats)
     if !ok {
       fmt.Println("failed to converter")
     }
     //fmt.Println("",f.ContainerId,f.AccountId)
    cpuDelta := float64(f.CPUStats.TotalUsage) -  float64(f.PreCPUStats.TotalUsage)
    systemDelta := float64(f.CPUStats.SystemCPUUsage) - float64(f.PreCPUStats.SystemCPUUsage)
    fmt.Println(cpuDelta, "RAM TotalUsage :", float64(f.MemoryUsage/1024.0/1024.0))
    fmt.Println(systemDelta, "RAM SystemCPUUsage :", float64(f.SystemMemory/1024/1024))
    fmt.Println(cpuDelta, "CPU TotalUsage :", float64(cpuDelta/1000/1000))
    fmt.Println(systemDelta, "CPU SystemCPUUsage :",float64((cpuDelta + systemDelta)/1000/1000),"\n" , float64(systemDelta/1000/1000))
    fmt.Println(f.CPUStats.PercpuUsage, "CPU Len PercpuUsage :", len(f.CPUStats.PercpuUsage))

/*
    if systemDelta > 0.0 && cpuDelta > 0.0 {
       RESULT_CPU_USAGE := (cpuDelta / systemDelta) * float64(len(stats.CPUStats.CPUUsage.PercpuUsage)) * 100.0
       fmt.Println("CPUStats   % ",RESULT_CPU_USAGE)
    }
     fmt.Println(f.CPUStats.TotalUsage, "CPU TotalUsage :", float64(f.CPUStats.TotalUsage/1024/1024))
     fmt.Println(f.CPUStats.SystemCPUUsage, "CPU SystemCPUUsage :", float64(f.CPUStats.SystemCPUUsage/1024/1024))
     fmt.Println(f.CPUStats.SystemCPUUsage, "CPU PercpuUsage :",       float64(f.CPUStats.SystemCPUUsage/1024/1024))
     fmt.Println("\n\n",f.PreCPUStats.TotalUsage,"PreCPU TotalUsage :", float64(f.PreCPUStats.TotalUsage/1024/1024))
     fmt.Println(f.PreCPUStats.SystemCPUUsage,"PreCPU SystemCPUUsage :", float64(f.PreCPUStats.SystemCPUUsage/1024/1024))
*/

   }
}

func parseContainerStats(d docker.APIContainers,stats *docker.Stats) (*Stats) {
        fmt.Println("**********container name :",d.Image)
	return &Stats{
		ContainerId: d.ID,
		Image: d.Image,
		MemoryUsage:  stats.MemoryStats.Usage,
		SystemMemory: stats.MemoryStats.Limit,
		CPUStats:    CPUStats{
			PercpuUsage: stats.CPUStats.CPUUsage.PercpuUsage,
			UsageInUsermode: stats.CPUStats.CPUUsage.UsageInUsermode,
			TotalUsage: stats.CPUStats.CPUUsage.TotalUsage,
			UsageInKernelmode: stats.CPUStats.CPUUsage.UsageInKernelmode,
		  SystemCPUUsage: stats.CPUStats.SystemCPUUsage,
		},
		PreCPUStats:  CPUStats{
			PercpuUsage: stats.PreCPUStats.CPUUsage.PercpuUsage,
			UsageInUsermode: stats.PreCPUStats.CPUUsage.UsageInUsermode,
			TotalUsage: stats.PreCPUStats.CPUUsage.TotalUsage,
			UsageInKernelmode: stats.PreCPUStats.CPUUsage.UsageInKernelmode,
		  SystemCPUUsage: stats.PreCPUStats.SystemCPUUsage,
		},
		AccountId:    d.Labels["account_id"],
		AssemblyId:   d.Labels["assembly_id"],
		AssembliesId: d.Labels["assemblies_id"],
		Status: d.State,
	}

}
