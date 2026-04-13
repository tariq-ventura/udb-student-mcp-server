package biblio_koha

import (
	"fmt"
	"os"
)

type KohaClient struct {
	baseUrl  string
	username string
	password string
	session  string
}

var SetupKoha = func() (*KohaClient, error) {
	baseUrl, find := os.LookupEnv("KOHA_BASE_URL")

	if !find {
		return nil, fmt.Errorf("KOHA_BASE_URL environment variable not found")
	}

	username, find := os.LookupEnv("KOHA_USERNAME")

	if !find {
		return nil, fmt.Errorf("KOHA_USERNAME environment variable not found")
	}

	password, find := os.LookupEnv("KOHA_PASSWORD")

	if !find {
		return nil, fmt.Errorf("KOHA_PASSWORD environment variable not found")
	}

	return &KohaClient{
		baseUrl:  baseUrl,
		username: username,
		password: password,
		session:  "",
	}, nil
}
