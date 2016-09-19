// This file was generated by counterfeiter
package opspec

import (
	"sync"

	"github.com/opspec-io/sdk-golang/models"
)

type fakeStartOpRunUseCase struct {
	ExecuteStub        func(req models.StartOpRunReq) (opRunId string, err error)
	executeMutex       sync.RWMutex
	executeArgsForCall []struct {
		req models.StartOpRunReq
	}
	executeReturns struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *fakeStartOpRunUseCase) Execute(req models.StartOpRunReq) (opRunId string, err error) {
	fake.executeMutex.Lock()
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
		req models.StartOpRunReq
	}{req})
	fake.recordInvocation("Execute", []interface{}{req})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		return fake.ExecuteStub(req)
	} else {
		return fake.executeReturns.result1, fake.executeReturns.result2
	}
}

func (fake *fakeStartOpRunUseCase) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *fakeStartOpRunUseCase) ExecuteArgsForCall(i int) models.StartOpRunReq {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return fake.executeArgsForCall[i].req
}

func (fake *fakeStartOpRunUseCase) ExecuteReturns(result1 string, result2 error) {
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *fakeStartOpRunUseCase) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return fake.invocations
}

func (fake *fakeStartOpRunUseCase) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
