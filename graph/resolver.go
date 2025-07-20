package graph

import (
	"github.com/andrewjamesmoore/andrew-projects-api/database"
	"github.com/andrewjamesmoore/andrew-projects-api/graph/model"
)

type Resolver struct {
	DB *database.DB
}

func toGraphQLProject(p *model.DBProject) *model.Project {
	return &model.Project{
		ID:          p.ID.Hex(),
		Title:       p.Title,
		Description: p.Description,
		Status:      p.Status,
		Link:        p.Link,
		URL:         p.URL,
		Giturl:      p.Giturl,
		Tags:        p.Tags,
	}
}
