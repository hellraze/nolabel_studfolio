package domain

type Image struct {
	imageID int
	name    string
	path    string
}

func (i *Image) ID() int      { return i.imageID }
func (i *Image) Name() string { return i.name }
func (i *Image) Path() string { return i.path }

func (i *Image) SetName(name string) { i.name = name }
func (i *Image) SetPath(path string) { i.path = path }

func NewImage(name string, path string) *Image {
	return &Image{
		name: name,
		path: path,
	}
}

type ImageRepository interface {
}
