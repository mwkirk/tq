// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: tq.proto

package pbuf

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

// ------------------------------------------------------------------
// Status request messages
// ------------------------------------------------------------------
type WorkerState int32

const (
	WorkerState_WORKER_STATE_UNAVAILABLE WorkerState = 0
	WorkerState_WORKER_STATE_AVAILABLE   WorkerState = 1
	WorkerState_WORKER_STATE_WORKING     WorkerState = 2
)

// Enum value maps for WorkerState.
var (
	WorkerState_name = map[int32]string{
		0: "WORKER_STATE_UNAVAILABLE",
		1: "WORKER_STATE_AVAILABLE",
		2: "WORKER_STATE_WORKING",
	}
	WorkerState_value = map[string]int32{
		"WORKER_STATE_UNAVAILABLE": 0,
		"WORKER_STATE_AVAILABLE":   1,
		"WORKER_STATE_WORKING":     2,
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
	return file_tq_proto_enumTypes[0].Descriptor()
}

func (WorkerState) Type() protoreflect.EnumType {
	return &file_tq_proto_enumTypes[0]
}

func (x WorkerState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorkerState.Descriptor instead.
func (WorkerState) EnumDescriptor() ([]byte, []int) {
	return file_tq_proto_rawDescGZIP(), []int{0}
}

type JobState int32

const (
	JobState_JOB_STATE_NONE        JobState = 0
	JobState_JOB_STATE_RUN         JobState = 1
	JobState_JOB_STATE_DONE_OK     JobState = 2
	JobState_JOB_STATE_DONE_ERR    JobState = 3
	JobState_JOB_STATE_DONE_CANCEL JobState = 4
)

// Enum value maps for JobState.
var (
	JobState_name = map[int32]string{
		0: "JOB_STATE_NONE",
		1: "JOB_STATE_RUN",
		2: "JOB_STATE_DONE_OK",
		3: "JOB_STATE_DONE_ERR",
		4: "JOB_STATE_DONE_CANCEL",
	}
	JobState_value = map[string]int32{
		"JOB_STATE_NONE":        0,
		"JOB_STATE_RUN":         1,
		"JOB_STATE_DONE_OK":     2,
		"JOB_STATE_DONE_ERR":    3,
		"JOB_STATE_DONE_CANCEL": 4,
	}
)

func (x JobState) Enum() *JobState {
	p := new(JobState)
	*p = x
	return p
}

func (x JobState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobState) Descriptor() protoreflect.EnumDescriptor {
	return file_tq_proto_enumTypes[1].Descriptor()
}

func (JobState) Type() protoreflect.EnumType {
	return &file_tq_proto_enumTypes[1]
}

func (x JobState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobState.Descriptor instead.
func (JobState) EnumDescriptor() ([]byte, []int) {
	return file_tq_proto_rawDescGZIP(), []int{1}
}

// ------------------------------------------------------------------
// Status response messages
// ------------------------------------------------------------------
type JobControl int32

const (
	JobControl_JOB_CONTROL_NONE     JobControl = 0
	JobControl_JOB_CONTROL_CONTINUE JobControl = 1
	JobControl_JOB_CONTROL_NEW      JobControl = 2
	JobControl_JOB_CONTROL_CANCEL   JobControl = 3
)

// Enum value maps for JobControl.
var (
	JobControl_name = map[int32]string{
		0: "JOB_CONTROL_NONE",
		1: "JOB_CONTROL_CONTINUE",
		2: "JOB_CONTROL_NEW",
		3: "JOB_CONTROL_CANCEL",
	}
	JobControl_value = map[string]int32{
		"JOB_CONTROL_NONE":     0,
		"JOB_CONTROL_CONTINUE": 1,
		"JOB_CONTROL_NEW":      2,
		"JOB_CONTROL_CANCEL":   3,
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
	return file_tq_proto_enumTypes[2].Descriptor()
}

func (JobControl) Type() protoreflect.EnumType {
	return &file_tq_proto_enumTypes[2]
}

func (x JobControl) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobControl.Descriptor instead.
func (JobControl) EnumDescriptor() ([]byte, []int) {
	return file_tq_proto_rawDescGZIP(), []int{2}
}

type JobKind int32

const (
	JobKind_JOB_KIND_NULL   JobKind = 0
	JobKind_JOB_KIND_TEST   JobKind = 1
	JobKind_JOB_KIND_SLEEP  JobKind = 2
	JobKind_JOB_KIND_FFMPEG JobKind = 3
)

// Enum value maps for JobKind.
var (
	JobKind_name = map[int32]string{
		0: "JOB_KIND_NULL",
		1: "JOB_KIND_TEST",
		2: "JOB_KIND_SLEEP",
		3: "JOB_KIND_FFMPEG",
	}
	JobKind_value = map[string]int32{
		"JOB_KIND_NULL":   0,
		"JOB_KIND_TEST":   1,
		"JOB_KIND_SLEEP":  2,
		"JOB_KIND_FFMPEG": 3,
	}
)

func (x JobKind) Enum() *JobKind {
	p := new(JobKind)
	*p = x
	return p
}

func (x JobKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobKind) Descriptor() protoreflect.EnumDescriptor {
	return file_tq_proto_enumTypes[3].Descriptor()
}

func (JobKind) Type() protoreflect.EnumType {
	return &file_tq_proto_enumTypes[3]
}

func (x JobKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobKind.Descriptor instead.
func (JobKind) EnumDescriptor() ([]byte, []int) {
	return file_tq_proto_rawDescGZIP(), []int{3}
}

// ------------------------------------------------------------------
// Register messages
// ------------------------------------------------------------------
type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_tq_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Registered bool   `protobuf:"varint,1,opt,name=registered,proto3" json:"registered,omitempty"`
	Id         string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tq_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tq_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_tq_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResponse) GetRegistered() bool {
	if x != nil {
		return x.Registered
	}
	return false
}

func (x *RegisterResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// ------------------------------------------------------------------
// Deregister messages
// ------------------------------------------------------------------
type DeregisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeregisterRequest) Reset() {
	*x = DeregisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tq_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeregisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeregisterRequest) ProtoMessage() {}

func (x *DeregisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tq_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeregisterRequest.ProtoReflect.Descriptor instead.
func (*DeregisterRequest) Descriptor() ([]byte, []int) {
	return file_tq_proto_rawDescGZIP(), []int{2}
}

func (x *DeregisterRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeregisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Registered bool `protobuf:"varint,1,opt,name=registered,proto3" json:"registered,omitempty"`
}

func (x *DeregisterResponse) Reset() {
	*x = DeregisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tq_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeregisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeregisterResponse) ProtoMessage() {}

func (x *DeregisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tq_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeregisterResponse.ProtoReflect.Descriptor instead.
func (*DeregisterResponse) Descriptor() ([]byte, []int) {
	return file_tq_proto_rawDescGZIP(), []int{3}
}

func (x *DeregisterResponse) GetRegistered() bool {
	if x != nil {
		return x.Registered
	}
	return false
}

type JobStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobState JobState `protobuf:"varint,1,opt,name=job_state,json=jobState,proto3,enum=pbuf.JobState" json:"job_state,omitempty"`
	Progress *float32 `protobuf:"fixed32,2,opt,name=progress,proto3,oneof" json:"progress,omitempty"`
	Msg      []string `protobuf:"bytes,3,rep,name=msg,proto3" json:"msg,omitempty"`
}

func (x *JobStatus) Reset() {
	*x = JobStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tq_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobStatus) ProtoMessage() {}

func (x *JobStatus) ProtoReflect() protoreflect.Message {
	mi := &file_tq_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobStatus.ProtoReflect.Descriptor instead.
func (*JobStatus) Descriptor() ([]byte, []int) {
	return file_tq_proto_rawDescGZIP(), []int{4}
}

func (x *JobStatus) GetJobState() JobState {
	if x != nil {
		return x.JobState
	}
	return JobState_JOB_STATE_NONE
}

func (x *JobStatus) GetProgress() float32 {
	if x != nil && x.Progress != nil {
		return *x.Progress
	}
	return 0
}

func (x *JobStatus) GetMsg() []string {
	if x != nil {
		return x.Msg
	}
	return nil
}

type StatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	WorkerState WorkerState `protobuf:"varint,2,opt,name=worker_state,json=workerState,proto3,enum=pbuf.WorkerState" json:"worker_state,omitempty"`
	JobStatus   *JobStatus  `protobuf:"bytes,3,opt,name=job_status,json=jobStatus,proto3,oneof" json:"job_status,omitempty"`
}

func (x *StatusRequest) Reset() {
	*x = StatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tq_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusRequest) ProtoMessage() {}

func (x *StatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tq_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusRequest.ProtoReflect.Descriptor instead.
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return file_tq_proto_rawDescGZIP(), []int{5}
}

func (x *StatusRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StatusRequest) GetWorkerState() WorkerState {
	if x != nil {
		return x.WorkerState
	}
	return WorkerState_WORKER_STATE_UNAVAILABLE
}

func (x *StatusRequest) GetJobStatus() *JobStatus {
	if x != nil {
		return x.JobStatus
	}
	return nil
}

type Job struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind  JobKind           `protobuf:"varint,1,opt,name=kind,proto3,enum=pbuf.JobKind" json:"kind,omitempty"`
	Num   int64             `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	Name  string            `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Parms map[string]string `protobuf:"bytes,4,rep,name=parms,proto3" json:"parms,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Job) Reset() {
	*x = Job{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tq_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_tq_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_tq_proto_rawDescGZIP(), []int{6}
}

func (x *Job) GetKind() JobKind {
	if x != nil {
		return x.Kind
	}
	return JobKind_JOB_KIND_NULL
}

func (x *Job) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *Job) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Job) GetParms() map[string]string {
	if x != nil {
		return x.Parms
	}
	return nil
}

type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobControl JobControl `protobuf:"varint,1,opt,name=job_control,json=jobControl,proto3,enum=pbuf.JobControl" json:"job_control,omitempty"`
	Job        *Job       `protobuf:"bytes,2,opt,name=job,proto3,oneof" json:"job,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tq_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tq_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_tq_proto_rawDescGZIP(), []int{7}
}

func (x *StatusResponse) GetJobControl() JobControl {
	if x != nil {
		return x.JobControl
	}
	return JobControl_JOB_CONTROL_NONE
}

func (x *StatusResponse) GetJob() *Job {
	if x != nil {
		return x.Job
	}
	return nil
}

var File_tq_proto protoreflect.FileDescriptor

var file_tq_proto_rawDesc = []byte{
	0x0a, 0x08, 0x74, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x62, 0x75, 0x66,
	0x22, 0x27, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x42, 0x0a, 0x10, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x23, 0x0a,
	0x11, 0x44, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x34, 0x0a, 0x12, 0x44, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x22, 0x78, 0x0a, 0x09, 0x4a, 0x6f, 0x62, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2b, 0x0a, 0x09, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x75, 0x66, 0x2e,
	0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x08, 0x6a, 0x6f, 0x62, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x1f, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x88, 0x01, 0x01, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x22, 0x99, 0x01, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x34, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x75,
	0x66, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0b, 0x77,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x6a, 0x6f,
	0x62, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x70, 0x62, 0x75, 0x66, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48,
	0x00, 0x52, 0x09, 0x6a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xb4,
	0x01, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x21, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x75, 0x66, 0x2e, 0x4a, 0x6f, 0x62, 0x4b,
	0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x2a, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x70, 0x62, 0x75, 0x66, 0x2e, 0x4a, 0x6f, 0x62, 0x2e, 0x50, 0x61, 0x72, 0x6d, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x70, 0x61, 0x72, 0x6d, 0x73, 0x1a, 0x38, 0x0a, 0x0a, 0x50,
	0x61, 0x72, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x6d, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0b, 0x6a, 0x6f, 0x62, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70,
	0x62, 0x75, 0x66, 0x2e, 0x4a, 0x6f, 0x62, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x0a,
	0x6a, 0x6f, 0x62, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x20, 0x0a, 0x03, 0x6a, 0x6f,
	0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x75, 0x66, 0x2e, 0x4a,
	0x6f, 0x62, 0x48, 0x00, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04,
	0x5f, 0x6a, 0x6f, 0x62, 0x2a, 0x61, 0x0a, 0x0b, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x57, 0x4f, 0x52, 0x4b, 0x45, 0x52, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10,
	0x00, 0x12, 0x1a, 0x0a, 0x16, 0x57, 0x4f, 0x52, 0x4b, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x18, 0x0a,
	0x14, 0x57, 0x4f, 0x52, 0x4b, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x57, 0x4f,
	0x52, 0x4b, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x2a, 0x7b, 0x0a, 0x08, 0x4a, 0x6f, 0x62, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x4a, 0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x4a, 0x4f, 0x42, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x55, 0x4e, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x4a, 0x4f,
	0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x44, 0x4f, 0x4e, 0x45, 0x5f, 0x4f, 0x4b, 0x10,
	0x02, 0x12, 0x16, 0x0a, 0x12, 0x4a, 0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x44,
	0x4f, 0x4e, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x4a, 0x4f, 0x42,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x44, 0x4f, 0x4e, 0x45, 0x5f, 0x43, 0x41, 0x4e, 0x43,
	0x45, 0x4c, 0x10, 0x04, 0x2a, 0x69, 0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x12, 0x14, 0x0a, 0x10, 0x4a, 0x4f, 0x42, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f,
	0x4c, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x4a, 0x4f, 0x42, 0x5f,
	0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x49, 0x4e, 0x55, 0x45,
	0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x4a, 0x4f, 0x42, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f,
	0x4c, 0x5f, 0x4e, 0x45, 0x57, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x4a, 0x4f, 0x42, 0x5f, 0x43,
	0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x10, 0x03, 0x2a,
	0x58, 0x0a, 0x07, 0x4a, 0x6f, 0x62, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x11, 0x0a, 0x0d, 0x4a, 0x4f,
	0x42, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x4e, 0x55, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x11, 0x0a,
	0x0d, 0x4a, 0x4f, 0x42, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x54, 0x45, 0x53, 0x54, 0x10, 0x01,
	0x12, 0x12, 0x0a, 0x0e, 0x4a, 0x4f, 0x42, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x53, 0x4c, 0x45,
	0x45, 0x50, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x4a, 0x4f, 0x42, 0x5f, 0x4b, 0x49, 0x4e, 0x44,
	0x5f, 0x46, 0x46, 0x4d, 0x50, 0x45, 0x47, 0x10, 0x03, 0x32, 0xbb, 0x01, 0x0a, 0x02, 0x74, 0x71,
	0x12, 0x3b, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x70,
	0x62, 0x75, 0x66, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a,
	0x0a, 0x44, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x70, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x65, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x35, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x70, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x74, 0x71, 0x2f, 0x70, 0x62,
	0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tq_proto_rawDescOnce sync.Once
	file_tq_proto_rawDescData = file_tq_proto_rawDesc
)

func file_tq_proto_rawDescGZIP() []byte {
	file_tq_proto_rawDescOnce.Do(func() {
		file_tq_proto_rawDescData = protoimpl.X.CompressGZIP(file_tq_proto_rawDescData)
	})
	return file_tq_proto_rawDescData
}

var file_tq_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_tq_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_tq_proto_goTypes = []interface{}{
	(WorkerState)(0),           // 0: pbuf.WorkerState
	(JobState)(0),              // 1: pbuf.JobState
	(JobControl)(0),            // 2: pbuf.JobControl
	(JobKind)(0),               // 3: pbuf.JobKind
	(*RegisterRequest)(nil),    // 4: pbuf.RegisterRequest
	(*RegisterResponse)(nil),   // 5: pbuf.RegisterResponse
	(*DeregisterRequest)(nil),  // 6: pbuf.DeregisterRequest
	(*DeregisterResponse)(nil), // 7: pbuf.DeregisterResponse
	(*JobStatus)(nil),          // 8: pbuf.JobStatus
	(*StatusRequest)(nil),      // 9: pbuf.StatusRequest
	(*Job)(nil),                // 10: pbuf.Job
	(*StatusResponse)(nil),     // 11: pbuf.StatusResponse
	nil,                        // 12: pbuf.Job.ParmsEntry
}
var file_tq_proto_depIdxs = []int32{
	1,  // 0: pbuf.JobStatus.job_state:type_name -> pbuf.JobState
	0,  // 1: pbuf.StatusRequest.worker_state:type_name -> pbuf.WorkerState
	8,  // 2: pbuf.StatusRequest.job_status:type_name -> pbuf.JobStatus
	3,  // 3: pbuf.Job.kind:type_name -> pbuf.JobKind
	12, // 4: pbuf.Job.parms:type_name -> pbuf.Job.ParmsEntry
	2,  // 5: pbuf.StatusResponse.job_control:type_name -> pbuf.JobControl
	10, // 6: pbuf.StatusResponse.job:type_name -> pbuf.Job
	4,  // 7: pbuf.tq.Register:input_type -> pbuf.RegisterRequest
	6,  // 8: pbuf.tq.Deregister:input_type -> pbuf.DeregisterRequest
	9,  // 9: pbuf.tq.Status:input_type -> pbuf.StatusRequest
	5,  // 10: pbuf.tq.Register:output_type -> pbuf.RegisterResponse
	7,  // 11: pbuf.tq.Deregister:output_type -> pbuf.DeregisterResponse
	11, // 12: pbuf.tq.Status:output_type -> pbuf.StatusResponse
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_tq_proto_init() }
func file_tq_proto_init() {
	if File_tq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_tq_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
		file_tq_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeregisterRequest); i {
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
		file_tq_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeregisterResponse); i {
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
		file_tq_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobStatus); i {
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
		file_tq_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusRequest); i {
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
		file_tq_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Job); i {
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
		file_tq_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResponse); i {
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
	file_tq_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_tq_proto_msgTypes[5].OneofWrappers = []interface{}{}
	file_tq_proto_msgTypes[7].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tq_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tq_proto_goTypes,
		DependencyIndexes: file_tq_proto_depIdxs,
		EnumInfos:         file_tq_proto_enumTypes,
		MessageInfos:      file_tq_proto_msgTypes,
	}.Build()
	File_tq_proto = out.File
	file_tq_proto_rawDesc = nil
	file_tq_proto_goTypes = nil
	file_tq_proto_depIdxs = nil
}
