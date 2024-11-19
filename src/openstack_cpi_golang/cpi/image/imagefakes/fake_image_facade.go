// Code generated by counterfeiter. DO NOT EDIT.
package imagefakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-openstack-cpi-release/src/openstack_cpi_golang/cpi/image"
	"github.com/cloudfoundry/bosh-openstack-cpi-release/src/openstack_cpi_golang/cpi/utils"
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
)

type FakeImageFacade struct {
	CreateImageStub        func(utils.ServiceClient, images.CreateOptsBuilder) (*images.Image, error)
	createImageMutex       sync.RWMutex
	createImageArgsForCall []struct {
		arg1 utils.ServiceClient
		arg2 images.CreateOptsBuilder
	}
	createImageReturns struct {
		result1 *images.Image
		result2 error
	}
	createImageReturnsOnCall map[int]struct {
		result1 *images.Image
		result2 error
	}
	DeleteImageStub        func(utils.RetryableServiceClient, string) error
	deleteImageMutex       sync.RWMutex
	deleteImageArgsForCall []struct {
		arg1 utils.RetryableServiceClient
		arg2 string
	}
	deleteImageReturns struct {
		result1 error
	}
	deleteImageReturnsOnCall map[int]struct {
		result1 error
	}
	GetImageStub        func(utils.RetryableServiceClient, string) (*images.Image, error)
	getImageMutex       sync.RWMutex
	getImageArgsForCall []struct {
		arg1 utils.RetryableServiceClient
		arg2 string
	}
	getImageReturns struct {
		result1 *images.Image
		result2 error
	}
	getImageReturnsOnCall map[int]struct {
		result1 *images.Image
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImageFacade) CreateImage(arg1 utils.ServiceClient, arg2 images.CreateOptsBuilder) (*images.Image, error) {
	fake.createImageMutex.Lock()
	ret, specificReturn := fake.createImageReturnsOnCall[len(fake.createImageArgsForCall)]
	fake.createImageArgsForCall = append(fake.createImageArgsForCall, struct {
		arg1 utils.ServiceClient
		arg2 images.CreateOptsBuilder
	}{arg1, arg2})
	stub := fake.CreateImageStub
	fakeReturns := fake.createImageReturns
	fake.recordInvocation("CreateImage", []interface{}{arg1, arg2})
	fake.createImageMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImageFacade) CreateImageCallCount() int {
	fake.createImageMutex.RLock()
	defer fake.createImageMutex.RUnlock()
	return len(fake.createImageArgsForCall)
}

func (fake *FakeImageFacade) CreateImageCalls(stub func(utils.ServiceClient, images.CreateOptsBuilder) (*images.Image, error)) {
	fake.createImageMutex.Lock()
	defer fake.createImageMutex.Unlock()
	fake.CreateImageStub = stub
}

func (fake *FakeImageFacade) CreateImageArgsForCall(i int) (utils.ServiceClient, images.CreateOptsBuilder) {
	fake.createImageMutex.RLock()
	defer fake.createImageMutex.RUnlock()
	argsForCall := fake.createImageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImageFacade) CreateImageReturns(result1 *images.Image, result2 error) {
	fake.createImageMutex.Lock()
	defer fake.createImageMutex.Unlock()
	fake.CreateImageStub = nil
	fake.createImageReturns = struct {
		result1 *images.Image
		result2 error
	}{result1, result2}
}

func (fake *FakeImageFacade) CreateImageReturnsOnCall(i int, result1 *images.Image, result2 error) {
	fake.createImageMutex.Lock()
	defer fake.createImageMutex.Unlock()
	fake.CreateImageStub = nil
	if fake.createImageReturnsOnCall == nil {
		fake.createImageReturnsOnCall = make(map[int]struct {
			result1 *images.Image
			result2 error
		})
	}
	fake.createImageReturnsOnCall[i] = struct {
		result1 *images.Image
		result2 error
	}{result1, result2}
}

func (fake *FakeImageFacade) DeleteImage(arg1 utils.RetryableServiceClient, arg2 string) error {
	fake.deleteImageMutex.Lock()
	ret, specificReturn := fake.deleteImageReturnsOnCall[len(fake.deleteImageArgsForCall)]
	fake.deleteImageArgsForCall = append(fake.deleteImageArgsForCall, struct {
		arg1 utils.RetryableServiceClient
		arg2 string
	}{arg1, arg2})
	stub := fake.DeleteImageStub
	fakeReturns := fake.deleteImageReturns
	fake.recordInvocation("DeleteImage", []interface{}{arg1, arg2})
	fake.deleteImageMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImageFacade) DeleteImageCallCount() int {
	fake.deleteImageMutex.RLock()
	defer fake.deleteImageMutex.RUnlock()
	return len(fake.deleteImageArgsForCall)
}

func (fake *FakeImageFacade) DeleteImageCalls(stub func(utils.RetryableServiceClient, string) error) {
	fake.deleteImageMutex.Lock()
	defer fake.deleteImageMutex.Unlock()
	fake.DeleteImageStub = stub
}

func (fake *FakeImageFacade) DeleteImageArgsForCall(i int) (utils.RetryableServiceClient, string) {
	fake.deleteImageMutex.RLock()
	defer fake.deleteImageMutex.RUnlock()
	argsForCall := fake.deleteImageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImageFacade) DeleteImageReturns(result1 error) {
	fake.deleteImageMutex.Lock()
	defer fake.deleteImageMutex.Unlock()
	fake.DeleteImageStub = nil
	fake.deleteImageReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageFacade) DeleteImageReturnsOnCall(i int, result1 error) {
	fake.deleteImageMutex.Lock()
	defer fake.deleteImageMutex.Unlock()
	fake.DeleteImageStub = nil
	if fake.deleteImageReturnsOnCall == nil {
		fake.deleteImageReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteImageReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageFacade) GetImage(arg1 utils.RetryableServiceClient, arg2 string) (*images.Image, error) {
	fake.getImageMutex.Lock()
	ret, specificReturn := fake.getImageReturnsOnCall[len(fake.getImageArgsForCall)]
	fake.getImageArgsForCall = append(fake.getImageArgsForCall, struct {
		arg1 utils.RetryableServiceClient
		arg2 string
	}{arg1, arg2})
	stub := fake.GetImageStub
	fakeReturns := fake.getImageReturns
	fake.recordInvocation("GetImage", []interface{}{arg1, arg2})
	fake.getImageMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImageFacade) GetImageCallCount() int {
	fake.getImageMutex.RLock()
	defer fake.getImageMutex.RUnlock()
	return len(fake.getImageArgsForCall)
}

func (fake *FakeImageFacade) GetImageCalls(stub func(utils.RetryableServiceClient, string) (*images.Image, error)) {
	fake.getImageMutex.Lock()
	defer fake.getImageMutex.Unlock()
	fake.GetImageStub = stub
}

func (fake *FakeImageFacade) GetImageArgsForCall(i int) (utils.RetryableServiceClient, string) {
	fake.getImageMutex.RLock()
	defer fake.getImageMutex.RUnlock()
	argsForCall := fake.getImageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImageFacade) GetImageReturns(result1 *images.Image, result2 error) {
	fake.getImageMutex.Lock()
	defer fake.getImageMutex.Unlock()
	fake.GetImageStub = nil
	fake.getImageReturns = struct {
		result1 *images.Image
		result2 error
	}{result1, result2}
}

func (fake *FakeImageFacade) GetImageReturnsOnCall(i int, result1 *images.Image, result2 error) {
	fake.getImageMutex.Lock()
	defer fake.getImageMutex.Unlock()
	fake.GetImageStub = nil
	if fake.getImageReturnsOnCall == nil {
		fake.getImageReturnsOnCall = make(map[int]struct {
			result1 *images.Image
			result2 error
		})
	}
	fake.getImageReturnsOnCall[i] = struct {
		result1 *images.Image
		result2 error
	}{result1, result2}
}

func (fake *FakeImageFacade) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createImageMutex.RLock()
	defer fake.createImageMutex.RUnlock()
	fake.deleteImageMutex.RLock()
	defer fake.deleteImageMutex.RUnlock()
	fake.getImageMutex.RLock()
	defer fake.getImageMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeImageFacade) recordInvocation(key string, args []interface{}) {
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

var _ image.ImageFacade = new(FakeImageFacade)