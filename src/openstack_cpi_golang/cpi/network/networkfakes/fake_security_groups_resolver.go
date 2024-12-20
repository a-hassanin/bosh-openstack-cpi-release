// Code generated by counterfeiter. DO NOT EDIT.
package networkfakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-openstack-cpi-release/src/openstack_cpi_golang/cpi/network"
)

type FakeSecurityGroupsResolver struct {
	ResolveStub        func([]string) ([]string, error)
	resolveMutex       sync.RWMutex
	resolveArgsForCall []struct {
		arg1 []string
	}
	resolveReturns struct {
		result1 []string
		result2 error
	}
	resolveReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSecurityGroupsResolver) Resolve(arg1 []string) ([]string, error) {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.resolveMutex.Lock()
	ret, specificReturn := fake.resolveReturnsOnCall[len(fake.resolveArgsForCall)]
	fake.resolveArgsForCall = append(fake.resolveArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	stub := fake.ResolveStub
	fakeReturns := fake.resolveReturns
	fake.recordInvocation("Resolve", []interface{}{arg1Copy})
	fake.resolveMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSecurityGroupsResolver) ResolveCallCount() int {
	fake.resolveMutex.RLock()
	defer fake.resolveMutex.RUnlock()
	return len(fake.resolveArgsForCall)
}

func (fake *FakeSecurityGroupsResolver) ResolveCalls(stub func([]string) ([]string, error)) {
	fake.resolveMutex.Lock()
	defer fake.resolveMutex.Unlock()
	fake.ResolveStub = stub
}

func (fake *FakeSecurityGroupsResolver) ResolveArgsForCall(i int) []string {
	fake.resolveMutex.RLock()
	defer fake.resolveMutex.RUnlock()
	argsForCall := fake.resolveArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSecurityGroupsResolver) ResolveReturns(result1 []string, result2 error) {
	fake.resolveMutex.Lock()
	defer fake.resolveMutex.Unlock()
	fake.ResolveStub = nil
	fake.resolveReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeSecurityGroupsResolver) ResolveReturnsOnCall(i int, result1 []string, result2 error) {
	fake.resolveMutex.Lock()
	defer fake.resolveMutex.Unlock()
	fake.ResolveStub = nil
	if fake.resolveReturnsOnCall == nil {
		fake.resolveReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.resolveReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeSecurityGroupsResolver) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.resolveMutex.RLock()
	defer fake.resolveMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSecurityGroupsResolver) recordInvocation(key string, args []interface{}) {
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

var _ network.SecurityGroupsResolver = new(FakeSecurityGroupsResolver)
