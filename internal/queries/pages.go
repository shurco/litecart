package queries

import (
	"database/sql"

	"github.com/shurco/litecart/internal/models"
	"github.com/shurco/litecart/pkg/errors"
	"github.com/shurco/litecart/pkg/security"
)

// PageQueries is ...
type PageQueries struct {
	*sql.DB
}

// ListPages is ...
func (q *PageQueries) ListPages(private bool, idList ...string) ([]models.Page, error) {
	pages := []models.Page{}

	queryPrivate := ` WHERE active = 1`
	query := `SELECT id, name, slug, position, active, strftime('%s', created), strftime('%s', updated) FROM page`

	if !private {
		query = query + queryPrivate
	}

	rows, err := q.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var updated sql.NullInt64

		page := models.Page{}
		err := rows.Scan(&page.ID, &page.Name, &page.Slug, &page.Position, &page.Active, &page.Created, &updated)
		if err != nil {
			return nil, err
		}

		if updated.Valid {
			page.Updated = updated.Int64
		}

		pages = append(pages, page)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return pages, nil
}

// Page is ...
func (q *PageQueries) Page(slug string) (*models.Page, error) {
	page := models.Page{
		Slug: slug,
	}

	err := q.DB.QueryRow(`SELECT id, name, content FROM page WHERE slug = ?`, slug).Scan(&page.ID, &page.Name, &page.Content)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrPageNotFound
		}
		return nil, err
	}

	return &page, nil
}

// AddPage is ...
func (q *PageQueries) AddPage(page *models.Page) (*models.Page, error) {
	page.ID = security.RandomString()
	page.Active = false

	sql := `INSERT INTO page (id, name, slug, position) VALUES (?, ?, ?, ?) RETURNING strftime('%s', created)`
	err := q.DB.QueryRow(sql, page.ID, page.Name, page.Slug, page.Position).Scan(&page.Created)
	if err != nil {
		return nil, err
	}

	return page, nil
}

// UpdatePage is ...
func (q *PageQueries) UpdatePage(page *models.Page) error {
	_, err := q.DB.Exec(`UPDATE page SET name = ?, slug = ?, position = ?, updated = datetime('now') WHERE id = ?`,
		page.Name,
		page.Slug,
		page.Position,
		page.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

// DeletePage is ...
func (q *PageQueries) DeletePage(id string) error {
	if _, err := q.DB.Exec(`DELETE FROM page WHERE id = ?`, id); err != nil {
		return err
	}

	return nil
}

// UpdatePageContent is ...
func (q *PageQueries) UpdatePageContent(page *models.Page) error {
	_, err := q.DB.Exec(`UPDATE page SET content = ?, updated = datetime('now') WHERE id = ? `, page.Content, page.ID)
	if err != nil {
		return err
	}
	return nil
}

// UpdatePageActive is ...
func (q *ProductQueries) UpdatePageActive(id string) error {
	var active bool
	query := `SELECT active FROM page WHERE id = ?`
	err := q.DB.QueryRow(query, id).Scan(&active)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if _, err := q.DB.Exec(`UPDATE page SET active = ?, updated = datetime('now') WHERE id = ?`, !active, id); err != nil {
		return err
	}

	return nil
}
