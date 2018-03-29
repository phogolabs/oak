// This file was generated by counterfeiter
package fake

import (
	"sync"

	"github.com/phogolabs/gom/migration"
)

type MigrationGenerator struct {
	CreateStub        func(m *migration.Item) (string, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		m *migration.Item
	}
	createReturns struct {
		result1 string
		result2 error
	}
	WriteStub        func(m *migration.Item, content *migration.Content) error
	writeMutex       sync.RWMutex
	writeArgsForCall []struct {
		m       *migration.Item
		content *migration.Content
	}
	writeReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *MigrationGenerator) Create(m *migration.Item) (string, error) {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		m *migration.Item
	}{m})
	fake.recordInvocation("Create", []interface{}{m})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(m)
	}
	return fake.createReturns.result1, fake.createReturns.result2
}

func (fake *MigrationGenerator) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *MigrationGenerator) CreateArgsForCall(i int) *migration.Item {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].m
}

func (fake *MigrationGenerator) CreateReturns(result1 string, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *MigrationGenerator) Write(m *migration.Item, content *migration.Content) error {
	fake.writeMutex.Lock()
	fake.writeArgsForCall = append(fake.writeArgsForCall, struct {
		m       *migration.Item
		content *migration.Content
	}{m, content})
	fake.recordInvocation("Write", []interface{}{m, content})
	fake.writeMutex.Unlock()
	if fake.WriteStub != nil {
		return fake.WriteStub(m, content)
	}
	return fake.writeReturns.result1
}

func (fake *MigrationGenerator) WriteCallCount() int {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return len(fake.writeArgsForCall)
}

func (fake *MigrationGenerator) WriteArgsForCall(i int) (*migration.Item, *migration.Content) {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return fake.writeArgsForCall[i].m, fake.writeArgsForCall[i].content
}

func (fake *MigrationGenerator) WriteReturns(result1 error) {
	fake.WriteStub = nil
	fake.writeReturns = struct {
		result1 error
	}{result1}
}

func (fake *MigrationGenerator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return fake.invocations
}

func (fake *MigrationGenerator) recordInvocation(key string, args []interface{}) {
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

var _ migration.FileGenerator = new(MigrationGenerator)
