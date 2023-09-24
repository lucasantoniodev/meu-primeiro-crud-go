package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasantoniodev/meu-primeiro-crud-go/src/configuration/validation"
	"github.com/lucasantoniodev/meu-primeiro-crud-go/src/controller/model/request"
	"net/http"
)

/*
CreateUser

(c *gin.Context) -> Guarda as informações recebidas da requisição

c.ShouldBindJSON() pega um JSON que é recebido no body e faz um bind para um struct/objeto preparado
para ser bindado por JSON, usandos o "&" pois queremos que o valor seja alterado na memória (referência)
*/
func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusCreated, userRequest)
}
