package aula_moodle

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (mc *MoodleClient) FetchCoursesPdf(courseId float64) ([]map[string]string, error) {
	apiURL := fmt.Sprintf("%s/webservice/rest/server.php?wstoken=%s&wsfunction=core_course_get_contents&courseid=%v&moodlewsrestformat=json", mc.baseUrl, mc.token, courseId)

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)

	if len(bodyBytes) > 0 && bodyBytes[0] == '{' {
		var apiErr MoodleAPIError
		json.Unmarshal(bodyBytes, &apiErr)
		return nil, fmt.Errorf("error de Moodle: %s", apiErr.Message)
	}

	var sections []MoodleSection
	if err := json.Unmarshal(bodyBytes, &sections); err != nil {
		return nil, fmt.Errorf("error leyendo contenido del curso: %v", err)
	}

	var pdfList []map[string]string
	for _, section := range sections {
		for _, mod := range section.Modules {
			if mod.ModName == "resource" || mod.ModName == "folder" {
				for _, file := range mod.Contents {
					if file.MimeType == "application/pdf" {
						pdfList = append(pdfList, map[string]string{
							"seccion":  section.Name,
							"recurso":  mod.Name,
							"archivo":  file.Filename,
							"file_url": file.FileURL,
						})
					}
				}
			}
		}
	}
	return pdfList, nil
}
