package usecase

import (
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/rwirdemann/3skills.time/usecase/mocks"

	"github.com/rwirdemann/3skills.time/domain"
)

func TestAddProjectWithFakeRepository(t *testing.T) {
	repository := NewFakeRepository()
	consumer := NewIdentityConsumer()
	addProject := NewAddProject(consumer, repository)

	p := domain.Project{Name: "Test Projekt"}
	addProject.Run(&p)

	if !repository.contains("Test Projekt") {
		t.Errorf("Expect project %s was not added", p.Name)
	}
}

func TestAddProjectWithMockRepository(t *testing.T) {
	repository := mocks.Repository{}
	consumer := NewIdentityConsumer()
	addProject := NewAddProject(consumer, &repository)

	repository.On("AddProject", mock.AnythingOfType("domain.Project")).Return(1)
	p := domain.Project{Name: "Test Projekt"}
	addProject.Run(&p)

	repository.AssertCalled(t, "AddProject", mock.AnythingOfType("domain.Project"))
}
