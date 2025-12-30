package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/shurco/litecart/internal/models"
	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/internal/testutil"
	"github.com/shurco/litecart/migrations"
	"github.com/shurco/litecart/pkg/jwtutil"
)

func setupAuthApp(t *testing.T) (*fiber.App, string, func()) {
	cleanup := testutil.WithCmdTestDir(t)
	if err := queries.New(migrations.Embed()); err != nil {
		t.Fatal(err)
	}
	app := fiber.New()

	// install and jwt
	db := queries.DB()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.Install(ctx, &models.Install{Email: "admin@example.com", Password: "secret", Domain: "example.com"}); err != nil {
		t.Fatal(err)
	}
	if err := db.UpdateSettingByGroup(ctx, &models.JWT{Secret: "testsecret", ExpireHours: 1}); err != nil {
		t.Fatal(err)
	}

	// generate token
	exp := time.Now().Add(time.Hour).Unix()
	tok, err := jwtutil.GenerateNewToken("testsecret", "id", exp, nil)
	if err != nil {
		t.Fatal(err)
	}

	cookie := "token=" + tok
	return app, cookie, func() { cleanup(); _ = os.Unsetenv("_") }
}

func Test_pages_crud(t *testing.T) {
	app, cookie, cleanup := setupAuthApp(t)
	defer cleanup()

	// routes
	app.Get("/api/_/pages", Pages)
	app.Post("/api/_/pages", AddPage)
	app.Patch("/api/_/pages/:page_id", UpdatePage)
	app.Patch("/api/_/pages/:page_id/content", UpdatePageContent)
	app.Delete("/api/_/pages/:page_id", DeletePage)

	// create
	body := map[string]any{"name": "P", "slug": "p", "position": "footer"}
	bb, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/api/_/pages", bytes.NewReader(bb))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", cookie)
	resp, _ := app.Test(req)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("create %d", resp.StatusCode)
	}
	var res struct {
		Result struct {
			ID string `json:"id"`
		} `json:"result"`
	}
	_ = json.NewDecoder(resp.Body).Decode(&res)
	_ = resp.Body.Close()
	if res.Result.ID == "" {
		t.Fatal("no id")
	}

	// list
	req = httptest.NewRequest(http.MethodGet, "/api/_/pages", nil)
	req.Header.Set("Cookie", cookie)
	resp, _ = app.Test(req)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("list %d", resp.StatusCode)
	}
	_ = resp.Body.Close()

	// update
	upd := map[string]any{"name": "P2", "slug": "p2", "position": "footer"}
	bb, _ = json.Marshal(upd)
	req = httptest.NewRequest(http.MethodPatch, "/api/_/pages/"+res.Result.ID, bytes.NewReader(bb))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", cookie)
	resp, _ = app.Test(req)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("update %d", resp.StatusCode)
	}
	_ = resp.Body.Close()

	// content
	cnt := map[string]any{"content": "hello"}
	bb, _ = json.Marshal(cnt)
	req = httptest.NewRequest(http.MethodPatch, "/api/_/pages/"+res.Result.ID+"/content", bytes.NewReader(bb))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", cookie)
	resp, _ = app.Test(req)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("content %d", resp.StatusCode)
	}
	_ = resp.Body.Close()

	// delete
	req = httptest.NewRequest(http.MethodDelete, "/api/_/pages/"+res.Result.ID, nil)
	req.Header.Set("Cookie", cookie)
	resp, _ = app.Test(req)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("delete %d", resp.StatusCode)
	}
	_ = resp.Body.Close()
}
