// This file was generated by counterfeiter
package opspec

import (
  "sync"
  "github.com/opspec-io/sdk-golang/models"
)

type FakeSdk struct {
  CreateCollectionStub                   func(req models.CreateCollectionReq) (err error)
  createCollectionMutex                  sync.RWMutex
  createCollectionArgsForCall            []struct {
    req models.CreateCollectionReq
  }
  createCollectionReturns                struct {
                                           result1 error
                                         }
  CreateOpStub                           func(req models.CreateOpReq) (err error)
  createOpMutex                          sync.RWMutex
  createOpArgsForCall                    []struct {
    req models.CreateOpReq
  }
  createOpReturns                        struct {
                                           result1 error
                                         }
  GetCollectionStub                      func(collectionBundlePath string) (collectionView models.CollectionView, err error)
  getCollectionMutex                     sync.RWMutex
  getCollectionArgsForCall               []struct {
    collectionBundlePath string
  }
  getCollectionReturns                   struct {
                                           result1 models.CollectionView
                                           result2 error
                                         }
  GetEventStreamStub                     func() (stream chan models.Event, err error)
  getEventStreamMutex                    sync.RWMutex
  getEventStreamArgsForCall              []struct{}
  getEventStreamReturns                  struct {
                                           result1 chan models.Event
                                           result2 error
                                         }
  GetOpStub                              func(opBundlePath string) (opView models.OpView, err error)
  getOpMutex                             sync.RWMutex
  getOpArgsForCall                       []struct {
    opBundlePath string
  }
  getOpReturns                           struct {
                                           result1 models.OpView
                                           result2 error
                                         }
  KillOpRunStub                          func(req models.KillOpRunReq) (err error)
  killOpRunMutex                         sync.RWMutex
  killOpRunArgsForCall                   []struct {
    req models.KillOpRunReq
  }
  killOpRunReturns                       struct {
                                           result1 error
                                         }
  SetCollectionDescriptionStub           func(req models.SetCollectionDescriptionReq) (err error)
  setCollectionDescriptionMutex          sync.RWMutex
  setCollectionDescriptionArgsForCall    []struct {
    req models.SetCollectionDescriptionReq
  }
  setCollectionDescriptionReturns        struct {
                                           result1 error
                                         }
  SetOpDescriptionStub                   func(req models.SetOpDescriptionReq) (err error)
  setOpDescriptionMutex                  sync.RWMutex
  setOpDescriptionArgsForCall            []struct {
    req models.SetOpDescriptionReq
  }
  setOpDescriptionReturns                struct {
                                           result1 error
                                         }
  StartOpRunStub                         func(req models.StartOpRunReq) (opRunId string, err error)
  startOpRunMutex                        sync.RWMutex
  startOpRunArgsForCall                  []struct {
    req models.StartOpRunReq
  }
  startOpRunReturns                      struct {
                                           result1 string
                                           result2 error
                                         }
  TryResolveDefaultCollectionStub        func(req models.TryResolveDefaultCollectionReq) (pathToDefaultCollection string, err error)
  tryResolveDefaultCollectionMutex       sync.RWMutex
  tryResolveDefaultCollectionArgsForCall []struct {
    req models.TryResolveDefaultCollectionReq
  }
  tryResolveDefaultCollectionReturns     struct {
                                           result1 string
                                           result2 error
                                         }
  invocations                            map[string][][]interface{}
  invocationsMutex                       sync.RWMutex
}

func (fake *FakeSdk) CreateCollection(req models.CreateCollectionReq) (err error) {
  fake.createCollectionMutex.Lock()
  fake.createCollectionArgsForCall = append(fake.createCollectionArgsForCall, struct {
    req models.CreateCollectionReq
  }{req})
  fake.recordInvocation("CreateCollection", []interface{}{req})
  fake.createCollectionMutex.Unlock()
  if fake.CreateCollectionStub != nil {
    return fake.CreateCollectionStub(req)
  } else {
    return fake.createCollectionReturns.result1
  }
}

func (fake *FakeSdk) CreateCollectionCallCount() int {
  fake.createCollectionMutex.RLock()
  defer fake.createCollectionMutex.RUnlock()
  return len(fake.createCollectionArgsForCall)
}

func (fake *FakeSdk) CreateCollectionArgsForCall(i int) models.CreateCollectionReq {
  fake.createCollectionMutex.RLock()
  defer fake.createCollectionMutex.RUnlock()
  return fake.createCollectionArgsForCall[i].req
}

