package core

import (
	"fmt"
	"github.com/opspec-io/opctl/pkg/containerengine"
	"github.com/opspec-io/opctl/util/eventbus"
	"github.com/opspec-io/sdk-golang/pkg/bundle"
	"github.com/opspec-io/sdk-golang/pkg/model"
	"time"
)

type containerCaller interface {
	Call(
		args map[string]*model.Data,
		containerId string,
		containerCall *model.ContainerCall,
		opRef string,
		opGraphId string,
	) (
		outputs map[string]*model.Data,
		err error,
	)
}

func newContainerCaller(
	bundle bundle.Bundle,
	containerEngine containerengine.ContainerEngine,
	eventBus eventbus.EventBus,
	nodeRepo nodeRepo,
) containerCaller {

	return &_containerCaller{
		bundle:          bundle,
		containerEngine: containerEngine,
		eventBus:        eventBus,
		nodeRepo:        nodeRepo,
	}

}

type _containerCaller struct {
	bundle          bundle.Bundle
	containerEngine containerengine.ContainerEngine
	eventBus        eventbus.EventBus
	nodeRepo        nodeRepo
}

func (this _containerCaller) Call(
	args map[string]*model.Data,
	containerId string,
	containerCall *model.ContainerCall,
	opRef string,
	opGraphId string,
) (
	outputs map[string]*model.Data,
	err error,
) {

	op, err := this.bundle.GetOp(
		opRef,
	)
	if nil != err {
		return
	}

	this.nodeRepo.add(
		&nodeDescriptor{
			Id:        containerId,
			OpRef:     opRef,
			OpGraphId: opGraphId,
			Container: &containerDescriptor{},
		},
	)

	containerStartedEvent := model.Event{
		Timestamp: time.Now().UTC(),
		ContainerStarted: &model.ContainerStartedEvent{
			ContainerId: containerId,
			OpRef:       opRef,
			OpGraphId:   opGraphId,
		},
	}
	this.eventBus.Publish(containerStartedEvent)

	err = this.containerEngine.StartContainer(
		newContainerStartReq(args, containerCall, containerId, op.Inputs, opGraphId),
		this.eventBus,
	)
	if nil != err {
		return
	}

	// construct outputs
	outputs = map[string]*model.Data{}
	container, err := this.containerEngine.InspectContainerIfExists(containerId)
	fmt.Println(opRef)
	fmt.Printf("containerCaller.container:\n %#v\n", container)
	// construct
	for _, fsEntry := range containerCall.Fs {
		for _, rawFsEntry := range container.Fs {
			if fsEntry.Path == rawFsEntry.Path {
				outputs[fsEntry.Bind] = &model.Data{File: rawFsEntry.SrcRef}
			}
		}
	}
	for _, envEntry := range containerCall.Env {
		for _, rawEnvEntry := range container.Env {
			if envEntry.Name == rawEnvEntry.Name {
				outputs[envEntry.Bind] = &model.Data{String: rawEnvEntry.Value}
			}
		}
	}
	fmt.Printf("containerCaller.outputs:\n %#v\n", outputs)
	// @todo: handle containerCall.Net

	defer func() {

		this.nodeRepo.deleteIfExists(containerId)

		this.eventBus.Publish(
			model.Event{
				Timestamp: time.Now().UTC(),
				ContainerExited: &model.ContainerExitedEvent{
					ContainerId: containerId,
					OpRef:       opRef,
					OpGraphId:   opGraphId,
				},
			},
		)

	}()

	return
}
