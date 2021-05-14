package main

import (
	"zamsil_church_offering_attendance_mangement/config"
	"zamsil_church_offering_attendance_mangement/mappings"
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/gin-contrib/cors"
)

// Main program.
func main() {
	conf := config.GetConf()
	models.InitDb(conf)
	mappings.CreateUrlMappings(conf)
	mappings.Router.Use(cors.Default())
	mappings.Router.Run()
}
