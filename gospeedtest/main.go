package main

import (
	"fmt"
	"github.com/showwin/speedtest-go/speedtest"
	"os"
	"time"
)

// getspeedtest get Latency Download speed ald Upload speed
func getspeedtest() (Lat time.Duration, Dl float64, Ul float64) {
	//test val
	if len(os.Getenv("asdasddasasd")) != 0 {
		user, _ := speedtest.FetchUserInfo()
		serverList, _ := speedtest.FetchServers(user)
		targets, _ := serverList.FindServer([]int{})
		for _, s := range targets {
			s.PingTest()
			s.DownloadTest(false)
			s.UploadTest(false)

			Lat, Dl, Ul = s.Latency, s.DLSpeed, s.ULSpeed
		}
	} else {
		Lat, Dl, Ul = (8 * time.Millisecond), 96.65129866971003, 25.25837866052983
	}
	return //8.12645ms 96.65129866971003 25.25837866052983
}
func main() {
	/*var (
		INFLUXDB_DB string = "INFLUXDB_HOST"
	)*/
	/*if os.Getenv("INFLUXDB_HOST") {
		INFLUXDB_HOST := "influxdb"
	}*/
	/*
		INFLUXDB_DB := 'speedtest'
		INFLUXDB_USERNAME := 'root'
		INFLUXDB_PASSWORD =: 'root'
		SPEEDTEST_HOST := : 'local'
		SPEEDTEST_INTERVAL := 3600
		os.Getenv("HOMEPATH")*/

	fmt.Println(getspeedtest())
}
