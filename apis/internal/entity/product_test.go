package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("product", 1000.0)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, "product", p.Name)
	assert.Equal(t, 1000.0, p.Price)
}

func TestProductWhenNameIsRequire(t *testing.T) {
	p, err := NewProduct("", 1000)
	assert.NotNil(t, err)
	assert.Nil(t, p)
	assert.Equal(t, ErrNameRequired, err)
}

func TestProductWhenPriceIsRequire(t *testing.T) {
	p, err := NewProduct("product", 0)
	assert.Nil(t, p)
	assert.Equal(t, ErrPriceRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	p, err := NewProduct("product", -1)
	assert.NotNil(t, err)
	assert.Nil(t, p)
	assert.Equal(t, ErrPriceInvalid, err)
}

func TestProductValidate(t *testing.T) {
	p, err := NewProduct("product", 1000)
	assert.Nil(t, err)
	assert.Nil(t, p.Validate())
	assert.NotNil(t, p)
}
