// Code generated by counterfeiter. DO NOT EDIT.
package containerruntime

import (
	"context"
	"io"
	"sync"

	"github.com/opctl/opctl/sdks/go/model"
	"github.com/opctl/opctl/sdks/go/util/pubsub"
)

type Fake struct {
	DeleteContainerIfExistsStub        func(containerID string) error
	deleteContainerIfExistsMutex       sync.RWMutex
	deleteContainerIfExistsArgsForCall []struct {
		containerID string
	}
	deleteContainerIfExistsReturns struct {
		result1 error
	}
	deleteContainerIfExistsReturnsOnCall map[int]struct {
		result1 error
	}
	RunContainerStub        func(ctx context.Context, req *model.DCGContainerCall, eventPublisher pubsub.EventPublisher, stdout io.WriteCloser, stderr io.WriteCloser) (*int64, error)
	runContainerMutex       sync.RWMutex
	runContainerArgsForCall []struct {
		ctx            context.Context
		req            *model.DCGContainerCall
		eventPublisher pubsub.EventPublisher
		stdout         io.WriteCloser
		stderr         io.WriteCloser
	}
	runContainerReturns struct {
		result1 *int64
		result2 error
	}
	runContainerReturnsOnCall map[int]struct {
		result1 *int64
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Fake) DeleteContainerIfExists(containerID string) error {
	fake.deleteContainerIfExistsMutex.Lock()
	ret, specificReturn := fake.deleteContainerIfExistsReturnsOnCall[len(fake.deleteContainerIfExistsArgsForCall)]
	fake.deleteContainerIfExistsArgsForCall = append(fake.deleteContainerIfExistsArgsForCall, struct {
		containerID string
	}{containerID})
	fake.recordInvocation("DeleteContainerIfExists", []interface{}{containerID})
	fake.deleteContainerIfExistsMutex.Unlock()
	if fake.DeleteContainerIfExistsStub != nil {
		return fake.DeleteContainerIfExistsStub(containerID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteContainerIfExistsReturns.result1
}

func (fake *Fake) DeleteContainerIfExistsCallCount() int {
	fake.deleteContainerIfExistsMutex.RLock()
	defer fake.deleteContainerIfExistsMutex.RUnlock()
	return len(fake.deleteContainerIfExistsArgsForCall)
}

func (fake *Fake) DeleteContainerIfExistsArgsForCall(i int) string {
	fake.deleteContainerIfExistsMutex.RLock()
	defer fake.deleteContainerIfExistsMutex.RUnlock()
	return fake.deleteContainerIfExistsArgsForCall[i].containerID
}

func (fake *Fake) DeleteContainerIfExistsReturns(result1 error) {
	fake.DeleteContainerIfExistsStub = nil
	fake.deleteContainerIfExistsReturns = struct {
		result1 error
	}{result1}
}

func (fake *Fake) DeleteContainerIfExistsReturnsOnCall(i int, result1 error) {
	fake.DeleteContainerIfExistsStub = nil
	if fake.deleteContainerIfExistsReturnsOnCall == nil {
		fake.deleteContainerIfExistsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteContainerIfExistsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Fake) RunContainer(ctx context.Context, req *model.DCGContainerCall, eventPublisher pubsub.EventPublisher, stdout io.WriteCloser, stderr io.WriteCloser) (*int64, error) {
	fake.runContainerMutex.Lock()
	ret, specificReturn := fake.runContainerReturnsOnCall[len(fake.runContainerArgsForCall)]
	fake.runContainerArgsForCall = append(fake.runContainerArgsForCall, struct {
		ctx            context.Context
		req            *model.DCGContainerCall
		eventPublisher pubsub.EventPublisher
		stdout         io.WriteCloser
		stderr         io.WriteCloser
	}{ctx, req, eventPublisher, stdout, stderr})
	fake.recordInvocation("RunContainer", []interface{}{ctx, req, eventPublisher, stdout, stderr})
	fake.runContainerMutex.Unlock()
	if fake.RunContainerStub != nil {
		return fake.RunContainerStub(ctx, req, eventPublisher, stdout, stderr)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.runContainerReturns.result1, fake.runContainerReturns.result2
}

func (fake *Fake) RunContainerCallCount() int {
	fake.runContainerMutex.RLock()
	defer fake.runContainerMutex.RUnlock()
	return len(fake.runContainerArgsForCall)
}

func (fake *Fake) RunContainerArgsForCall(i int) (context.Context, *model.DCGContainerCall, pubsub.EventPublisher, io.WriteCloser, io.WriteCloser) {
	fake.runContainerMutex.RLock()
	defer fake.runContainerMutex.RUnlock()
	return fake.runContainerArgsForCall[i].ctx, fake.runContainerArgsForCall[i].req, fake.runContainerArgsForCall[i].eventPublisher, fake.runContainerArgsForCall[i].stdout, fake.runContainerArgsForCall[i].stderr
}

func (fake *Fake) RunContainerReturns(result1 *int64, result2 error) {
	fake.RunContainerStub = nil
	fake.runContainerReturns = struct {
		result1 *int64
		result2 error
	}{result1, result2}
}

func (fake *Fake) RunContainerReturnsOnCall(i int, result1 *int64, result2 error) {
	fake.RunContainerStub = nil
	if fake.runContainerReturnsOnCall == nil {
		fake.runContainerReturnsOnCall = make(map[int]struct {
			result1 *int64
			result2 error
		})
	}
	fake.runContainerReturnsOnCall[i] = struct {
		result1 *int64
		result2 error
	}{result1, result2}
}

func (fake *Fake) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteContainerIfExistsMutex.RLock()
	defer fake.deleteContainerIfExistsMutex.RUnlock()
	fake.runContainerMutex.RLock()
	defer fake.runContainerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Fake) recordInvocation(key string, args []interface{}) {
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

var _ ContainerRuntime = new(Fake)
