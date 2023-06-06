package errorsResponse


import (
	"net/http"
	"log"
	"encoding/json"
)


func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}
	
}


func Error(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `"json:erro"`
	}{
		Erro: erro.Error(),
	})
}