package gotrue

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/abc3354/CODEV-back/local"
	"github.com/abc3354/CODEV-back/services/env"
	"github.com/supabase-community/gotrue-go"
)

var client gotrue.Client = nil

func Get() gotrue.Client {
	if client == nil {
		vars := env.Get()
		client = gotrue.New(vars.SupabaseProjectRef, vars.SupabaseAPIKey)
	}
	return client
}

type gotrueError struct {
	StatusCode  int    `json:"-"`
	Code        string `json:"error"`
	Description string `json:"error_description"`
}

func (e gotrueError) Error() string {
	return fmt.Sprintf("%d - %s: %s", e.StatusCode, e.Code, e.Description)
}

func ParseError(err error) *local.Error {
	if err == nil {
		return nil
	}

	if strings.HasPrefix(err.Error(), "response status code") {
		var content gotrueError
		data := strings.TrimPrefix(err.Error(), "response status code ")
		split := strings.Split(data, ": ")
		content.StatusCode, err = strconv.Atoi(split[0])
		if err != nil {
			return &local.Error{
				StatusCode: 500,
				Origin:     "services/gotrue",
				Code:       "status_atoi",
				From:       err,
			}
		}
		err = json.Unmarshal([]byte(split[1]), &content)
		if err != nil {
			return &local.Error{
				StatusCode: 500,
				Origin:     "services/gotrue",
				Code:       "json_unmarshal",
				From:       err,
			}
		}
		return &local.Error{
			StatusCode: 400,
			Origin:     "supabase",
			Code:       content.Code,
			From:       content,
		}
	} else {
		return &local.Error{
			StatusCode: 500,
			Origin:     "services/gotrue",
			Code:       "unknown",
			From:       err,
		}
	}
}
