package queries

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/shurco/litecart/internal/models"
	"github.com/shurco/litecart/pkg/security"
)

// ProductQueries is ...
type ProductQueries struct {
	*sql.DB
}

// ListProducts is ...
func (q *ProductQueries) ListProducts(private bool) (*models.Products, error) {
	products := &models.Products{}

	queryPrivate := ` WHERE deleted = 0 AND active = 1 AND json_extract(stripe, '$.product.id') != '' AND json_extract(stripe, '$.product.valid') = 1`
	query := `
			SELECT 
				id, 
				name, 
				desc, 
				url, 
				active,
				json_extract(stripe, '$.product.id') as product_id, 
				json_extract(stripe, '$.product.valid') as product_valid, 
				json_extract(stripe, '$.price.id') as price_id, 
				json_extract(stripe, '$.price.currency') as currency, 
				json_extract(stripe, '$.price.amount') as amount,
				strftime('%s', created) 
			FROM product
		`
	queryTotal := `SELECT COUNT(id) FROM product`

	if !private {
		query = query + queryPrivate
		queryTotal = queryTotal + queryPrivate
	}

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
			&product.Stripe.Product.Valid,
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
	err = q.DB.QueryRow(queryTotal).Scan(&products.Total)
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
				json_group_array(json_object('id', product_image.id, 'name', product_image.name, 'ext', product_image.ext)) as images,
				strftime('%s', product.created), 
				strftime('%s', product.updated)
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
			&images,
			&product.Created,
			&updated,
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

	if images.Valid && images.String != `[{"id":null,"name":null,"ext":null}]` {
		json.Unmarshal([]byte(images.String), &product.Images)
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

// ProductImages is ...
func (q *ProductQueries) ProductImages(id string) (*[]models.Images, error) {
	images := &[]models.Images{}

	query := `
			SELECT 
				json_group_array(json_object('id', id, 'name', name, 'ext', ext)) as images
			FROM product_image 
			WHERE product_id = ?
	`
	var imgs sql.NullString
	err := q.DB.QueryRow(query, id).Scan(&imgs)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	if imgs.Valid && imgs.String != `[{"id":null,"name":null,"ext":null}]` {
		json.Unmarshal([]byte(imgs.String), &images)
	}

	return images, nil
}

// AddImage is ...
func (q *ProductQueries) AddImage(productID, fileUUID, fileExt string) (*models.Images, error) {
	file := &models.Images{
		ID:   security.RandomString(),
		Name: fileUUID,
		Ext:  fileExt,
	}

	// add db record
	_, err := q.DB.Exec(`INSERT INTO product_image (id, product_id, name, ext) VALUES (?, ?, ?, ?)`, file.ID, productID, fileUUID, fileExt)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// DeleteImage is ...
func (q *ProductQueries) DeleteImage(productID, imageID string) error {
	var name, ext string
	err := q.DB.QueryRow(`SELECT name, ext FROM product_image WHERE id = ?`, imageID).Scan(&name, &ext)
	if err != nil {
		return err
	}

	if _, err := q.DB.Exec(`DELETE FROM product_image WHERE id = ? AND product_id = ?`, imageID, productID); err != nil {
		return err
	}

	filePaths := []string{
		fmt.Sprintf("./uploads/%s.%s", name, ext),
		fmt.Sprintf("./uploads/%s_sm.%s", name, ext),
	}

	for _, filePath := range filePaths {
		if err := os.Remove(filePath); err != nil {
			return fmt.Errorf("failed to remove file %s: %w", filePath, err)
		}
	}

	return nil
}
