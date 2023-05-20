package service

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewUserService,
	wire.Bind(new(UserService), new(*userService)),
)
