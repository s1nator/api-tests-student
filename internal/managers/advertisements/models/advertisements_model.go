package models

import authModels "api-tests-template/internal/managers/auth/models"

type AdvertisementsList struct {
	HasNext    bool            `json:"has_next"`
	Items      []Advertisement `json:"items"`
	NextCursor string          `json:"next_cursor"`
	Total      int             `json:"total"`
}

type Advertisement struct {
	CreatedAt   string          `json:"created_at"`
	Description string          `json:"description"`
	Id          string          `json:"id"`
	Photos      []Photo         `json:"photos"`
	Price       float64         `json:"price"`
	Quantity    int             `json:"quantity"`
	Title       string          `json:"title"`
	UpdatedAt   string          `json:"updated_at"`
	User        authModels.User `json:"user"`
	UserId      string          `json:"user_id"`
}

type Photo struct {
	AdvertisementId string `json:"advertisement_id"`
	CreatedAt       string `json:"created_at"`
	Id              string `json:"id"`
	SortOrder       int    `json:"sort_order"`
	Url             string `json:"url"`
}
