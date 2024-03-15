package validator

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

type ProductGetParams struct {
	ID string `param:"id"`
}

func (p *ProductGetParams) Validate(ctx echo.Context) (bool, error) {
	err := ctx.Bind(p)
	if err != nil {
		return false, err
	}

	if len(p.ID) <= 5 {
		return false, fmt.Errorf("invalid 'Id' provided")
	}

	return true, nil
}

type ProductCreateParams struct {
	Name     string  `param:"name"`
	Slug     string  `param:"slug"`
	Category string  `param:"category"`
	Seller   string  `param:"seller"`
	Price    float64 `param:"price"`
}

func (p *ProductCreateParams) Validate(ctx echo.Context) (bool, error) {

	err := ctx.Bind(p)
	if err != nil {
		return false, err
	}

	if len(p.Name) <= 5 {
		return false, fmt.Errorf("invalid 'Name' provided")
	}

	if len(p.Slug) <= 3 {
		return false, fmt.Errorf("invalid 'Slug' provided")
	}

	if len(p.Category) <= 3 {
		return false, fmt.Errorf("invalid 'Category' provided")
	}

	if len(p.Seller) <= 3 {
		return false, fmt.Errorf("invalid 'Seller' provided")
	}

	if p.Price == 0 {
		return false, fmt.Errorf("invalid 'Price' provided")
	}

	return true, nil
}
