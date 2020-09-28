package handler

import (
	"fmt"
	"net/http"
	"up_time_monitor/database"
	mt "up_time_monitor/monitor"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// Deactivater : Stops crawling a activated url
//if activated => stop crawling and update status
//else return StatusNotAcceptable with error Massage
func Deactivater(c *gin.Context) {
	sid := c.Param("id")
	id, err := uuid.FromString(sid)
	if err != nil {
		fmt.Println("Error while Converting UID", err.Error())
	}
	dataBase := database.GetDatabase()
	info, _ := dataBase.GetURLByID(id)
	monitor := mt.GetMonitor()
	if info.Crawling == true {
		dataBase.UpdateColumnInDatabase(id, "crawling", false)
		monitor.StopMonitoring(id)
		c.String(http.StatusOK, "Deactivated ")
	} else {
		msg := fmt.Sprintf("Error!!!!! \nAlready deactivated %s", sid)
		c.String(http.StatusNotAcceptable, msg)
	}
}
