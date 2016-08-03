package mocks

import (
	"github.com/opencontrol/compliance-masonry/lib"
	"github.com/opencontrol/compliance-masonry/lib/common"
	"github.com/stretchr/testify/mock"
)

// Workspace is an autogenerated mock type for the Workspace type
type Workspace struct {
	mock.Mock
}

// GetAllComponents provides a mock function with given fields:
func (_m *Workspace) GetAllComponents() []common.Component {
	ret := _m.Called()

	var r0 []common.Component
	if rf, ok := ret.Get(0).(func() []common.Component); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.Component)
		}
	}

	return r0
}

// GetCertification provides a mock function with given fields:
func (_m *Workspace) GetCertification() common.Certification {
	ret := _m.Called()

	var r0 common.Certification
	if rf, ok := ret.Get(0).(func() common.Certification); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Certification)
		}
	}

	return r0
}

// GetComponent provides a mock function with given fields: _a0
func (_m *Workspace) GetComponent(_a0 string) common.Component {
	ret := _m.Called(_a0)

	var r0 common.Component
	if rf, ok := ret.Get(0).(func(string) common.Component); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Component)
		}
	}

	return r0
}

// GetJustification provides a mock function with given fields: _a0, _a1
func (_m *Workspace) GetJustification(_a0 string, _a1 string) lib.Verifications {
	ret := _m.Called(_a0, _a1)

	var r0 lib.Verifications
	if rf, ok := ret.Get(0).(func(string, string) lib.Verifications); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(lib.Verifications)
		}
	}

	return r0
}

// GetStandard provides a mock function with given fields: _a0
func (_m *Workspace) GetStandard(_a0 string) common.Standard {
	ret := _m.Called(_a0)

	var r0 common.Standard
	if rf, ok := ret.Get(0).(func(string) common.Standard); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Standard)
		}
	}

	return r0
}

// GetStandards provides a mock function with given fields:
func (_m *Workspace) GetStandards() []common.Standard {
	ret := _m.Called()

	var r0 []common.Standard
	if rf, ok := ret.Get(0).(func() []common.Standard); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.Standard)
		}
	}

	return r0
}

// LoadCertification provides a mock function with given fields: _a0
func (_m *Workspace) LoadCertification(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoadComponents provides a mock function with given fields: _a0
func (_m *Workspace) LoadComponents(_a0 string) []error {
	ret := _m.Called(_a0)

	var r0 []error
	if rf, ok := ret.Get(0).(func(string) []error); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]error)
		}
	}

	return r0
}

// LoadStandards provides a mock function with given fields: _a0
func (_m *Workspace) LoadStandards(_a0 string) []error {
	ret := _m.Called(_a0)

	var r0 []error
	if rf, ok := ret.Get(0).(func(string) []error); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]error)
		}
	}

	return r0
}
