package mysql

import (
	"apiserver/internal/repo"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewUserRepo,
	wire.Bind(new(repo.UserRepo), new(*userRepo)),
)
