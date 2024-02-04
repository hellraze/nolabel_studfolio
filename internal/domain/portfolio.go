package domain

import (
	"github.com/gofrs/uuid"
	"time"
)

type Portfolio struct {
	portfolioID  uuid.UUID
	name         string
	about        string
	createdAt    time.Time
	urls         []string
	technologies []string
	views        int
	imageID      int
	tagIDs       []int
	authorID     uuid.UUID
}

func (p *Portfolio) PortfolioID() uuid.UUID { return p.portfolioID }
func (p *Portfolio) Name() string           { return p.name }
func (p *Portfolio) About() string          { return p.about }
func (p *Portfolio) CreatedAt() time.Time   { return p.createdAt }
func (p *Portfolio) Urls() []string         { return p.urls }
func (p *Portfolio) Technologies() []string { return p.technologies }
func (p *Portfolio) Views() int             { return p.views }
func (p *Portfolio) ImageID() int           { return p.imageID }
func (p *Portfolio) TagIDs() []int          { return p.tagIDs }
func (p *Portfolio) AuthorID() uuid.UUID    { return p.authorID }

func (p *Portfolio) SetName(name string)              { p.name = name }
func (p *Portfolio) SetAbout(about string)            { p.about = about }
func (p *Portfolio) SetCreatedAt(createdAt time.Time) { p.createdAt = createdAt }
func (p *Portfolio) SetImageID(imageID int)           { p.imageID = imageID }
func (p *Portfolio) AddToUrls(url string)             { p.urls = append(p.urls, url) }
func (p *Portfolio) AddToTechnologies(technology string) {
	p.technologies = append(p.technologies, technology)
}
func (p *Portfolio) IncreaseViews()                 { p.views++ }
func (p *Portfolio) AddToTagsIDs(tagID int)         { p.tagIDs = append(p.tagIDs, tagID) }
func (p *Portfolio) SetAuthorID(authorID uuid.UUID) { p.authorID = authorID }

func NewPortfolio(portfolioID uuid.UUID, name string, about string, createdAt time.Time, urls []string, technologies []string, imageID int, tagsIDs []int, authorID uuid.UUID) *Portfolio {
	return &Portfolio{
		portfolioID:  portfolioID,
		name:         name,
		about:        about,
		createdAt:    createdAt,
		urls:         urls,
		technologies: technologies,
		imageID:      imageID,
		tagIDs:       tagsIDs,
		authorID:     authorID,
	}
}

type PortfolioRepository interface {
}
