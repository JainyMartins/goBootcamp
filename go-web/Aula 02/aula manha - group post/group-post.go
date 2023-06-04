package main

import (
	"github.com/gin-gonic/gin"
)

type produto struct {
	ID    int    `json:"id"`
	Nome     string `json:"nome"`
	Tipo     string  `json:"tipo"`  
	Quantidade    int    `json:"quantidade"`
	Preco   float64     `json:"preco"`
}

func main(){
	r := gin.Default()

	pr := r.Group("/produtos")
	pr.POST("/salvar", Salvar())

	r.Run()

}

func Salvar() gin.HandlerFunc{
	return func (c *gin.Context) {
		var prod produto
		//Estamos pegando JSON que colocamos na requisição no BODY e transformando na nossa ESTRUTURA produto e devolvendo se deu certo 
		if err := c.ShouldBindJSON(&prod); err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		prod.ID = 4
		c.JSON(200, prod)
	}
}