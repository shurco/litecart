package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/shurco/litecart/internal/models"
	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/migrations"
	"github.com/shurco/litecart/pkg/jwtutil"
)

func TestJWTProtected_BearerFlow(t *testing.T) {
	// init temp DB
	app := fiber.New()
	if err := queries.New(migrations.Embed()); err != nil {
		t.Fatalf("init queries: %v", err)
	}
	db := queries.DB()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// write JWT settings
	setting := &models.JWT{Secret: "secret", ExpireHours: 1}
	if err := db.UpdateSettingByGroup(ctx, setting); err != nil {
		t.Fatalf("update jwt setting: %v", err)
	}

	app.Use(JWTProtected())
	app.Get("/api/test", func(c *fiber.Ctx) error { return c.SendStatus(200) })

	// no token → 401
	req := httptest.NewRequest(http.MethodGet, "/api/test", nil)
	resp, _ := app.Test(req)
	if resp.StatusCode != http.StatusUnauthorized && resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 401/400, got %d", resp.StatusCode)
	}

	// valid token
	userID := uuid.NewString()
	exp := time.Now().Add(time.Hour).Unix()
	tok, err := jwtutil.GenerateNewToken("secret", userID, exp, nil)
	if err != nil {
		t.Fatalf("token gen: %v", err)
	}

	req2 := httptest.NewRequest(http.MethodGet, "/api/test", nil)
	req2.Header.Set("Authorization", "Bearer "+tok)
	resp2, _ := app.Test(req2)
	if resp2.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp2.StatusCode)
	}
}
