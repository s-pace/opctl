package loop

//go:generate counterfeiter -o ./fakeDeScoper.go --fake-name FakeDeScoper ./ DeScoper

import (
	"github.com/opctl/sdk-golang/model"
)

type DeScoper interface {
	// DeScope de-scopes loop vars (index, key, value)
	DeScope(
		parentScope map[string]*model.Value,
		scgLoopRange interface{},
		scgLoopVars *model.SCGLoopVars,
		iterationScope map[string]*model.Value,
	) map[string]*model.Value
}

func NewDeScoper() DeScoper {
	return _deScoper{}
}

type _deScoper struct{}

func (ds _deScoper) DeScope(
	parentScope map[string]*model.Value,
	scgLoopRange interface{},
	scgLoopVars *model.SCGLoopVars,
	iterationScope map[string]*model.Value,
) map[string]*model.Value {
	if nil == scgLoopVars {
		return parentScope
	}

	outboundScope := map[string]*model.Value{}
	for varName, varData := range iterationScope {
		outboundScope[varName] = varData
	}

	// restore vars shadowed in the loop
	if nil != scgLoopVars.Index {
		outboundScope[*scgLoopVars.Index] = parentScope[*scgLoopVars.Index]
	}

	if nil != scgLoopRange {
		return outboundScope
	}

	if nil != scgLoopVars.Key {
		outboundScope[*scgLoopVars.Key] = parentScope[*scgLoopVars.Key]
	}
	if nil != scgLoopVars.Value {
		outboundScope[*scgLoopVars.Value] = parentScope[*scgLoopVars.Value]
	}

	return outboundScope
}
