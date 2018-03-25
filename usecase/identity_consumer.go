package usecase

type IdentityConsumer struct {
}

func NewIdentityConsumer() IdentityConsumer {
	return IdentityConsumer{}
}

func (this IdentityConsumer) Consume(i interface{}) interface{} {
	return i
}
