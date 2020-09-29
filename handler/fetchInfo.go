package handler

import (
	"net/http"
	"up_time_monitor/database"

	"github.com/gin-gonic/gin"
)

// FetchInfo : Fetches information by id (Provided  in reqeusted url) from database
// returns json responce of URL information
func FetchInfo(c *gin.Context) {
	id := stringToUUID(c.Param("id"))
	dataBase := database.GetDatabase()
	info, err := dataBase.GetURLByID(id)
	if err != nil {			// if url record not found in database
		handleDataBaseError(c, err)
	} else {
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
}
