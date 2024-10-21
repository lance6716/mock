// Code generated by MockGen. DO NOT EDIT.
// Source: external.go
//
// Generated by this command:
//
//	mockgen --source=external.go --destination=source/mock_external_mock.go --package source
//

// Package source is a generated GoMock package.
package source

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	generics "go.uber.org/mock/mockgen/internal/tests/generics"
	other "go.uber.org/mock/mockgen/internal/tests/generics/other"
	constraints "golang.org/x/exp/constraints"
)

// MockExternalConstraint is a mock of ExternalConstraint interface.
type MockExternalConstraint[I constraints.Integer, F any] struct {
	ctrl     *gomock.Controller
	recorder *MockExternalConstraintMockRecorder[I, F]
	isgomock struct{}
}

// MockExternalConstraintMockRecorder is the mock recorder for MockExternalConstraint.
type MockExternalConstraintMockRecorder[I constraints.Integer, F any] struct {
	mock *MockExternalConstraint[I, F]
}

// NewMockExternalConstraint creates a new mock instance.
func NewMockExternalConstraint[I constraints.Integer, F any](ctrl *gomock.Controller) *MockExternalConstraint[I, F] {
	mock := &MockExternalConstraint[I, F]{ctrl: ctrl}
	mock.recorder = &MockExternalConstraintMockRecorder[I, F]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExternalConstraint[I, F]) EXPECT() *MockExternalConstraintMockRecorder[I, F] {
	return m.recorder
}

// Eight mocks base method.
func (m *MockExternalConstraint[I, F]) Eight(arg0 F) other.Two[I, F] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Eight", arg0)
	ret0, _ := ret[0].(other.Two[I, F])
	return ret0
}

// Eight indicates an expected call of Eight.
func (mr *MockExternalConstraintMockRecorder[I, F]) Eight(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Eight", reflect.TypeOf((*MockExternalConstraint[I, F])(nil).Eight), arg0)
}

// Eleven mocks base method.
func (m *MockExternalConstraint[I, F]) Eleven() map[string]I {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Eleven")
	ret0, _ := ret[0].(map[string]I)
	return ret0
}

// Eleven indicates an expected call of Eleven.
func (mr *MockExternalConstraintMockRecorder[I, F]) Eleven() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Eleven", reflect.TypeOf((*MockExternalConstraint[I, F])(nil).Eleven))
}

// Five mocks base method.
func (m *MockExternalConstraint[I, F]) Five(arg0 I) generics.Baz[F] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Five", arg0)
	ret0, _ := ret[0].(generics.Baz[F])
	return ret0
}

// Five indicates an expected call of Five.
func (mr *MockExternalConstraintMockRecorder[I, F]) Five(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Five", reflect.TypeOf((*MockExternalConstraint[I, F])(nil).Five), arg0)
}

// Four mocks base method.
func (m *MockExternalConstraint[I, F]) Four(arg0 I) generics.Foo[I, F] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Four", arg0)
	ret0, _ := ret[0].(generics.Foo[I, F])
	return ret0
}

// Four indicates an expected call of Four.
func (mr *MockExternalConstraintMockRecorder[I, F]) Four(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Four", reflect.TypeOf((*MockExternalConstraint[I, F])(nil).Four), arg0)
}

// Nine mocks base method.
func (m *MockExternalConstraint[I, F]) Nine(arg0 generics.Iface[I]) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Nine", arg0)
}

// Nine indicates an expected call of Nine.
func (mr *MockExternalConstraintMockRecorder[I, F]) Nine(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nine", reflect.TypeOf((*MockExternalConstraint[I, F])(nil).Nine), arg0)
}

// One mocks base method.
func (m *MockExternalConstraint[I, F]) One(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// One indicates an expected call of One.
func (mr *MockExternalConstraintMockRecorder[I, F]) One(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockExternalConstraint[I, F])(nil).One), arg0)
}

// Seven mocks base method.
func (m *MockExternalConstraint[I, F]) Seven(arg0 I) other.One[I] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Seven", arg0)
	ret0, _ := ret[0].(other.One[I])
	return ret0
}

// Seven indicates an expected call of Seven.
func (mr *MockExternalConstraintMockRecorder[I, F]) Seven(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seven", reflect.TypeOf((*MockExternalConstraint[I, F])(nil).Seven), arg0)
}

// Six mocks base method.
func (m *MockExternalConstraint[I, F]) Six(arg0 I) *generics.Baz[F] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Six", arg0)
	ret0, _ := ret[0].(*generics.Baz[F])
	return ret0
}

