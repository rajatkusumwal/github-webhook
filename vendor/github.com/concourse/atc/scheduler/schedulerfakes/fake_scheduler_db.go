// This file was generated by counterfeiter
package schedulerfakes

import (
	"sync"
	"time"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/db/algorithm"
	"github.com/concourse/atc/scheduler"
)

type FakeSchedulerDB struct {
	AcquireSchedulingLockStub        func(lager.Logger, time.Duration) (db.Lock, bool, error)
	acquireSchedulingLockMutex       sync.RWMutex
	acquireSchedulingLockArgsForCall []struct {
		arg1 lager.Logger
		arg2 time.Duration
	}
	acquireSchedulingLockReturns struct {
		result1 db.Lock
		result2 bool
		result3 error
	}
	LoadVersionsDBStub        func() (*algorithm.VersionsDB, error)
	loadVersionsDBMutex       sync.RWMutex
	loadVersionsDBArgsForCall []struct{}
	loadVersionsDBReturns     struct {
		result1 *algorithm.VersionsDB
		result2 error
	}
	GetPipelineNameStub        func() string
	getPipelineNameMutex       sync.RWMutex
	getPipelineNameArgsForCall []struct{}
	getPipelineNameReturns     struct {
		result1 string
	}
	ReloadStub        func() (bool, error)
	reloadMutex       sync.RWMutex
	reloadArgsForCall []struct{}
	reloadReturns     struct {
		result1 bool
		result2 error
	}
	ConfigStub        func() atc.Config
	configMutex       sync.RWMutex
	configArgsForCall []struct{}
	configReturns     struct {
		result1 atc.Config
	}
	CreateJobBuildStub        func(job string) (db.Build, error)
	createJobBuildMutex       sync.RWMutex
	createJobBuildArgsForCall []struct {
		job string
	}
	createJobBuildReturns struct {
		result1 db.Build
		result2 error
	}
	EnsurePendingBuildExistsStub        func(jobName string) error
	ensurePendingBuildExistsMutex       sync.RWMutex
	ensurePendingBuildExistsArgsForCall []struct {
		jobName string
	}
	ensurePendingBuildExistsReturns struct {
		result1 error
	}
	AcquireResourceCheckingForJobLockStub        func(logger lager.Logger, job string) (db.Lock, bool, error)
	acquireResourceCheckingForJobLockMutex       sync.RWMutex
	acquireResourceCheckingForJobLockArgsForCall []struct {
		logger lager.Logger
		job    string
	}
	acquireResourceCheckingForJobLockReturns struct {
		result1 db.Lock
		result2 bool
		result3 error
	}
	GetAllPendingBuildsStub        func() (map[string][]db.Build, error)
	getAllPendingBuildsMutex       sync.RWMutex
	getAllPendingBuildsArgsForCall []struct{}
	getAllPendingBuildsReturns     struct {
		result1 map[string][]db.Build
		result2 error
	}
	GetPendingBuildsForJobStub        func(jobName string) ([]db.Build, error)
	getPendingBuildsForJobMutex       sync.RWMutex
	getPendingBuildsForJobArgsForCall []struct {
		jobName string
	}
	getPendingBuildsForJobReturns struct {
		result1 []db.Build
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSchedulerDB) AcquireSchedulingLock(arg1 lager.Logger, arg2 time.Duration) (db.Lock, bool, error) {
	fake.acquireSchedulingLockMutex.Lock()
	fake.acquireSchedulingLockArgsForCall = append(fake.acquireSchedulingLockArgsForCall, struct {
		arg1 lager.Logger
		arg2 time.Duration
	}{arg1, arg2})
	fake.recordInvocation("AcquireSchedulingLock", []interface{}{arg1, arg2})
	fake.acquireSchedulingLockMutex.Unlock()
	if fake.AcquireSchedulingLockStub != nil {
		return fake.AcquireSchedulingLockStub(arg1, arg2)
	} else {
		return fake.acquireSchedulingLockReturns.result1, fake.acquireSchedulingLockReturns.result2, fake.acquireSchedulingLockReturns.result3
	}
}

func (fake *FakeSchedulerDB) AcquireSchedulingLockCallCount() int {
	fake.acquireSchedulingLockMutex.RLock()
	defer fake.acquireSchedulingLockMutex.RUnlock()
	return len(fake.acquireSchedulingLockArgsForCall)
}

func (fake *FakeSchedulerDB) AcquireSchedulingLockArgsForCall(i int) (lager.Logger, time.Duration) {
	fake.acquireSchedulingLockMutex.RLock()
	defer fake.acquireSchedulingLockMutex.RUnlock()
	return fake.acquireSchedulingLockArgsForCall[i].arg1, fake.acquireSchedulingLockArgsForCall[i].arg2
}

func (fake *FakeSchedulerDB) AcquireSchedulingLockReturns(result1 db.Lock, result2 bool, result3 error) {
	fake.AcquireSchedulingLockStub = nil
	fake.acquireSchedulingLockReturns = struct {
		result1 db.Lock
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSchedulerDB) LoadVersionsDB() (*algorithm.VersionsDB, error) {
	fake.loadVersionsDBMutex.Lock()
	fake.loadVersionsDBArgsForCall = append(fake.loadVersionsDBArgsForCall, struct{}{})
	fake.recordInvocation("LoadVersionsDB", []interface{}{})
	fake.loadVersionsDBMutex.Unlock()
	if fake.LoadVersionsDBStub != nil {
		return fake.LoadVersionsDBStub()
	} else {
		return fake.loadVersionsDBReturns.result1, fake.loadVersionsDBReturns.result2
	}
}

func (fake *FakeSchedulerDB) LoadVersionsDBCallCount() int {
	fake.loadVersionsDBMutex.RLock()
	defer fake.loadVersionsDBMutex.RUnlock()
	return len(fake.loadVersionsDBArgsForCall)
}

func (fake *FakeSchedulerDB) LoadVersionsDBReturns(result1 *algorithm.VersionsDB, result2 error) {
	fake.LoadVersionsDBStub = nil
	fake.loadVersionsDBReturns = struct {
		result1 *algorithm.VersionsDB
		result2 error
	}{result1, result2}
}

func (fake *FakeSchedulerDB) GetPipelineName() string {
	fake.getPipelineNameMutex.Lock()
	fake.getPipelineNameArgsForCall = append(fake.getPipelineNameArgsForCall, struct{}{})
	fake.recordInvocation("GetPipelineName", []interface{}{})
	fake.getPipelineNameMutex.Unlock()
	if fake.GetPipelineNameStub != nil {
		return fake.GetPipelineNameStub()
	} else {
		return fake.getPipelineNameReturns.result1
	}
}

func (fake *FakeSchedulerDB) GetPipelineNameCallCount() int {
	fake.getPipelineNameMutex.RLock()
	defer fake.getPipelineNameMutex.RUnlock()
	return len(fake.getPipelineNameArgsForCall)
}

func (fake *FakeSchedulerDB) GetPipelineNameReturns(result1 string) {
	fake.GetPipelineNameStub = nil
	fake.getPipelineNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeSchedulerDB) Reload() (bool, error) {
	fake.reloadMutex.Lock()
	fake.reloadArgsForCall = append(fake.reloadArgsForCall, struct{}{})
	fake.recordInvocation("Reload", []interface{}{})
	fake.reloadMutex.Unlock()
	if fake.ReloadStub != nil {
		return fake.ReloadStub()
	} else {
		return fake.reloadReturns.result1, fake.reloadReturns.result2
	}
}

func (fake *FakeSchedulerDB) ReloadCallCount() int {
	fake.reloadMutex.RLock()
	defer fake.reloadMutex.RUnlock()
	return len(fake.reloadArgsForCall)
}

func (fake *FakeSchedulerDB) ReloadReturns(result1 bool, result2 error) {
	fake.ReloadStub = nil
	fake.reloadReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeSchedulerDB) Config() atc.Config {
	fake.configMutex.Lock()
	fake.configArgsForCall = append(fake.configArgsForCall, struct{}{})
	fake.recordInvocation("Config", []interface{}{})
	fake.configMutex.Unlock()
	if fake.ConfigStub != nil {
		return fake.ConfigStub()
	} else {
		return fake.configReturns.result1
	}
}

func (fake *FakeSchedulerDB) ConfigCallCount() int {
	fake.configMutex.RLock()
	defer fake.configMutex.RUnlock()
	return len(fake.configArgsForCall)
}

func (fake *FakeSchedulerDB) ConfigReturns(result1 atc.Config) {
	fake.ConfigStub = nil
	fake.configReturns = struct {
		result1 atc.Config
	}{result1}
}

func (fake *FakeSchedulerDB) CreateJobBuild(job string) (db.Build, error) {
	fake.createJobBuildMutex.Lock()
	fake.createJobBuildArgsForCall = append(fake.createJobBuildArgsForCall, struct {
		job string
	}{job})
	fake.recordInvocation("CreateJobBuild", []interface{}{job})
	fake.createJobBuildMutex.Unlock()
	if fake.CreateJobBuildStub != nil {
		return fake.CreateJobBuildStub(job)
	} else {
		return fake.createJobBuildReturns.result1, fake.createJobBuildReturns.result2
	}
}

func (fake *FakeSchedulerDB) CreateJobBuildCallCount() int {
	fake.createJobBuildMutex.RLock()
	defer fake.createJobBuildMutex.RUnlock()
	return len(fake.createJobBuildArgsForCall)
}

func (fake *FakeSchedulerDB) CreateJobBuildArgsForCall(i int) string {
	fake.createJobBuildMutex.RLock()
	defer fake.createJobBuildMutex.RUnlock()
	return fake.createJobBuildArgsForCall[i].job
}

func (fake *FakeSchedulerDB) CreateJobBuildReturns(result1 db.Build, result2 error) {
	fake.CreateJobBuildStub = nil
	fake.createJobBuildReturns = struct {
		result1 db.Build
		result2 error
	}{result1, result2}
}

func (fake *FakeSchedulerDB) EnsurePendingBuildExists(jobName string) error {
	fake.ensurePendingBuildExistsMutex.Lock()
	fake.ensurePendingBuildExistsArgsForCall = append(fake.ensurePendingBuildExistsArgsForCall, struct {
		jobName string
	}{jobName})
	fake.recordInvocation("EnsurePendingBuildExists", []interface{}{jobName})
	fake.ensurePendingBuildExistsMutex.Unlock()
	if fake.EnsurePendingBuildExistsStub != nil {
		return fake.EnsurePendingBuildExistsStub(jobName)
	} else {
		return fake.ensurePendingBuildExistsReturns.result1
	}
}

func (fake *FakeSchedulerDB) EnsurePendingBuildExistsCallCount() int {
	fake.ensurePendingBuildExistsMutex.RLock()
	defer fake.ensurePendingBuildExistsMutex.RUnlock()
	return len(fake.ensurePendingBuildExistsArgsForCall)
}

func (fake *FakeSchedulerDB) EnsurePendingBuildExistsArgsForCall(i int) string {
	fake.ensurePendingBuildExistsMutex.RLock()
	defer fake.ensurePendingBuildExistsMutex.RUnlock()
	return fake.ensurePendingBuildExistsArgsForCall[i].jobName
}

func (fake *FakeSchedulerDB) EnsurePendingBuildExistsReturns(result1 error) {
	fake.EnsurePendingBuildExistsStub = nil
	fake.ensurePendingBuildExistsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSchedulerDB) AcquireResourceCheckingForJobLock(logger lager.Logger, job string) (db.Lock, bool, error) {
	fake.acquireResourceCheckingForJobLockMutex.Lock()
	fake.acquireResourceCheckingForJobLockArgsForCall = append(fake.acquireResourceCheckingForJobLockArgsForCall, struct {
		logger lager.Logger
		job    string
	}{logger, job})
	fake.recordInvocation("AcquireResourceCheckingForJobLock", []interface{}{logger, job})
	fake.acquireResourceCheckingForJobLockMutex.Unlock()
	if fake.AcquireResourceCheckingForJobLockStub != nil {
		return fake.AcquireResourceCheckingForJobLockStub(logger, job)
	} else {
		return fake.acquireResourceCheckingForJobLockReturns.result1, fake.acquireResourceCheckingForJobLockReturns.result2, fake.acquireResourceCheckingForJobLockReturns.result3
	}
}

func (fake *FakeSchedulerDB) AcquireResourceCheckingForJobLockCallCount() int {
	fake.acquireResourceCheckingForJobLockMutex.RLock()
	defer fake.acquireResourceCheckingForJobLockMutex.RUnlock()
	return len(fake.acquireResourceCheckingForJobLockArgsForCall)
}

func (fake *FakeSchedulerDB) AcquireResourceCheckingForJobLockArgsForCall(i int) (lager.Logger, string) {
	fake.acquireResourceCheckingForJobLockMutex.RLock()
	defer fake.acquireResourceCheckingForJobLockMutex.RUnlock()
	return fake.acquireResourceCheckingForJobLockArgsForCall[i].logger, fake.acquireResourceCheckingForJobLockArgsForCall[i].job
}

func (fake *FakeSchedulerDB) AcquireResourceCheckingForJobLockReturns(result1 db.Lock, result2 bool, result3 error) {
	fake.AcquireResourceCheckingForJobLockStub = nil
	fake.acquireResourceCheckingForJobLockReturns = struct {
		result1 db.Lock
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSchedulerDB) GetAllPendingBuilds() (map[string][]db.Build, error) {
	fake.getAllPendingBuildsMutex.Lock()
	fake.getAllPendingBuildsArgsForCall = append(fake.getAllPendingBuildsArgsForCall, struct{}{})
	fake.recordInvocation("GetAllPendingBuilds", []interface{}{})
	fake.getAllPendingBuildsMutex.Unlock()
	if fake.GetAllPendingBuildsStub != nil {
		return fake.GetAllPendingBuildsStub()
	} else {
		return fake.getAllPendingBuildsReturns.result1, fake.getAllPendingBuildsReturns.result2
	}
}

func (fake *FakeSchedulerDB) GetAllPendingBuildsCallCount() int {
	fake.getAllPendingBuildsMutex.RLock()
	defer fake.getAllPendingBuildsMutex.RUnlock()
	return len(fake.getAllPendingBuildsArgsForCall)
}

func (fake *FakeSchedulerDB) GetAllPendingBuildsReturns(result1 map[string][]db.Build, result2 error) {
	fake.GetAllPendingBuildsStub = nil
	fake.getAllPendingBuildsReturns = struct {
		result1 map[string][]db.Build
		result2 error
	}{result1, result2}
}

func (fake *FakeSchedulerDB) GetPendingBuildsForJob(jobName string) ([]db.Build, error) {
	fake.getPendingBuildsForJobMutex.Lock()
	fake.getPendingBuildsForJobArgsForCall = append(fake.getPendingBuildsForJobArgsForCall, struct {
		jobName string
	}{jobName})
	fake.recordInvocation("GetPendingBuildsForJob", []interface{}{jobName})
	fake.getPendingBuildsForJobMutex.Unlock()
	if fake.GetPendingBuildsForJobStub != nil {
		return fake.GetPendingBuildsForJobStub(jobName)
	} else {
		return fake.getPendingBuildsForJobReturns.result1, fake.getPendingBuildsForJobReturns.result2
	}
}

func (fake *FakeSchedulerDB) GetPendingBuildsForJobCallCount() int {
	fake.getPendingBuildsForJobMutex.RLock()
	defer fake.getPendingBuildsForJobMutex.RUnlock()
	return len(fake.getPendingBuildsForJobArgsForCall)
}

func (fake *FakeSchedulerDB) GetPendingBuildsForJobArgsForCall(i int) string {
	fake.getPendingBuildsForJobMutex.RLock()
	defer fake.getPendingBuildsForJobMutex.RUnlock()
	return fake.getPendingBuildsForJobArgsForCall[i].jobName
}

func (fake *FakeSchedulerDB) GetPendingBuildsForJobReturns(result1 []db.Build, result2 error) {
	fake.GetPendingBuildsForJobStub = nil
	fake.getPendingBuildsForJobReturns = struct {
		result1 []db.Build
		result2 error
	}{result1, result2}
}

func (fake *FakeSchedulerDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.acquireSchedulingLockMutex.RLock()
	defer fake.acquireSchedulingLockMutex.RUnlock()
	fake.loadVersionsDBMutex.RLock()
	defer fake.loadVersionsDBMutex.RUnlock()
	fake.getPipelineNameMutex.RLock()
	defer fake.getPipelineNameMutex.RUnlock()
	fake.reloadMutex.RLock()
	defer fake.reloadMutex.RUnlock()
	fake.configMutex.RLock()
	defer fake.configMutex.RUnlock()
	fake.createJobBuildMutex.RLock()
	defer fake.createJobBuildMutex.RUnlock()
	fake.ensurePendingBuildExistsMutex.RLock()
	defer fake.ensurePendingBuildExistsMutex.RUnlock()
	fake.acquireResourceCheckingForJobLockMutex.RLock()
	defer fake.acquireResourceCheckingForJobLockMutex.RUnlock()
	fake.getAllPendingBuildsMutex.RLock()
	defer fake.getAllPendingBuildsMutex.RUnlock()
	fake.getPendingBuildsForJobMutex.RLock()
	defer fake.getPendingBuildsForJobMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeSchedulerDB) recordInvocation(key string, args []interface{}) {
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

var _ scheduler.SchedulerDB = new(FakeSchedulerDB)
