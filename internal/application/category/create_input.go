package category

import (
	"github.com/vitormoschetta/go/internal/application/common"
	"github.com/vitormoschetta/go/internal/domain/category"
)

type CreateCategoryInput struct {
	Name string `json:"name"`
}

func (c *CreateCategoryInput) Validate() (output common.Output) {
	if c.Name == "" {
		output.Errors = append(output.Errors, "Name is required")
	}
	return
}

func (c *CreateCategoryInput) ToCategoryEntity() category.Category {
	return category.NewCategory(c.Name)
}