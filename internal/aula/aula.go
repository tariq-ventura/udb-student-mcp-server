package aula

import (
	"context"
	"fmt"
	"os"

	aula_moodle "github.com/tariq-ventura/udb-mcp/internal/aula/moodle"
	"github.com/tariq-ventura/udb-mcp/internal/domain"
)

type IAuth interface {
	GetToken() error
	FetchCourses() ([]domain.Course, error)
	GetUserId() error
	FetchCoursesPdf(courseId float64) ([]map[string]string, error)
	DownloadPdf(fileUrl, filename string) (string, error)
}

var NewAula = func(ctx context.Context) (IAuth, error) {
	authType, find := os.LookupEnv("AUTH_TYPE")

	if !find {
		return nil, fmt.Errorf("AUTH_TYPE environment variable not found")
	}

	switch authType {
	case "moodle":
		return aula_moodle.SetupMoodle(ctx)
	default:
		return nil, fmt.Errorf("unsupported AUTH_TYPE: %s", authType)
	}
}
