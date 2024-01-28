package router

import "github.com/gin-gonic/gin"

func Init() {
	// Inicializa o Router utilizando as configurações padrão do Gin
	router := gin.Default()

	// Define uma rota para responder a requisição GET
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Inicia o servidor HTTP e escuta as requisições na porta 8080
	router.Run(":8080")
}
