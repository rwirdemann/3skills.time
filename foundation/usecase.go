package foundation

type Usecase interface {
	Run(i ...interface{}) interface{}
}
