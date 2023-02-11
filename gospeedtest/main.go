package main

import (
	"context"
	"fmt"
	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/showwin/speedtest-go/speedtest"
	"os"
	"time"
)

// getspeedtest get Latency Download speed ald Upload speed
func getspeedtest() (Lat int64, Dl float64, Ul float64) {
	//test val
	if len(os.Getenv("asdasddasasd")) != 0 {
		user, _ := speedtest.FetchUserInfo()
		serverList, _ := speedtest.FetchServers(user)
		targets, _ := serverList.FindServer([]int{})
		for _, s := range targets {
			s.PingTest()
			s.DownloadTest(false)
			s.UploadTest(false)
			Lat, Dl, Ul = s.Latency.Milliseconds(), s.DLSpeed, s.ULSpeed
		}
	} else {
		dur := 8 * time.Millisecond //, _ := time.ParseDuration(string(8 * time.Millisecond))
		Lat, Dl, Ul = dur.Milliseconds(), 96.65129866971003, 25.25837866052983
	}
	return //8.12645ms 96.65129866971003 25.25837866052983
}
func writetoinfluxdb(Lat int64, Dl float64, Ul float64) {
	//todo host, password, INFLUXDB_DB = speedtest
	client := influxdb2.NewClient("http://localhost:8086", "password")
	writeAPI := client.WriteAPIBlocking("speedtest", "my-bucket")
	line := fmt.Sprintf("speedtest,Latency=%df Download=%.2f,Upload=%.2f", Lat, Dl, Ul)
	err := writeAPI.WriteRecord(context.Background(), line)
	if err != nil {
		panic(err)
	}
}
func main() {
	var Lat, Dl, Ul = getspeedtest()
	writetoinfluxdb(Lat, Dl, Ul)
}

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
	os.Getenv("HOMEPATH")
*/
