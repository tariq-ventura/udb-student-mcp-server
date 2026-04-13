package biblio_koha

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func (kc *KohaClient) FetchBiblioAccount() (string, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s/cgi-bin/koha/opac-user.pl", kc.baseUrl)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Cookie", fmt.Sprintf("CGISESSID=%s", kc.biblioSession))

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	html := string(bodyBytes)

	reScripts := regexp.MustCompile(`(?is)<script.*?>.*?</script>`)
	reStyles := regexp.MustCompile(`(?is)<style.*?>.*?</style>`)
	html = reScripts.ReplaceAllString(html, "")
	html = reStyles.ReplaceAllString(html, "")

	return html, nil
}
