package handler

import (
	"fmt"
	"net/http"
	"up_time_monitor/database"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// FetchInfo : Fetches information by id (Provided  in reqeusted url) from database
// returns json responce of URL information
func FetchInfo(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		fmt.Println("its nil", err.Error())
	}
	dataBase := database.GetDatabase()
	info, _ := dataBase.GetURLByID(id)
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
