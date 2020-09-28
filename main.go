package main

import (
	"up_time_monitor/database"
	handler "up_time_monitor/handler"
	mt "up_time_monitor/monitor"
	st "up_time_monitor/structures"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
)

func main() {
	dataBase := &database.GormDatabase{} // Get a instance of Database
	dataBase.ConnectToDatabase()         //Establish Connection  to Database
	dataBase.Migrate(&st.URLInfo{})      //Create table
	database.SetDatabase(dataBase)       // Set database Instance for database package

	monitor := &mt.MyMonitor{MonitorStp: make(chan uuid.UUID)} // Get a instance of Moniter
	mt.SetMonitor(monitor)                                     //Set Moniter Instance for moniter package
	monitor.StartMonitoringFromDatabase()                      //Start Monitoring of Already stored URLS

	//Define routers
	router := gin.Default()
	router.POST("/urls", handler.URLRegister)
	router.GET("/urls/:id", handler.FetchInfo)
	router.PATCH("/urls/:id", handler.PatchInfo)
	router.POST("/urls/:id/activate", handler.Activater)
	router.POST("/urls/:id/deactivate", handler.Deactivater)
	router.DELETE("/urls/:id", handler.DeleteURL)
	router.Run() //Default port is 8080
}
