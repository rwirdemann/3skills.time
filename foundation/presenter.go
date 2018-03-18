package foundation

type Presenter interface {
	Present(i interface{}) interface{}
}
