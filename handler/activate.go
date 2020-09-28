package handler

import (
	"fmt"
	"net/http"
	"up_time_monitor/database"
	mt "up_time_monitor/monitor"

	"github.com/gin-gonic/gin"

	uuid "github.com/satori/go.uuid"
)

// Activater : Starts crawling to a Deactivated or inactive url
// if not activated start crawling
// else return StatusNotAcceptable with error Massage
func Activater(c *gin.Context) {
	sid := c.Param("id")
	id, err := uuid.FromString(sid)
	if err != nil {
		fmt.Println("Error while Converting UID", err.Error())
	}
	dataBase := database.GetDatabase()
	info, _ := dataBase.GetURLByID(id)
	monitor := mt.GetMonitor()
	if info.Crawling == false {
		info.Crawling = true
		info.FailureCount = 0
		info.Status = "active"
		dataBase.UpdateDatabase(info)
		monitor.StartMonitoring(info)
		c.String(http.StatusOK, "Activated ")
	} else {
		msg := fmt.Sprintf("Error!!!!! Already activated ")
		c.String(http.StatusNotAcceptable, msg)
	}
}
