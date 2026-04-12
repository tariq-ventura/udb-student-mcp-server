package aula_moodle

import (
	"context"
	"fmt"
	"os"
)

type MoodleClient struct {
	baseUrl  string
	username string
	password string
	token    string
	userId   int
	ctx      context.Context
}

var SetupMoodle = func(ctx context.Context) (*MoodleClient, error) {
	baseUrl, find := os.LookupEnv("MOODLE_BASE_URL")

	if !find {
		return nil, fmt.Errorf("MOODLE_BASE_URL environment variable not found")
	}

	username, find := os.LookupEnv("MOODLE_USERNAME")

	if !find {
		return nil, fmt.Errorf("MOODLE_USERNAME environment variable not found")
	}

	password, find := os.LookupEnv("MOODLE_PASSWORD")

	if !find {
		return nil, fmt.Errorf("MOODLE_PASSWORD environment variable not found")
	}

	return &MoodleClient{
		baseUrl:  baseUrl,
		username: username,
		password: password,
		token:    "",
		userId:   0,
		ctx:      ctx,
	}, nil
}
