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
func (q *ProductQueries) ListProducts(private bool, idList ...string) (*models.Products, error) {
	products := &models.Products{
		Currency: "USD",
	}

	queryPrivate := ` WHERE deleted = 0 AND active = 1`
	query := `
			SELECT 
				id, 
				name, 
				desc, 
				url,
				amount,
				active,
				(SELECT json_group_array(json_object('id', product_image.id, 'name', product_image.name, 'ext', product_image.ext)) as images FROM product_image WHERE product_id = product.id LIMIT 3) as image,
				strftime('%s', created)
			FROM product
		`
	queryTotal := `SELECT COUNT(id) FROM product`

	queryList := ""
	if len(idList) > 0 {
		var queryType = map[bool]string{
			false: "AND",
			true:  "WHERE",
		}

		list := ""
		for _, id := range idList {
			list += list + fmt.Sprintf(",'%s'", id)
		}

		queryList = fmt.Sprintf(" %s id IN (%s)", queryType[private], list[1:])
	}

	if !private {
		query = query + queryPrivate
		queryTotal = queryTotal + queryPrivate
	}

	rows, err := q.DB.Query(query + queryList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var image sql.NullString
		product := models.Product{}

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Url,
			&product.Amount,
			&product.Active,
			&image,
			&product.Created,
		)
		if err != nil {
			return nil, err
		}

		if image.Valid && image.String != `[{"id":null,"name":null,"ext":null}]` {
			json.Unmarshal([]byte(image.String), &product.Images)
		}

		products.Products = append(products.Products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// total records
	err = q.DB.QueryRow(queryTotal + queryList).Scan(&products.Total)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return products, nil
}

// Product is ...
func (q *ProductQueries) Product(id string, private bool) (*models.Product, error) {
	product := &models.Product{}

	query := `
			SELECT 
				product.id,
				product.name, 
				product.desc, 
				product.url, 
				product.amount,
				product.active,
				product.metadata, 
				product.attribute, 
				json_group_array(json_object('id', product_image.id, 'name', product_image.name, 'ext', product_image.ext)) as images,
				strftime('%s', product.created), 
				strftime('%s', product.updated)
			FROM product 
			LEFT JOIN product_image ON product.id = product_image.product_id
	`
	if private {
		query = query + `WHERE product.id = ?`
	} else {
		query = query + `WHERE product.url = ? AND product.active = 1`
	}

	// stripeID
	var images, metadata, attributes sql.NullString
	var updated sql.NullInt64

	err := q.DB.QueryRow(query, id).
		Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Url,
			&product.Amount,
			&product.Active,
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
func (q *ProductQueries) AddProduct(product *models.Product) (*models.Product, error) {
	product.ID = security.RandomString()
	metadata, _ := json.Marshal(product.Metadata)
	attributes, _ := json.Marshal(product.Attributes)

	sql := `INSERT INTO product (id, name, amount, url, metadata, attribute, desc) VALUES (?, ?, ?, ?, ?, ?, ?) RETURNING strftime('%s', created)`
	err := q.DB.QueryRow(sql, product.ID, product.Name, product.Amount, product.Url, metadata, attributes, product.Description).Scan(&product.Created)
	if err != nil {
		return nil, err
	}

	return product, nil
}

// UpdateProduct is ...
func (q *ProductQueries) UpdateProduct(product *models.Product) error {
	metadata, _ := json.Marshal(product.Metadata)
	attributes, _ := json.Marshal(product.Attributes)

	_, err := q.DB.Exec(`UPDATE product SET name = ?, desc = ?, url = ?, amount = ?, metadata = ?, attribute=?, updated = datetime('now') WHERE id = ?`,
		product.Name,
		product.Description,
		product.Url,
		product.Amount,
		metadata,
		attributes,
		product.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

// DeleteProduct is ...
func (q *ProductQueries) DeleteProduct(id string) error {
	if _, err := q.DB.Exec(`DELETE FROM product WHERE id = ?`, id); err != nil {
		return err
	}

	return nil
}

// IsProduct is ...
func (q *ProductQueries) IsProduct(url string) bool {
	var id string
	query := `
			SELECT 
				id
			FROM product 
			WHERE url = ? AND active = 1
	`
	err := q.DB.QueryRow(query, url).Scan(&id)
	if err != nil {
		return false
	}
	if id != "" {
		return true
	}

	return false
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
		fmt.Sprintf("./lc_uploads/%s.%s", name, ext),
		fmt.Sprintf("./lc_uploads/%s_sm.%s", name, ext),
	}

	for _, filePath := range filePaths {
		if err := os.Remove(filePath); err != nil {
			return fmt.Errorf("failed to remove file %s: %w", filePath, err)
		}
	}

	return nil
}
