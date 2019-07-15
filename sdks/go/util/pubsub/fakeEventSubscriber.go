// Code generated by counterfeiter. DO NOT EDIT.
package pubsub

import (
	"context"
	"sync"

	"github.com/opctl/opctl/sdks/go/model"
)

type FakeEventSubscriber struct {
	SubscribeStub        func(ctx context.Context, filter model.EventFilter) (<-chan model.Event, <-chan error)
	subscribeMutex       sync.RWMutex
	subscribeArgsForCall []struct {
		ctx    context.Context
		filter model.EventFilter
	}
	subscribeReturns struct {
		result1 <-chan model.Event
		result2 <-chan error
	}
	subscribeReturnsOnCall map[int]struct {
		result1 <-chan model.Event
		result2 <-chan error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEventSubscriber) Subscribe(ctx context.Context, filter model.EventFilter) (<-chan model.Event, <-chan error) {
	fake.subscribeMutex.Lock()
	ret, specificReturn := fake.subscribeReturnsOnCall[len(fake.subscribeArgsForCall)]
	fake.subscribeArgsForCall = append(fake.subscribeArgsForCall, struct {
		ctx    context.Context
		filter model.EventFilter
	}{ctx, filter})
	fake.recordInvocation("Subscribe", []interface{}{ctx, filter})
	fake.subscribeMutex.Unlock()
	if fake.SubscribeStub != nil {
		return fake.SubscribeStub(ctx, filter)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.subscribeReturns.result1, fake.subscribeReturns.result2
}

func (fake *FakeEventSubscriber) SubscribeCallCount() int {
	fake.subscribeMutex.RLock()
	defer fake.subscribeMutex.RUnlock()
	return len(fake.subscribeArgsForCall)
}

func (fake *FakeEventSubscriber) SubscribeArgsForCall(i int) (context.Context, model.EventFilter) {
	fake.subscribeMutex.RLock()
	defer fake.subscribeMutex.RUnlock()
	return fake.subscribeArgsForCall[i].ctx, fake.subscribeArgsForCall[i].filter
}

func (fake *FakeEventSubscriber) SubscribeReturns(result1 <-chan model.Event, result2 <-chan error) {
	fake.SubscribeStub = nil
	fake.subscribeReturns = struct {
		result1 <-chan model.Event
		result2 <-chan error
	}{result1, result2}
}

func (fake *FakeEventSubscriber) SubscribeReturnsOnCall(i int, result1 <-chan model.Event, result2 <-chan error) {
	fake.SubscribeStub = nil
	if fake.subscribeReturnsOnCall == nil {
		fake.subscribeReturnsOnCall = make(map[int]struct {
			result1 <-chan model.Event
			result2 <-chan error
		})
	}
	fake.subscribeReturnsOnCall[i] = struct {
		result1 <-chan model.Event
		result2 <-chan error
	}{result1, result2}
}

func (fake *FakeEventSubscriber) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.subscribeMutex.RLock()
	defer fake.subscribeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEventSubscriber) recordInvocation(key string, args []interface{}) {
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

var _ EventSubscriber = new(FakeEventSubscriber)
