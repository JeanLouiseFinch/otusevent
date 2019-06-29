package otusevent

import (
	"io"
	"log"

	"github.com/google/uuid"
)

type hwSubmitted struct {
	id      uuid.UUID
	code    string
	comment string
}
//NewHwSubmitted создает новый объект работы студента
func NewHwSubmitted(code, comment string) *hwSubmitted {
	return &hwSubmitted{id: uuid.New(), code: code, comment: comment}
}
//GetID возвращает ID
func (hwa *hwSubmitted) GetID() uuid.UUID{
	return hwa.id
}
//Log функция логирования
func (hws *hwSubmitted) Log(w io.Writer) {
	logger := log.New(w,"",log.Ldate)
	logger.Printf("ID=%v\tComment=%s",hws.id,hws.comment)

}