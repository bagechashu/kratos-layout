package greeter

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewData,
	NewGreeterRepo,
	NewGreeterUsecase,
	NewGreeterService,
)
