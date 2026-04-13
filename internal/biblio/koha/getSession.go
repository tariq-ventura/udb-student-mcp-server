package biblio_koha

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

func (kc *KohaClient) GetSession() error {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	indexURL := fmt.Sprintf("%s/biblioteca/web/index.php", kc.baseUrl)
	respGet, err := client.Get(indexURL)
	if err != nil {
		return fmt.Errorf("error conectando a la biblioteca: %v", err)
	}
	defer respGet.Body.Close()

	bodyBytes, _ := io.ReadAll(respGet.Body)
	htmlContent := string(bodyBytes)

	re := regexp.MustCompile(`name="csrf_token" value="([^"]+)"`)
	matches := re.FindStringSubmatch(htmlContent)

	csrfToken := ""
	if len(matches) > 1 {
		csrfToken = matches[1]
	} else {
		return fmt.Errorf("no se pudo encontrar el token CSRF inicial")
	}

	initialCookie := ""
	for _, c := range respGet.Header["Set-Cookie"] {
		if strings.Contains(c, "MRBS_SESSID") {
			parts := strings.Split(c, ";")
			initialCookie = strings.TrimPrefix(parts[0], "MRBS_SESSID=")
		}
	}

	loginURL := fmt.Sprintf("%s/biblioteca/web/search.php", kc.baseUrl)

	formData := url.Values{}
	formData.Set("csrf_token", csrfToken)
	formData.Set("returl", "")
	formData.Set("target_url", "search.php")
	formData.Set("action", "SetName")
	formData.Set("username", kc.username)
	formData.Set("password", kc.password)
	formData.Set("datatable", "1")

	reqPost, err := http.NewRequest("POST", loginURL, strings.NewReader(formData.Encode()))
	if err != nil {
		return err
	}

	reqPost.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if initialCookie != "" {
		reqPost.Header.Add("Cookie", fmt.Sprintf("MRBS_SESSID=%s", initialCookie))
	}

	respPost, err := client.Do(reqPost)
	if err != nil {
		return err
	}
	defer respPost.Body.Close()

	for _, c := range respPost.Header["Set-Cookie"] {
		if strings.Contains(c, "MRBS_SESSID") {
			parts := strings.Split(c, ";")
			kc.session = strings.TrimPrefix(parts[0], "MRBS_SESSID=")
			return nil
		}
	}

	if initialCookie != "" && (respPost.StatusCode == 302 || respPost.StatusCode == 200) {
		kc.session = initialCookie
		return nil
	}

	return fmt.Errorf("login fallido. Verifica tu usuario y contraseña de la biblioteca")
}
