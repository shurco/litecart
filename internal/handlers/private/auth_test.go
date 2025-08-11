package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/shurco/litecart/internal/models"
	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/internal/testutil"
	"github.com/shurco/litecart/migrations"
)

func setupApp(t *testing.T) (*fiber.App, func()) {
	t.Helper()
	cleanup := testutil.WithCmdTestDir(t)

	if err := queries.New(migrations.Embed()); err != nil {
		t.Fatal(err)
	}
	app := fiber.New()
	return app, func() { cleanup(); _ = os.Unsetenv("_") }
}

func TestAuth_SignInOut(t *testing.T) {
	app, cleanup := setupApp(t)
	defer cleanup()

	// prepare settings (install + jwt)
	db := queries.DB()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// install minimal creds
	inst := &models.Install{Email: "admin@example.com", Password: "secret", Domain: "example.com"}
	if err := db.Install(ctx, inst); err != nil {
		t.Fatal(err)
	}

	// jwt expiry shorter
	if err := db.UpdateSettingByGroup(ctx, &models.JWT{Secret: "secretjwt", ExpireHours: 1}); err != nil {
		t.Fatal(err)
	}

	app.Post("/api/sign/in", SignIn)
	app.Post("/api/sign/out", SignOut)

	// sign in
	body := `{"email":"admin@example.com","password":"secret"}`
	req := httptest.NewRequest(http.MethodPost, "/api/sign/in", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("signin status %d", resp.StatusCode)
	}

	// extract cookie
	cookie := resp.Header.Get("Set-Cookie")
	if cookie == "" {
		t.Fatalf("expected token cookie")
	}

	// sign out
	req2 := httptest.NewRequest(http.MethodPost, "/api/sign/out", nil)
	req2.Header.Set("Cookie", cookie)
	resp2, err := app.Test(req2)
	if err != nil {
		t.Fatal(err)
	}
	if resp2.StatusCode != http.StatusNoContent {
		t.Fatalf("signout status %d", resp2.StatusCode)
	}
}
