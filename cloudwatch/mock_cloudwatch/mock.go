// Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/awslabs/amazon-cloudwatch-logs-for-fluent-bit/cloudwatch (interfaces: CloudWatchLogsClient)

// Package mock_cloudwatch is a generated GoMock package.
package mock_cloudwatch

import (
	reflect "reflect"

	cloudwatchlogs "github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	gomock "github.com/golang/mock/gomock"
)

// MockCloudWatchLogsClient is a mock of CloudWatchLogsClient interface
type MockCloudWatchLogsClient struct {
	ctrl     *gomock.Controller
	recorder *MockCloudWatchLogsClientMockRecorder
}

// MockCloudWatchLogsClientMockRecorder is the mock recorder for MockCloudWatchLogsClient
type MockCloudWatchLogsClientMockRecorder struct {
	mock *MockCloudWatchLogsClient
}

// NewMockCloudWatchLogsClient creates a new mock instance
func NewMockCloudWatchLogsClient(ctrl *gomock.Controller) *MockCloudWatchLogsClient {
	mock := &MockCloudWatchLogsClient{ctrl: ctrl}
	mock.recorder = &MockCloudWatchLogsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCloudWatchLogsClient) EXPECT() *MockCloudWatchLogsClientMockRecorder {
	return m.recorder
}

// CreateLogGroup mocks base method
func (m *MockCloudWatchLogsClient) CreateLogGroup(arg0 *cloudwatchlogs.CreateLogGroupInput) (*cloudwatchlogs.CreateLogGroupOutput, error) {
	ret := m.ctrl.Call(m, "CreateLogGroup", arg0)
	ret0, _ := ret[0].(*cloudwatchlogs.CreateLogGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLogGroup indicates an expected call of CreateLogGroup
func (mr *MockCloudWatchLogsClientMockRecorder) CreateLogGroup(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLogGroup", reflect.TypeOf((*MockCloudWatchLogsClient)(nil).CreateLogGroup), arg0)
}

// CreateLogStream mocks base method
func (m *MockCloudWatchLogsClient) CreateLogStream(arg0 *cloudwatchlogs.CreateLogStreamInput) (*cloudwatchlogs.CreateLogStreamOutput, error) {
	ret := m.ctrl.Call(m, "CreateLogStream", arg0)
	ret0, _ := ret[0].(*cloudwatchlogs.CreateLogStreamOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLogStream indicates an expected call of CreateLogStream
func (mr *MockCloudWatchLogsClientMockRecorder) CreateLogStream(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLogStream", reflect.TypeOf((*MockCloudWatchLogsClient)(nil).CreateLogStream), arg0)
}

// PutLogEvents mocks base method
func (m *MockCloudWatchLogsClient) PutLogEvents(arg0 *cloudwatchlogs.PutLogEventsInput) (*cloudwatchlogs.PutLogEventsOutput, error) {
	ret := m.ctrl.Call(m, "PutLogEvents", arg0)
	ret0, _ := ret[0].(*cloudwatchlogs.PutLogEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutLogEvents indicates an expected call of PutLogEvents
func (mr *MockCloudWatchLogsClientMockRecorder) PutLogEvents(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutLogEvents", reflect.TypeOf((*MockCloudWatchLogsClient)(nil).PutLogEvents), arg0)
}
