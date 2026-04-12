package tools_mcpgo

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

func (tc *ToolsClient) GetPdfs() {
	tool := mcp.NewTool("get_course_pdfs",
		mcp.WithDescription("Obtiene la lista de PDFs de un curso. REGLA ESTRICTA: NUNCA uses read_moodle_pdf inmediatamente después de esto. Primero, muéstrale al usuario los nombres de los archivos encontrados y pregúntale exactamente cuál de ellos quiere que leas."),
		mcp.WithNumber("course_id",
			mcp.Required(),
			mcp.Description("El ID del curso (obtenido de get_my_courses)"),
		),
	)

	tc.server.AddTool(tool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]interface{})
		if !ok {
			return mcp.NewToolResultError("los argumentos no tienen el formato esperado"), nil
		}

		courseId, ok := args["course_id"].(float64)
		if !ok {
			return mcp.NewToolResultError("course_id es requerido y debe ser un número"), nil
		}

		err := tc.aula.GetToken()

		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Error de login: %v", err)), nil
		}

		pdfs, err := tc.aula.FetchCoursesPdf(courseId)

		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Error al obtener PDFs del curso: %v", err)), nil
		}

		resultJSON, _ := json.MarshalIndent(pdfs, "", "  ")
		return mcp.NewToolResultText(string(resultJSON)), nil
	})
}
