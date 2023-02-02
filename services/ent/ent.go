package ent

import (
	"context"
	"log"

	"github.com/abc3354/CODEV-back/ent"
	"github.com/abc3354/CODEV-back/services/env"

	"github.com/abc3354/CODEV-back/ent/enttest"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

var client *ent.Client

func Get() *ent.Client {
	if client == nil {
		log.Println("Creating new ent client")
		openingClient, err := ent.Open("postgres", env.Get().Datasource)
		if err != nil {
			log.Fatal(err)
		}
		if err = openingClient.Schema.Create(context.Background()); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
		client = openingClient
	}
	return client
}

func Close() {
	err := client.Close()
	if err != nil {
		log.Fatal(err)
	}
	client = nil
}

func Mock(t *testing.T) {
	if client != nil {
		log.Fatal("ent client already exists")
	}
	client = enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
}
