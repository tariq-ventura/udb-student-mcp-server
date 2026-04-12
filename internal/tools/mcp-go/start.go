package tools_mcpgo

import (
	"fmt"

	"github.com/mark3labs/mcp-go/server"
)

func (tc *ToolsClient) StartServer() error {
	if err := server.ServeStdio(tc.server); err != nil {
		return fmt.Errorf("Error iniciando el servidor MCP: %v\n", err)
	}

	return nil
}
