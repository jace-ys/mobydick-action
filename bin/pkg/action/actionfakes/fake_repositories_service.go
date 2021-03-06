// Code generated by counterfeiter. DO NOT EDIT.
package actionfakes

import (
	"context"
	"sync"

	"github.com/google/go-github/v29/github"
	"github.com/jace-ys/mobydick-action/bin/pkg/action"
)

type FakeRepositoriesService struct {
	CreateFileStub        func(context.Context, string, string, string, *github.RepositoryContentFileOptions) (*github.RepositoryContentResponse, *github.Response, error)
	createFileMutex       sync.RWMutex
	createFileArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
		arg5 *github.RepositoryContentFileOptions
	}
	createFileReturns struct {
		result1 *github.RepositoryContentResponse
		result2 *github.Response
		result3 error
	}
	createFileReturnsOnCall map[int]struct {
		result1 *github.RepositoryContentResponse
		result2 *github.Response
		result3 error
	}
	ListByOrgStub        func(context.Context, string, *github.RepositoryListByOrgOptions) ([]*github.Repository, *github.Response, error)
	listByOrgMutex       sync.RWMutex
	listByOrgArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 *github.RepositoryListByOrgOptions
	}
	listByOrgReturns struct {
		result1 []*github.Repository
		result2 *github.Response
		result3 error
	}
	listByOrgReturnsOnCall map[int]struct {
		result1 []*github.Repository
		result2 *github.Response
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRepositoriesService) CreateFile(arg1 context.Context, arg2 string, arg3 string, arg4 string, arg5 *github.RepositoryContentFileOptions) (*github.RepositoryContentResponse, *github.Response, error) {
	fake.createFileMutex.Lock()
	ret, specificReturn := fake.createFileReturnsOnCall[len(fake.createFileArgsForCall)]
	fake.createFileArgsForCall = append(fake.createFileArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
		arg5 *github.RepositoryContentFileOptions
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("CreateFile", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.createFileMutex.Unlock()
	if fake.CreateFileStub != nil {
		return fake.CreateFileStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.createFileReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeRepositoriesService) CreateFileCallCount() int {
	fake.createFileMutex.RLock()
	defer fake.createFileMutex.RUnlock()
	return len(fake.createFileArgsForCall)
}

func (fake *FakeRepositoriesService) CreateFileCalls(stub func(context.Context, string, string, string, *github.RepositoryContentFileOptions) (*github.RepositoryContentResponse, *github.Response, error)) {
	fake.createFileMutex.Lock()
	defer fake.createFileMutex.Unlock()
	fake.CreateFileStub = stub
}

func (fake *FakeRepositoriesService) CreateFileArgsForCall(i int) (context.Context, string, string, string, *github.RepositoryContentFileOptions) {
	fake.createFileMutex.RLock()
	defer fake.createFileMutex.RUnlock()
	argsForCall := fake.createFileArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeRepositoriesService) CreateFileReturns(result1 *github.RepositoryContentResponse, result2 *github.Response, result3 error) {
	fake.createFileMutex.Lock()
	defer fake.createFileMutex.Unlock()
	fake.CreateFileStub = nil
	fake.createFileReturns = struct {
		result1 *github.RepositoryContentResponse
		result2 *github.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRepositoriesService) CreateFileReturnsOnCall(i int, result1 *github.RepositoryContentResponse, result2 *github.Response, result3 error) {
	fake.createFileMutex.Lock()
	defer fake.createFileMutex.Unlock()
	fake.CreateFileStub = nil
	if fake.createFileReturnsOnCall == nil {
		fake.createFileReturnsOnCall = make(map[int]struct {
			result1 *github.RepositoryContentResponse
			result2 *github.Response
			result3 error
		})
	}
	fake.createFileReturnsOnCall[i] = struct {
		result1 *github.RepositoryContentResponse
		result2 *github.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRepositoriesService) ListByOrg(arg1 context.Context, arg2 string, arg3 *github.RepositoryListByOrgOptions) ([]*github.Repository, *github.Response, error) {
	fake.listByOrgMutex.Lock()
	ret, specificReturn := fake.listByOrgReturnsOnCall[len(fake.listByOrgArgsForCall)]
	fake.listByOrgArgsForCall = append(fake.listByOrgArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 *github.RepositoryListByOrgOptions
	}{arg1, arg2, arg3})
	fake.recordInvocation("ListByOrg", []interface{}{arg1, arg2, arg3})
	fake.listByOrgMutex.Unlock()
	if fake.ListByOrgStub != nil {
		return fake.ListByOrgStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.listByOrgReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeRepositoriesService) ListByOrgCallCount() int {
	fake.listByOrgMutex.RLock()
	defer fake.listByOrgMutex.RUnlock()
	return len(fake.listByOrgArgsForCall)
}

func (fake *FakeRepositoriesService) ListByOrgCalls(stub func(context.Context, string, *github.RepositoryListByOrgOptions) ([]*github.Repository, *github.Response, error)) {
	fake.listByOrgMutex.Lock()
	defer fake.listByOrgMutex.Unlock()
	fake.ListByOrgStub = stub
}

func (fake *FakeRepositoriesService) ListByOrgArgsForCall(i int) (context.Context, string, *github.RepositoryListByOrgOptions) {
	fake.listByOrgMutex.RLock()
	defer fake.listByOrgMutex.RUnlock()
	argsForCall := fake.listByOrgArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeRepositoriesService) ListByOrgReturns(result1 []*github.Repository, result2 *github.Response, result3 error) {
	fake.listByOrgMutex.Lock()
	defer fake.listByOrgMutex.Unlock()
	fake.ListByOrgStub = nil
	fake.listByOrgReturns = struct {
		result1 []*github.Repository
		result2 *github.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRepositoriesService) ListByOrgReturnsOnCall(i int, result1 []*github.Repository, result2 *github.Response, result3 error) {
	fake.listByOrgMutex.Lock()
	defer fake.listByOrgMutex.Unlock()
	fake.ListByOrgStub = nil
	if fake.listByOrgReturnsOnCall == nil {
		fake.listByOrgReturnsOnCall = make(map[int]struct {
			result1 []*github.Repository
			result2 *github.Response
			result3 error
		})
	}
	fake.listByOrgReturnsOnCall[i] = struct {
		result1 []*github.Repository
		result2 *github.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRepositoriesService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createFileMutex.RLock()
	defer fake.createFileMutex.RUnlock()
	fake.listByOrgMutex.RLock()
	defer fake.listByOrgMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRepositoriesService) recordInvocation(key string, args []interface{}) {
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

var _ action.RepositoriesService = new(FakeRepositoriesService)
