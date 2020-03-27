package interactor

import "github.com/cgressang/go-api-skeleton/usecase/interactor"

const message = "You are Home!"

type home struct {
}

func NewHomeInteractor() interactor.Home {
	return &home{}
}

func (hu *home) GetMessage() string {
	return message
}
