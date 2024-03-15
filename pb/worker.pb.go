// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: worker.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WorkerState int32

const (
	WorkerState_WORKER_STATE_UNSPECIFIED WorkerState = 0
	WorkerState_WORKER_STATE_UNAVAILABLE WorkerState = 1
	WorkerState_WORKER_STATE_AVAILABLE   WorkerState = 2
	WorkerState_WORKER_STATE_WORKING     WorkerState = 3
)

// Enum value maps for WorkerState.
var (
	WorkerState_name = map[int32]string{
		0: "WORKER_STATE_UNSPECIFIED",
		1: "WORKER_STATE_UNAVAILABLE",
		2: "WORKER_STATE_AVAILABLE",
		3: "WORKER_STATE_WORKING",
	}
	WorkerState_value = map[string]int32{
		"WORKER_STATE_UNSPECIFIED": 0,
		"WORKER_STATE_UNAVAILABLE": 1,
		"WORKER_STATE_AVAILABLE":   2,
		"WORKER_STATE_WORKING":     3,
	}
)

func (x WorkerState) Enum() *WorkerState {
	p := new(WorkerState)
	*p = x
	return p
}

func (x WorkerState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorkerState) Descriptor() protoreflect.EnumDescriptor {
	return file_worker_proto_enumTypes[0].Descriptor()
}

func (WorkerState) Type() protoreflect.EnumType {
	return &file_worker_proto_enumTypes[0]
}

func (x WorkerState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorkerState.Descriptor instead.
func (WorkerState) EnumDescriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{0}
}

type JobControl int32

const (
	JobControl_JOB_CONTROL_NONE     JobControl = 0
	JobControl_JOB_CONTROL_CONTINUE JobControl = 1
	JobControl_JOB_CONTROL_NEW      JobControl = 2
	JobControl_JOB_CONTROL_CANCEL   JobControl = 4
)

// Enum value maps for JobControl.
var (
	JobControl_name = map[int32]string{
		0: "JOB_CONTROL_NONE",
		1: "JOB_CONTROL_CONTINUE",
		2: "JOB_CONTROL_NEW",
		4: "JOB_CONTROL_CANCEL",
	}
	JobControl_value = map[string]int32{
		"JOB_CONTROL_NONE":     0,
		"JOB_CONTROL_CONTINUE": 1,
		"JOB_CONTROL_NEW":      2,
		"JOB_CONTROL_CANCEL":   4,
	}
)

func (x JobControl) Enum() *JobControl {
	p := new(JobControl)
	*p = x
	return p
}

func (x JobControl) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobControl) Descriptor() protoreflect.EnumDescriptor {
	return file_worker_proto_enumTypes[1].Descriptor()
}

func (JobControl) Type() protoreflect.EnumType {
	return &file_worker_proto_enumTypes[1]
}

func (x JobControl) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobControl.Descriptor instead.
func (JobControl) EnumDescriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{1}
}

type RegisterOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *RegisterOptions) Reset() {
	*x = RegisterOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterOptions) ProtoMessage() {}

func (x *RegisterOptions) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterOptions.ProtoReflect.Descriptor instead.
func (*RegisterOptions) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterOptions) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type RegisterResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Registered bool   `protobuf:"varint,1,opt,name=registered,proto3" json:"registered,omitempty"`
	WorkerId   string `protobuf:"bytes,2,opt,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty"`
}

func (x *RegisterResult) Reset() {
	*x = RegisterResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResult) ProtoMessage() {}

func (x *RegisterResult) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResult.ProtoReflect.Descriptor instead.
func (*RegisterResult) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResult) GetRegistered() bool {
	if x != nil {
		return x.Registered
	}
	return false
}

func (x *RegisterResult) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

type DeregisterOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkerId string `protobuf:"bytes,1,opt,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty"`
}

func (x *DeregisterOptions) Reset() {
	*x = DeregisterOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeregisterOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeregisterOptions) ProtoMessage() {}

func (x *DeregisterOptions) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeregisterOptions.ProtoReflect.Descriptor instead.
func (*DeregisterOptions) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{2}
}

func (x *DeregisterOptions) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

type DeregisterResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deregistered bool `protobuf:"varint,1,opt,name=deregistered,proto3" json:"deregistered,omitempty"`
}

func (x *DeregisterResult) Reset() {
	*x = DeregisterResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeregisterResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeregisterResult) ProtoMessage() {}

func (x *DeregisterResult) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeregisterResult.ProtoReflect.Descriptor instead.
func (*DeregisterResult) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{3}
}

func (x *DeregisterResult) GetDeregistered() bool {
	if x != nil {
		return x.Deregistered
	}
	return false
}

type StatusOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkerId    string       `protobuf:"bytes,1,opt,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty"`
	WorkerState WorkerState  `protobuf:"varint,2,opt,name=worker_state,json=workerState,proto3,enum=pb.WorkerState" json:"worker_state,omitempty"`
	JobStatus   []*JobStatus `protobuf:"bytes,3,rep,name=job_status,json=jobStatus,proto3" json:"job_status,omitempty"`
}

func (x *StatusOptions) Reset() {
	*x = StatusOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusOptions) ProtoMessage() {}

