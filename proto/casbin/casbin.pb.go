// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/casbin/casbin.proto

package casbin

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type Permission struct {
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Method  string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
}

func (m *Permission) Reset()         { *m = Permission{} }
func (m *Permission) String() string { return proto.CompactTextString(m) }
func (*Permission) ProtoMessage()    {}
func (*Permission) Descriptor() ([]byte, []int) {
	return fileDescriptor_f835afd114115d9e, []int{0}
}
func (m *Permission) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Permission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Permission.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Permission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permission.Merge(m, src)
}
func (m *Permission) XXX_Size() int {
	return m.Size()
}
func (m *Permission) XXX_DiscardUnknown() {
	xxx_messageInfo_Permission.DiscardUnknown(m)
}

var xxx_messageInfo_Permission proto.InternalMessageInfo

func (m *Permission) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Permission) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

type Request struct {
	Label       string        `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	Role        string        `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	Roles       []string      `protobuf:"bytes,3,rep,name=roles,proto3" json:"roles,omitempty"`
	Permission  *Permission   `protobuf:"bytes,4,opt,name=permission,proto3" json:"permission,omitempty"`
	Permissions []*Permission `protobuf:"bytes,5,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_f835afd114115d9e, []int{1}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Request.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return m.Size()
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Request) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *Request) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *Request) GetPermission() *Permission {
	if m != nil {
		return m.Permission
	}
	return nil
}

func (m *Request) GetPermissions() []*Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

type Response struct {
	Valid       bool          `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	Label       string        `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Role        string        `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
	Roles       []string      `protobuf:"bytes,4,rep,name=roles,proto3" json:"roles,omitempty"`
	Permission  *Permission   `protobuf:"bytes,5,opt,name=permission,proto3" json:"permission,omitempty"`
	Permissions []*Permission `protobuf:"bytes,6,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_f835afd114115d9e, []int{2}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Response.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return m.Size()
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Response) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Response) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *Response) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *Response) GetPermission() *Permission {
	if m != nil {
		return m.Permission
	}
	return nil
}

func (m *Response) GetPermissions() []*Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func init() {
	proto.RegisterType((*Permission)(nil), "casbin.Permission")
	proto.RegisterType((*Request)(nil), "casbin.Request")
	proto.RegisterType((*Response)(nil), "casbin.Response")
}

func init() { proto.RegisterFile("proto/casbin/casbin.proto", fileDescriptor_f835afd114115d9e) }

var fileDescriptor_f835afd114115d9e = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcd, 0x6a, 0xf2, 0x40,
	0x14, 0x86, 0x33, 0x46, 0xa3, 0xdf, 0x91, 0xaf, 0x3f, 0x43, 0x29, 0xd3, 0x2e, 0x42, 0xc8, 0xca,
	0x95, 0x42, 0x6a, 0xa1, 0xab, 0x82, 0x6d, 0xc1, 0x6d, 0x19, 0x68, 0xf7, 0xd1, 0x39, 0xd0, 0x40,
	0x34, 0x69, 0x66, 0xea, 0x75, 0xf4, 0x5e, 0x7a, 0x13, 0xa5, 0x2b, 0x97, 0x5d, 0x16, 0x73, 0x23,
	0x25, 0x93, 0xb1, 0x71, 0xa1, 0x10, 0x5d, 0x8d, 0xef, 0xf8, 0x3e, 0x70, 0x9e, 0x9c, 0x04, 0x2e,
	0xd2, 0x2c, 0x51, 0xc9, 0x60, 0x1a, 0xca, 0x49, 0x34, 0x37, 0x47, 0x5f, 0xdf, 0x51, 0xa7, 0x4c,
	0xfe, 0x2d, 0xc0, 0x23, 0x66, 0xb3, 0x48, 0xca, 0x28, 0x99, 0x53, 0x06, 0x6d, 0x89, 0xd9, 0x22,
	0x9a, 0x22, 0x23, 0x1e, 0xe9, 0xfd, 0xe3, 0xeb, 0x48, 0xcf, 0xc1, 0x99, 0xa1, 0x7a, 0x49, 0x04,
	0x6b, 0xe8, 0x3f, 0x4c, 0xf2, 0x3f, 0x08, 0xb4, 0x39, 0xbe, 0xbe, 0xa1, 0x54, 0xf4, 0x0c, 0x5a,
	0x71, 0x38, 0xc1, 0xd8, 0xb0, 0x65, 0xa0, 0x14, 0x9a, 0x59, 0x12, 0xa3, 0xe1, 0xf4, 0xef, 0xa2,
	0x59, 0x9c, 0x92, 0xd9, 0x9e, 0x5d, 0x34, 0x75, 0xa0, 0x01, 0x40, 0xfa, 0x37, 0x0b, 0x6b, 0x7a,
	0xa4, 0xd7, 0x0d, 0x68, 0xdf, 0x8c, 0x5d, 0x4d, 0xc9, 0x37, 0x5a, 0x74, 0x08, 0xdd, 0x2a, 0x49,
	0xd6, 0xf2, 0xec, 0x1d, 0xd0, 0x66, 0xcd, 0xff, 0x22, 0xd0, 0xe1, 0x28, 0xd3, 0x64, 0x2e, 0xf5,
	0x30, 0x8b, 0x30, 0x8e, 0x84, 0x1e, 0xbb, 0xc3, 0xcb, 0x50, 0xc9, 0x34, 0xb6, 0xc9, 0xd8, 0xdb,
	0x64, 0x9a, 0xbb, 0x65, 0x5a, 0x87, 0xc8, 0x38, 0xb5, 0x64, 0x82, 0xdc, 0x06, 0xe7, 0x5e, 0x57,
	0xe8, 0x10, 0xfe, 0x8f, 0x84, 0xd8, 0x58, 0xe8, 0xf1, 0x1a, 0x36, 0x3b, 0xba, 0x3c, 0xa9, 0x2e,
	0x4a, 0x7d, 0xdf, 0xa2, 0x37, 0x70, 0xfa, 0x80, 0x31, 0x2a, 0xac, 0x40, 0x59, 0x9b, 0x7c, 0x4a,
	0x45, 0x78, 0x00, 0x79, 0x0d, 0x47, 0x63, 0x54, 0x7b, 0x63, 0x7d, 0x68, 0x8f, 0x84, 0xe0, 0xc5,
	0x63, 0xaf, 0xd5, 0x0f, 0xa0, 0x5b, 0xaa, 0x71, 0xbd, 0x94, 0xba, 0x4c, 0x29, 0xb5, 0x07, 0x33,
	0x80, 0xce, 0x18, 0xd5, 0x7e, 0xc0, 0x73, 0xf1, 0x9e, 0x85, 0xaa, 0x9e, 0xc9, 0x1d, 0xfb, 0x5c,
	0xb9, 0x64, 0xb9, 0x72, 0xc9, 0xcf, 0xca, 0x25, 0xef, 0xb9, 0x6b, 0x2d, 0x73, 0xd7, 0xfa, 0xce,
	0x5d, 0x6b, 0xe2, 0xe8, 0x2f, 0xfa, 0xea, 0x37, 0x00, 0x00, 0xff, 0xff, 0xb3, 0xc6, 0x60, 0x7b,
	0xee, 0x03, 0x00, 0x00,
}

func (m *Permission) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Permission) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Permission) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Method) > 0 {
		i -= len(m.Method)
		copy(dAtA[i:], m.Method)
		i = encodeVarintCasbin(dAtA, i, uint64(len(m.Method)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Service) > 0 {
		i -= len(m.Service)
		copy(dAtA[i:], m.Service)
		i = encodeVarintCasbin(dAtA, i, uint64(len(m.Service)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Permissions) > 0 {
		for iNdEx := len(m.Permissions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Permissions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCasbin(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.Permission != nil {
		{
			size, err := m.Permission.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCasbin(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Roles) > 0 {
		for iNdEx := len(m.Roles) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Roles[iNdEx])
			copy(dAtA[i:], m.Roles[iNdEx])
			i = encodeVarintCasbin(dAtA, i, uint64(len(m.Roles[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Role) > 0 {
		i -= len(m.Role)
		copy(dAtA[i:], m.Role)
		i = encodeVarintCasbin(dAtA, i, uint64(len(m.Role)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Label) > 0 {
		i -= len(m.Label)
		copy(dAtA[i:], m.Label)
		i = encodeVarintCasbin(dAtA, i, uint64(len(m.Label)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Response) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Permissions) > 0 {
		for iNdEx := len(m.Permissions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Permissions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCasbin(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.Permission != nil {
		{
			size, err := m.Permission.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCasbin(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Roles) > 0 {
		for iNdEx := len(m.Roles) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Roles[iNdEx])
			copy(dAtA[i:], m.Roles[iNdEx])
			i = encodeVarintCasbin(dAtA, i, uint64(len(m.Roles[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Role) > 0 {
		i -= len(m.Role)
		copy(dAtA[i:], m.Role)
		i = encodeVarintCasbin(dAtA, i, uint64(len(m.Role)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Label) > 0 {
		i -= len(m.Label)
		copy(dAtA[i:], m.Label)
		i = encodeVarintCasbin(dAtA, i, uint64(len(m.Label)))
		i--
		dAtA[i] = 0x12
	}
	if m.Valid {
		i--
		if m.Valid {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCasbin(dAtA []byte, offset int, v uint64) int {
	offset -= sovCasbin(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Permission) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Service)
	if l > 0 {
		n += 1 + l + sovCasbin(uint64(l))
	}
	l = len(m.Method)
	if l > 0 {
		n += 1 + l + sovCasbin(uint64(l))
	}
	return n
}

func (m *Request) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Label)
	if l > 0 {
		n += 1 + l + sovCasbin(uint64(l))
	}
	l = len(m.Role)
	if l > 0 {
		n += 1 + l + sovCasbin(uint64(l))
	}
	if len(m.Roles) > 0 {
		for _, s := range m.Roles {
			l = len(s)
			n += 1 + l + sovCasbin(uint64(l))
		}
	}
	if m.Permission != nil {
		l = m.Permission.Size()
		n += 1 + l + sovCasbin(uint64(l))
	}
	if len(m.Permissions) > 0 {
		for _, e := range m.Permissions {
			l = e.Size()
			n += 1 + l + sovCasbin(uint64(l))
		}
	}
	return n
}

func (m *Response) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Valid {
		n += 2
	}
	l = len(m.Label)
	if l > 0 {
		n += 1 + l + sovCasbin(uint64(l))
	}
	l = len(m.Role)
	if l > 0 {
		n += 1 + l + sovCasbin(uint64(l))
	}
	if len(m.Roles) > 0 {
		for _, s := range m.Roles {
			l = len(s)
			n += 1 + l + sovCasbin(uint64(l))
		}
	}
	if m.Permission != nil {
		l = m.Permission.Size()
		n += 1 + l + sovCasbin(uint64(l))
	}
	if len(m.Permissions) > 0 {
		for _, e := range m.Permissions {
			l = e.Size()
			n += 1 + l + sovCasbin(uint64(l))
		}
	}
	return n
}

func sovCasbin(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCasbin(x uint64) (n int) {
	return sovCasbin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Permission) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCasbin
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
			return fmt.Errorf("proto: Permission: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Permission: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Service", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCasbin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCasbin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCasbin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Service = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCasbin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCasbin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCasbin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Method = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCasbin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCasbin
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCasbin
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
func (m *Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCasbin
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
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Label", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCasbin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCasbin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCasbin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Label = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCasbin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCasbin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCasbin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Role = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Roles", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCasbin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCasbin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCasbin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Roles = append(m.Roles, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permission", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCasbin
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
				return ErrInvalidLengthCasbin
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCasbin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Permission == nil {
				m.Permission = &Permission{}
			}
			if err := m.Permission.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCasbin
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
				return ErrInvalidLengthCasbin
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCasbin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Permissions = append(m.Permissions, &Permission{})
			if err := m.Permissions[len(m.Permissions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCasbin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCasbin
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCasbin
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
func (m *Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCasbin
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
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Valid", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCasbin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Valid = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Label", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCasbin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCasbin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCasbin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Label = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCasbin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCasbin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCasbin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Role = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Roles", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCasbin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCasbin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCasbin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Roles = append(m.Roles, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permission", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCasbin
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
				return ErrInvalidLengthCasbin
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCasbin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Permission == nil {
				m.Permission = &Permission{}
			}
			if err := m.Permission.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCasbin
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
				return ErrInvalidLengthCasbin
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCasbin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Permissions = append(m.Permissions, &Permission{})
			if err := m.Permissions[len(m.Permissions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCasbin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCasbin
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCasbin
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
func skipCasbin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCasbin
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
					return 0, ErrIntOverflowCasbin
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
					return 0, ErrIntOverflowCasbin
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
				return 0, ErrInvalidLengthCasbin
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCasbin
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCasbin
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCasbin        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCasbin          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCasbin = fmt.Errorf("proto: unexpected end of group")
)
