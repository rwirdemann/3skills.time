package usecase

type Presenter interface {
	Present(i interface{}) interface{}
}
