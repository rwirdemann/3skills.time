package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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
	Add(p Project) int
	All() []Project
	Get(id int) Project
	Delete(id int)
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

func (this *DefaultProjectRepository) Add(p Project) int {
	p.Id = this.nextId()
	this.projects[p.Id] = p
	return p.Id
}

func (this *DefaultProjectRepository) All() []Project {
	projects := make([]Project, 0)
	for _, p := range this.projects {
		projects = append(projects, p)
	}
	return projects
}

func (this *DefaultProjectRepository) nextId() int {
	id := 1
	for _, p := range this.projects {
		if p.Id >= id {
			id = p.Id + 1
		}
	}
	return id
}

func (this *DefaultProjectRepository) Get(id int) Project {
	return this.projects[id]
}

func (this *DefaultProjectRepository) Delete(id int) {
	delete(this.projects, id)
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, _ := ioutil.ReadAll(r.Body)
		var p Project
		json.Unmarshal(body, &p)
		id := projectRepository.Add(p)
		w.Header().Set("Location", fmt.Sprintf("%s/%d", r.URL.String(), id))
		w.WriteHeader(http.StatusCreated)
	} else if r.Method == "GET" {
		b, _ := json.Marshal(projectRepository.All())
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func singleProjectsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		idFromUrl := r.URL.Path[len("/projects/"):]
		projectId, _ := strconv.Atoi(idFromUrl)
		projectRepository.Delete(projectId)
		w.WriteHeader(http.StatusNoContent)
	}
}

var projectRepository ProjectRepository

func main() {
	projectRepository = NewDefaultProjectRepository()
	http.HandleFunc("/projects", projectsHandler)
	http.HandleFunc("/projects/", singleProjectsHandler)
	http.ListenAndServe(":8080", nil)
}
