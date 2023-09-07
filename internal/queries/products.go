package queries

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/shurco/litecart/internal/models"
	"github.com/shurco/litecart/pkg/errors"
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
				slug,
				amount,
				active,
				digital,
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
		var image, digitalType sql.NullString
		product := models.Product{}

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Slug,
			&product.Amount,
			&product.Active,
			&digitalType,
			&image,
			&product.Created,
		)
		if err != nil {
			return nil, err
		}

		if image.Valid && image.String != `[{"id":null,"name":null,"ext":null}]` {
			json.Unmarshal([]byte(image.String), &product.Images)
		}

		if digitalType.Valid {
			product.Digital.Type = digitalType.String
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
func (q *ProductQueries) Product(private bool, id string) (*models.Product, error) {
	product := &models.Product{}

	query := `
			SELECT 
				product.id,
				product.name, 
				product.desc, 
				product.slug, 
				product.amount,
				product.active,
				product.metadata, 
				product.attribute, 
				product.digital,
				json_group_array(json_object('id', product_image.id, 'name', product_image.name, 'ext', product_image.ext)) as images,
				strftime('%s', product.created), 
				strftime('%s', product.updated)
			FROM product 
			LEFT JOIN product_image ON product.id = product_image.product_id
	`
	if private {
		query = query + `WHERE product.id = ?`
	} else {
		query = query + `WHERE product.slug = ? AND product.active = 1`
	}

	// stripeID
	var images, metadata, attributes, digitalType sql.NullString
	var updated sql.NullInt64

	err := q.DB.QueryRow(query, id).
		Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Slug,
			&product.Amount,
			&product.Active,
			&metadata,
			&attributes,
			&digitalType,
			&images,
			&product.Created,
			&updated,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrProductNotFound
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

	if digitalType.Valid {
		product.Digital.Type = digitalType.String
	}

	return product, nil
}

// AddProduct is ...
func (q *ProductQueries) AddProduct(product *models.Product) (*models.Product, error) {
	product.ID = security.RandomString()
	metadata, _ := json.Marshal(product.Metadata)
	attributes, _ := json.Marshal(product.Attributes)

	sql := `INSERT INTO product (id, name, amount, slug, metadata, attribute, desc, digital) VALUES (?, ?, ?, ?, ?, ?, ?, ?) RETURNING strftime('%s', created)`
	err := q.DB.QueryRow(sql, product.ID, product.Name, product.Amount, product.Slug, metadata, attributes, product.Description, product.Digital.Type).Scan(&product.Created)
	if err != nil {
		return nil, err
	}

	return product, nil
}

// UpdateProduct is ...
func (q *ProductQueries) UpdateProduct(product *models.Product) error {
	metadata, _ := json.Marshal(product.Metadata)
	attributes, _ := json.Marshal(product.Attributes)

	_, err := q.DB.Exec(`UPDATE product SET name = ?, desc = ?, slug = ?, amount = ?, metadata = ?, attribute=?, updated = datetime('now') WHERE id = ?`,
		product.Name,
		product.Description,
		product.Slug,
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
func (q *ProductQueries) IsProduct(slug string) bool {
	var id string
	query := `
			SELECT 
				id
			FROM product 
			WHERE slug = ? AND active = 1
	`
	err := q.DB.QueryRow(query, slug).Scan(&id)
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
func (q *ProductQueries) ProductImages(id string) (*[]models.File, error) {
	images := &[]models.File{}

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
			return nil, errors.ErrProductNotFound
		}
		return nil, err
	}

	if imgs.Valid && imgs.String != `[{"id":null,"name":null,"ext":null}]` {
		json.Unmarshal([]byte(imgs.String), &images)
	}

	return images, nil
}

// AddImage is ...
func (q *ProductQueries) AddImage(productID, fileUUID, fileExt, origName string) (*models.File, error) {
	file := &models.File{
		ID:   security.RandomString(),
		Name: fileUUID,
		Ext:  fileExt,
	}

	// add db record
	_, err := q.DB.Exec(`INSERT INTO product_image (id, product_id, name, ext, orig_name) VALUES (?, ?, ?, ?, ?)`, file.ID, productID, file.Name, file.Ext, origName)
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

// ProductDigitals is ...
func (q *ProductQueries) ProductDigital(productID string) (*models.Digital, error) {
	digital := &models.Digital{}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tx, err := q.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil || err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	// digital type
	var digitalType sql.NullString
	err = q.DB.QueryRow(`SELECT digital FROM product WHERE id = ?`, productID).Scan(&digitalType)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrProductNotFound
		}
		return nil, err
	}

	if digitalType.Valid {
		digital.Type = digitalType.String
	} else {
		return nil, nil
	}

	// digital file
	rows, err := q.DB.Query(`SELECT id, name, ext FROM digital_file WHERE product_id = ?`, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		file := models.File{}

		err := rows.Scan(
			&file.ID,
			&file.Name,
			&file.Ext,
		)
		if err != nil {
			return nil, err
		}

		digital.Files = append(digital.Files, file)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// digital data
	rows, err = q.DB.Query(`SELECT id, content, active FROM digital_data WHERE product_id = ?`, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		data := models.Data{}

		err := rows.Scan(
			&data.ID,
			&data.Content,
			&data.Active,
		)
		if err != nil {
			return nil, err
		}

		digital.Data = append(digital.Data, data)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return digital, nil
}

// AddDigitalFile is ...
func (q *ProductQueries) AddDigitalFile(productID, fileUUID, fileExt, origName string) (*models.File, error) {
	file := &models.File{
		ID:   security.RandomString(),
		Name: fileUUID,
		Ext:  fileExt,
	}

	// add db record
	_, err := q.DB.Exec(`INSERT INTO digital_file (id, product_id, name, ext, orig_name) VALUES (?, ?, ?, ?, ?)`, file.ID, productID, file.Name, file.Ext, origName)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// AddDigitalData
func (q *ProductQueries) AddDigitalData(productID, content string) (*models.Data, error) {
	file := &models.Data{
		ID:      security.RandomString(),
		Content: content,
		Active:  true,
	}

	// add db record
	_, err := q.DB.Exec(`INSERT INTO digital_data (id, product_id, content) VALUES (?, ?, ?)`, file.ID, productID, file.Content)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// UpdateDigital is ...
func (q *ProductQueries) UpdateDigital(digital *models.Data) error {
	_, err := q.DB.Exec(`UPDATE digital_data SET content = ?, active = ? WHERE id = ?`,
		digital.Content,
		digital.Active,
		digital.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

// DeleteDigital is ...
func (q *ProductQueries) DeleteDigital(productID, digitalID string) error {
	var digitalType string
	err := q.DB.QueryRow(`SELECT digital FROM product WHERE id = ?`, productID).Scan(&digitalType)
	if err != nil {
		return err
	}

	switch digitalType {
	case "file":
		var name, ext string
		err := q.DB.QueryRow(`SELECT name, ext FROM digital_file WHERE id = ?`, digitalID).Scan(&name, &ext)
		if err != nil {
			return err
		}

		if _, err := q.DB.Exec(`DELETE FROM digital_file WHERE id = ? AND product_id = ?`, digitalID, productID); err != nil {
			return err
		}

		filePaths := []string{
			fmt.Sprintf("./lc_digitals/%s.%s", name, ext),
		}

		for _, filePath := range filePaths {
			if err := os.Remove(filePath); err != nil {
				return fmt.Errorf("failed to remove file %s: %w", filePath, err)
			}
		}

	case "data":
		if _, err := q.DB.Exec(`DELETE FROM digital_data WHERE id = ? AND product_id = ?`, digitalID, productID); err != nil {
			return err
		}
	}

	return nil
}
