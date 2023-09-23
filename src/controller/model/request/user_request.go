package request

/*
	Mapeamento de Objeto para JSON:
		-> TAG "json" é utilizada para nomear o campo
			-> Primeiro valor é o nome do campo
			-> Segundo valor é escolher se o campo deve ser omitido ou visível com valor nulo

	Validaçõoes:
			-> TAG "binding" é utilizada no lugar da TAG "validate" pois o Ging Gonic já utilizada o validator;
*/

type UserRequest struct {
	Email    string `json:"email,omitempty" binding:"required,email"`
	Password string `json:"password,omitempty" binding:"required,min=6,containsany=!@#$%*"`
	Name     string `json:"name,omitempty" binding:"required"`
	Age      int8   `json:"age,omitempty" binding:"required,min=18"`
}