// Six indicates an expected call of Six.
func (mr *MockExternalConstraintMockRecorder[I, F]) Six(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Six", reflect.TypeOf((*MockExternalConstraint[I, F])(nil).Six), arg0)
}

// Ten mocks base method.
func (m *MockExternalConstraint[I, F]) Ten(arg0 *I) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Ten", arg0)
}

// Ten indicates an expected call of Ten.
func (mr *MockExternalConstraintMockRecorder[I, F]) Ten(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ten", reflect.TypeOf((*MockExternalConstraint[I, F])(nil).Ten), arg0)
}

// Thirteen mocks base method.
func (m *MockExternalConstraint[I, F]) Thirteen(arg0 ...I) *F {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Thirteen", varargs...)
	ret0, _ := ret[0].(*F)
	return ret0
}

// Thirteen indicates an expected call of Thirteen.
func (mr *MockExternalConstraintMockRecorder[I, F]) Thirteen(arg0 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Thirteen", reflect.TypeOf((*MockExternalConstraint[I, F])(nil).Thirteen), arg0...)
}

// Three mocks base method.
func (m *MockExternalConstraint[I, F]) Three(arg0 I) F {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Three", arg0)
	ret0, _ := ret[0].(F)
	return ret0
}

// Three indicates an expected call of Three.
func (mr *MockExternalConstraintMockRecorder[I, F]) Three(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Three", reflect.TypeOf((*MockExternalConstraint[I, F])(nil).Three), arg0)
}

// Twelve mocks base method.
func (m *MockExternalConstraint[I, F]) Twelve(ctx context.Context) <-chan []I {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Twelve", ctx)
	ret0, _ := ret[0].(<-chan []I)
	return ret0
}

// Twelve indicates an expected call of Twelve.
func (mr *MockExternalConstraintMockRecorder[I, F]) Twelve(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Twelve", reflect.TypeOf((*MockExternalConstraint[I, F])(nil).Twelve), ctx)
}

// Two mocks base method.
func (m *MockExternalConstraint[I, F]) Two(arg0 I) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Two", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Two indicates an expected call of Two.
func (mr *MockExternalConstraintMockRecorder[I, F]) Two(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Two", reflect.TypeOf((*MockExternalConstraint[I, F])(nil).Two), arg0)
}

// MockEmbeddingIface is a mock of EmbeddingIface interface.
type MockEmbeddingIface[T constraints.Integer, R constraints.Float] struct {
	ctrl     *gomock.Controller
	recorder *MockEmbeddingIfaceMockRecorder[T, R]
	isgomock struct{}
}

// MockEmbeddingIfaceMockRecorder is the mock recorder for MockEmbeddingIface.
type MockEmbeddingIfaceMockRecorder[T constraints.Integer, R constraints.Float] struct {
	mock *MockEmbeddingIface[T, R]
}

// NewMockEmbeddingIface creates a new mock instance.
func NewMockEmbeddingIface[T constraints.Integer, R constraints.Float](ctrl *gomock.Controller) *MockEmbeddingIface[T, R] {
	mock := &MockEmbeddingIface[T, R]{ctrl: ctrl}
	mock.recorder = &MockEmbeddingIfaceMockRecorder[T, R]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmbeddingIface[T, R]) EXPECT() *MockEmbeddingIfaceMockRecorder[T, R] {
	return m.recorder
}

// Eight mocks base method.
func (m *MockEmbeddingIface[T, R]) Eight(arg0 R) other.Two[T, R] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Eight", arg0)
	ret0, _ := ret[0].(other.Two[T, R])
	return ret0
}

// Eight indicates an expected call of Eight.
func (mr *MockEmbeddingIfaceMockRecorder[T, R]) Eight(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Eight", reflect.TypeOf((*MockEmbeddingIface[T, R])(nil).Eight), arg0)
}

// Eleven mocks base method.
func (m *MockEmbeddingIface[T, R]) Eleven() map[string]T {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Eleven")
	ret0, _ := ret[0].(map[string]T)
	return ret0
}

// Eleven indicates an expected call of Eleven.
func (mr *MockEmbeddingIfaceMockRecorder[T, R]) Eleven() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Eleven", reflect.TypeOf((*MockEmbeddingIface[T, R])(nil).Eleven))
}

// First mocks base method.
func (m *MockEmbeddingIface[T, R]) First() R {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "First")
	ret0, _ := ret[0].(R)
	return ret0
}

// First indicates an expected call of First.
func (mr *MockEmbeddingIfaceMockRecorder[T, R]) First() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "First", reflect.TypeOf((*MockEmbeddingIface[T, R])(nil).First))
}

