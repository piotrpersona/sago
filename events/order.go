package events

type OrderHandler interface {
	Handle(payload string) (err error)
}
