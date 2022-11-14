package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
)

func TestAppShowingInfos(t *testing.T) {
	app := Generate()

	assert.Equal(t, "SIS", app.Name)
	assert.Equal(t, "Busca IPs e nomes de servidores na internet", app.Usage)
	assert.Equal(t, "1.0.0", app.Version)
	assert.Equal(t, "Paulo Vieira", app.Authors[0].Name)
	assert.Equal(t, "junior.vieira.1990@outlook.com", app.Authors[0].Email)
}

func TestAppCommandsOfIp(t *testing.T) {
	app := Generate()

	assert.Equal(t, "ip", app.Commands[0].Name)
	assert.Equal(t, "Busca endereços IPs na internet", app.Commands[0].Usage)
	assert.Equal(t, "host", app.Commands[0].Flags[0].GetName())
	assert.Equal(t, "www.github.com", app.Commands[0].Flags[0].(cli.StringFlag).Value)
}

func TestAppCommandsOfServer(t *testing.T) {
	app := Generate()

	assert.Equal(t, "server", app.Commands[1].Name)
	assert.Equal(t, "Busca nomes de servidores na internet", app.Commands[1].Usage)
	assert.Equal(t, "host", app.Commands[1].Flags[0].GetName())
	assert.Equal(t, "www.github.com", app.Commands[1].Flags[0].(cli.StringFlag).Value)
}

func TestAppToInstantiateItCantBeNil(t *testing.T) {
	app := Generate()

	assert.NotNil(t, app.Commands[0].Action)
}
