package gotrue

import (
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
