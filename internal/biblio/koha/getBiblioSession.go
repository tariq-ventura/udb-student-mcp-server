package biblio_koha

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

func (kc *KohaClient) GetBiblioSession() error {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	loginURL := fmt.Sprintf("%s/cgi-bin/koha/opac-user.pl", kc.baseUrl)
	respGet, err := client.Get(loginURL)
	if err != nil {
		return err
	}
	defer respGet.Body.Close()

	bodyBytes, _ := io.ReadAll(respGet.Body)
	htmlContent := string(bodyBytes)

	re := regexp.MustCompile(`name="csrf_token"\s*value="([^"]+)"`)
	matches := re.FindStringSubmatch(htmlContent)
	csrfToken := ""
	if len(matches) > 1 {
		csrfToken = matches[1]
	}

	initialCookie := ""
	for _, c := range respGet.Header["Set-Cookie"] {
		if strings.Contains(c, "CGISESSID") {
			parts := strings.Split(c, ";")
			initialCookie = strings.TrimPrefix(parts[0], "CGISESSID=")
		}
	}

	formData := url.Values{}
	formData.Set("csrf_token", csrfToken)
	formData.Set("koha_login_context", "opac")
	formData.Set("op", "cud-login")
	formData.Set("login_userid", kc.username)
	formData.Set("login_password", kc.password)

	reqPost, _ := http.NewRequest("POST", loginURL, strings.NewReader(formData.Encode()))
	reqPost.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if initialCookie != "" {
		reqPost.Header.Add("Cookie", fmt.Sprintf("CGISESSID=%s", initialCookie))
	}

	respPost, err := client.Do(reqPost)
	if err != nil {
		return err
	}
	defer respPost.Body.Close()

	for _, c := range respPost.Header["Set-Cookie"] {
		if strings.Contains(c, "CGISESSID") {
			parts := strings.Split(c, ";")
			kc.biblioSession = strings.TrimPrefix(parts[0], "CGISESSID=")
			return nil
		}
	}

	if respPost.StatusCode == 302 {
		kc.biblioSession = initialCookie
		return nil
	}
	return fmt.Errorf("login fallido en Koha")
}
