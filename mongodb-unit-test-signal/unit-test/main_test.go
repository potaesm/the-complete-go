package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAdd(t *testing.T) {
	t.Run("Add func valid case", func(t *testing.T) {
		result := add(1, 9)
		assert.Equal(t, 10, result)
	})
}

func TestNewUser(t *testing.T) {
	t.Run("Valid case", func(t *testing.T) {
		u := newUser("Ahmed", 10)
		assert.NotNil(t, u)
	})
}

type mockProduct struct {
	mock.Mock
}

func (p *mockProduct) buy() error {
	args := p.Called()
	return args.Error(0)
}

func (p *mockProduct) sell() error {
	args := p.Called()
	return args.Error(0)
}

func TestBuyProduct(t *testing.T) {
	u := newUser("Suthinan", 10)
	assert.NotNil(t, u)

	t.Run("Test success buy", func(t *testing.T) {
		p := new(mockProduct)
		p.On("buy").Return(nil)
		result := u.buyProduct(p)
		assert.NoError(t, result)
	})

	t.Run("Test fail buy", func(t *testing.T) {
		p := new(mockProduct)
		e := errors.New("Error buy product")
		p.On("buy").Return(e)
		result := u.buyProduct(p)
		assert.Error(t, result)
		assert.EqualError(t, e, result.Error())
	})
}
