package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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

func addProjectHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var p Project
	json.Unmarshal(body, &p)
	id := projectRepository.Add(p)
	w.Header().Set("Location", fmt.Sprintf("%s/%d", r.URL.String(), id))
	w.WriteHeader(http.StatusCreated)
}

func getProjectsHandler(w http.ResponseWriter, r *http.Request) {
	b, _ := json.Marshal(projectRepository.All())
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func deleteProjectHandler(w http.ResponseWriter, r *http.Request) {
	projectId, _ := strconv.Atoi(mux.Vars(r)["id"])
	projectRepository.Delete(projectId)
	w.WriteHeader(http.StatusNoContent)
}

var projectRepository ProjectRepository

func main() {
	projectRepository = NewDefaultProjectRepository()
	r := mux.NewRouter()
	r.HandleFunc("/projects", addProjectHandler).Methods("POST")
	r.HandleFunc("/projects", getProjectsHandler).Methods("GET")
	r.HandleFunc("/projects/{id}", deleteProjectHandler).Methods("DELETE")
	http.ListenAndServe(":8080", r)
}
