package aula_moodle

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (mc *MoodleClient) GetUserId() error {
	apiUrl := fmt.Sprintf("%s/webservice/rest/server.php?wstoken=%s&wsfunction=core_webservice_get_site_info&moodlewsrestformat=json", mc.baseUrl, mc.token)

	resp, err := http.Get(apiUrl)

	if err != nil {
		return fmt.Errorf("failed to get user ID: %w", err)
	}
	defer resp.Body.Close()

	var res MoodleSiteInfo

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return fmt.Errorf("failed to decode site info response: %w", err)
	}

	if res.ErrorCode != "" {
		return fmt.Errorf("error from Moodle: %s - %s", res.ErrorCode, res.Message)
	}

	mc.userId = res.UserID

	return nil
}
