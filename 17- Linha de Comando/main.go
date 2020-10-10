package main

import (
	"log"
	"os"
	"linha-de-comando/app"
)


func main() {

	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	} //Reconhecimento terminal
	
	
}
