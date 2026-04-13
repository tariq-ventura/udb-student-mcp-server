package tools_mcpgo

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

func (tc *ToolsClient) GetBorrowedBooks() {
	tool := mcp.NewTool("get_borrowed_books",
		mcp.WithDescription("Obtiene la lista de libros actualmente prestados al usuario, sus fechas de vencimiento y los IDs necesarios para renovarlos. Analiza el HTML devuelto para extraer esta información."),
	)

	tc.server.AddTool(tool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		err := tc.biblio.GetBiblioSession()

		if err != nil {
			return mcp.NewToolResultError("Error al iniciar sesión en el sistema bibliotecario: " + err.Error()), nil
		}

		htmlContent, err := tc.biblio.FetchBiblioAccount()

		if err != nil {
			return mcp.NewToolResultError("Error al obtener la cuenta bibliotecaria: " + err.Error()), nil
		}

		prompt := fmt.Sprintf("Aquí está el HTML de la cuenta del usuario. Busca la tabla de préstamos (usualmente con id='checkoutst' o similar). Lista los libros, fechas de vencimiento y fíjate en los 'value' de los checkboxes de renovación para saber los IDs de los items.\n\n%s", htmlContent)
		return mcp.NewToolResultText(prompt), nil
	})
}
