package confluenceService

import "unified-atlassian-platform/db"

// Page represents a Confluence page model.
type Page struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

// GetPages retrieves all pages from the database.
func GetPages() ([]Page, error) {
	var pages []Page
	query := "SELECT id, title, body FROM confluence_pages"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var page Page
		if err := rows.Scan(&page.ID, &page.Title, &page.Body); err != nil {
			return nil, err
		}
		pages = append(pages, page)
	}
	return pages, nil
}

// CreatePage inserts a new page into the database.
func CreatePage(page Page) error {
	query := "INSERT INTO confluence_pages (title, body) VALUES (?, ?)"
	_, err := db.DB.Exec(query, page.Title, page.Body)
	return err
}
