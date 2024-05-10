// Code generated by mockery v2.36.1. DO NOT EDIT.

package mock_envoy

import mock "github.com/stretchr/testify/mock"

// BufferInstance is an autogenerated mock type for the BufferInstance type
type BufferInstance struct {
	mock.Mock
}

type BufferInstance_Expecter struct {
	mock *mock.Mock
}

func (_m *BufferInstance) EXPECT() *BufferInstance_Expecter {
	return &BufferInstance_Expecter{mock: &_m.Mock}
}

// Append provides a mock function with given fields: data
func (_m *BufferInstance) Append(data []byte) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BufferInstance_Append_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Append'
type BufferInstance_Append_Call struct {
	*mock.Call
}

// Append is a helper method to define mock.On call
//   - data []byte
func (_e *BufferInstance_Expecter) Append(data interface{}) *BufferInstance_Append_Call {
	return &BufferInstance_Append_Call{Call: _e.mock.On("Append", data)}
}

func (_c *BufferInstance_Append_Call) Run(run func(data []byte)) *BufferInstance_Append_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *BufferInstance_Append_Call) Return(_a0 error) *BufferInstance_Append_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BufferInstance_Append_Call) RunAndReturn(run func([]byte) error) *BufferInstance_Append_Call {
	_c.Call.Return(run)
	return _c
}

// AppendString provides a mock function with given fields: s
func (_m *BufferInstance) AppendString(s string) error {
	ret := _m.Called(s)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BufferInstance_AppendString_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AppendString'
type BufferInstance_AppendString_Call struct {
	*mock.Call
}

// AppendString is a helper method to define mock.On call
//   - s string
func (_e *BufferInstance_Expecter) AppendString(s interface{}) *BufferInstance_AppendString_Call {
	return &BufferInstance_AppendString_Call{Call: _e.mock.On("AppendString", s)}
}

func (_c *BufferInstance_AppendString_Call) Run(run func(s string)) *BufferInstance_AppendString_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *BufferInstance_AppendString_Call) Return(_a0 error) *BufferInstance_AppendString_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BufferInstance_AppendString_Call) RunAndReturn(run func(string) error) *BufferInstance_AppendString_Call {
	_c.Call.Return(run)
	return _c
}

// Bytes provides a mock function with given fields:
func (_m *BufferInstance) Bytes() []byte {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// BufferInstance_Bytes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Bytes'
type BufferInstance_Bytes_Call struct {
	*mock.Call
}

// Bytes is a helper method to define mock.On call
func (_e *BufferInstance_Expecter) Bytes() *BufferInstance_Bytes_Call {
	return &BufferInstance_Bytes_Call{Call: _e.mock.On("Bytes")}
}

func (_c *BufferInstance_Bytes_Call) Run(run func()) *BufferInstance_Bytes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BufferInstance_Bytes_Call) Return(_a0 []byte) *BufferInstance_Bytes_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BufferInstance_Bytes_Call) RunAndReturn(run func() []byte) *BufferInstance_Bytes_Call {
	_c.Call.Return(run)
	return _c
}

// Drain provides a mock function with given fields: offset
func (_m *BufferInstance) Drain(offset int) {
	_m.Called(offset)
}

// BufferInstance_Drain_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Drain'
type BufferInstance_Drain_Call struct {
	*mock.Call
}

// Drain is a helper method to define mock.On call
//   - offset int
func (_e *BufferInstance_Expecter) Drain(offset interface{}) *BufferInstance_Drain_Call {
	return &BufferInstance_Drain_Call{Call: _e.mock.On("Drain", offset)}
}

func (_c *BufferInstance_Drain_Call) Run(run func(offset int)) *BufferInstance_Drain_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *BufferInstance_Drain_Call) Return() *BufferInstance_Drain_Call {
	_c.Call.Return()
	return _c
}

func (_c *BufferInstance_Drain_Call) RunAndReturn(run func(int)) *BufferInstance_Drain_Call {
	_c.Call.Return(run)
	return _c
}

