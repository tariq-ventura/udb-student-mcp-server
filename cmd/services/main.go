package main

import (
	"context"

	"github.com/tariq-ventura/udb-mcp/internal/tools"
)

func main() {
	ctx := context.Background()
	tools, err := tools.NewTools(ctx)

	if err != nil {
		panic(err)
	}

	tools.GetCourses()
	tools.GetPdfs()
	tools.DownloadPdf()

	if err := tools.StartServer(); err != nil {
		panic(err)
	}
}
