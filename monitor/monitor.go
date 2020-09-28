package monitor

import (
	"fmt"
	"time"

	"up_time_monitor/database"
	st "up_time_monitor/structures"

	uuid "github.com/satori/go.uuid"
)

// Monitor : Monitors a record from database
type Monitor interface {
	StartMonitoringFromDatabase()
	StartMonitoring(st.URLInfo)
	StopMonitoring(uuid.UUID)
}

var monitorImpl Monitor

// SetMonitor : Sets Monitor Instance
func SetMonitor(monitor Monitor) {
	monitorImpl = monitor
}

// GetMonitor : Returns Moniter Instance
func GetMonitor() Monitor {
	return monitorImpl
}

// MyMonitor : Implemention of Moniter interface
type MyMonitor struct {
	MonitorStp chan uuid.UUID
}

// StartMonitoringFromDatabase : It starts monitoring of all records from database which are activated for crawling
func (r *MyMonitor) StartMonitoringFromDatabase() {
	dataBase := database.GetDatabase()
	records := dataBase.GetAllRecords()
	for _, v := range records {
		if v.Crawling == true {
			r.StartMonitoring(v)
		}
	}
}

// StartMonitoring : StartMonitering of a given URL
// It sends request to url according to frequency
// It stops crawling a URL if it is inactive or deactivated
// If responce is not 200 OK it increases Failure count
// Once Failure count reaches to Failure Threshold it stops crawling it
func (r *MyMonitor) StartMonitoring(urlinfo st.URLInfo) {

	go func() {
		ticker := time.NewTicker(time.Duration(urlinfo.Frequency) * time.Second)
		requestStatus := make(chan string)
		dataBase := database.GetDatabase()
		for {
			select {
			case idStop := <-r.MonitorStp:
				if idStop == urlinfo.ID {
					return // stop monitering
				}
			case <-ticker.C: // at Frequency time
				fmt.Printf("Request to %s\t", urlinfo.URL)
				go Request(urlinfo.URL, urlinfo.CrawlTimeout, requestStatus)
			case st := <-requestStatus:
				fmt.Println("Status: ", st)
				if st != "200 OK" {
					urlinfo.FailureCount++
					dataBase.UpdateColumnInDatabase(urlinfo.ID, "failure_count", urlinfo.FailureCount)
					if urlinfo.FailureCount == urlinfo.FailureThreshold {
						dataBase.UpdateColumnInDatabase(urlinfo.ID, "status", "inactive")
						dataBase.UpdateColumnInDatabase(urlinfo.ID, "crawling", false)
						return //Stop Monitering
					}
				}
			}
		}
	}()

}

// StopMonitoring :  It sends id through a channel of URL which is being monitered by StartMonitoring
// As StrtMoniter Gets id, it Immediately  stops crawling of particular url
func (r *MyMonitor) StopMonitoring(id uuid.UUID) {
	r.MonitorStp <- id
}
