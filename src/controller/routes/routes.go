package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasantoniodev/meu-primeiro-crud-go/src/controller"
	"log"
)

/*
RunServer

gin.New() -> Cria um novo roteador gin, porém sem nenhum middleware;
gin.Default() -> Cria um novo roteador gin, porém com logger e middlewares de recovery.
*/
func RunServer(port string) {
	router := gin.Default()
	initRoutes(&router.RouterGroup)
	if err := router.Run(port); err != nil {
		log.Fatalf("Whoops! There was a problem! %s", err)
	}
}

func initRoutes(r *gin.RouterGroup) {
	r.POST("/users", controller.CreateUser)
	r.GET("/getUserById/:id", controller.FindUserById)
	r.GET("/getUserByEmail/:email", controller.FindUserByEmail)
	r.PUT("/users/:id", controller.UpdateUser)
	r.DELETE("/users/:id", controller.DeleteUser)
}
