package dto

import (
	"errors"
	"strconv"
	"time"

	"github.com/onemgvv/WB_L2/develop/dev11/internal/domain"
)

const (
	LayoutISO = "2006-01-02"
)

type CreateEventDTO struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserId      int       `json:"userId"`
	Date        time.Time `json:"date"`
}

func (u *CreateEventDTO) ToModel() *domain.Event {
	return domain.NewEvent(u.Title, u.Description, u.Date, u.UserId)
}

func ParseCreateEvDto(title, userId, description, date string) (*CreateEventDTO, error) {
	var cTime time.Time
	uId, err := strconv.Atoi(userId)
	if err != nil {
		return nil, errors.New("Id is not valid")
	}

	cTime, _ = time.Parse(LayoutISO, date)

	if title == "" || description == "" {
		return nil, errors.New("not enough parameters")
	}
	
	return &CreateEventDTO{
		Title: title,
		Description: description,
		UserId: uId,
		Date: cTime,
	}, nil
}