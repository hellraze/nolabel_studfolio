package domain

type Status struct {
	statusID int
	name     string
	path     string
}

func (s *Status) ID() int      { return s.statusID }
func (s *Status) Name() string { return s.name }
func (s *Status) Path() string { return s.path }

func (s *Status) SetName(name string) { s.name = name }
func (s *Status) SetPath(path string) { s.path = path }

func NewStatus(name string, path string) *Status {
	return &Status{
		name: name,
		path: path,
	}
}

type StatusRepository interface {
}
