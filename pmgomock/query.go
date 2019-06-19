// Automatically generated by MockGen. DO NOT EDIT!
// Source: /home/karl/golang/src/github.com/percona/pmgo/query.go

package pmgomock

import (
	time "time"

	"github.com/bukalapak/mgo"
	. "github.com/bukalapak/pmgo"
	gomock "github.com/golang/mock/gomock"
)

// Mock of QueryManager interface
type MockQueryManager struct {
	ctrl     *gomock.Controller
	recorder *_MockQueryManagerRecorder
}

// Recorder for MockQueryManager (not exported)
type _MockQueryManagerRecorder struct {
	mock *MockQueryManager
}

func NewMockQueryManager(ctrl *gomock.Controller) *MockQueryManager {
	mock := &MockQueryManager{ctrl: ctrl}
	mock.recorder = &_MockQueryManagerRecorder{mock}
	return mock
}

func (_m *MockQueryManager) EXPECT() *_MockQueryManagerRecorder {
	return _m.recorder
}

func (_m *MockQueryManager) All(result interface{}) error {
	ret := _m.ctrl.Call(_m, "All", result)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockQueryManagerRecorder) All(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "All", arg0)
}

func (_m *MockQueryManager) Apply(change mgo.Change, result interface{}) (*mgo.ChangeInfo, error) {
	ret := _m.ctrl.Call(_m, "Apply", change, result)
	ret0, _ := ret[0].(*mgo.ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockQueryManagerRecorder) Apply(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Apply", arg0, arg1)
}

func (_m *MockQueryManager) Batch(n int) QueryManager {
	ret := _m.ctrl.Call(_m, "Batch", n)
	ret0, _ := ret[0].(QueryManager)
	return ret0
}

func (_mr *_MockQueryManagerRecorder) Batch(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Batch", arg0)
}

func (_m *MockQueryManager) Comment(comment string) QueryManager {
	ret := _m.ctrl.Call(_m, "Comment", comment)
	ret0, _ := ret[0].(QueryManager)
	return ret0
}

func (_mr *_MockQueryManagerRecorder) Comment(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Comment", arg0)
}

func (_m *MockQueryManager) Count() (int, error) {
	ret := _m.ctrl.Call(_m, "Count")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockQueryManagerRecorder) Count() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Count")
}

func (_m *MockQueryManager) Distinct(key string, result interface{}) error {
	ret := _m.ctrl.Call(_m, "Distinct", key, result)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockQueryManagerRecorder) Distinct(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Distinct", arg0, arg1)
}

func (_m *MockQueryManager) Explain(result interface{}) error {
	ret := _m.ctrl.Call(_m, "Explain", result)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockQueryManagerRecorder) Explain(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Explain", arg0)
}

func (_m *MockQueryManager) For(result interface{}, f func() error) error {
	ret := _m.ctrl.Call(_m, "For", result, f)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockQueryManagerRecorder) For(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "For", arg0, arg1)
}

func (_m *MockQueryManager) Hint(indexKey ...string) QueryManager {
	_s := []interface{}{}
	for _, _x := range indexKey {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Hint", _s...)
	ret0, _ := ret[0].(QueryManager)
	return ret0
}

func (_mr *_MockQueryManagerRecorder) Hint(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Hint", arg0...)
}

func (_m *MockQueryManager) Iter() *mgo.Iter {
	ret := _m.ctrl.Call(_m, "Iter")
	ret0, _ := ret[0].(*mgo.Iter)
	return ret0
}

func (_mr *_MockQueryManagerRecorder) Iter() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Iter")
}

func (_m *MockQueryManager) Limit(n int) QueryManager {
	ret := _m.ctrl.Call(_m, "Limit", n)
	ret0, _ := ret[0].(QueryManager)
	return ret0
}

func (_mr *_MockQueryManagerRecorder) Limit(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Limit", arg0)
}

func (_m *MockQueryManager) LogReplay() QueryManager {
	ret := _m.ctrl.Call(_m, "LogReplay")
	ret0, _ := ret[0].(QueryManager)
	return ret0
}

func (_mr *_MockQueryManagerRecorder) LogReplay() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LogReplay")
}

func (_m *MockQueryManager) MapReduce(job *mgo.MapReduce, result interface{}) (*mgo.MapReduceInfo, error) {
	ret := _m.ctrl.Call(_m, "MapReduce", job, result)
	ret0, _ := ret[0].(*mgo.MapReduceInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockQueryManagerRecorder) MapReduce(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MapReduce", arg0, arg1)
}

func (_m *MockQueryManager) One(result interface{}) error {
	ret := _m.ctrl.Call(_m, "One", result)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockQueryManagerRecorder) One(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "One", arg0)
}

func (_m *MockQueryManager) Prefetch(p float64) QueryManager {
	ret := _m.ctrl.Call(_m, "Prefetch", p)
	ret0, _ := ret[0].(QueryManager)
	return ret0
}

func (_mr *_MockQueryManagerRecorder) Prefetch(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Prefetch", arg0)
}

func (_m *MockQueryManager) Select(selector interface{}) QueryManager {
	ret := _m.ctrl.Call(_m, "Select", selector)
	ret0, _ := ret[0].(QueryManager)
	return ret0
}

func (_mr *_MockQueryManagerRecorder) Select(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Select", arg0)
}

func (_m *MockQueryManager) SetMaxScan(n int) QueryManager {
	ret := _m.ctrl.Call(_m, "SetMaxScan", n)
	ret0, _ := ret[0].(QueryManager)
	return ret0
}

func (_mr *_MockQueryManagerRecorder) SetMaxScan(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetMaxScan", arg0)
}

func (_m *MockQueryManager) SetMaxTime(d time.Duration) QueryManager {
	ret := _m.ctrl.Call(_m, "SetMaxTime", d)
	ret0, _ := ret[0].(QueryManager)
	return ret0
}

func (_mr *_MockQueryManagerRecorder) SetMaxTime(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetMaxTime", arg0)
}

func (_m *MockQueryManager) Skip(n int) QueryManager {
	ret := _m.ctrl.Call(_m, "Skip", n)
	ret0, _ := ret[0].(QueryManager)
	return ret0
}

func (_mr *_MockQueryManagerRecorder) Skip(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Skip", arg0)
}

func (_m *MockQueryManager) Snapshot() QueryManager {
	ret := _m.ctrl.Call(_m, "Snapshot")
	ret0, _ := ret[0].(QueryManager)
	return ret0
}

func (_mr *_MockQueryManagerRecorder) Snapshot() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Snapshot")
}

func (_m *MockQueryManager) Sort(fields ...string) QueryManager {
	_s := []interface{}{}
	for _, _x := range fields {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Sort", _s...)
	ret0, _ := ret[0].(QueryManager)
	return ret0
}

func (_mr *_MockQueryManagerRecorder) Sort(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Sort", arg0...)
}

func (_m *MockQueryManager) Tail(timeout time.Duration) IterManager {
	ret := _m.ctrl.Call(_m, "Tail", timeout)
	ret0, _ := ret[0].(IterManager)
	return ret0
}

func (_mr *_MockQueryManagerRecorder) Tail(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Tail", arg0)
}