// Len provides a mock function with given fields:
func (_m *BufferInstance) Len() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// BufferInstance_Len_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Len'
type BufferInstance_Len_Call struct {
	*mock.Call
}

// Len is a helper method to define mock.On call
func (_e *BufferInstance_Expecter) Len() *BufferInstance_Len_Call {
	return &BufferInstance_Len_Call{Call: _e.mock.On("Len")}
}

func (_c *BufferInstance_Len_Call) Run(run func()) *BufferInstance_Len_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BufferInstance_Len_Call) Return(_a0 int) *BufferInstance_Len_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BufferInstance_Len_Call) RunAndReturn(run func() int) *BufferInstance_Len_Call {
	_c.Call.Return(run)
	return _c
}

// Prepend provides a mock function with given fields: data
func (_m *BufferInstance) Prepend(data []byte) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BufferInstance_Prepend_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Prepend'
type BufferInstance_Prepend_Call struct {
	*mock.Call
}

// Prepend is a helper method to define mock.On call
//   - data []byte
func (_e *BufferInstance_Expecter) Prepend(data interface{}) *BufferInstance_Prepend_Call {
	return &BufferInstance_Prepend_Call{Call: _e.mock.On("Prepend", data)}
}

func (_c *BufferInstance_Prepend_Call) Run(run func(data []byte)) *BufferInstance_Prepend_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *BufferInstance_Prepend_Call) Return(_a0 error) *BufferInstance_Prepend_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BufferInstance_Prepend_Call) RunAndReturn(run func([]byte) error) *BufferInstance_Prepend_Call {
	_c.Call.Return(run)
	return _c
}

// PrependString provides a mock function with given fields: s
func (_m *BufferInstance) PrependString(s string) error {
	ret := _m.Called(s)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BufferInstance_PrependString_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PrependString'
type BufferInstance_PrependString_Call struct {
	*mock.Call
}

// PrependString is a helper method to define mock.On call
//   - s string
func (_e *BufferInstance_Expecter) PrependString(s interface{}) *BufferInstance_PrependString_Call {
	return &BufferInstance_PrependString_Call{Call: _e.mock.On("PrependString", s)}
}

func (_c *BufferInstance_PrependString_Call) Run(run func(s string)) *BufferInstance_PrependString_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *BufferInstance_PrependString_Call) Return(_a0 error) *BufferInstance_PrependString_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BufferInstance_PrependString_Call) RunAndReturn(run func(string) error) *BufferInstance_PrependString_Call {
	_c.Call.Return(run)
	return _c
}

// Reset provides a mock function with given fields:
func (_m *BufferInstance) Reset() {
	_m.Called()
}

// BufferInstance_Reset_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Reset'
type BufferInstance_Reset_Call struct {
	*mock.Call
}

// Reset is a helper method to define mock.On call
func (_e *BufferInstance_Expecter) Reset() *BufferInstance_Reset_Call {
	return &BufferInstance_Reset_Call{Call: _e.mock.On("Reset")}
}

func (_c *BufferInstance_Reset_Call) Run(run func()) *BufferInstance_Reset_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BufferInstance_Reset_Call) Return() *BufferInstance_Reset_Call {
	_c.Call.Return()
	return _c
}

func (_c *BufferInstance_Reset_Call) RunAndReturn(run func()) *BufferInstance_Reset_Call {
	_c.Call.Return(run)
	return _c
}

// Set provides a mock function with given fields: _a0
func (_m *BufferInstance) Set(_a0 []byte) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BufferInstance_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type BufferInstance_Set_Call struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - _a0 []byte
func (_e *BufferInstance_Expecter) Set(_a0 interface{}) *BufferInstance_Set_Call {
	return &BufferInstance_Set_Call{Call: _e.mock.On("Set", _a0)}
}

func (_c *BufferInstance_Set_Call) Run(run func(_a0 []byte)) *BufferInstance_Set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *BufferInstance_Set_Call) Return(_a0 error) *BufferInstance_Set_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BufferInstance_Set_Call) RunAndReturn(run func([]byte) error) *BufferInstance_Set_Call {
	_c.Call.Return(run)
	return _c
}

