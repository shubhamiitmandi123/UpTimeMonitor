package handler

import (
	"fmt"
	"net/http"
	"up_time_monitor/database"
	mt "up_time_monitor/monitor"

	"github.com/gin-gonic/gin"
)

// Deactivater : Stops crawling a activated url
//if activated => stop crawling and update status
//else return StatusNotAcceptable with error Massage
func Deactivater(c *gin.Context) {
	id := stringToUUID(c.Param("id"))
	dataBase := database.GetDatabase()
	info, err := dataBase.GetURLByID(id)
	if err != nil { //if url not found in database
		handleDataBaseError(c, err)
	} else {
		monitor := mt.GetMonitor()
		if info.Crawling == true {
			err := dataBase.UpdateColumnInDatabase(id, "crawling", false)
			if err != nil { //if fail to update in database
				handleDataBaseError(c, err)
			} else {
				monitor.StopMonitoring(id) //Stop Monitoring
				c.String(http.StatusOK, "Deactivated ")
			}
		} else {
			msg := fmt.Sprintf("Error!!!!! Already deactivated\n")
			c.String(http.StatusNotAcceptable, msg)
		}
	}
}
