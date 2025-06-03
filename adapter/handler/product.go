package handler

import (
	"strconv"

	"github.com/WatkornFeng/go-hexa/core/domain"
	"github.com/WatkornFeng/go-hexa/core/port"
	"github.com/gofiber/fiber/v2"
)

type productHandlder struct {
	productSrv port.ProductService
}

func NewProductHandler(productSrv port.ProductService) *productHandlder {
	return &productHandlder{productSrv: productSrv}
}

type productRequest struct {
	Name  string  `json:"name" validate:"required,min=3,max=20,alphanum"`
	Price float64 `json:"price" validate:"required,gt=0,lte=5000"`
}

func (h *productHandlder) CreateProduct(c *fiber.Ctx) error {
	var req productRequest

	idParam := c.Params("userId")
	idUint64, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return parameterError(c)
	}

	if err := c.BodyParser(&req); err != nil {
		return bodyParseError(c)
	}

	if err := validate.Struct(req); err != nil {
		return validationError(c, err)
	}
	ctx := c.UserContext()
	newProductReq := domain.Product{
		Name:   req.Name,
		Price:  req.Price,
		UserID: uint(idUint64),
	}
	product, err := h.productSrv.CreateNewProduct(ctx, &newProductReq)
	if err != nil {
		return handleError(c, err)
	}
	return handleSuccess(c, "Create new product success", product)

}
