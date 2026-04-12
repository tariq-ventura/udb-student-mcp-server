package tools_mcpgo

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

func (tc *ToolsClient) DownloadPdf() {
	tool := mcp.NewTool("download_moodle_pdf",
		mcp.WithDescription("Descarga un PDF de Moodle a local. USA ESTA RUTA con tu capacidad nativa para leer el documento (tablas, fechas, porcentajes)."),
		mcp.WithString("file_url",
			mcp.Required(),
			mcp.Description("La URL del archivo PDF obtenida de get_course_pdfs"),
		),
		mcp.WithString("filename",
			mcp.Required(),
			mcp.Description("Nombre del archivo"),
		),
	)

	tc.server.AddTool(tool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]interface{})
		if !ok {
			return mcp.NewToolResultError("los argumentos no tienen el formato esperado"), nil
		}

		fileURL, ok := args["file_url"].(string)

		if !ok {
			return mcp.NewToolResultError("file_url es requerido y debe ser una cadena"), nil
		}

		filename, ok := args["filename"].(string)

		if !ok {
			return mcp.NewToolResultError("filename es requerido y debe ser una cadena"), nil
		}

		err := tc.aula.GetToken()

		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Error de login: %v", err)), nil
		}

		absPath, err := tc.aula.DownloadPdf(fileURL, filename)

		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		resp := fmt.Sprintf(
			"✅ Archivo descargado en el contenedor con éxito.\n\n"+
				"Ruta absoluta: %s\n\n"+
				"INSTRUCCIÓN PARA EL AGENTE: Utiliza tu herramienta interna para leer archivos locales y analiza el PDF en esta ruta para responder a la solicitud del usuario.", absPath)

		return mcp.NewToolResultText(resp), nil
	})
}
