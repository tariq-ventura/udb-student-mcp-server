package tools_mcpgo

import (
	"context"

	"github.com/mark3labs/mcp-go/server"
	"github.com/tariq-ventura/udb-mcp/internal/aula"
	"github.com/tariq-ventura/udb-mcp/internal/biblio"
)

type ToolsClient struct {
	aula   aula.IAuth
	biblio biblio.IBiblio
	server *server.MCPServer
}

var SetupMcp = func(ctx context.Context) (*ToolsClient, error) {
	s := server.NewMCPServer("udb-mcp", "0.1.0")

	aula, err := aula.NewAula(ctx)

	if err != nil {
		return nil, err
	}

	biblio, err := biblio.NewBiblio()

	if err != nil {
		return nil, err
	}

	return &ToolsClient{
		aula:   aula,
		biblio: biblio,
		server: s,
	}, nil
}
