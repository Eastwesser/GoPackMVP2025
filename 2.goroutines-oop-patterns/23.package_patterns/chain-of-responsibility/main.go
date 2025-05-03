package main

import "fmt"

// Handler интерфейс для обработчиков.
type Handler interface {
	SetNext(handler Handler)
	Handle(request string)
}

// BaseHandler базовая структура для обработчиков.
type BaseHandler struct {
	next Handler
}

// SetNext устанавливает следующий обработчик в цепочке.
func (b *BaseHandler) SetNext(handler Handler) {
	b.next = handler
}

// Handle передает запрос следующему обработчику, если он существует.
func (b *BaseHandler) Handle(request string) {
	if b.next != nil {
		b.next.Handle(request)
	}
}

// ConcreteHandlerA конкретный обработчик A.
type ConcreteHandlerA struct {
	BaseHandler
}

// Handle обрабатывает запрос, если может, или передает следующему обработчику.
func (c *ConcreteHandlerA) Handle(request string) {
	if request == "A" {
		fmt.Println("ConcreteHandlerA обработал запрос")
	} else {
		c.BaseHandler.Handle(request)
	}
}

// ConcreteHandlerB конкретный обработчик B.
type ConcreteHandlerB struct {
	BaseHandler
}

// Handle обрабатывает запрос, если может, или передает следующему обработчику.
func (c *ConcreteHandlerB) Handle(request string) {
	if request == "B" {
		fmt.Println("ConcreteHandlerB обработал запрос")
	} else {
		c.BaseHandler.Handle(request)
	}
}

func main() {
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}

	handlerA.SetNext(handlerB)

	handlerA.Handle("A") // Обработает ConcreteHandlerA
	handlerA.Handle("B") // Обработает ConcreteHandlerB
	handlerA.Handle("C") // Никто не обработает
}
