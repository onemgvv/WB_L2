package dto

import (
	"errors"
	"strconv"
	"time"

	"github.com/onemgvv/WB_L2/develop/dev11/internal/domain"
)

type UpdateEventDTO struct {
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	UserId      int       `json:"userId,omitempty"`
	Date        time.Time `json:"date,omitempty"`
}

func (u *UpdateEventDTO) ToModel() *domain.Event {
	return domain.NewEvent(u.Title, u.Description, u.Date, u.UserId)
}

func ParseUpdateEvDto(title, userId, description, date string) (*UpdateEventDTO, error) {
	uId, err := strconv.Atoi(userId)
	if err != nil {
		return nil, errors.New("Id is not valid")
	}

	cTime, err := time.Parse(LayoutISO, date)
	if date != "" && err != nil {
		return nil, errors.New("Date is not valid")
	}
	
	return &UpdateEventDTO{
		Title: title,
		Description: description,
		UserId: uId,
		Date: cTime,
	}, nil
}