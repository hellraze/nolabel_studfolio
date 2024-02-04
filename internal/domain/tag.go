package domain

type Tag struct {
	tagID int
	name  string
}

func (t *Tag) ID() int      { return t.tagID }
func (t *Tag) Name() string { return t.name }

func (t *Tag) SetName(name string) { t.name = name }

func NewTag(name string) *Tag {
	return &Tag{
		name: name,
	}
}

type TagRepository interface {
}
