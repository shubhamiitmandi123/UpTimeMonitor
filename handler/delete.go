package handler

import (
	"fmt"
	"net/http"
	"up_time_monitor/database"
	mt "up_time_monitor/monitor"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// DeleteURL : Deletes a URL from database
// if url was active then stops crawling before delete
func DeleteURL(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		fmt.Println("Error while Converting UID", err.Error())
	}

	dataBase := database.GetDatabase()
	monitor := mt.GetMonitor()

	info, _ := dataBase.GetURLByID(id)
	if info.Crawling == true {
		monitor.StopMonitoring(id)
	}
	dataBase.DeleteFromDatabase(id)
	c.String(http.StatusNoContent, "Deleted ")
}
