package app

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
)

// Gerar vai retornar a aplicação de linha de comando pronta
func Gerar()*cli.App {
  app := cli.NewApp()
  app.Name = "Aplicação de Linha de Comando "
  app.Usage = "Busca IPs e Nomes de Servidor na Internet"
  
  flags := []cli.Flag{ //Parametros
	cli.StringFlag{
		Name: "host",
		Value: "devbook.com.br",
	},
},

  app.Commands = []cli.Command{
	  {
		  Name: "ip",
		  Usage: "Busca Ips de endereços na internet",
		  Flags: flags,
		  Action: buscarIps,
	  }
	  
	  {
		  Name: "servidores",
		  Usage: "Buscar Servidores",
		  Flags: flags,
		  Action: buscarServidores,
	  }
  }
  return app
}

func buscarIps( c *cli.Context){
	 host := c.String("host")

	 ips, erro := net.LookupIP(host)
	 if erro != nil {
		 log.Fatal(erro)
	 }

	 for  _, ip := range ips {
		 fmt.Println(ip)
	 }

}

func buscarServidores(c *cli.Context){
	host := c.String("host")

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidore  := range servidores {
		fmt.Println(servidores)
	}
}


