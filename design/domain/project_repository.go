package domain

type ProjectRepository interface {
	All() []Project
	Add(p Project)
}
