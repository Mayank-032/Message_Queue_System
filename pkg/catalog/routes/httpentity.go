package routes

import (
	"errors"
)

type CreateProductReq struct {
	UserId      int      `json:"userId"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Images      []string `json:"images"`
	Price       int      `json:"price"`
}

func (req CreateProductReq) Validate() error {
	if req.UserId == 0 || len(req.Name) == 0 || req.Price == 0 {
		return errors.New("mandatory details missing")
	}
	return nil
}
