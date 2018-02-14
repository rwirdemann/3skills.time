package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Project struct {
	Id   int
	Name string
}

type Activity struct {
	Id          int
	ProjectId   int
	Name        string
	Description string
	TimeSpent   float32
	CreatedAt   time.Time
}

type ProjectRepository interface {
	Add(p Project)
	All() []Project
}

type ActivityRepository interface {
	Add(projectId int, p Activity)
	All(projectId int) []Activity
}

type DefaultProjectRepository struct {
	projects map[int]Project
}

func NewDefaultProjectRepository() *DefaultProjectRepository {
	return &DefaultProjectRepository{projects: make(map[int]Project)}
}

func (this *DefaultProjectRepository) Add(p Project) {
	this.projects[p.Id] = p
}

func (this *DefaultProjectRepository) All() []Project {
	projects := make([]Project, 0)
	for _, p := range this.projects {
		projects = append(projects, p)
	}
	return projects
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, _ := ioutil.ReadAll(r.Body)
		var p Project
		json.Unmarshal(body, &p)
		projectRepository.Add(p)
		w.Header().Set("Location", fmt.Sprintf("%s/%d", r.URL.String(), p.Id))
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	if r.Method == "GET" {
		b, _ := json.Marshal(projectRepository.All())
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}
}

var projectRepository ProjectRepository

func main() {
	projectRepository = NewDefaultProjectRepository()
	http.HandleFunc("/projects", projectsHandler)
	http.ListenAndServe(":8080", nil)
}
