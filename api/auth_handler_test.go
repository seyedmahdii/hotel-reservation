package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/seyedmahdii/hotel-reservation/db"
	"github.com/seyedmahdii/hotel-reservation/types"
)

func insertTestUser(t *testing.T, userStore db.UserStore) *types.User {
	user, err := types.NewUserFromParams(types.CreateUserParams{
		Email:     "mahdi@gmail.com",
		FirstName: "mahdi",
		LastName:  "jalali",
		Password:  "test1234",
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = userStore.InsertUser(context.TODO(), user)
	if err != nil {
		t.Fatal(err)
	}

	return user
}

func TestAuthenticateSuccess(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)
	insertedUser := insertTestUser(t, tdb.UserStore)

	app := fiber.New()
	authHandler := NewAuthHandler(tdb.UserStore)
	app.Post("/auth", authHandler.HandleAuthenticate)

	params := AuthParams{
		Email:    "mahdi@gmail.com",
		Password: "test1234",
	}
	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST", "/auth", bytes.NewReader(b))
	req.Header.Add("content-type", "application/json")
	res, err := app.Test(req)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}

	var authRes AuthResponse
	if err := json.NewDecoder(res.Body).Decode(&authRes); err != nil {
		t.Fatal(err)
	}

	if authRes.Token == "" {
		t.Fatalf("expected non-empty token")
	}

	// Set the encrypted password to an empty string, because we don't return that in the response
	insertedUser.EncryptedPassword = ""

	if !reflect.DeepEqual(authRes.User, insertedUser) {
		t.Fatalf("expected user %+v, got %+v", insertedUser, authRes.User)
	}
}

func TestAuthenticateWithWrongPassword(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	app := fiber.New()
	authHandler := NewAuthHandler(tdb.UserStore)
	app.Post("/auth", authHandler.HandleAuthenticate)

	params := AuthParams{
		Email:    "mahdi@gmail.com",
		Password: "WRONG_PASSWORD",
	}
	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST", "/auth", bytes.NewReader(b))
	req.Header.Add("content-type", "application/json")
	res, err := app.Test(req)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}

	var genRes genericResp
	if err := json.NewDecoder(res.Body).Decode(&genRes); err != nil {
		t.Fatal(err)
	}

	if genRes.Type != "error" || genRes.Msg != "invalid credentials" {
		t.Fatalf("expected error response with message 'invalid credentials', got %+v", genRes)
	}
}
