// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/opctl/opctl/cli/internal/cliparamsatisfier/inputsrc"
)

type FakeInputSrc struct {
	ReadStringStub        func(string) (*string, bool)
	readStringMutex       sync.RWMutex
	readStringArgsForCall []struct {
		arg1 string
	}
	readStringReturns struct {
		result1 *string
		result2 bool
	}
	readStringReturnsOnCall map[int]struct {
		result1 *string
		result2 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInputSrc) ReadString(arg1 string) (*string, bool) {
	fake.readStringMutex.Lock()
	ret, specificReturn := fake.readStringReturnsOnCall[len(fake.readStringArgsForCall)]
	fake.readStringArgsForCall = append(fake.readStringArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ReadString", []interface{}{arg1})
	fake.readStringMutex.Unlock()
	if fake.ReadStringStub != nil {
		return fake.ReadStringStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.readStringReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInputSrc) ReadStringCallCount() int {
	fake.readStringMutex.RLock()
	defer fake.readStringMutex.RUnlock()
	return len(fake.readStringArgsForCall)
}

func (fake *FakeInputSrc) ReadStringCalls(stub func(string) (*string, bool)) {
	fake.readStringMutex.Lock()
	defer fake.readStringMutex.Unlock()
	fake.ReadStringStub = stub
}

func (fake *FakeInputSrc) ReadStringArgsForCall(i int) string {
	fake.readStringMutex.RLock()
	defer fake.readStringMutex.RUnlock()
	argsForCall := fake.readStringArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeInputSrc) ReadStringReturns(result1 *string, result2 bool) {
	fake.readStringMutex.Lock()
	defer fake.readStringMutex.Unlock()
	fake.ReadStringStub = nil
	fake.readStringReturns = struct {
		result1 *string
		result2 bool
	}{result1, result2}
}

func (fake *FakeInputSrc) ReadStringReturnsOnCall(i int, result1 *string, result2 bool) {
	fake.readStringMutex.Lock()
	defer fake.readStringMutex.Unlock()
	fake.ReadStringStub = nil
	if fake.readStringReturnsOnCall == nil {
		fake.readStringReturnsOnCall = make(map[int]struct {
			result1 *string
			result2 bool
		})
	}
	fake.readStringReturnsOnCall[i] = struct {
		result1 *string
		result2 bool
	}{result1, result2}
}

func (fake *FakeInputSrc) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.readStringMutex.RLock()
	defer fake.readStringMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeInputSrc) recordInvocation(key string, args []interface{}) {
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

var _ inputsrc.InputSrc = new(FakeInputSrc)
