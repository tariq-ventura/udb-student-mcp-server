package tools_mcpgo

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

func (tc *ToolsClient) GetCourses() {
	tool := mcp.NewTool("get_my_courses", mcp.WithDescription("Obtiene la lista de cursos en los que el estudiante está inscrito"))

	tc.server.AddTool(tool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		err := tc.aula.GetToken()

		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Error de login: %v", err)), nil
		}

		errId := tc.aula.GetUserId()

		if errId != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Error al obtener ID de usuario: %v", err)), nil
		}

		courses, err := tc.aula.FetchCourses()

		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Error al obtener cursos: %v", err)), nil
		}

		coursesJSON, _ := json.MarshalIndent(courses, "", "  ")

		return mcp.NewToolResultText(string(coursesJSON)), nil
	})

}
