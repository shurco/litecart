package queries

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/shurco/litecart/internal/models"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/client"
)

var (
	StripeProductNotFound = errors.New("stripe product not found")
)

func stripeClient() (*client.API, error) {
	db := DB()
	stripe, err := db.SettingStripe()
	if err != nil {
		return nil, err
	}

	client := &client.API{}
	client.Init(stripe.SecretKey, nil)

	return client, nil
}

// AddStripeProduct is ...
func (q *ProductQueries) AddStripeProduct(productID string) (*models.StripeInfo, error) {
	stripeInfo := &models.StripeInfo{}
	product, err := q.Product(productID, true)
	if err != nil {
		return nil, err
	}

	var stripeProductID string
	err = q.DB.QueryRow(`SELECT json_extract(stripe, '$.product.id') as product_id 
			FROM product 
			WHERE id = ? 
				AND json_extract(stripe, '$.product.valid') = 1`, productID).Scan(&stripeProductID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	stripeInfo.ProductID = stripeProductID
	if stripeProductID != "" {
		return stripeInfo, nil
	}

	images := []string{}
	for _, image := range product.Images {
		path := fmt.Sprintf("https://%s/uploads/%s_md.%s", db.GetDomain(), image.Name, image.Ext)
		images = append(images, path)
	}

	newProduct := &stripe.ProductParams{
		Active:      stripe.Bool(true),
		Name:        stripe.String(product.Name),
		Description: stripe.String(product.Description),
		URL:         stripe.String(product.URL),
		Images:      stripe.StringSlice(images),
		Shippable:   stripe.Bool(true),
		DefaultPriceData: &stripe.ProductDefaultPriceDataParams{
			Currency:   stripe.String(product.Stripe.Price.Currency),
			UnitAmount: stripe.Int64(int64(product.Stripe.Price.Amount)),
		},
	}
	for key, value := range product.Metadata {
		newProduct.AddMetadata(key, value)
	}

	client, err := stripeClient()
	if err != nil {
		return nil, err
	}

	stripeProduct, err := client.Products.New(newProduct)
	if err != nil {
		return nil, err
	}

	stripeInfo.ProductID = stripeProduct.ID
	stripeInfo.PriceID = stripeProduct.DefaultPrice.ID

	return stripeInfo, nil
}

// UpdateStripeProduct is ...
func (q *ProductQueries) UpdateStripeProduct(productID string, stripe *models.StripeInfo) error {
	query := `
			UPDATE product SET 
				stripe = json_replace(stripe, '$.price.id', ?, '$.product.id', ?, '$.product.valid', 1), 
				updated = datetime('now') 
			WHERE id = ?
	`
	_, err := q.DB.Exec(query, stripe.PriceID, stripe.ProductID, productID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteStripeProduct is ...
func (q *ProductQueries) DeleteStripeProduct(productID string) error {
	var stripeProductID, stripePriceID string
	query := `
			SELECT 
				json_extract(stripe, '$.product.id') as product_id, 
				json_extract(stripe, '$.price.id') as price_id 
			FROM product 
			WHERE id = ?
	`

	err := q.DB.QueryRow(query, productID).Scan(&stripeProductID, &stripePriceID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if !q.IsStripeProduct(stripeProductID) {
		return StripeProductNotFound
	}

	client, err := stripeClient()
	if err != nil {
		return err
	}

	if stripeProductID != "" {
		if stripePriceID != "" {
			if _, err = client.Products.Update(stripeProductID, &stripe.ProductParams{Active: stripe.Bool(false)}); err != nil {
				return err
			}
		} else {
			if _, err = client.Products.Del(stripeProductID, nil); err != nil {
				return err
			}
		}
	}

	if stripePriceID != "" {
		if _, err := client.Prices.Update(stripePriceID, &stripe.PriceParams{Active: stripe.Bool(false)}); err != nil {
			return err
		}
	}

	return nil
}

// IsStripeProduct is ...
func (q *ProductQueries) IsStripeProduct(stripeProductID string) bool {
	client, err := stripeClient()
	if err != nil {
		return false
	}

	if stripeProductID != "" {
		if _, err := client.Products.Get(stripeProductID, nil); err != nil {
			return false
		}
		return true
	}

	return false
}