func (x *StatusOptions) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusOptions.ProtoReflect.Descriptor instead.
func (*StatusOptions) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{4}
}

func (x *StatusOptions) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

func (x *StatusOptions) GetWorkerState() WorkerState {
	if x != nil {
		return x.WorkerState
	}
	return WorkerState_WORKER_STATE_UNSPECIFIED
}

func (x *StatusOptions) GetJobStatus() []*JobStatus {
	if x != nil {
		return x.JobStatus
	}
	return nil
}

type StatusResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobControl JobControl `protobuf:"varint,1,opt,name=job_control,json=jobControl,proto3,enum=pb.JobControl" json:"job_control,omitempty"`
	Job        *JobSpec   `protobuf:"bytes,2,opt,name=job,proto3,oneof" json:"job,omitempty"`
}

func (x *StatusResult) Reset() {
	*x = StatusResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResult) ProtoMessage() {}

func (x *StatusResult) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResult.ProtoReflect.Descriptor instead.
func (*StatusResult) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{5}
}

func (x *StatusResult) GetJobControl() JobControl {
	if x != nil {
		return x.JobControl
	}
	return JobControl_JOB_CONTROL_NONE
}

func (x *StatusResult) GetJob() *JobSpec {
	if x != nil {
		return x.Job
	}
	return nil
}

var File_worker_proto protoreflect.FileDescriptor

var file_worker_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x1a, 0x09, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a,
	0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x4d, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x11, 0x44, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x22, 0x36, 0x0a, 0x10, 0x44, 0x65, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x64,
	0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0c, 0x64, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x22,
	0x8e, 0x01, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x12, 0x32,
	0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x2c, 0x0a, 0x0a, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x09, 0x6a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x6b, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x2f, 0x0a, 0x0b, 0x6a, 0x6f, 0x62, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x0a, 0x6a, 0x6f, 0x62, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x12, 0x22, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x70, 0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x70, 0x65, 0x63, 0x48, 0x00, 0x52, 0x03, 0x6a,
	0x6f, 0x62, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6a, 0x6f, 0x62, 0x2a, 0x7f, 0x0a,
	0x0b, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x18,
	0x57, 0x4f, 0x52, 0x4b, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x57, 0x4f,
	0x52, 0x4b, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x41, 0x56, 0x41,
	0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x57, 0x4f, 0x52, 0x4b,
	0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42,
	0x4c, 0x45, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x57, 0x4f, 0x52, 0x4b, 0x45, 0x52, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x2a, 0x69,
	0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x14, 0x0a, 0x10,
	0x4a, 0x4f, 0x42, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x5f, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x4a, 0x4f, 0x42, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f,
	0x4c, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x49, 0x4e, 0x55, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f,
	0x4a, 0x4f, 0x42, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x5f, 0x4e, 0x45, 0x57, 0x10,
	0x02, 0x12, 0x16, 0x0a, 0x12, 0x4a, 0x4f, 0x42, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c,
	0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x10, 0x04, 0x42, 0x07, 0x5a, 0x05, 0x74, 0x71, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_worker_proto_rawDescOnce sync.Once
	file_worker_proto_rawDescData = file_worker_proto_rawDesc
)

func file_worker_proto_rawDescGZIP() []byte {
	file_worker_proto_rawDescOnce.Do(func() {
		file_worker_proto_rawDescData = protoimpl.X.CompressGZIP(file_worker_proto_rawDescData)
	})
	return file_worker_proto_rawDescData
}

var file_worker_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_worker_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_worker_proto_goTypes = []interface{}{
	(WorkerState)(0),          // 0: pb.WorkerState
	(JobControl)(0),           // 1: pb.JobControl
	(*RegisterOptions)(nil),   // 2: pb.RegisterOptions
	(*RegisterResult)(nil),    // 3: pb.RegisterResult
	(*DeregisterOptions)(nil), // 4: pb.DeregisterOptions
	(*DeregisterResult)(nil),  // 5: pb.DeregisterResult
	(*StatusOptions)(nil),     // 6: pb.StatusOptions
	(*StatusResult)(nil),      // 7: pb.StatusResult
	(*JobStatus)(nil),         // 8: pb.JobStatus
	(*JobSpec)(nil),           // 9: pb.JobSpec
}
var file_worker_proto_depIdxs = []int32{
	0, // 0: pb.StatusOptions.worker_state:type_name -> pb.WorkerState
	8, // 1: pb.StatusOptions.job_status:type_name -> pb.JobStatus
	1, // 2: pb.StatusResult.job_control:type_name -> pb.JobControl
	9, // 3: pb.StatusResult.job:type_name -> pb.JobSpec
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_worker_proto_init() }
func file_worker_proto_init() {
	if File_worker_proto != nil {
		return
	}
	file_job_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_worker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_worker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_worker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeregisterOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_worker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeregisterResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_worker_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_worker_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_worker_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_worker_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_worker_proto_goTypes,
		DependencyIndexes: file_worker_proto_depIdxs,
		EnumInfos:         file_worker_proto_enumTypes,
		MessageInfos:      file_worker_proto_msgTypes,
	}.Build()
	File_worker_proto = out.File
	file_worker_proto_rawDesc = nil
	file_worker_proto_goTypes = nil
	file_worker_proto_depIdxs = nil
}
