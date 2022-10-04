package dto

import (
	"log"
	"strconv"
	"time"

	"github.com/onemgvv/WB_L2/develop/dev11/internal/domain"
)

type CreateEventDTO struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserId      int       `json:"userId"`
	Date        time.Time `json:"date"`
}

func (u *CreateEventDTO) ToModel() *domain.Event {
	return domain.NewEvent(u.Title, u.Description, u.UserId)
}

func ParseCreateEvDto(title, userId, description, date string) *CreateEventDTO {
	uId, err := strconv.Atoi(userId)
	if err != nil {
		log.Printf("%s.\n", err.Error())
		return nil
	}

	cTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Printf("%s.\n", err.Error())
		return nil
	}
	
	return &CreateEventDTO{
		Title: title,
		Description: description,
		UserId: uId,
		Date: cTime,
	}
}