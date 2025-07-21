package database

import "github.com/andrewjamesmoore/andrew-projects-api/graph/model"

func ConvertProjectToGraphQL(db *DBProject) *model.Project {
	return &model.Project{
		ID:          db.ID.Hex(),
		Title:       db.Title,
		Description: db.Description,
		Status:      db.Status,
		Link:        db.Link,
		URL:         db.URL,
		Giturl:      db.Giturl,
		Tags:        db.Tags,
	}
}

func ConvertNoteToGraphQL(db *DBNote) *model.Note {
	return &model.Note{
		ID:          db.ID.Hex(),
		Title:       db.Title,
		Description: db.Description,
		URL:         db.URL,
		Date:		 db.Date,
	}
}
