package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"bitbucket.org/rwirdemann/go-tracker/security/middleware"
	"github.com/joho/godotenv"

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

type ProjectRepository struct {
	projects map[int]Project
}

func NewProjectRepository() *ProjectRepository {
	return &ProjectRepository{projects: make(map[int]Project)}
}

func (this *ProjectRepository) Add(p Project) int {
	p.Id = this.nextId()
	this.projects[p.Id] = p
	return p.Id
}

func (this *ProjectRepository) All() []Project {
	projects := make([]Project, 0)
	for _, p := range this.projects {
		projects = append(projects, p)
	}
	return projects
}

func (this *ProjectRepository) nextId() int {
	id := 1
	for _, p := range this.projects {
		if p.Id >= id {
			id = p.Id + 1
		}
	}
	return id
}

func (this *ProjectRepository) Get(id int) Project {
	return this.projects[id]
}

func (this *ProjectRepository) Delete(id int) {
	delete(this.projects, id)
}

func NewActivityRepository() *ActivityRepository {
	return &ActivityRepository{activities: make(map[int][]Activity)}
}

type ActivityRepository struct {
	activities map[int][]Activity
}

func (this *ActivityRepository) Add(projectId int, a Activity) int {
	a.Id = this.nextId()
	a.ProjectId = projectId
	this.activities[projectId] = append(this.activities[projectId], a)
	return a.Id
}

func (this *ActivityRepository) nextId() int {
	id := 1
	for _, activities := range this.activities {
		for _, a := range activities {
			if a.Id > id {
				id = a.Id + 1
			}
		}
	}
	return id
}

func (this *ActivityRepository) All(projectId int) []Activity {
	return this.activities[projectId]
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
	if !basicAuth(r) {
		http.Error(w, "authorization failed", http.StatusUnauthorized)
		return
	}

	b, _ := json.Marshal(projectRepository.All())
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func deleteProjectHandler(w http.ResponseWriter, r *http.Request) {
	projectId, _ := strconv.Atoi(mux.Vars(r)["id"])
	projectRepository.Delete(projectId)
	w.WriteHeader(http.StatusNoContent)
}

func addActivityHandler(w http.ResponseWriter, r *http.Request) {
	projectId, _ := strconv.Atoi(mux.Vars(r)["id"])
	body, _ := ioutil.ReadAll(r.Body)
	var a Activity
	json.Unmarshal(body, &a)
	id := activityRepository.Add(projectId, a)
	w.Header().Set("Location", fmt.Sprintf("%s/%d", r.URL.String(), id))
	w.WriteHeader(http.StatusCreated)
}

func getActivitiesHandler(w http.ResponseWriter, r *http.Request) {
	projectId, _ := strconv.Atoi(mux.Vars(r)["id"])
	b, _ := json.Marshal(activityRepository.All(projectId))
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func basicAuth(r *http.Request) bool {
	if username, password, ok := r.BasicAuth(); ok {
		if username == os.Getenv("USERNAME") && password == os.Getenv("PASSWORD") {
			return true
		}
	}

	return false
}

var projectRepository *ProjectRepository
var activityRepository *ActivityRepository

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	projectRepository = NewProjectRepository()
	activityRepository = NewActivityRepository()
	r := mux.NewRouter()
	r.HandleFunc("/projects",
		middleware.BasicAuth(addProjectHandler)).Methods("POST")
	r.HandleFunc("/projects",
		middleware.BasicAuth(getProjectsHandler)).Methods("GET")
	r.HandleFunc("/projects/{id}", middleware.BasicAuth(deleteProjectHandler)).Methods("DELETE")
	r.HandleFunc("/projects/{id}/activities", middleware.BasicAuth(addActivityHandler)).Methods("POST")
	r.HandleFunc("/projects/{id}/activities", middleware.BasicAuth(getActivitiesHandler)).Methods("GET")
	http.ListenAndServe(":8080", r)
}
