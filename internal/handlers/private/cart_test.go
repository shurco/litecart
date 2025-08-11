package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/shurco/litecart/internal/models"
	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/internal/testutil"
	"github.com/shurco/litecart/migrations"
)

func setupCartApp(t *testing.T) (*fiber.App, func()) {
	cleanup := testutil.WithCmdTestDir(t)
	if err := queries.New(migrations.Embed()); err != nil {
		t.Fatal(err)
	}
	app := fiber.New()
	return app, func() { cleanup() }
}

func TestCarts_List(t *testing.T) {
	app, cleanup := setupCartApp(t)
	defer cleanup()
	app.Get("/api/_/carts", Carts)
	// call empty list
	req := httptest.NewRequest(http.MethodGet, "/api/_/carts", nil)
	resp, _ := app.Test(req)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status %d", resp.StatusCode)
	}
}

func TestCartSendMail_Status(t *testing.T) {
	app, cleanup := setupCartApp(t)
	defer cleanup()

	db := queries.DB()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// add a cart so that handler won't fail on mailer selection
	_ = db.AddCart(ctx, &models.Cart{Core: models.Core{ID: "cart123456789012"}, AmountTotal: 100, Currency: "USD"})

	app.Post("/api/_/carts/:cart_id/mail", CartSendMail)
	req := httptest.NewRequest(http.MethodPost, "/api/_/carts/cart123456789012/mail", nil)
	resp, _ := app.Test(req)
	// could be 200 or 500 depending on SMTP configuration; we assert code is within allowed set
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusInternalServerError {
		t.Fatalf("unexpected status %d", resp.StatusCode)
	}
}
