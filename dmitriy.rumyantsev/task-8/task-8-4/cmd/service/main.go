package main

import (
	"fmt"

	"github.com/dmitriy.rumyantsev/task-8/task-8-4/internal/mocks"

	"go.uber.org/mock/gomock"
)

//go:generate go generate ../../internal/user/user.go

func main() {
	ctrl := gomock.NewController(nil)
    defer ctrl.Finish()

    mock := mocks.NewMockUser(ctrl)
    mock.EXPECT().GetUser(1).Return("Dima", nil)

    name, _ := mock.GetUser(1)
    fmt.Println("User:", name)
}
