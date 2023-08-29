package queries

import (
	"database/sql"
	"errors"

	"github.com/shurco/litecart/internal/models"
)

// PageQueries is ...
type PageQueries struct {
	*sql.DB
}

// ListPages is ...
func (q *PageQueries) ListPages(private bool, idList ...string) ([]models.Page, error) {
	pages := []models.Page{}

	query := `SELECT id, name, url FROM page WHERE content != ""`
	rows, err := q.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		page := models.Page{}
		err := rows.Scan(&page.ID, &page.Name, &page.Url)
		if err != nil {
			return nil, err
		}

		pages = append(pages, page)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return pages, nil
}

// Page is ...
func (q *PageQueries) Page(url string) (*models.Page, error) {
	page := models.Page{
		Url: url,
	}

	err := q.DB.QueryRow(`SELECT id, name, content FROM page WHERE url = ?`, url).Scan(&page.ID, &page.Name, &page.Content)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("page not found")
		}
		return nil, err
	}

	return &page, nil
}

// SystemQueries is ...
func (q *PageQueries) UpdatePage(page *models.Page) error {
	_, err := q.DB.Exec(`UPDATE page SET content = ?, updated = datetime('now') WHERE id = ? `, page.Content, page.ID)
	if err != nil {
		return err
	}
	return nil
}
