package ops

//go:generate counterfeiter -o ./fakeHandler.go --fake-name FakeHandler ./ Handler

import (
	"net/http"

	"github.com/opctl/opctl/sdks/go/node/api/handler/ops/kills"
	"github.com/opctl/opctl/sdks/go/node/api/handler/ops/starts"
	"github.com/opctl/opctl/sdks/go/node/core"
	"github.com/opctl/opctl/sdks/go/util/urlpath"
)

type Handler interface {
	Handle(
		httpResp http.ResponseWriter,
		httpReq *http.Request,
	)
}

// NewHandler returns an initialized Handler instance
func NewHandler(
	core core.Core,
) Handler {
	return _handler{
		startsHandler: starts.NewHandler(core),
		killsHandler:  kills.NewHandler(core),
	}
}

type _handler struct {
	startsHandler starts.Handler
	killsHandler  kills.Handler
}

func (hdlr _handler) Handle(
	httpResp http.ResponseWriter,
	httpReq *http.Request,
) {
	pathSegment, err := urlpath.NextSegment(httpReq.URL)
	if nil != err {
		http.Error(httpResp, err.Error(), http.StatusBadRequest)
		return
	}

	switch pathSegment {
	case "kills":
		hdlr.killsHandler.Handle(
			httpResp,
			httpReq,
		)
	case "starts":
		hdlr.startsHandler.Handle(
			httpResp,
			httpReq,
		)
	default:
		http.NotFoundHandler().ServeHTTP(httpResp, httpReq)
		return
	}
}
