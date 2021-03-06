// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"

import mock "github.com/stretchr/testify/mock"
import models "github.com/siddhiparekh11/GoChallenge/models"

// IPlayer is an autogenerated mock type for the IPlayer type
type IPlayer struct {
	mock.Mock
}

// CreatePlayer provides a mock function with given fields: ctx, player
func (_m *IPlayer) CreatePlayer(ctx context.Context, player models.Player) (bool, *models.QError) {
	ret := _m.Called(ctx, player)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, models.Player) bool); ok {
		r0 = rf(ctx, player)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 *models.QError
	if rf, ok := ret.Get(1).(func(context.Context, models.Player) *models.QError); ok {
		r1 = rf(ctx, player)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*models.QError)
		}
	}

	return r0, r1
}

// GetPlayer provides a mock function with given fields: ctx, playerId
func (_m *IPlayer) GetPlayer(ctx context.Context, playerId int) (*models.Player, *models.QError) {
	ret := _m.Called(ctx, playerId)

	var r0 *models.Player
	if rf, ok := ret.Get(0).(func(context.Context, int) *models.Player); ok {
		r0 = rf(ctx, playerId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Player)
		}
	}

	var r1 *models.QError
	if rf, ok := ret.Get(1).(func(context.Context, int) *models.QError); ok {
		r1 = rf(ctx, playerId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*models.QError)
		}
	}

	return r0, r1
}

// GetPlayers provides a mock function with given fields: ctx
func (_m *IPlayer) GetPlayers(ctx context.Context) ([]*models.Player, *models.QError) {
	ret := _m.Called(ctx)

	var r0 []*models.Player
	if rf, ok := ret.Get(0).(func(context.Context) []*models.Player); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Player)
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
