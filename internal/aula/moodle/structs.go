package aula_moodle

type MoodleTokenResponse struct {
	Token string `json:"token"`
	Error string `json:"error"`
}

type MoodleSiteInfo struct {
	UserID    int    `json:"userid"`
	ErrorCode string `json:"errorcode,omitempty"`
	Message   string `json:"message,omitempty"`
}

type MoodleAPIError struct {
	Exception string `json:"exception"`
	ErrorCode string `json:"errorcode"`
	Message   string `json:"message"`
}

type MoodleFile struct {
	Filename string `json:"filename"`
	FileURL  string `json:"fileurl"`
	MimeType string `json:"mimetype,omitempty"`
}

type MoodleModule struct {
	Name     string       `json:"name"`
	ModName  string       `json:"modname"`
	Contents []MoodleFile `json:"contents,omitempty"`
}

type MoodleSection struct {
	Name    string         `json:"name"`
	Modules []MoodleModule `json:"modules,omitempty"`
}
