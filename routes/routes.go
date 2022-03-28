package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/luisolimpio/api-go-gin/controllers"
)

func Router() {
	router := gin.Default()

	router.GET("/students", controllers.GetStudents)
	router.GET("/students/:id", controllers.GetStudent)
	router.GET("/students/cpf/:cpf", controllers.GetStudentByCpf)
	router.POST("/students", controllers.PostStudent)
	router.PATCH("/students/:id", controllers.PatchStudent)
	router.DELETE("/students/:id", controllers.DeleteStudent)

	router.Run(":3333")
}