func (fake *FakeSdk) CreateCollectionReturns(result1 error) {
  fake.CreateCollectionStub = nil
  fake.createCollectionReturns = struct {
    result1 error
  }{result1}
}

func (fake *FakeSdk) CreateOp(req models.CreateOpReq) (err error) {
  fake.createOpMutex.Lock()
  fake.createOpArgsForCall = append(fake.createOpArgsForCall, struct {
    req models.CreateOpReq
  }{req})
  fake.recordInvocation("CreateOp", []interface{}{req})
  fake.createOpMutex.Unlock()
  if fake.CreateOpStub != nil {
    return fake.CreateOpStub(req)
  } else {
    return fake.createOpReturns.result1
  }
}

func (fake *FakeSdk) CreateOpCallCount() int {
  fake.createOpMutex.RLock()
  defer fake.createOpMutex.RUnlock()
  return len(fake.createOpArgsForCall)
}

func (fake *FakeSdk) CreateOpArgsForCall(i int) models.CreateOpReq {
  fake.createOpMutex.RLock()
  defer fake.createOpMutex.RUnlock()
  return fake.createOpArgsForCall[i].req
}

func (fake *FakeSdk) CreateOpReturns(result1 error) {
  fake.CreateOpStub = nil
  fake.createOpReturns = struct {
    result1 error
  }{result1}
}

func (fake *FakeSdk) GetCollection(collectionBundlePath string) (collectionView models.CollectionView, err error) {
  fake.getCollectionMutex.Lock()
  fake.getCollectionArgsForCall = append(fake.getCollectionArgsForCall, struct {
    collectionBundlePath string
  }{collectionBundlePath})
  fake.recordInvocation("GetCollection", []interface{}{collectionBundlePath})
  fake.getCollectionMutex.Unlock()
  if fake.GetCollectionStub != nil {
    return fake.GetCollectionStub(collectionBundlePath)
  } else {
    return fake.getCollectionReturns.result1, fake.getCollectionReturns.result2
  }
}

func (fake *FakeSdk) GetCollectionCallCount() int {
  fake.getCollectionMutex.RLock()
  defer fake.getCollectionMutex.RUnlock()
  return len(fake.getCollectionArgsForCall)
}

func (fake *FakeSdk) GetCollectionArgsForCall(i int) string {
  fake.getCollectionMutex.RLock()
  defer fake.getCollectionMutex.RUnlock()
  return fake.getCollectionArgsForCall[i].collectionBundlePath
}

func (fake *FakeSdk) GetCollectionReturns(result1 models.CollectionView, result2 error) {
  fake.GetCollectionStub = nil
  fake.getCollectionReturns = struct {
    result1 models.CollectionView
    result2 error
  }{result1, result2}
}

func (fake *FakeSdk) GetEventStream() (stream chan models.Event, err error) {
  fake.getEventStreamMutex.Lock()
  fake.getEventStreamArgsForCall = append(fake.getEventStreamArgsForCall, struct{}{})
  fake.recordInvocation("GetEventStream", []interface{}{})
  fake.getEventStreamMutex.Unlock()
  if fake.GetEventStreamStub != nil {
    return fake.GetEventStreamStub()
  } else {
    return fake.getEventStreamReturns.result1, fake.getEventStreamReturns.result2
  }
}

func (fake *FakeSdk) GetEventStreamCallCount() int {
  fake.getEventStreamMutex.RLock()
  defer fake.getEventStreamMutex.RUnlock()
  return len(fake.getEventStreamArgsForCall)
}

func (fake *FakeSdk) GetEventStreamReturns(result1 chan models.Event, result2 error) {
  fake.GetEventStreamStub = nil
  fake.getEventStreamReturns = struct {
    result1 chan models.Event
    result2 error
  }{result1, result2}
}

func (fake *FakeSdk) GetOp(opBundlePath string) (opView models.OpView, err error) {
  fake.getOpMutex.Lock()
  fake.getOpArgsForCall = append(fake.getOpArgsForCall, struct {
    opBundlePath string
  }{opBundlePath})
  fake.recordInvocation("GetOp", []interface{}{opBundlePath})
  fake.getOpMutex.Unlock()
  if fake.GetOpStub != nil {
    return fake.GetOpStub(opBundlePath)
  } else {
    return fake.getOpReturns.result1, fake.getOpReturns.result2
  }
}

func (fake *FakeSdk) GetOpCallCount() int {
  fake.getOpMutex.RLock()
  defer fake.getOpMutex.RUnlock()
  return len(fake.getOpArgsForCall)
}

