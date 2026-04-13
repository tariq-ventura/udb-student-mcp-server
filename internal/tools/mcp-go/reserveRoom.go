package tools_mcpgo

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

func (tc *ToolsClient) ReserveRoom() {
	tool := mcp.NewTool("reserve_library_room",
		mcp.WithDescription("Reserva un cubículo en la biblioteca (MRBS). ATENCIÓN: Exige al usuario la fecha, hora de inicio y hora de fin exactas antes de proceder."),
		mcp.WithString("date", mcp.Required(), mcp.Description("Fecha de reserva (YYYY-MM-DD), ej: 2026-04-13")),
		mcp.WithString("start_time", mcp.Required(), mcp.Description("Hora de inicio en formato 24h (HH:MM), ej: 14:00")),
		mcp.WithString("end_time", mcp.Required(), mcp.Description("Hora de fin en formato 24h (HH:MM), ej: 15:00")),
		mcp.WithString("room_id", mcp.Required(), mcp.Description("ID numérico del cubículo (ej: 1, 2, 3)")),
	)

	tc.server.AddTool(tool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]interface{})
		if !ok {
			return mcp.NewToolResultError("los argumentos no tienen el formato esperado"), nil
		}

		dateStr := args["date"].(string)
		startStr := args["start_time"].(string)
		endStr := args["end_time"].(string)
		roomID := args["room_id"].(string)

		err := tc.biblio.GetSession()

		if err != nil {
			return mcp.NewToolResultError("error de login: " + err.Error()), nil
		}

		var h, m int

		fmt.Sscanf(startStr, "%d:%d", &h, &m)
		startSecs := fmt.Sprintf("%d", (h*3600)+(m*60))

		fmt.Sscanf(endStr, "%d:%d", &h, &m)
		endSecs := fmt.Sprintf("%d", (h*3600)+(m*60))

		result, err := tc.biblio.ReserveMRBS(dateStr, startSecs, endSecs, roomID)

		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		return mcp.NewToolResultText(result), nil
	})
}
