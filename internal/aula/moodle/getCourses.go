package aula_moodle

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/tariq-ventura/udb-mcp/internal/domain"
)

func (mc *MoodleClient) FetchCourses() ([]domain.Course, error) {
	apiUrl := fmt.Sprintf("%s/webservice/rest/server.php?wstoken=%s&wsfunction=core_enrol_get_users_courses&userid=%d&moodlewsrestformat=json", mc.baseUrl, mc.token, mc.userId)
	resp, err := http.Get(apiUrl)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if len(bodyBytes) > 0 && bodyBytes[0] == '{' {
		var apiErr MoodleAPIError
		json.Unmarshal(bodyBytes, &apiErr)
		return nil, fmt.Errorf("error de la API de Moodle: %s (%s)", apiErr.Message, apiErr.ErrorCode)
	}

	var courses []domain.Course
	if err := json.Unmarshal(bodyBytes, &courses); err != nil {
		return nil, fmt.Errorf("error decodificando cursos: %v", err)
	}

	return courses, nil
}