func (fake *FakeSdk) GetOpArgsForCall(i int) string {
  fake.getOpMutex.RLock()
  defer fake.getOpMutex.RUnlock()
  return fake.getOpArgsForCall[i].opBundlePath
}

func (fake *FakeSdk) GetOpReturns(result1 models.OpView, result2 error) {
  fake.GetOpStub = nil
  fake.getOpReturns = struct {
    result1 models.OpView
    result2 error
  }{result1, result2}
}

func (fake *FakeSdk) KillOpRun(req models.KillOpRunReq) (err error) {
  fake.killOpRunMutex.Lock()
  fake.killOpRunArgsForCall = append(fake.killOpRunArgsForCall, struct {
    req models.KillOpRunReq
  }{req})
  fake.recordInvocation("KillOpRun", []interface{}{req})
  fake.killOpRunMutex.Unlock()
  if fake.KillOpRunStub != nil {
    return fake.KillOpRunStub(req)
  } else {
    return fake.killOpRunReturns.result1
  }
}

func (fake *FakeSdk) KillOpRunCallCount() int {
  fake.killOpRunMutex.RLock()
  defer fake.killOpRunMutex.RUnlock()
  return len(fake.killOpRunArgsForCall)
}

func (fake *FakeSdk) KillOpRunArgsForCall(i int) models.KillOpRunReq {
  fake.killOpRunMutex.RLock()
  defer fake.killOpRunMutex.RUnlock()
  return fake.killOpRunArgsForCall[i].req
}

func (fake *FakeSdk) KillOpRunReturns(result1 error) {
  fake.KillOpRunStub = nil
  fake.killOpRunReturns = struct {
    result1 error
  }{result1}
}

func (fake *FakeSdk) SetCollectionDescription(req models.SetCollectionDescriptionReq) (err error) {
  fake.setCollectionDescriptionMutex.Lock()
  fake.setCollectionDescriptionArgsForCall = append(fake.setCollectionDescriptionArgsForCall, struct {
    req models.SetCollectionDescriptionReq
  }{req})
  fake.recordInvocation("SetCollectionDescription", []interface{}{req})
  fake.setCollectionDescriptionMutex.Unlock()
  if fake.SetCollectionDescriptionStub != nil {
    return fake.SetCollectionDescriptionStub(req)
  } else {
    return fake.setCollectionDescriptionReturns.result1
  }
}

func (fake *FakeSdk) SetCollectionDescriptionCallCount() int {
  fake.setCollectionDescriptionMutex.RLock()
  defer fake.setCollectionDescriptionMutex.RUnlock()
  return len(fake.setCollectionDescriptionArgsForCall)
}

func (fake *FakeSdk) SetCollectionDescriptionArgsForCall(i int) models.SetCollectionDescriptionReq {
  fake.setCollectionDescriptionMutex.RLock()
  defer fake.setCollectionDescriptionMutex.RUnlock()
  return fake.setCollectionDescriptionArgsForCall[i].req
}

func (fake *FakeSdk) SetCollectionDescriptionReturns(result1 error) {
  fake.SetCollectionDescriptionStub = nil
  fake.setCollectionDescriptionReturns = struct {
    result1 error
  }{result1}
}

func (fake *FakeSdk) SetOpDescription(req models.SetOpDescriptionReq) (err error) {
  fake.setOpDescriptionMutex.Lock()
  fake.setOpDescriptionArgsForCall = append(fake.setOpDescriptionArgsForCall, struct {
    req models.SetOpDescriptionReq
  }{req})
  fake.recordInvocation("SetOpDescription", []interface{}{req})
  fake.setOpDescriptionMutex.Unlock()
  if fake.SetOpDescriptionStub != nil {
    return fake.SetOpDescriptionStub(req)
  } else {
    return fake.setOpDescriptionReturns.result1
  }
}

func (fake *FakeSdk) SetOpDescriptionCallCount() int {
  fake.setOpDescriptionMutex.RLock()
  defer fake.setOpDescriptionMutex.RUnlock()
  return len(fake.setOpDescriptionArgsForCall)
}

func (fake *FakeSdk) SetOpDescriptionArgsForCall(i int) models.SetOpDescriptionReq {
  fake.setOpDescriptionMutex.RLock()
  defer fake.setOpDescriptionMutex.RUnlock()
  return fake.setOpDescriptionArgsForCall[i].req
}

func (fake *FakeSdk) SetOpDescriptionReturns(result1 error) {
  fake.SetOpDescriptionStub = nil
  fake.setOpDescriptionReturns = struct {
    result1 error
  }{result1}
}

