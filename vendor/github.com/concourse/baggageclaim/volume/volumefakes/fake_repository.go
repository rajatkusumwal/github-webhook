// This file was generated by counterfeiter
package volumefakes

import (
	"sync"

	"github.com/concourse/baggageclaim/volume"
)

type FakeRepository struct {
	ListVolumesStub        func(queryProperties volume.Properties) (volume.Volumes, error)
	listVolumesMutex       sync.RWMutex
	listVolumesArgsForCall []struct {
		queryProperties volume.Properties
	}
	listVolumesReturns struct {
		result1 volume.Volumes
		result2 error
	}
	GetVolumeStub        func(handle string) (volume.Volume, bool, error)
	getVolumeMutex       sync.RWMutex
	getVolumeArgsForCall []struct {
		handle string
	}
	getVolumeReturns struct {
		result1 volume.Volume
		result2 bool
		result3 error
	}
	GetVolumeStatsStub        func(handle string) (volume.VolumeStats, bool, error)
	getVolumeStatsMutex       sync.RWMutex
	getVolumeStatsArgsForCall []struct {
		handle string
	}
	getVolumeStatsReturns struct {
		result1 volume.VolumeStats
		result2 bool
		result3 error
	}
	CreateVolumeStub        func(strategy volume.Strategy, properties volume.Properties, ttlInSeconds uint) (volume.Volume, error)
	createVolumeMutex       sync.RWMutex
	createVolumeArgsForCall []struct {
		strategy     volume.Strategy
		properties   volume.Properties
		ttlInSeconds uint
	}
	createVolumeReturns struct {
		result1 volume.Volume
		result2 error
	}
	DestroyVolumeStub        func(handle string) error
	destroyVolumeMutex       sync.RWMutex
	destroyVolumeArgsForCall []struct {
		handle string
	}
	destroyVolumeReturns struct {
		result1 error
	}
	SetPropertyStub        func(handle string, propertyName string, propertyValue string) error
	setPropertyMutex       sync.RWMutex
	setPropertyArgsForCall []struct {
		handle        string
		propertyName  string
		propertyValue string
	}
	setPropertyReturns struct {
		result1 error
	}
	SetTTLStub        func(handle string, ttl uint) error
	setTTLMutex       sync.RWMutex
	setTTLArgsForCall []struct {
		handle string
		ttl    uint
	}
	setTTLReturns struct {
		result1 error
	}
	VolumeParentStub        func(handle string) (volume.Volume, bool, error)
	volumeParentMutex       sync.RWMutex
	volumeParentArgsForCall []struct {
		handle string
	}
	volumeParentReturns struct {
		result1 volume.Volume
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRepository) ListVolumes(queryProperties volume.Properties) (volume.Volumes, error) {
	fake.listVolumesMutex.Lock()
	fake.listVolumesArgsForCall = append(fake.listVolumesArgsForCall, struct {
		queryProperties volume.Properties
	}{queryProperties})
	fake.recordInvocation("ListVolumes", []interface{}{queryProperties})
	fake.listVolumesMutex.Unlock()
	if fake.ListVolumesStub != nil {
		return fake.ListVolumesStub(queryProperties)
	} else {
		return fake.listVolumesReturns.result1, fake.listVolumesReturns.result2
	}
}

func (fake *FakeRepository) ListVolumesCallCount() int {
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	return len(fake.listVolumesArgsForCall)
}

func (fake *FakeRepository) ListVolumesArgsForCall(i int) volume.Properties {
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	return fake.listVolumesArgsForCall[i].queryProperties
}

func (fake *FakeRepository) ListVolumesReturns(result1 volume.Volumes, result2 error) {
	fake.ListVolumesStub = nil
	fake.listVolumesReturns = struct {
		result1 volume.Volumes
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) GetVolume(handle string) (volume.Volume, bool, error) {
	fake.getVolumeMutex.Lock()
	fake.getVolumeArgsForCall = append(fake.getVolumeArgsForCall, struct {
		handle string
	}{handle})
	fake.recordInvocation("GetVolume", []interface{}{handle})
	fake.getVolumeMutex.Unlock()
	if fake.GetVolumeStub != nil {
		return fake.GetVolumeStub(handle)
	} else {
		return fake.getVolumeReturns.result1, fake.getVolumeReturns.result2, fake.getVolumeReturns.result3
	}
}

func (fake *FakeRepository) GetVolumeCallCount() int {
	fake.getVolumeMutex.RLock()
	defer fake.getVolumeMutex.RUnlock()
	return len(fake.getVolumeArgsForCall)
}

func (fake *FakeRepository) GetVolumeArgsForCall(i int) string {
	fake.getVolumeMutex.RLock()
	defer fake.getVolumeMutex.RUnlock()
	return fake.getVolumeArgsForCall[i].handle
}

func (fake *FakeRepository) GetVolumeReturns(result1 volume.Volume, result2 bool, result3 error) {
	fake.GetVolumeStub = nil
	fake.getVolumeReturns = struct {
		result1 volume.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRepository) GetVolumeStats(handle string) (volume.VolumeStats, bool, error) {
	fake.getVolumeStatsMutex.Lock()
	fake.getVolumeStatsArgsForCall = append(fake.getVolumeStatsArgsForCall, struct {
		handle string
	}{handle})
	fake.recordInvocation("GetVolumeStats", []interface{}{handle})
	fake.getVolumeStatsMutex.Unlock()
	if fake.GetVolumeStatsStub != nil {
		return fake.GetVolumeStatsStub(handle)
	} else {
		return fake.getVolumeStatsReturns.result1, fake.getVolumeStatsReturns.result2, fake.getVolumeStatsReturns.result3
	}
}

func (fake *FakeRepository) GetVolumeStatsCallCount() int {
	fake.getVolumeStatsMutex.RLock()
	defer fake.getVolumeStatsMutex.RUnlock()
	return len(fake.getVolumeStatsArgsForCall)
}

func (fake *FakeRepository) GetVolumeStatsArgsForCall(i int) string {
	fake.getVolumeStatsMutex.RLock()
	defer fake.getVolumeStatsMutex.RUnlock()
	return fake.getVolumeStatsArgsForCall[i].handle
}

func (fake *FakeRepository) GetVolumeStatsReturns(result1 volume.VolumeStats, result2 bool, result3 error) {
	fake.GetVolumeStatsStub = nil
	fake.getVolumeStatsReturns = struct {
		result1 volume.VolumeStats
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRepository) CreateVolume(strategy volume.Strategy, properties volume.Properties, ttlInSeconds uint) (volume.Volume, error) {
	fake.createVolumeMutex.Lock()
	fake.createVolumeArgsForCall = append(fake.createVolumeArgsForCall, struct {
		strategy     volume.Strategy
		properties   volume.Properties
		ttlInSeconds uint
	}{strategy, properties, ttlInSeconds})
	fake.recordInvocation("CreateVolume", []interface{}{strategy, properties, ttlInSeconds})
	fake.createVolumeMutex.Unlock()
	if fake.CreateVolumeStub != nil {
		return fake.CreateVolumeStub(strategy, properties, ttlInSeconds)
	} else {
		return fake.createVolumeReturns.result1, fake.createVolumeReturns.result2
	}
}

func (fake *FakeRepository) CreateVolumeCallCount() int {
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	return len(fake.createVolumeArgsForCall)
}

func (fake *FakeRepository) CreateVolumeArgsForCall(i int) (volume.Strategy, volume.Properties, uint) {
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	return fake.createVolumeArgsForCall[i].strategy, fake.createVolumeArgsForCall[i].properties, fake.createVolumeArgsForCall[i].ttlInSeconds
}

func (fake *FakeRepository) CreateVolumeReturns(result1 volume.Volume, result2 error) {
	fake.CreateVolumeStub = nil
	fake.createVolumeReturns = struct {
		result1 volume.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) DestroyVolume(handle string) error {
	fake.destroyVolumeMutex.Lock()
	fake.destroyVolumeArgsForCall = append(fake.destroyVolumeArgsForCall, struct {
		handle string
	}{handle})
	fake.recordInvocation("DestroyVolume", []interface{}{handle})
	fake.destroyVolumeMutex.Unlock()
	if fake.DestroyVolumeStub != nil {
		return fake.DestroyVolumeStub(handle)
	} else {
		return fake.destroyVolumeReturns.result1
	}
}

func (fake *FakeRepository) DestroyVolumeCallCount() int {
	fake.destroyVolumeMutex.RLock()
	defer fake.destroyVolumeMutex.RUnlock()
	return len(fake.destroyVolumeArgsForCall)
}

func (fake *FakeRepository) DestroyVolumeArgsForCall(i int) string {
	fake.destroyVolumeMutex.RLock()
	defer fake.destroyVolumeMutex.RUnlock()
	return fake.destroyVolumeArgsForCall[i].handle
}

func (fake *FakeRepository) DestroyVolumeReturns(result1 error) {
	fake.DestroyVolumeStub = nil
	fake.destroyVolumeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) SetProperty(handle string, propertyName string, propertyValue string) error {
	fake.setPropertyMutex.Lock()
	fake.setPropertyArgsForCall = append(fake.setPropertyArgsForCall, struct {
		handle        string
		propertyName  string
		propertyValue string
	}{handle, propertyName, propertyValue})
	fake.recordInvocation("SetProperty", []interface{}{handle, propertyName, propertyValue})
	fake.setPropertyMutex.Unlock()
	if fake.SetPropertyStub != nil {
		return fake.SetPropertyStub(handle, propertyName, propertyValue)
	} else {
		return fake.setPropertyReturns.result1
	}
}

func (fake *FakeRepository) SetPropertyCallCount() int {
	fake.setPropertyMutex.RLock()
	defer fake.setPropertyMutex.RUnlock()
	return len(fake.setPropertyArgsForCall)
}

func (fake *FakeRepository) SetPropertyArgsForCall(i int) (string, string, string) {
	fake.setPropertyMutex.RLock()
	defer fake.setPropertyMutex.RUnlock()
	return fake.setPropertyArgsForCall[i].handle, fake.setPropertyArgsForCall[i].propertyName, fake.setPropertyArgsForCall[i].propertyValue
}

func (fake *FakeRepository) SetPropertyReturns(result1 error) {
	fake.SetPropertyStub = nil
	fake.setPropertyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) SetTTL(handle string, ttl uint) error {
	fake.setTTLMutex.Lock()
	fake.setTTLArgsForCall = append(fake.setTTLArgsForCall, struct {
		handle string
		ttl    uint
	}{handle, ttl})
	fake.recordInvocation("SetTTL", []interface{}{handle, ttl})
	fake.setTTLMutex.Unlock()
	if fake.SetTTLStub != nil {
		return fake.SetTTLStub(handle, ttl)
	} else {
		return fake.setTTLReturns.result1
	}
}

func (fake *FakeRepository) SetTTLCallCount() int {
	fake.setTTLMutex.RLock()
	defer fake.setTTLMutex.RUnlock()
	return len(fake.setTTLArgsForCall)
}

func (fake *FakeRepository) SetTTLArgsForCall(i int) (string, uint) {
	fake.setTTLMutex.RLock()
	defer fake.setTTLMutex.RUnlock()
	return fake.setTTLArgsForCall[i].handle, fake.setTTLArgsForCall[i].ttl
}

func (fake *FakeRepository) SetTTLReturns(result1 error) {
	fake.SetTTLStub = nil
	fake.setTTLReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) VolumeParent(handle string) (volume.Volume, bool, error) {
	fake.volumeParentMutex.Lock()
	fake.volumeParentArgsForCall = append(fake.volumeParentArgsForCall, struct {
		handle string
	}{handle})
	fake.recordInvocation("VolumeParent", []interface{}{handle})
	fake.volumeParentMutex.Unlock()
	if fake.VolumeParentStub != nil {
		return fake.VolumeParentStub(handle)
	} else {
		return fake.volumeParentReturns.result1, fake.volumeParentReturns.result2, fake.volumeParentReturns.result3
	}
}

func (fake *FakeRepository) VolumeParentCallCount() int {
	fake.volumeParentMutex.RLock()
	defer fake.volumeParentMutex.RUnlock()
	return len(fake.volumeParentArgsForCall)
}

func (fake *FakeRepository) VolumeParentArgsForCall(i int) string {
	fake.volumeParentMutex.RLock()
	defer fake.volumeParentMutex.RUnlock()
	return fake.volumeParentArgsForCall[i].handle
}

func (fake *FakeRepository) VolumeParentReturns(result1 volume.Volume, result2 bool, result3 error) {
	fake.VolumeParentStub = nil
	fake.volumeParentReturns = struct {
		result1 volume.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	fake.getVolumeMutex.RLock()
	defer fake.getVolumeMutex.RUnlock()
	fake.getVolumeStatsMutex.RLock()
	defer fake.getVolumeStatsMutex.RUnlock()
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	fake.destroyVolumeMutex.RLock()
	defer fake.destroyVolumeMutex.RUnlock()
	fake.setPropertyMutex.RLock()
	defer fake.setPropertyMutex.RUnlock()
	fake.setTTLMutex.RLock()
	defer fake.setTTLMutex.RUnlock()
	fake.volumeParentMutex.RLock()
	defer fake.volumeParentMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeRepository) recordInvocation(key string, args []interface{}) {
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

var _ volume.Repository = new(FakeRepository)
