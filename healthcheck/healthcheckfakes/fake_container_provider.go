// Code generated by counterfeiter. DO NOT EDIT.
package healthcheckfakes

import (
	"context"
	"sync"

	"github.com/cirocosta/concourse-worker-health-checker/healthcheck"
)

type FakeContainerProvider struct {
	CreateStub        func(context.Context, string, string) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}
	createReturns struct {
		result1 error
	}
	createReturnsOnCall map[int]struct {
		result1 error
	}
	DestroyStub        func(context.Context, string) error
	destroyMutex       sync.RWMutex
	destroyArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	destroyReturns struct {
		result1 error
	}
	destroyReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeContainerProvider) Create(arg1 context.Context, arg2 string, arg3 string) error {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("Create", []interface{}{arg1, arg2, arg3})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createReturns
	return fakeReturns.result1
}

func (fake *FakeContainerProvider) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeContainerProvider) CreateCalls(stub func(context.Context, string, string) error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeContainerProvider) CreateArgsForCall(i int) (context.Context, string, string) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeContainerProvider) CreateReturns(result1 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainerProvider) CreateReturnsOnCall(i int, result1 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainerProvider) Destroy(arg1 context.Context, arg2 string) error {
	fake.destroyMutex.Lock()
	ret, specificReturn := fake.destroyReturnsOnCall[len(fake.destroyArgsForCall)]
	fake.destroyArgsForCall = append(fake.destroyArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Destroy", []interface{}{arg1, arg2})
	fake.destroyMutex.Unlock()
	if fake.DestroyStub != nil {
		return fake.DestroyStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.destroyReturns
	return fakeReturns.result1
}

func (fake *FakeContainerProvider) DestroyCallCount() int {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return len(fake.destroyArgsForCall)
}

func (fake *FakeContainerProvider) DestroyCalls(stub func(context.Context, string) error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = stub
}

func (fake *FakeContainerProvider) DestroyArgsForCall(i int) (context.Context, string) {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	argsForCall := fake.destroyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeContainerProvider) DestroyReturns(result1 error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = nil
	fake.destroyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainerProvider) DestroyReturnsOnCall(i int, result1 error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = nil
	if fake.destroyReturnsOnCall == nil {
		fake.destroyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.destroyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainerProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeContainerProvider) recordInvocation(key string, args []interface{}) {
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

var _ healthcheck.ContainerProvider = new(FakeContainerProvider)
