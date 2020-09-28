package handler

import (
	"fmt"
	"net/http"
	"up_time_monitor/database"
	mt "up_time_monitor/monitor"
	st "up_time_monitor/structures"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// PatchInfo : Updates url information into database
// Only FailureThreshold, Frequency, CrawlTimeout can be updated
// Update if any or all of them is given in PATCH request
// After update set Status to "active" and  FailureCount to 0
// If url was activated start crawling it otherwise leave it
// Returns json responce of Updated information
func PatchInfo(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		fmt.Println("Error while Converting UID", err.Error())
	}
	dataBase := database.GetDatabase()
	info, _ := dataBase.GetURLByID(id)
	monitor := mt.GetMonitor()

	var req st.Request
	req.FailureThreshold = -1
	req.CrawlTimeout = -1
	req.Frequency = -1
	c.BindJSON(&req)

	if req.FailureThreshold != -1 {
		info.FailureThreshold = req.FailureThreshold
	}
	if req.Frequency != -1 {
		info.Frequency = req.Frequency
	}
	if req.CrawlTimeout != -1 {
		info.CrawlTimeout = req.CrawlTimeout
	}
	info.Status = "active"
	info.FailureCount = 0
	dataBase.UpdateDatabase(info)

	if info.Crawling == true {
		monitor.StopMonitoring(id)
		monitor.StartMonitoring(info)
	}
	c.JSON(http.StatusOK, gin.H{
		"id":                info.ID,
		"url":               info.URL,
		"crawl_timeout":     info.CrawlTimeout,
		"frequency":         info.Frequency,
		"failure_threshold": info.FailureThreshold,
		"status":            info.Status,
		"failure_count":     info.FailureCount,
	})
}
