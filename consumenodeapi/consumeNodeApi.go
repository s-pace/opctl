package consumenodeapi

//go:generate counterfeiter -o ./fake.go --fake-name Fake ./ ConsumeNodeApi

import (
	"github.com/opspec-io/sdk-golang/model"
	"github.com/opspec-io/sdk-golang/util/http"
	"github.com/sethgrid/pester"
)

type ConsumeNodeApi interface {
	GetEventStream(
		req *model.GetEventStreamReq,
	) (
		stream chan model.Event,
		err error,
	)

	KillOp(
		req model.KillOpReq,
	) (
		err error,
	)

	StartOp(
		req model.StartOpReq,
	) (
		opId string,
		err error,
	)
}

func New() ConsumeNodeApi {

	httpClient := pester.New()
	httpClient.Backoff = pester.ExponentialBackoff

	return &consumeNodeApi{
		httpClient: httpClient,
	}
}

type consumeNodeApi struct {
	httpClient http.Client
}
