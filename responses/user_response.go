package responses

import "github.com/gofiber/fiber/v2"

type UserResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    *fiber.Map `json:"data"`
}

/*
json:"status", json:"message", and json:"data" are known as struct tags. Struct tags allow us to attach meta-information to corresponding struct properties. In other words, we use them to reformat the JSON response returned by the API.
*/
