package bitbucketService

import "unified-atlassian-platform/db"

// Repository represents a Bitbucket repository model.
type Repository struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

// GetRepositories retrieves all repositories from the database.
func GetRepositories() ([]Repository, error) {
	var repos []Repository
	query := "SELECT id, name, url FROM bitbucket_repos"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var repo Repository
		if err := rows.Scan(&repo.ID, &repo.Name, &repo.URL); err != nil {
			return nil, err
		}
		repos = append(repos, repo)
	}
	return repos, nil
}

// CreateRepository inserts a new repository into the database.
func CreateRepository(repo Repository) error {
	query := "INSERT INTO bitbucket_repos (name, url) VALUES (?, ?)"
	_, err := db.DB.Exec(query, repo.Name, repo.URL)
	return err
}
