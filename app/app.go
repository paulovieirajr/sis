package app

import (
	"log"
	"net"

	"github.com/urfave/cli"
)

/*
	Generate vai retornar uma aplicação cli com as configurações para serem executadas.
	A aplicação vai ter dois comandos: ip e server.

	O comando ip vai ter um flag chamado host que vai receber um valor padrão de www.github.com.
	O comando server vai ter um flag chamado host que vai receber um valor padrão de www.github.com.
*/
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "SIS"
	app.Usage = "Busca IPs e nomes de servidores na internet"
	app.Version = "1.0.0"
	app.Authors = []cli.Author{
		{
			Name:  "Paulo Vieira",
			Email: "junior.vieira.1990@outlook.com",
		},
	}

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "www.github.com",

		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca endereços IPs na internet",
			Flags: flags,
			Action: searchIps,
		},
		{
			Name:  "server",
			Usage: "Busca nomes de servidores na internet",
			Flags: flags,
			Action: searchServers,
		},
	}

	return app
}


/*
	searchIps vai buscar os endereços IPs de um host na internet.
*/
func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, error := net.LookupIP(host)
	
	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		log.Println(ip)
	}
}

/*
	searchServers vai buscar os nomes de servidores de um host na internet.
*/
func searchServers(c *cli.Context) {
	host := c.String("host")

	servers, error := net.LookupNS(host)
	
	if error != nil {
		log.Fatal(error)
	}

	for _, server := range servers {
		log.Println(server.Host)
	}
}
