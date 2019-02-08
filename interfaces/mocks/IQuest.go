// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"

import mock "github.com/stretchr/testify/mock"
import models "github.com/siddhiparekh11/GoChallenge/models"

// IQuest is an autogenerated mock type for the IQuest type
type IQuest struct {
	mock.Mock
}

// CreateQuest provides a mock function with given fields: ctx, quest
func (_m *IQuest) CreateQuest(ctx context.Context, quest models.Quest) (bool, *models.QError) {
	ret := _m.Called(ctx, quest)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, models.Quest) bool); ok {
		r0 = rf(ctx, quest)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 *models.QError
	if rf, ok := ret.Get(1).(func(context.Context, models.Quest) *models.QError); ok {
		r1 = rf(ctx, quest)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*models.QError)
		}
	}

	return r0, r1
}

// GetQuest provides a mock function with given fields: ctx, questId
func (_m *IQuest) GetQuest(ctx context.Context, questId int) (*models.Quest, *models.QError) {
	ret := _m.Called(ctx, questId)

	var r0 *models.Quest
	if rf, ok := ret.Get(0).(func(context.Context, int) *models.Quest); ok {
		r0 = rf(ctx, questId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Quest)
		}
	}

	var r1 *models.QError
	if rf, ok := ret.Get(1).(func(context.Context, int) *models.QError); ok {
		r1 = rf(ctx, questId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*models.QError)
		}
	}

	return r0, r1
}

// GetQuests provides a mock function with given fields: ctx
func (_m *IQuest) GetQuests(ctx context.Context) ([]*models.Quest, *models.QError) {
	ret := _m.Called(ctx)

	var r0 []*models.Quest
	if rf, ok := ret.Get(0).(func(context.Context) []*models.Quest); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Quest)
		}
	}

	var r1 *models.QError
	if rf, ok := ret.Get(1).(func(context.Context) *models.QError); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*models.QError)
		}
	}

	return r0, r1
}