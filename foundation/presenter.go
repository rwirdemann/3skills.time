package adapter

type Presenter interface {
	Present(i interface{}) interface{}
}
