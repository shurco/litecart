package queries

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/shurco/litecart/internal/models"
)

type ProductQueries struct {
	*sql.DB
}

func (q *ProductQueries) ListProducts() ([]models.Product, error) {
	products := []models.Product{}

	query := `
			SELECT p."id", p."name", p."desc", p."url", strftime('%s', p."created"), pp."id", pp."stripe_id", pp."currency", pp."amount"
      FROM "product" AS p
      JOIN "product_price" AS pp ON p."id" = pp."product_id"`

	rows, err := q.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		product := models.Product{}
		err := rows.Scan(&product.ID,
			&product.Name,
			&product.Description,
			&product.URL,
			&product.Created,
			&product.Price.ID,
			&product.Price.StripeID,
			&product.Price.Currency,
			&product.Price.Amount,
		)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (q *ProductQueries) Product(id string) (*models.Product, error) {
	product := &models.Product{
		ID: id,
	}

	query := `
			SELECT 
					"product"."name", 
					"product"."desc", 
					"product"."url", 
					strftime('%s', "product"."created"), 
					strftime('%s', "product"."updated"),
					"product_price"."id",
					"product_price"."stripe_id",
					"product_price"."currency",
					"product_price"."amount",
					group_concat(DISTINCT "product_image"."name" || '.' || "product_image"."ext") as images,
					group_concat(DISTINCT "product_metadata"."key" || ':' || "product_metadata"."value") as metadata,
					group_concat(DISTINCT "product_attribute"."name") as attributes
			FROM "product" 
			LEFT JOIN "product_price" ON "product"."id" = "product_price"."product_id"
			LEFT JOIN "product_image" ON "product"."id" = "product_image"."product_id"
			LEFT JOIN "product_metadata" ON "product"."id" = "product_metadata"."product_id"
			LEFT JOIN "product_attribute" ON "product"."id" = "product_attribute"."product_id"
			WHERE "product"."id" = ?
			GROUP BY "product"."id"
	`
	var updated int64
	var images, metadata, attributes string

	err := q.DB.QueryRow(query, id).
		Scan(
			&product.Name,
			&product.Description,
			&product.URL,
			&product.Created,
			&updated,
			&product.Price.ID,
			&product.Price.StripeID,
			&product.Price.Currency,
			&product.Price.Amount,
			&images,
			&metadata,
			&attributes,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	if updated != product.Created {
		product.Updated = updated
	}

	if images != "" {
		product.Images = strings.Split(images, ",")
	}

	if attributes != "" {
		product.Attributes = strings.Split(attributes, ",")
	}

	if metadata != "" {
		meta := strings.Split(metadata, ",")
		product.Metadata = make(map[string]string, len(meta))
		for _, kv := range meta {
			parts := strings.SplitN(kv, ":", 2)
			if len(parts) > 1 {
				product.Metadata[parts[0]] = parts[1]
			}
		}
	}

	return product, nil
}

/*
func (b *Base) GetStripeProducts() []models.Product {
	products := []models.Product{}
	list := b.Stripe.Products.List(nil)
	for _, val := range list.ProductList().Data {
		price, _ := b.Stripe.Prices.Get(val.DefaultPrice.ID, nil)

		product := models.Product{
			ID:          val.ID,
			Name:        val.Name,
			Description: val.Description,
			Price: models.Price{
				ID:       val.DefaultPrice.ID,
				Currency: string(price.Currency),
				Amount:   int(price.UnitAmount),
			},
			Images:     val.Images,
			URL:        val.URL,
			Metadata:   val.Metadata,
			Attributes: val.Attributes,
			Created:    val.Created,
			Updated:    val.Updated,
		}

		products = append(products, product)
	}

	return products
}

func (b *Base) AddStripeProducts() *stripe.Product {
	newProduct := &stripe.ProductParams{
		Active:      stripe.Bool(true),
		Name:        stripe.String("Test Name"),
		Description: stripe.String("This is a description"),
		URL:         stripe.String("http://example.com"),
		Images: stripe.StringSlice([]string{
			"https://werbot.com/assets/img/steps-1.png",
			"https://werbot.com/assets/img/steps-2.png",
			"https://werbot.com/assets/img/steps-3.png",
		}),
		Shippable: stripe.Bool(true),
		DefaultPriceData: &stripe.ProductDefaultPriceDataParams{
			Currency:   stripe.String(string(stripe.CurrencyEUR)),
			UnitAmount: stripe.Int64(5000),
		},
	}
	newProduct.AddMetadata("key", "value")

	recponse, err := b.Stripe.Products.New(newProduct)
	if err != nil {
		b.Log.Fatal().Err(err).Send()
	}

	return recponse
}
*/
