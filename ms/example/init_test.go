package example

import (
	"testing"

	"github.com/powerman/check"

	"github.com/powerman/go-monolith-example/internal/apix"
	"github.com/powerman/go-monolith-example/internal/dom"
	"github.com/powerman/go-monolith-example/ms/example/internal/app"
	"github.com/powerman/go-monolith-example/ms/example/internal/config"
	"github.com/powerman/go-monolith-example/ms/example/internal/dal"
	"github.com/powerman/go-monolith-example/ms/example/internal/srv/jsonrpc2"
	"github.com/powerman/go-monolith-example/pkg/def"
)

func TestMain(m *testing.M) {
	def.Init()
	dal.InitMetrics(reg)
	app.InitMetrics(reg)
	jsonrpc2.InitMetrics(reg)
	cfg = config.MustGetServeTest()
	check.TestMain(m)
}

type tLogger check.C

func (l tLogger) Print(v ...interface{}) { l.Log(v...) }

// Const shared by tests. Recommended naming scheme: <dataType><Variant>.
var (
	cfg        *config.ServeConfig
	ctx        = def.NewContext(app.ServiceName)
	tokenAdmin = apix.AccessToken("admin")
	tokenUser  = apix.AccessToken("user")
	authAdmin  = dom.Auth{
		UserID: 1,
		Admin:  true,
	}
	authUser = dom.Auth{
		UserID: 2,
		Admin:  false,
	}
)