package api

import (
	"fmt"
	"lambda-func/database"
	"lambda-func/types"
)

type ApiHandler struct {
	dbStore database.DynamoDBClient
}

func NewApiHandler(dbStore database.DynamoDBClient) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}

func (api ApiHandler) RegisterUserHandler(event types.RegisterUser) error {
	if event.Username == "" || event.Password == "" {
		return fmt.Errorf("missing required parameter")
	}
	// does a user with this username already exist?

	exists, err := api.dbStore.DoesUserExist(event.Username)
	if err != nil {
		return fmt.Errorf("error checking user exists: %w", err)
	}
	if exists {
		return fmt.Errorf("cannot register username")
	}
	err = api.dbStore.InsertUser(event)
	if err != nil {
		return fmt.Errorf("error registering user: %w", err)
	}
	return nil
}