// Five mocks base method.
func (m *MockEmbeddingIface[T, R]) Five(arg0 T) generics.Baz[R] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Five", arg0)
	ret0, _ := ret[0].(generics.Baz[R])
	return ret0
}

// Five indicates an expected call of Five.
func (mr *MockEmbeddingIfaceMockRecorder[T, R]) Five(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Five", reflect.TypeOf((*MockEmbeddingIface[T, R])(nil).Five), arg0)
}

// Four mocks base method.
func (m *MockEmbeddingIface[T, R]) Four(arg0 T) generics.Foo[T, R] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Four", arg0)
	ret0, _ := ret[0].(generics.Foo[T, R])
	return ret0
}

// Four indicates an expected call of Four.
func (mr *MockEmbeddingIfaceMockRecorder[T, R]) Four(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Four", reflect.TypeOf((*MockEmbeddingIface[T, R])(nil).Four), arg0)
}

// Fourth mocks base method.
func (m *MockEmbeddingIface[T, R]) Fourth() generics.Generator[T] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fourth")
	ret0, _ := ret[0].(generics.Generator[T])
	return ret0
}

// Fourth indicates an expected call of Fourth.
func (mr *MockEmbeddingIfaceMockRecorder[T, R]) Fourth() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fourth", reflect.TypeOf((*MockEmbeddingIface[T, R])(nil).Fourth))
}

// Generate mocks base method.
func (m *MockEmbeddingIface[T, R]) Generate() R {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate")
	ret0, _ := ret[0].(R)
	return ret0
}

// Generate indicates an expected call of Generate.
func (mr *MockEmbeddingIfaceMockRecorder[T, R]) Generate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockEmbeddingIface[T, R])(nil).Generate))
}

// Nine mocks base method.
func (m *MockEmbeddingIface[T, R]) Nine(arg0 generics.Iface[T]) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Nine", arg0)
}

// Nine indicates an expected call of Nine.
func (mr *MockEmbeddingIfaceMockRecorder[T, R]) Nine(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nine", reflect.TypeOf((*MockEmbeddingIface[T, R])(nil).Nine), arg0)
}

// One mocks base method.
func (m *MockEmbeddingIface[T, R]) One(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// One indicates an expected call of One.
func (mr *MockEmbeddingIfaceMockRecorder[T, R]) One(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockEmbeddingIface[T, R])(nil).One), arg0)
}

// Read mocks base method.
func (m *MockEmbeddingIface[T, R]) Read(p []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", p)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockEmbeddingIfaceMockRecorder[T, R]) Read(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockEmbeddingIface[T, R])(nil).Read), p)
}

// Second mocks base method.
func (m *MockEmbeddingIface[T, R]) Second() generics.StructType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Second")
	ret0, _ := ret[0].(generics.StructType)
	return ret0
}

// Second indicates an expected call of Second.
func (mr *MockEmbeddingIfaceMockRecorder[T, R]) Second() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Second", reflect.TypeOf((*MockEmbeddingIface[T, R])(nil).Second))
}

// Seven mocks base method.
func (m *MockEmbeddingIface[T, R]) Seven(arg0 T) other.One[T] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Seven", arg0)
	ret0, _ := ret[0].(other.One[T])
	return ret0
}

// Seven indicates an expected call of Seven.
func (mr *MockEmbeddingIfaceMockRecorder[T, R]) Seven(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seven", reflect.TypeOf((*MockEmbeddingIface[T, R])(nil).Seven), arg0)
}

// Six mocks base method.
func (m *MockEmbeddingIface[T, R]) Six(arg0 T) *generics.Baz[R] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Six", arg0)
	ret0, _ := ret[0].(*generics.Baz[R])
	return ret0
}

// Six indicates an expected call of Six.
func (mr *MockEmbeddingIfaceMockRecorder[T, R]) Six(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Six", reflect.TypeOf((*MockEmbeddingIface[T, R])(nil).Six), arg0)
}

// Ten mocks base method.
func (m *MockEmbeddingIface[T, R]) Ten(arg0 *T) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Ten", arg0)
}

// Ten indicates an expected call of Ten.
func (mr *MockEmbeddingIfaceMockRecorder[T, R]) Ten(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ten", reflect.TypeOf((*MockEmbeddingIface[T, R])(nil).Ten), arg0)
}

// Third mocks base method.
func (m *MockEmbeddingIface[T, R]) Third() other.Five {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Third")
	ret0, _ := ret[0].(other.Five)
	return ret0
}

