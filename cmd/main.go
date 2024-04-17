package main

import (
	"log"

	_ "github.com/josue/chalenge_golang/docs"
	"github.com/josue/chalenge_golang/internal/app/api"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

//	@title			User API
//	@version		1.0
//	@description	APIs to manage Users.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:8080
// @BasePath	/api/v1
func main() {
	log.Println("Starting service")
	api.LoadApiV1()
}
