package view

type ViewMessage interface {
	ForwardTo() ViewName
}