// SetString provides a mock function with given fields: _a0
func (_m *BufferInstance) SetString(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BufferInstance_SetString_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetString'
type BufferInstance_SetString_Call struct {
	*mock.Call
}

// SetString is a helper method to define mock.On call
//   - _a0 string
func (_e *BufferInstance_Expecter) SetString(_a0 interface{}) *BufferInstance_SetString_Call {
	return &BufferInstance_SetString_Call{Call: _e.mock.On("SetString", _a0)}
}

func (_c *BufferInstance_SetString_Call) Run(run func(_a0 string)) *BufferInstance_SetString_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *BufferInstance_SetString_Call) Return(_a0 error) *BufferInstance_SetString_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BufferInstance_SetString_Call) RunAndReturn(run func(string) error) *BufferInstance_SetString_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *BufferInstance) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// BufferInstance_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type BufferInstance_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *BufferInstance_Expecter) String() *BufferInstance_String_Call {
	return &BufferInstance_String_Call{Call: _e.mock.On("String")}
}

func (_c *BufferInstance_String_Call) Run(run func()) *BufferInstance_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BufferInstance_String_Call) Return(_a0 string) *BufferInstance_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BufferInstance_String_Call) RunAndReturn(run func() string) *BufferInstance_String_Call {
	_c.Call.Return(run)
	return _c
}

// Write provides a mock function with given fields: p
func (_m *BufferInstance) Write(p []byte) (int, error) {
	ret := _m.Called(p)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) (int, error)); ok {
		return rf(p)
	}
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BufferInstance_Write_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Write'
type BufferInstance_Write_Call struct {
	*mock.Call
}

// Write is a helper method to define mock.On call
//   - p []byte
func (_e *BufferInstance_Expecter) Write(p interface{}) *BufferInstance_Write_Call {
	return &BufferInstance_Write_Call{Call: _e.mock.On("Write", p)}
}

func (_c *BufferInstance_Write_Call) Run(run func(p []byte)) *BufferInstance_Write_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *BufferInstance_Write_Call) Return(n int, err error) *BufferInstance_Write_Call {
	_c.Call.Return(n, err)
	return _c
}

func (_c *BufferInstance_Write_Call) RunAndReturn(run func([]byte) (int, error)) *BufferInstance_Write_Call {
	_c.Call.Return(run)
	return _c
}

// WriteByte provides a mock function with given fields: p
func (_m *BufferInstance) WriteByte(p byte) error {
	ret := _m.Called(p)

	var r0 error
	if rf, ok := ret.Get(0).(func(byte) error); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BufferInstance_WriteByte_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteByte'
type BufferInstance_WriteByte_Call struct {
	*mock.Call
}

// WriteByte is a helper method to define mock.On call
//   - p byte
func (_e *BufferInstance_Expecter) WriteByte(p interface{}) *BufferInstance_WriteByte_Call {
	return &BufferInstance_WriteByte_Call{Call: _e.mock.On("WriteByte", p)}
}

func (_c *BufferInstance_WriteByte_Call) Run(run func(p byte)) *BufferInstance_WriteByte_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(byte))
	})
	return _c
}

func (_c *BufferInstance_WriteByte_Call) Return(_a0 error) *BufferInstance_WriteByte_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BufferInstance_WriteByte_Call) RunAndReturn(run func(byte) error) *BufferInstance_WriteByte_Call {
	_c.Call.Return(run)
	return _c
}

// WriteString provides a mock function with given fields: s
func (_m *BufferInstance) WriteString(s string) (int, error) {
	ret := _m.Called(s)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (int, error)); ok {
		return rf(s)
	}
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(s)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BufferInstance_WriteString_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteString'
type BufferInstance_WriteString_Call struct {
	*mock.Call
}

