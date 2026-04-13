package tools

import (
	"context"
	"fmt"
	"os"

	tools_mcpgo "github.com/tariq-ventura/udb-mcp/internal/tools/mcp-go"
)

type IToolsMcp interface {
	StartServer() error
	GetCourses()
	GetPdfs()
	DownloadPdf()
	ReserveRoom()
}

var NewTools = func(ctx context.Context) (IToolsMcp, error) {
	mcpType, find := os.LookupEnv("MCP_TYPE")

	if !find {
		return nil, fmt.Errorf("MCP_TYPE environment variable not found")
	}

	switch mcpType {
	case "mcp-go":
		return tools_mcpgo.SetupMcp(ctx)
	default:
		return nil, fmt.Errorf("unsupported MCP_TYPE: %s", mcpType)
	}
}
