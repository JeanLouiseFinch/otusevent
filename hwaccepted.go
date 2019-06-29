package otusevent

import (
	"io"
	"log"

	"github.com/google/uuid"
)

type hwAccepted struct {
	id    uuid.UUID
	grade int
}
//NewHwAccepted возвращает новый объект проверенной работы
func NewHwAccepted(id uuid.UUID, grade int) *hwAccepted {
	return &hwAccepted{id: id, grade: grade}
}
//GetID возвращает ID
func (hwa *hwAccepted) GetID() uuid.UUID{
	return hwa.id
}
//Log функция логирования
func (hwa *hwAccepted) Log(w io.Writer) {
	logger :=  log.New(w,"",log.Ldate)
	logger.Printf("ID=%v\tGrade=%d",hwa.id,hwa.grade)
}