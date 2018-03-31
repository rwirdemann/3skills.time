package usecase

type IdentityPresenter struct {
}

func NewIdentityPresenter() IdentityPresenter {
	return IdentityPresenter{}
}

func (p IdentityPresenter) Present(i interface{}) interface{} {
	return i
}
