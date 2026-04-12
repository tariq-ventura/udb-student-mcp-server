package tools_mcpgo

import (
	"context"

	"github.com/mark3labs/mcp-go/server"
	"github.com/tariq-ventura/udb-mcp/internal/aula"
)

type ToolsClient struct {
	server *server.MCPServer
	aula   aula.IAuth
}

var SetupMcp = func(ctx context.Context) (*ToolsClient, error) {
	s := server.NewMCPServer("udb-mcp", "0.1.0")

	aula, err := aula.NewAula(ctx)

	if err != nil {
		return nil, err
	}

	return &ToolsClient{
		server: s,
		aula:   aula,
	}, nil
}
