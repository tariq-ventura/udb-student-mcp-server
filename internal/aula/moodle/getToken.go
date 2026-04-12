package aula_moodle

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func (mc *MoodleClient) GetToken() error {
	loginUrl := fmt.Sprintf("%s/login/token.php", mc.baseUrl)

	data := url.Values{}
	data.Set("username", mc.username)
	data.Set("password", mc.password)
	data.Set("service", "moodle_mobile_app")

	resp, err := http.PostForm(loginUrl, data)

	if err != nil {
		return fmt.Errorf("failed to get token: %w", err)
	}
	defer resp.Body.Close()

	var res MoodleTokenResponse

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return fmt.Errorf("failed to decode token response: %w", err)
	}

	if res.Error != "" {
		return fmt.Errorf("error from Moodle: %s", res.Error)
	}

	mc.token = res.Token

	return nil
}
