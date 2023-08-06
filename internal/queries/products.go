package queries

import (
	"database/sql"
	"encoding/json"
	"errors"
	"strings"

	"github.com/shurco/litecart/internal/models"
)

// ProductQueries is ...
type ProductQueries struct {
	*sql.DB
}

// ListProducts is ...
func (q *ProductQueries) ListProducts() (*models.Products, error) {
	products := &models.Products{}

	query := `
			SELECT 
				id, 
				name, 
				desc, 
				url, 
				active,
				json_extract(stripe, '$.product.id') as product_id, 
				json_extract(stripe, '$.price.id') as price_id, 
				json_extract(stripe, '$.price.currency') as currency, 
				json_extract(stripe, '$.price.amount') as amount, 
				strftime('%s', created) 
			FROM product
			WHERE deleted = 0
		`

	rows, err := q.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		product := models.Product{}
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.URL,
			&product.Active,
			&product.Stripe.Product.ID,
			&product.Stripe.Price.ID,
			&product.Stripe.Price.Currency,
			&product.Stripe.Price.Amount,
			&product.Created,
		)
		if err != nil {
			return nil, err
		}

		products.Products = append(products.Products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// total records
	err = q.DB.QueryRow(`SELECT COUNT(id) FROM product`).Scan(&products.Total)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return products, nil
}

// Product is ...
func (q *ProductQueries) Product(id string) (*models.Product, error) {
	product := &models.Product{
		ID: id,
	}

	query := `
			SELECT 
				product.name, 
				product.desc, 
				product.url, 
				product.active,
				json_extract(stripe, '$.product.id') as product_id, 
				json_extract(stripe, '$.product.valid') as product_valid, 
				json_extract(stripe, '$.price.id') as price_id, 
				json_extract(stripe, '$.price.currency') as currency, 
				json_extract(stripe, '$.price.amount') as amount, 
				product.metadata, 
				product.attribute, 
				strftime('%s', product.created), 
				strftime('%s', product.updated),
				group_concat(DISTINCT product_image.name || '.' || product_image.ext) as images
			FROM product 
			LEFT JOIN product_image ON product.id = product_image.product_id
			WHERE product.id = ?
			GROUP BY product.id
	`
	// stripeID
	var images, metadata, attributes sql.NullString
	var updated sql.NullInt64

	err := q.DB.QueryRow(query, id).
		Scan(
			&product.Name,
			&product.Description,
			&product.URL,
			&product.Active,
			&product.Stripe.Product.ID,
			&product.Stripe.Product.Valid,
			&product.Stripe.Price.ID,
			&product.Stripe.Price.Currency,
			&product.Stripe.Price.Amount,
			&metadata,
			&attributes,
			&product.Created,
			&updated,
			&images,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	if updated.Valid {
		product.Updated = updated.Int64
	}

	if images.Valid {
		product.Images = strings.Split(images.String, ",")
	}

	if attributes.Valid {
		json.Unmarshal([]byte(attributes.String), &product.Attributes)
	}

	if metadata.Valid {
		json.Unmarshal([]byte(metadata.String), &product.Metadata)
	}

	return product, nil
}

// AddProduct is ...
func (q *ProductQueries) AddProduct(product *models.Product) error {
	return nil
}

// UpdateProduct is ...
func (q *ProductQueries) UpdateProduct(product *models.Product) error {
	return nil
}

// DeleteProduct is ...
func (q *ProductQueries) DeleteProduct(id string) error {
	if err := q.DeleteStripeProduct(id); err != nil && err != StripeProductNotFound {
		return err
	}

	var stripeProductID, stripePriceID string
	query := `
			SELECT 
				json_extract(stripe, '$.product.id') as product_id ,
				json_extract(stripe, '$.price.id') as price_id 
			FROM product 
			WHERE id = ?
	`
	err := q.DB.QueryRow(query, id).Scan(&stripeProductID, &stripePriceID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if !q.IsStripeProduct(stripeProductID) {
		if _, err := q.DB.Exec(`DELETE FROM product WHERE id = ?`, id); err != nil {
			return err
		}
	}

	if stripePriceID != "" {
		if _, err := q.DB.Exec(`UPDATE product SET active = 0, deleted = 1, updated = datetime('now') WHERE id = ?`, id); err != nil {
			return err
		}
	} else {
		if _, err := q.DB.Exec(`DELETE FROM product WHERE id = ?`, id); err != nil {
			return err
		}
	}

	return nil
}

// UpdateActive is ...
func (q *ProductQueries) UpdateActive(id string) error {
	var active bool
	query := `
			SELECT 
				active
			FROM product 
			WHERE id = ?
	`
	err := q.DB.QueryRow(query, id).Scan(&active)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if _, err := q.DB.Exec(`UPDATE product SET active = ?, updated = datetime('now') WHERE id = ?`, !active, id); err != nil {
		return err
	}

	return nil
}
