package biblio_koha

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

func (kc *KohaClient) ReserveMRBS(dateStr, startSeconds, endSeconds, roomId string) (string, error) {
	client := &http.Client{}

	formURL := fmt.Sprintf("%s/biblioteca/web/edit_entry.php?area=1&room=%s", kc.baseUrl, roomId)
	reqGet, err := http.NewRequest("GET", formURL, nil)
	if err != nil {
		return "", err
	}

	reqGet.Header.Add("Cookie", fmt.Sprintf("MRBS_SESSID=%s", kc.session))

	respGet, err := client.Do(reqGet)
	if err != nil {
		return "", err
	}
	defer respGet.Body.Close()

	bodyBytes, _ := io.ReadAll(respGet.Body)
	htmlContent := string(bodyBytes)

	re := regexp.MustCompile(`name="csrf_token" value="([^"]+)"`)
	matches := re.FindStringSubmatch(htmlContent)

	if len(matches) < 2 {
		return "", fmt.Errorf("no se pudo encontrar el token CSRF para reservar")
	}
	csrfToken := matches[1]

	formData := url.Values{}
	formData.Set("csrf_token", csrfToken)
	formData.Set("returl", "")
	formData.Set("rep_id", "0")
	formData.Set("edit_type", "series")
	formData.Set("create_by", kc.username)
	formData.Set("name", kc.username)
	formData.Set("description", "Reserva vía OpenClaw MCP")
	formData.Set("start_date", dateStr)
	formData.Set("start_seconds", startSeconds)
	formData.Set("end_date", dateStr)
	formData.Set("end_seconds", endSeconds)
	formData.Set("area", "1")
	formData.Set("rooms[]", roomId)
	formData.Set("type", "I")
	formData.Set("rep_type", "0")

	postURL := fmt.Sprintf("%s/biblioteca/web/edit_entry_handler.php", kc.baseUrl)
	reqPost, err := http.NewRequest("POST", postURL, strings.NewReader(formData.Encode()))
	if err != nil {
		return "", err
	}

	reqPost.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	reqPost.Header.Add("Cookie", fmt.Sprintf("MRBS_SESSID=%s", kc.session))

	respPost, err := client.Do(reqPost)
	if err != nil {
		return "", err
	}
	defer respPost.Body.Close()

	if respPost.StatusCode == http.StatusOK || respPost.StatusCode == http.StatusFound {
		return fmt.Sprintf("✅ ¡Reserva confirmada con éxito!\nFecha: %s\nCubículo ID: %s", dateStr, roomId), nil
	}

	return "", fmt.Errorf("el servidor rechazó la reserva (Código: %d)", respPost.StatusCode)
}
