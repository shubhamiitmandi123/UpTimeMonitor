package handler

import (
	"net/http"
	"up_time_monitor/database"
	mt "up_time_monitor/monitor"
	st "up_time_monitor/structures"

	"github.com/gin-gonic/gin"
)

// URLRegister : Creates a new record into database
// Set initial Status to "active" and Failurecount to 0
// Start Crawling it
// Returns URL information with assigned id
func URLRegister(c *gin.Context) {
	var req st.Request
	c.BindJSON(&req)
	dataBase := database.GetDatabase()
	id, err := dataBase.CreateDataBase(req)
	if err != nil { // if unable to create record
		handleDataBaseError(c, err)
	} else {
		urlinfo := st.URLInfo{
			Base:             st.Base{ID: id},
			URL:              req.URL,
			CrawlTimeout:     req.CrawlTimeout,
			Frequency:        req.Frequency,
			FailureThreshold: req.FailureThreshold,
			FailureCount:     0,
			Status:           "active",
			Crawling:         true,
		}
		monitor := mt.GetMonitor()
		monitor.StartMonitoring(urlinfo) //start Monitoring URL
		c.JSON(http.StatusOK, gin.H{
			"id":                id,
			"url":               req.URL,
			"crawl_timeout":     req.CrawlTimeout,
			"frequency":         req.Frequency,
			"failure_threshold": req.FailureThreshold,
			"status":            "active",
			"failure_count":     0,
		})
	}
}
