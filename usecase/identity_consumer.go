package usecase

type IdentityConsumer struct {
}

func NewIdentityConsumer() IdentityConsumer {
	return IdentityConsumer{}
}

func (c IdentityConsumer) Consume(i interface{}) interface{} {
	return i
}