// WriteString is a helper method to define mock.On call
//   - s string
func (_e *BufferInstance_Expecter) WriteString(s interface{}) *BufferInstance_WriteString_Call {
	return &BufferInstance_WriteString_Call{Call: _e.mock.On("WriteString", s)}
}

func (_c *BufferInstance_WriteString_Call) Run(run func(s string)) *BufferInstance_WriteString_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *BufferInstance_WriteString_Call) Return(n int, err error) *BufferInstance_WriteString_Call {
	_c.Call.Return(n, err)
	return _c
}

func (_c *BufferInstance_WriteString_Call) RunAndReturn(run func(string) (int, error)) *BufferInstance_WriteString_Call {
	_c.Call.Return(run)
	return _c
}

// WriteUint16 provides a mock function with given fields: p
func (_m *BufferInstance) WriteUint16(p uint16) error {
	ret := _m.Called(p)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint16) error); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BufferInstance_WriteUint16_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteUint16'
type BufferInstance_WriteUint16_Call struct {
	*mock.Call
}

// WriteUint16 is a helper method to define mock.On call
//   - p uint16
func (_e *BufferInstance_Expecter) WriteUint16(p interface{}) *BufferInstance_WriteUint16_Call {
	return &BufferInstance_WriteUint16_Call{Call: _e.mock.On("WriteUint16", p)}
}

func (_c *BufferInstance_WriteUint16_Call) Run(run func(p uint16)) *BufferInstance_WriteUint16_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint16))
	})
	return _c
}

func (_c *BufferInstance_WriteUint16_Call) Return(_a0 error) *BufferInstance_WriteUint16_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BufferInstance_WriteUint16_Call) RunAndReturn(run func(uint16) error) *BufferInstance_WriteUint16_Call {
	_c.Call.Return(run)
	return _c
}

// WriteUint32 provides a mock function with given fields: p
func (_m *BufferInstance) WriteUint32(p uint32) error {
	ret := _m.Called(p)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint32) error); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BufferInstance_WriteUint32_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteUint32'
type BufferInstance_WriteUint32_Call struct {
	*mock.Call
}

// WriteUint32 is a helper method to define mock.On call
//   - p uint32
func (_e *BufferInstance_Expecter) WriteUint32(p interface{}) *BufferInstance_WriteUint32_Call {
	return &BufferInstance_WriteUint32_Call{Call: _e.mock.On("WriteUint32", p)}
}

func (_c *BufferInstance_WriteUint32_Call) Run(run func(p uint32)) *BufferInstance_WriteUint32_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint32))
	})
	return _c
}

func (_c *BufferInstance_WriteUint32_Call) Return(_a0 error) *BufferInstance_WriteUint32_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BufferInstance_WriteUint32_Call) RunAndReturn(run func(uint32) error) *BufferInstance_WriteUint32_Call {
	_c.Call.Return(run)
	return _c
}

// WriteUint64 provides a mock function with given fields: p
func (_m *BufferInstance) WriteUint64(p uint64) error {
	ret := _m.Called(p)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BufferInstance_WriteUint64_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteUint64'
type BufferInstance_WriteUint64_Call struct {
	*mock.Call
}

// WriteUint64 is a helper method to define mock.On call
//   - p uint64
func (_e *BufferInstance_Expecter) WriteUint64(p interface{}) *BufferInstance_WriteUint64_Call {
	return &BufferInstance_WriteUint64_Call{Call: _e.mock.On("WriteUint64", p)}
}

func (_c *BufferInstance_WriteUint64_Call) Run(run func(p uint64)) *BufferInstance_WriteUint64_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint64))
	})
	return _c
}

func (_c *BufferInstance_WriteUint64_Call) Return(_a0 error) *BufferInstance_WriteUint64_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BufferInstance_WriteUint64_Call) RunAndReturn(run func(uint64) error) *BufferInstance_WriteUint64_Call {
	_c.Call.Return(run)
	return _c
}

// NewBufferInstance creates a new instance of BufferInstance. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBufferInstance(t interface {
	mock.TestingT
	Cleanup(func())
}) *BufferInstance {
	mock := &BufferInstance{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