func (fake *FakeSdk) StartOpRun(req models.StartOpRunReq) (opRunId string, err error) {
  fake.startOpRunMutex.Lock()
  fake.startOpRunArgsForCall = append(fake.startOpRunArgsForCall, struct {
    req models.StartOpRunReq
  }{req})
  fake.recordInvocation("StartOpRun", []interface{}{req})
  fake.startOpRunMutex.Unlock()
  if fake.StartOpRunStub != nil {
    return fake.StartOpRunStub(req)
  } else {
    return fake.startOpRunReturns.result1, fake.startOpRunReturns.result2
  }
}

func (fake *FakeSdk) StartOpRunCallCount() int {
  fake.startOpRunMutex.RLock()
  defer fake.startOpRunMutex.RUnlock()
  return len(fake.startOpRunArgsForCall)
}

func (fake *FakeSdk) StartOpRunArgsForCall(i int) models.StartOpRunReq {
  fake.startOpRunMutex.RLock()
  defer fake.startOpRunMutex.RUnlock()
  return fake.startOpRunArgsForCall[i].req
}

func (fake *FakeSdk) StartOpRunReturns(result1 string, result2 error) {
  fake.StartOpRunStub = nil
  fake.startOpRunReturns = struct {
    result1 string
    result2 error
  }{result1, result2}
}

func (fake *FakeSdk) TryResolveDefaultCollection(req models.TryResolveDefaultCollectionReq) (pathToDefaultCollection string, err error) {
  fake.tryResolveDefaultCollectionMutex.Lock()
  fake.tryResolveDefaultCollectionArgsForCall = append(fake.tryResolveDefaultCollectionArgsForCall, struct {
    req models.TryResolveDefaultCollectionReq
  }{req})
  fake.recordInvocation("TryResolveDefaultCollection", []interface{}{req})
  fake.tryResolveDefaultCollectionMutex.Unlock()
  if fake.TryResolveDefaultCollectionStub != nil {
    return fake.TryResolveDefaultCollectionStub(req)
  } else {
    return fake.tryResolveDefaultCollectionReturns.result1, fake.tryResolveDefaultCollectionReturns.result2
  }
}

func (fake *FakeSdk) TryResolveDefaultCollectionCallCount() int {
  fake.tryResolveDefaultCollectionMutex.RLock()
  defer fake.tryResolveDefaultCollectionMutex.RUnlock()
  return len(fake.tryResolveDefaultCollectionArgsForCall)
}

func (fake *FakeSdk) TryResolveDefaultCollectionArgsForCall(i int) models.TryResolveDefaultCollectionReq {
  fake.tryResolveDefaultCollectionMutex.RLock()
  defer fake.tryResolveDefaultCollectionMutex.RUnlock()
  return fake.tryResolveDefaultCollectionArgsForCall[i].req
}

func (fake *FakeSdk) TryResolveDefaultCollectionReturns(result1 string, result2 error) {
  fake.TryResolveDefaultCollectionStub = nil
  fake.tryResolveDefaultCollectionReturns = struct {
    result1 string
    result2 error
  }{result1, result2}
}

func (fake *FakeSdk) Invocations() map[string][][]interface{} {
  fake.invocationsMutex.RLock()
  defer fake.invocationsMutex.RUnlock()
  fake.createCollectionMutex.RLock()
  defer fake.createCollectionMutex.RUnlock()
  fake.createOpMutex.RLock()
  defer fake.createOpMutex.RUnlock()
  fake.getCollectionMutex.RLock()
  defer fake.getCollectionMutex.RUnlock()
  fake.getEventStreamMutex.RLock()
  defer fake.getEventStreamMutex.RUnlock()
  fake.getOpMutex.RLock()
  defer fake.getOpMutex.RUnlock()
  fake.killOpRunMutex.RLock()
  defer fake.killOpRunMutex.RUnlock()
  fake.setCollectionDescriptionMutex.RLock()
  defer fake.setCollectionDescriptionMutex.RUnlock()
  fake.setOpDescriptionMutex.RLock()
  defer fake.setOpDescriptionMutex.RUnlock()
  fake.startOpRunMutex.RLock()
  defer fake.startOpRunMutex.RUnlock()
  fake.tryResolveDefaultCollectionMutex.RLock()
  defer fake.tryResolveDefaultCollectionMutex.RUnlock()
  return fake.invocations
}

func (fake *FakeSdk) recordInvocation(key string, args []interface{}) {
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

var _ Sdk = new(FakeSdk)
