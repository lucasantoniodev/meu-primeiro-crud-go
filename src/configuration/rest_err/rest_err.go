package rest_err

import "net/http"

/*
	Em GO se você criar uma struct com um método implementado a esta struct e existir uma
	interface com a mesma caracteristica, o GO vai assumir que a struct está implementando
	a interface.

	Interface a ser implementada:

	type error interface {
		Error() string
	}
*/

// Causes -> Estrutura para causas de erro
type Causes struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

// RestErr -> Estrutura JSON para erros
type RestErr struct {
	Message string   `json:"message,omitempty"`
	Err     string   `json:"error,omitempty"`
	Code    int      `json:"code,omitempty"`
	Causes  []Causes `json:"causes,omitempty"`
}

// Error -> Implementação de interface para sobre-escrever método Error() do GO;
func (r *RestErr) Error() string {
	return r.Message
}

// NewRestErr -> Método para gerar estrutura de Erro
func NewRestErr(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *RestErr {
	return NewRestErr(message, "bad_request", http.StatusBadRequest, nil)
}

func NewBadRequestValidationError(message string, causes []Causes) *RestErr {
	return NewRestErr(message, "bad_request", http.StatusBadRequest, causes)
}

func NewInternalServerError(message string) *RestErr {
	return NewRestErr(message, "internal_server_error", http.StatusInternalServerError, nil)
}

func NewNotFoundError(message string) *RestErr {
	return NewRestErr(message, "not_found", http.StatusNotFound, nil)
}

func NewForbiddenError(message string) *RestErr {
	return NewRestErr(message, "forbidden", http.StatusForbidden, nil)
}
