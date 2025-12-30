package handlers

import (
	"bytes"
	"context"
	"image"
	"image/color"
	"image/png"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/shurco/litecart/internal/models"
	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/internal/testutil"
	"github.com/shurco/litecart/migrations"
)

func setupProductEnv(t *testing.T) (*fiber.App, func()) {
	cleanup := testutil.WithCmdTestDir(t)
	if err := queries.New(migrations.Embed()); err != nil {
		t.Fatal(err)
	}
	app := fiber.New()
	return app, func() { cleanup() }
}

func Test_add_product_image_upload_and_resize(t *testing.T) {
	app, cleanup := setupProductEnv(t)
	defer cleanup()

	db := queries.DB()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// prepare product
	p, err := db.AddProduct(ctx, &models.Product{
		Name:    "p1",
		Amount:  100,
		Slug:    "p1",
		Digital: models.Digital{Type: "file"},
	})
	if err != nil {
		t.Fatal(err)
	}

	// build PNG 400x400 in memory
	img := image.NewRGBA(image.Rect(0, 0, 400, 400))
	for y := 0; y < 400; y++ {
		for x := 0; x < 400; x++ {
			img.Set(x, y, color.RGBA{255, 0, 0, 255})
		}
	}
	var imgBuf bytes.Buffer
	if err := png.Encode(&imgBuf, img); err != nil {
		t.Fatal(err)
	}

	var body bytes.Buffer
	w := multipart.NewWriter(&body)
	hdr := make(textproto.MIMEHeader)
	hdr.Set("Content-Disposition", `form-data; name="document"; filename="test.png"`)
	hdr.Set("Content-Type", "image/png")
	fw, err := w.CreatePart(hdr)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := fw.Write(imgBuf.Bytes()); err != nil {
		t.Fatal(err)
	}
	_ = w.Close()

	req := httptest.NewRequest(http.MethodPost, "/api/_/products/"+p.ID+"/image", &body)
	req.Header.Set("Content-Type", w.FormDataContentType())

	// register route directly to handler
	app.Post("/api/_/products/:product_id/image", AddProductImage)

	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status %d", resp.StatusCode)
	}

	// assert that original and sm/md resized files exist
	entries, err := os.ReadDir("./lc_uploads")
	if err != nil {
		t.Fatal(err)
	}
	var haveOrig, haveSm, haveMd bool
	for _, e := range entries {
		name := e.Name()
		if filepath.Ext(name) == ".png" {
			if len(name) > 4 && name[len(name)-4:] == ".png" {
				haveOrig = haveOrig || (len(name) > 4 && (name[len(name)-7:] != "_sm.png" && name[len(name)-7:] != "_md.png"))
			}
			haveSm = haveSm || (len(name) > 7 && name[len(name)-7:] == "_sm.png")
			haveMd = haveMd || (len(name) > 7 && name[len(name)-7:] == "_md.png")
		}
	}
	if !haveSm || !haveMd || !haveOrig {
		t.Fatalf("expected original+sm+md images")
	}
}
