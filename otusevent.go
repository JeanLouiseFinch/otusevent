package otusevent

import (
	"io"

	"github.com/google/uuid"
)
// OtusEvent интерфейс. содержит метод логирования и получения id.
type OtusEvent interface {
	Log(w io.Writer)
	GetID() uuid.UUID
}
// LogOtusEvent вызывает метод логирования у объектов
func LogOtusEvent(e OtusEvent, w io.Writer) {
	e.Log(w)
}