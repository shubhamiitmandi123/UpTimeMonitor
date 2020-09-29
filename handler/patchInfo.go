package handler

import (
	"net/http"
	"up_time_monitor/database"
	mt "up_time_monitor/monitor"

	"github.com/gin-gonic/gin"
)

// PatchInfo : Updates url information into database
// Only FailureThreshold, Frequency, CrawlTimeout can be updated
// Update if any or all of them is given in PATCH request
// After update set Status to "active" and  FailureCount to 0
// If url was activated start crawling it otherwise leave it
// Returns json responce of Updated information
func PatchInfo(c *gin.Context) {
	id := stringToUUID(c.Param("id"))
	dataBase := database.GetDatabase()
	fetchedInfo, err := dataBase.GetURLByID(id)
	if err != nil { // if url record is not found in database
		handleDataBaseError(c, err)
	} else {
		monitor := mt.GetMonitor()
		updatedInfo := getUpdatedURLInfo(c, fetchedInfo)
		err := dataBase.UpdateDatabase(updatedInfo)
		if err != nil { // if unable to update Information
			handleDataBaseError(c, err)
		} else {
			// if url was being Crawling then stop crawling it
			// And start crawling with updated Parameters
			if fetchedInfo.Crawling == true {
				monitor.StopMonitoring(id)
				monitor.StartMonitoring(updatedInfo)
			}
			c.JSON(http.StatusOK, gin.H{
				"id":                updatedInfo.ID,
				"url":               updatedInfo.URL,
				"crawl_timeout":     updatedInfo.CrawlTimeout,
				"frequency":         updatedInfo.Frequency,
				"failure_threshold": updatedInfo.FailureThreshold,
				"status":            updatedInfo.Status,
				"failure_count":     updatedInfo.FailureCount,
			})
		}
	}
}