// Third indicates an expected call of Third.
func (mr *MockEmbeddingIfaceMockRecorder[T, R]) Third() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Third", reflect.TypeOf((*MockEmbeddingIface[T, R])(nil).Third))
}

// Thirteen mocks base method.
func (m *MockEmbeddingIface[T, R]) Thirteen(arg0 ...T) *R {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Thirteen", varargs...)
	ret0, _ := ret[0].(*R)
	return ret0
}

// Thirteen indicates an expected call of Thirteen.
func (mr *MockEmbeddingIfaceMockRecorder[T, R]) Thirteen(arg0 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Thirteen", reflect.TypeOf((*MockEmbeddingIface[T, R])(nil).Thirteen), arg0...)
}

// Three mocks base method.
func (m *MockEmbeddingIface[T, R]) Three(arg0 T) R {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Three", arg0)
	ret0, _ := ret[0].(R)
	return ret0
}

// Three indicates an expected call of Three.
func (mr *MockEmbeddingIfaceMockRecorder[T, R]) Three(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Three", reflect.TypeOf((*MockEmbeddingIface[T, R])(nil).Three), arg0)
}

// Twelve mocks base method.
func (m *MockEmbeddingIface[T, R]) Twelve(ctx context.Context) <-chan []T {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Twelve", ctx)
	ret0, _ := ret[0].(<-chan []T)
	return ret0
}

// Twelve indicates an expected call of Twelve.
func (mr *MockEmbeddingIfaceMockRecorder[T, R]) Twelve(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Twelve", reflect.TypeOf((*MockEmbeddingIface[T, R])(nil).Twelve), ctx)
}

// Two mocks base method.
func (m *MockEmbeddingIface[T, R]) Two(arg0 T) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Two", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Two indicates an expected call of Two.
func (mr *MockEmbeddingIfaceMockRecorder[T, R]) Two(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Two", reflect.TypeOf((*MockEmbeddingIface[T, R])(nil).Two), arg0)
}

// Water mocks base method.
func (m *MockEmbeddingIface[T, R]) Water(arg0 generics.Generator[T]) []generics.Generator[T] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Water", arg0)
	ret0, _ := ret[0].([]generics.Generator[T])
	return ret0
}

// Water indicates an expected call of Water.
func (mr *MockEmbeddingIfaceMockRecorder[T, R]) Water(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Water", reflect.TypeOf((*MockEmbeddingIface[T, R])(nil).Water), arg0)
}

// MockGenerator is a mock of Generator interface.
type MockGenerator[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockGeneratorMockRecorder[T]
	isgomock struct{}
}

// MockGeneratorMockRecorder is the mock recorder for MockGenerator.
type MockGeneratorMockRecorder[T any] struct {
	mock *MockGenerator[T]
}

// NewMockGenerator creates a new mock instance.
func NewMockGenerator[T any](ctrl *gomock.Controller) *MockGenerator[T] {
	mock := &MockGenerator[T]{ctrl: ctrl}
	mock.recorder = &MockGeneratorMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGenerator[T]) EXPECT() *MockGeneratorMockRecorder[T] {
	return m.recorder
}

// Generate mocks base method.
func (m *MockGenerator[T]) Generate() T {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate")
	ret0, _ := ret[0].(T)
	return ret0
}

// Generate indicates an expected call of Generate.
func (mr *MockGeneratorMockRecorder[T]) Generate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockGenerator[T])(nil).Generate))
}

// MockGroup is a mock of Group interface.
type MockGroup[T generics.Generator[any]] struct {
	ctrl     *gomock.Controller
	recorder *MockGroupMockRecorder[T]
	isgomock struct{}
}

// MockGroupMockRecorder is the mock recorder for MockGroup.
type MockGroupMockRecorder[T generics.Generator[any]] struct {
	mock *MockGroup[T]
}

// NewMockGroup creates a new mock instance.
func NewMockGroup[T generics.Generator[any]](ctrl *gomock.Controller) *MockGroup[T] {
	mock := &MockGroup[T]{ctrl: ctrl}
	mock.recorder = &MockGroupMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGroup[T]) EXPECT() *MockGroupMockRecorder[T] {
	return m.recorder
}

// Join mocks base method.
func (m *MockGroup[T]) Join(ctx context.Context) []T {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Join", ctx)
	ret0, _ := ret[0].([]T)
	return ret0
}

// Join indicates an expected call of Join.
func (mr *MockGroupMockRecorder[T]) Join(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Join", reflect.TypeOf((*MockGroup[T])(nil).Join), ctx)
}
