package controller

import (
	"fmt"
	"net/http"
	"produx/domain/entity"
	"produx/domain/product/dto"
	"produx/infrastructure/container"
	"produx/infrastructure/response"
	"produx/infrastructure/validator"

	"github.com/labstack/echo/v4"
	"github.com/mindstand/gogm/v2"
)

type Validator interface {
	Validate(context echo.Context) (bool, error)
}

type Product struct{}

func (a Product) Get(context echo.Context, c container.Container) error {
	params := validator.ProductGetParams{}
	isValid, err := params.Validate(context)
	if !isValid {
		return err
	}

	prod := c.GetProductService().Get(params.ID)
	if prod == nil {
		return fmt.Errorf("invalid product ID provided")
	}

	return context.JSON(http.StatusOK, dto.NewProductResponse(prod))
}

func (a Product) Create(context echo.Context, c container.Container) error {
	params := validator.ProductCreateParams{}
	isValid, err := params.Validate(context)
	if !isValid {
		return err
	}

	product := entity.Product{
		BaseUUIDNode: gogm.BaseUUIDNode{},
		Name:         params.Name,
		Slug:         params.Slug,
		Price:        params.Price,
	}

	// find and link category
	category := c.GetCategoryService().Get(params.Category)
	if category == nil {
		return fmt.Errorf("invalid category ID provided for the product %s", product.Name)
	}

	product.Category = category

	// find and link seller
	seller := c.GetSellerService().Get(params.Seller)
	if seller == nil {
		return fmt.Errorf("invalid seller ID provided for the product %s", product.Name)
	}
	product.Seller = seller

	result := c.GetProductService().Add(product)
	if result == nil {
		return fmt.Errorf("error adding product %s", product.Name)
	}

	return context.JSON(http.StatusOK, dto.NewProductResponse(result))
}

func (a Product) Update(context echo.Context, c container.Container) error {
	return context.JSON(http.StatusOK, dto.NewProductResponse(&entity.Product{}))
}

func (a Product) Delete(context echo.Context, c container.Container) error {
	return context.JSON(http.StatusOK, response.NewSuccessResponse("Product deleted."))
}

func (a Product) List(context echo.Context, c container.Container) error {
	products := c.GetProductService().List()

	return context.JSON(http.StatusOK, dto.NewProductListResponse(products))
}
