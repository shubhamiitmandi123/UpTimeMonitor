package handler

import (
	"net/http"
	"up_time_monitor/database"
	mt "up_time_monitor/monitor"

	"github.com/gin-gonic/gin"
)

// DeleteURL : Deletes a URL from database
// if url was active then stops crawling before delete
func DeleteURL(c *gin.Context) {
	id := stringToUUID(c.Param("id"))

	dataBase := database.GetDatabase()
	monitor := mt.GetMonitor()

	info, err := dataBase.GetURLByID(id)
	if err != nil { // if url not found in database
		handleDataBaseError(c, err)
	} else {
		if info.Crawling == true { // stop Monitoring Before Deleteing if it is being monitered
			monitor.StopMonitoring(id)
		}
		err := dataBase.DeleteFromDatabase(id)
		if err != nil { //if unable to Delete
			handleDataBaseError(c, err)
		} else {
			c.String(http.StatusNoContent, "Deleted ")
		}
	}
}
