package biblio

import (
	"fmt"
	"os"

	biblio_koha "github.com/tariq-ventura/udb-mcp/internal/biblio/koha"
)

type IBiblio interface {
	GetSession() error
	ReserveMRBS(dateStr, startSeconds, endSeconds, roomId string) (string, error)
	GetBiblioSession() error
	FetchBiblioAccount() (string, error)
}

var NewBiblio = func() (IBiblio, error) {
	biblio, find := os.LookupEnv("BIBLIO_TYPE")

	if !find {
		return nil, fmt.Errorf("BIBLIO_TYPE environment variable not found")
	}

	switch biblio {
	case "koha":
		return biblio_koha.SetupKoha()
	default:
		return nil, fmt.Errorf("unsupported BIBLIO_TYPE: %s", biblio)
	}
}
