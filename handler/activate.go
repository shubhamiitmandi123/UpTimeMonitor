package handler

import (
	"fmt"
	"net/http"
	"up_time_monitor/database"
	mt "up_time_monitor/monitor"

	"github.com/gin-gonic/gin"
)

// Activater : Starts crawling to a Deactivated or inactive url
// if not activated then starts crawling
// else return StatusNotAcceptable with error Massage
func Activater(c *gin.Context) {
	id := stringToUUID(c.Param("id"))
	dataBase := database.GetDatabase()
	info, err := dataBase.GetURLByID(id)

	if err != nil { // if record not found
		handleDataBaseError(c, err)
	} else {
		monitor := mt.GetMonitor()
		if info.Crawling == false { // if url was Inactivated
			info.Crawling = true
			info.FailureCount = 0
			info.Status = "active"
			err := dataBase.UpdateDatabase(info)
			if err != nil { //database update fails
				handleDataBaseError(c, err)
			} else {
				monitor.StartMonitoring(info) //start Monitoring
				c.String(http.StatusOK, "Activated ")
			}
		} else { // else error already activated
			msg := fmt.Sprintf("Error!!!!! Already activated ")
			c.String(http.StatusNotAcceptable, msg)
		}
	}
}
