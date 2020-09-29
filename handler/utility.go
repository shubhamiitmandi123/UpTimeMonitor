package handler

import (
	"fmt"
	"net/http"
	st "up_time_monitor/structures"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func handleDataBaseError(c *gin.Context, err error) {
	fmt.Println("Error :", err.Error())
	c.String(http.StatusInternalServerError, err.Error())
}

// convert an string uuid to uuid.UUID
func stringToUUID(st string) uuid.UUID {
	id, err := uuid.FromString(st)
	if err != nil {
		fmt.Println("Error while Converting UID", err.Error())
	}
	return id
}

// decode data from url and Updates info Parameter
// returns updated information
func getUpdatedURLInfo(c *gin.Context, info st.URLInfo) st.URLInfo {
	var req st.Request
	// if any of below remain -1 after binding request then it was not provided in request data
	req.FailureThreshold = -1
	req.CrawlTimeout = -1
	req.Frequency = -1
	c.BindJSON(&req)

	if req.FailureThreshold != -1 { // if FailureThreshold need to Update
		info.FailureThreshold = req.FailureThreshold
	}
	if req.Frequency != -1 { //if Frequency need to Update
		info.Frequency = req.Frequency
	}
	if req.CrawlTimeout != -1 { //if CrawlTimeout need to Update
		info.CrawlTimeout = req.CrawlTimeout
	}
	info.Status = "active" //after any update consider url as active and set failure count to 0
	info.FailureCount = 0
	return info
}
