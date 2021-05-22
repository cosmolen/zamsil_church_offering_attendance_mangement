package mappings

import (
	"io"
	"os"
	"time"
	"zamsil_church_offering_attendance_mangement/config"
	"zamsil_church_offering_attendance_mangement/controllers"
	"zamsil_church_offering_attendance_mangement/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func CreateUrlMappings(conf *config.Config) {
	// save log
	now := time.Now()
	logFilePath, _ := os.Create("data/log/gin_" + now.Format("2006-01-02") + ".log")
	errorLogFilePath, _ := os.Create("data/log/gin_error_" + now.Format("2006-01-02") + ".log")
	gin.DefaultWriter = io.MultiWriter(logFilePath, os.Stdout)
	gin.DefaultErrorWriter = io.MultiWriter(errorLogFilePath, os.Stdout)

	// gin mode
	Router = gin.Default()

	// middelware for cors
	Router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{conf.CORS.ALLOWORIGINS},
		AllowCredentials: conf.CORS.ACCESSCONTROLALLOWCREDENTIALS,
		AllowHeaders:     []string{"content-type"},
	}))

	//logging to a file

	Router.POST("/Youth/login", controllers.Login(conf))

	// Router Get with middlewares AUthorize
	Youth := Router.Group("Youth").Use(middlewares.Authorize(conf))

	Youth.GET("/attendances", controllers.GetAttendanceInfoByDepartment)
	Youth.POST("/attendances", controllers.PostAttendanceInfo)

	Youth.GET("/attendance/view", controllers.GetAttendanceView)
	Youth.GET("/attendance/view/excel", controllers.GetAttendanceExcel)

	Youth.GET("/students", controllers.GetStudents)
	Youth.POST("/students", controllers.SaveStudents)

	Youth.GET("/offering/type", controllers.GetOfferingType)
	Youth.POST("/offering", controllers.PostOffering)

	Youth.GET("/offering/view", controllers.GetOfferingView)
	Youth.GET("/offering/view/excel", controllers.GetOfferingExcel)

	Youth.GET("/classes", controllers.GetClasses)

}