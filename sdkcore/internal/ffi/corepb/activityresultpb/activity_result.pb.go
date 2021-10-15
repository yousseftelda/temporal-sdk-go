// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: activity_result.proto

package coresdk_activity_result

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	v1 "go.temporal.io/api/failure/v1"
	commonpb "go.temporal.io/sdk/sdkcore/internal/ffi/corepb/commonpb"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//*
// Used to report activity completion to core and to resolve the activity in a workflow activation
type ActivityResult struct {
	// Types that are valid to be assigned to Status:
	//	*ActivityResult_Completed
	//	*ActivityResult_Failed
	//	*ActivityResult_Cancelled
	Status isActivityResult_Status `protobuf_oneof:"status"`
}

func (m *ActivityResult) Reset()      { *m = ActivityResult{} }
func (*ActivityResult) ProtoMessage() {}
func (*ActivityResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_4bdd21ad78dc5785, []int{0}
}
func (m *ActivityResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ActivityResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ActivityResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ActivityResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivityResult.Merge(m, src)
}
func (m *ActivityResult) XXX_Size() int {
	return m.Size()
}
func (m *ActivityResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivityResult.DiscardUnknown(m)
}

var xxx_messageInfo_ActivityResult proto.InternalMessageInfo

type isActivityResult_Status interface {
	isActivityResult_Status()
	Equal(interface{}) bool
	MarshalTo([]byte) (int, error)
	Size() int
}

type ActivityResult_Completed struct {
	Completed *Success `protobuf:"bytes,1,opt,name=completed,proto3,oneof" json:"completed,omitempty"`
}
type ActivityResult_Failed struct {
	Failed *Failure `protobuf:"bytes,2,opt,name=failed,proto3,oneof" json:"failed,omitempty"`
}
type ActivityResult_Cancelled struct {
	Cancelled *Cancellation `protobuf:"bytes,3,opt,name=cancelled,proto3,oneof" json:"cancelled,omitempty"`
}

func (*ActivityResult_Completed) isActivityResult_Status() {}
func (*ActivityResult_Failed) isActivityResult_Status()    {}
func (*ActivityResult_Cancelled) isActivityResult_Status() {}

func (m *ActivityResult) GetStatus() isActivityResult_Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ActivityResult) GetCompleted() *Success {
	if x, ok := m.GetStatus().(*ActivityResult_Completed); ok {
		return x.Completed
	}
	return nil
}

func (m *ActivityResult) GetFailed() *Failure {
	if x, ok := m.GetStatus().(*ActivityResult_Failed); ok {
		return x.Failed
	}
	return nil
}

func (m *ActivityResult) GetCancelled() *Cancellation {
	if x, ok := m.GetStatus().(*ActivityResult_Cancelled); ok {
		return x.Cancelled
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ActivityResult) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ActivityResult_Completed)(nil),
		(*ActivityResult_Failed)(nil),
		(*ActivityResult_Cancelled)(nil),
	}
}

