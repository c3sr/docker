package docker

import (
	"testing"

	"github.com/c3sr/config"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	config.Init(
		config.VerboseMode(true),
		config.DebugMode(true),
	)

	goleak.VerifyTestMain(m,
		goleak.IgnoreTopFunction("github.com/patrickmn/go-cache.(*janitor).Run"),
		goleak.IgnoreTopFunction("github.com/c3sr/docker/vendor/github.com/patrickmn/go-cache.(*janitor).Run"),
	)

}