//* Used in ActivityResult to report successful completion
type Success struct {
	Result *commonpb.Payload `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (m *Success) Reset()      { *m = Success{} }
func (*Success) ProtoMessage() {}
func (*Success) Descriptor() ([]byte, []int) {
	return fileDescriptor_4bdd21ad78dc5785, []int{1}
}
func (m *Success) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Success) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Success.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Success) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Success.Merge(m, src)
}
func (m *Success) XXX_Size() int {
	return m.Size()
}
func (m *Success) XXX_DiscardUnknown() {
	xxx_messageInfo_Success.DiscardUnknown(m)
}

var xxx_messageInfo_Success proto.InternalMessageInfo

func (m *Success) GetResult() *commonpb.Payload {
	if m != nil {
		return m.Result
	}
	return nil
}

//* Used in ActivityResult to report failure
type Failure struct {
	Failure *v1.Failure `protobuf:"bytes,1,opt,name=failure,proto3" json:"failure,omitempty"`
}

func (m *Failure) Reset()      { *m = Failure{} }
func (*Failure) ProtoMessage() {}
func (*Failure) Descriptor() ([]byte, []int) {
	return fileDescriptor_4bdd21ad78dc5785, []int{2}
}
func (m *Failure) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Failure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Failure.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Failure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Failure.Merge(m, src)
}
func (m *Failure) XXX_Size() int {
	return m.Size()
}
func (m *Failure) XXX_DiscardUnknown() {
	xxx_messageInfo_Failure.DiscardUnknown(m)
}

var xxx_messageInfo_Failure proto.InternalMessageInfo

func (m *Failure) GetFailure() *v1.Failure {
	if m != nil {
		return m.Failure
	}
	return nil
}

//*
// Used in ActivityResult to report cancellation from both Core and Lang.
// When Lang reports a cancelled ActivityResult, it must put a CancelledFailure in the failure field.
// When Core reports a cancelled ActivityResult, it must put an ActivityFailure with CancelledFailure
// as the cause in the failure field.
type Cancellation struct {
	Failure *v1.Failure `protobuf:"bytes,1,opt,name=failure,proto3" json:"failure,omitempty"`
}

func (m *Cancellation) Reset()      { *m = Cancellation{} }
func (*Cancellation) ProtoMessage() {}
func (*Cancellation) Descriptor() ([]byte, []int) {
	return fileDescriptor_4bdd21ad78dc5785, []int{3}
}
func (m *Cancellation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Cancellation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Cancellation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Cancellation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cancellation.Merge(m, src)
}
func (m *Cancellation) XXX_Size() int {
	return m.Size()
}
func (m *Cancellation) XXX_DiscardUnknown() {
	xxx_messageInfo_Cancellation.DiscardUnknown(m)
}

var xxx_messageInfo_Cancellation proto.InternalMessageInfo

func (m *Cancellation) GetFailure() *v1.Failure {
	if m != nil {
		return m.Failure
	}
	return nil
}

func init() {
	proto.RegisterType((*ActivityResult)(nil), "coresdk.activity_result.ActivityResult")
	proto.RegisterType((*Success)(nil), "coresdk.activity_result.Success")
	proto.RegisterType((*Failure)(nil), "coresdk.activity_result.Failure")
	proto.RegisterType((*Cancellation)(nil), "coresdk.activity_result.Cancellation")
}

func init() { proto.RegisterFile("activity_result.proto", fileDescriptor_4bdd21ad78dc5785) }

var fileDescriptor_4bdd21ad78dc5785 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0x67, 0xfe, 0x1f, 0x52, 0x3b, 0x16, 0x17, 0x01, 0x69, 0xe9, 0xe2, 0x52, 0x02, 0x05,
	0x57, 0x13, 0xaa, 0xae, 0xba, 0xd2, 0x4a, 0xa5, 0xb8, 0x92, 0xf8, 0x00, 0x32, 0x4e, 0x46, 0x09,
	0x26, 0x9d, 0x90, 0x99, 0x14, 0xba, 0xf3, 0x11, 0x7c, 0x0c, 0x1f, 0xc5, 0x65, 0x97, 0xdd, 0x69,
	0x27, 0x1b, 0x97, 0x7d, 0x04, 0x69, 0x32, 0xa9, 0x22, 0x14, 0x17, 0xae, 0xef, 0xf9, 0xce, 0x3d,
	0xf7, 0x5c, 0x72, 0xc8, 0xb8, 0x8e, 0x66, 0x91, 0x9e, 0xdf, 0x66, 0x42, 0xe5, 0xb1, 0xa6, 0x69,
	0x26, 0xb5, 0x74, 0xdb, 0x5c, 0x66, 0x42, 0x85, 0x8f, 0xf4, 0xc7, 0xb8, 0xdb, 0xe2, 0x32, 0x49,
	0xe4, 0xb4, 0x92, 0x75, 0xfb, 0x5a, 0x24, 0xa9, 0xcc, 0x58, 0xec, 0xb3, 0x34, 0xf2, 0xef, 0x59,
	0x14, 0xe7, 0x99, 0xf0, 0x67, 0x03, 0x3f, 0x11, 0x4a, 0xb1, 0x07, 0x51, 0xc9, 0xbc, 0x37, 0x4c,
	0x0e, 0xce, 0xad, 0x51, 0x50, 0xfa, 0xb8, 0x67, 0xa4, 0xc9, 0x65, 0x92, 0xc6, 0x42, 0x8b, 0xb0,
	0x83, 0x7b, 0xf8, 0x68, 0xff, 0xb8, 0x47, 0x77, 0x2c, 0xa5, 0x37, 0x39, 0xe7, 0x42, 0xa9, 0x09,
	0x0a, 0xbe, 0x20, 0x77, 0x48, 0x9c, 0xcd, 0x42, 0x11, 0x76, 0xfe, 0xfd, 0x82, 0x5f, 0x56, 0xb9,
	0x26, 0x28, 0xb0, 0x84, 0x3b, 0x26, 0x4d, 0xce, 0xa6, 0x5c, 0xc4, 0x1b, 0xfc, 0x7f, 0x89, 0xf7,
	0x77, 0xe2, 0x17, 0x95, 0x92, 0xe9, 0x48, 0x4e, 0xcb, 0x08, 0x35, 0x39, 0xda, 0x23, 0x8e, 0xd2,
	0x4c, 0xe7, 0xca, 0x1b, 0x92, 0x86, 0x0d, 0xe9, 0xfa, 0xc4, 0xa9, 0x40, 0x7b, 0x56, 0x7b, 0x6b,
	0x6c, 0xab, 0xbb, 0x66, 0xf3, 0x58, 0xb2, 0x30, 0xb0, 0x32, 0x6f, 0x4c, 0x1a, 0x36, 0xa1, 0x3b,
	0x24, 0x0d, 0x5b, 0xe2, 0xb6, 0x93, 0xba, 0x61, 0xca, 0xd2, 0x88, 0xda, 0x21, 0x9d, 0x0d, 0xea,
	0xa3, 0x82, 0x1a, 0xf0, 0xae, 0x48, 0xeb, 0x7b, 0xd2, 0xbf, 0x78, 0x8d, 0x4e, 0x17, 0x2b, 0x40,
	0xcb, 0x15, 0xa0, 0xf5, 0x0a, 0xf0, 0x93, 0x01, 0xfc, 0x62, 0x00, 0xbf, 0x1a, 0xc0, 0x0b, 0x03,
	0xf8, 0xdd, 0x00, 0xfe, 0x30, 0x80, 0xd6, 0x06, 0xf0, 0x73, 0x01, 0x68, 0x51, 0x00, 0x5a, 0x16,
	0x80, 0xee, 0x9c, 0xf2, 0xdb, 0x27, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x05, 0xbe, 0xe6,
	0x54, 0x02, 0x00, 0x00,
}

func (this *ActivityResult) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ActivityResult)
	if !ok {
		that2, ok := that.(ActivityResult)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.Status == nil {
		if this.Status != nil {
			return false
		}
	} else if this.Status == nil {
		return false
	} else if !this.Status.Equal(that1.Status) {
		return false
	}
	return true
}
func (this *ActivityResult_Completed) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ActivityResult_Completed)
	if !ok {
		that2, ok := that.(ActivityResult_Completed)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Completed.Equal(that1.Completed) {
		return false
	}
	return true
}
func (this *ActivityResult_Failed) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ActivityResult_Failed)
	if !ok {
		that2, ok := that.(ActivityResult_Failed)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Failed.Equal(that1.Failed) {
		return false
	}
	return true
}
func (this *ActivityResult_Cancelled) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ActivityResult_Cancelled)
	if !ok {
		that2, ok := that.(ActivityResult_Cancelled)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Cancelled.Equal(that1.Cancelled) {
		return false
	}
	return true
}
func (this *Success) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Success)
	if !ok {
		that2, ok := that.(Success)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Result.Equal(that1.Result) {
		return false
	}
	return true
}
func (this *Failure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Failure)
	if !ok {
		that2, ok := that.(Failure)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Failure.Equal(that1.Failure) {
		return false
	}
	return true
}
func (this *Cancellation) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Cancellation)
	if !ok {
		that2, ok := that.(Cancellation)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Failure.Equal(that1.Failure) {
		return false
	}
	return true
}
func (this *ActivityResult) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&coresdk_activity_result.ActivityResult{")
	if this.Status != nil {
		s = append(s, "Status: "+fmt.Sprintf("%#v", this.Status)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ActivityResult_Completed) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&coresdk_activity_result.ActivityResult_Completed{` +
		`Completed:` + fmt.Sprintf("%#v", this.Completed) + `}`}, ", ")
	return s
}
func (this *ActivityResult_Failed) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&coresdk_activity_result.ActivityResult_Failed{` +
		`Failed:` + fmt.Sprintf("%#v", this.Failed) + `}`}, ", ")
	return s
}
func (this *ActivityResult_Cancelled) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&coresdk_activity_result.ActivityResult_Cancelled{` +
		`Cancelled:` + fmt.Sprintf("%#v", this.Cancelled) + `}`}, ", ")
	return s
}
func (this *Success) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&coresdk_activity_result.Success{")
	if this.Result != nil {
		s = append(s, "Result: "+fmt.Sprintf("%#v", this.Result)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Failure) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&coresdk_activity_result.Failure{")
	if this.Failure != nil {
		s = append(s, "Failure: "+fmt.Sprintf("%#v", this.Failure)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Cancellation) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&coresdk_activity_result.Cancellation{")
	if this.Failure != nil {
		s = append(s, "Failure: "+fmt.Sprintf("%#v", this.Failure)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringActivityResult(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *ActivityResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ActivityResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ActivityResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != nil {
		{
			size := m.Status.Size()
			i -= size
			if _, err := m.Status.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *ActivityResult_Completed) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ActivityResult_Completed) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Completed != nil {
		{
			size, err := m.Completed.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintActivityResult(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *ActivityResult_Failed) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ActivityResult_Failed) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Failed != nil {
		{
			size, err := m.Failed.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintActivityResult(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *ActivityResult_Cancelled) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ActivityResult_Cancelled) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Cancelled != nil {
		{
			size, err := m.Cancelled.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintActivityResult(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *Success) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Success) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Success) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Result != nil {
		{
			size, err := m.Result.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintActivityResult(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Failure) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Failure) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Failure) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Failure != nil {
		{
			size, err := m.Failure.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintActivityResult(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Cancellation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Cancellation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Cancellation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Failure != nil {
		{
			size, err := m.Failure.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintActivityResult(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintActivityResult(dAtA []byte, offset int, v uint64) int {
	offset -= sovActivityResult(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ActivityResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != nil {
		n += m.Status.Size()
	}
	return n
}

func (m *ActivityResult_Completed) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Completed != nil {
		l = m.Completed.Size()
		n += 1 + l + sovActivityResult(uint64(l))
	}
	return n
}
func (m *ActivityResult_Failed) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Failed != nil {
		l = m.Failed.Size()
		n += 1 + l + sovActivityResult(uint64(l))
	}
	return n
}
func (m *ActivityResult_Cancelled) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Cancelled != nil {
		l = m.Cancelled.Size()
		n += 1 + l + sovActivityResult(uint64(l))
	}
	return n
}
func (m *Success) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Result != nil {
		l = m.Result.Size()
		n += 1 + l + sovActivityResult(uint64(l))
	}
	return n
}

func (m *Failure) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Failure != nil {
		l = m.Failure.Size()
		n += 1 + l + sovActivityResult(uint64(l))
	}
	return n
}

func (m *Cancellation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Failure != nil {
		l = m.Failure.Size()
		n += 1 + l + sovActivityResult(uint64(l))
	}
	return n
}

func sovActivityResult(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozActivityResult(x uint64) (n int) {
	return sovActivityResult(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ActivityResult) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ActivityResult{`,
		`Status:` + fmt.Sprintf("%v", this.Status) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ActivityResult_Completed) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ActivityResult_Completed{`,
		`Completed:` + strings.Replace(fmt.Sprintf("%v", this.Completed), "Success", "Success", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ActivityResult_Failed) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ActivityResult_Failed{`,
		`Failed:` + strings.Replace(fmt.Sprintf("%v", this.Failed), "Failure", "Failure", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ActivityResult_Cancelled) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ActivityResult_Cancelled{`,
		`Cancelled:` + strings.Replace(fmt.Sprintf("%v", this.Cancelled), "Cancellation", "Cancellation", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Success) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Success{`,
		`Result:` + strings.Replace(fmt.Sprintf("%v", this.Result), "Payload", "commonpb.Payload", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Failure) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Failure{`,
		`Failure:` + strings.Replace(fmt.Sprintf("%v", this.Failure), "Failure", "v1.Failure", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Cancellation) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Cancellation{`,
		`Failure:` + strings.Replace(fmt.Sprintf("%v", this.Failure), "Failure", "v1.Failure", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringActivityResult(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ActivityResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowActivityResult
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ActivityResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ActivityResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Completed", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActivityResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthActivityResult
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthActivityResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Success{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Status = &ActivityResult_Completed{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Failed", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActivityResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthActivityResult
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthActivityResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Failure{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Status = &ActivityResult_Failed{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cancelled", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActivityResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthActivityResult
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthActivityResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Cancellation{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Status = &ActivityResult_Cancelled{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipActivityResult(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthActivityResult
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Success) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowActivityResult
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Success: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Success: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActivityResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthActivityResult
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthActivityResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Result == nil {
				m.Result = &commonpb.Payload{}
			}
			if err := m.Result.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipActivityResult(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthActivityResult
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Failure) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowActivityResult
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Failure: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Failure: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Failure", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActivityResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthActivityResult
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthActivityResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Failure == nil {
				m.Failure = &v1.Failure{}
			}
			if err := m.Failure.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipActivityResult(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthActivityResult
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Cancellation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowActivityResult
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Cancellation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Cancellation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Failure", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActivityResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthActivityResult
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthActivityResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Failure == nil {
				m.Failure = &v1.Failure{}
			}
			if err := m.Failure.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipActivityResult(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthActivityResult
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipActivityResult(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowActivityResult
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowActivityResult
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowActivityResult
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthActivityResult
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupActivityResult
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthActivityResult
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthActivityResult        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowActivityResult          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupActivityResult = fmt.Errorf("proto: unexpected end of group")
)
