// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: layer/reporter/delegation.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Delegation is a type that represents a delegator's delegation to a reporter
type Delegation struct {
	// reporter is the address of the reporter being delegated to
	Reporter []byte `protobuf:"bytes,1,opt,name=reporter,proto3" json:"reporter,omitempty"`
	// amount is the amount of tokens delegated
	Amount cosmossdk_io_math.Int `protobuf:"bytes,2,opt,name=amount,proto3,customtype=cosmossdk.io/math.Int" json:"amount"`
}

func (m *Delegation) Reset()         { *m = Delegation{} }
func (m *Delegation) String() string { return proto.CompactTextString(m) }
func (*Delegation) ProtoMessage()    {}
func (*Delegation) Descriptor() ([]byte, []int) {
	return fileDescriptor_66d928de90a17af7, []int{0}
}
func (m *Delegation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Delegation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Delegation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Delegation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Delegation.Merge(m, src)
}
func (m *Delegation) XXX_Size() int {
	return m.Size()
}
func (m *Delegation) XXX_DiscardUnknown() {
	xxx_messageInfo_Delegation.DiscardUnknown(m)
}

var xxx_messageInfo_Delegation proto.InternalMessageInfo

func (m *Delegation) GetReporter() []byte {
	if m != nil {
		return m.Reporter
	}
	return nil
}

func init() {
	proto.RegisterType((*Delegation)(nil), "layer.reporter.Delegation")
}

func init() { proto.RegisterFile("layer/reporter/delegation.proto", fileDescriptor_66d928de90a17af7) }

var fileDescriptor_66d928de90a17af7 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcf, 0x49, 0xac, 0x4c,
	0x2d, 0xd2, 0x2f, 0x4a, 0x2d, 0xc8, 0x2f, 0x2a, 0x49, 0x2d, 0xd2, 0x4f, 0x49, 0xcd, 0x49, 0x4d,
	0x4f, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x03, 0x2b, 0xd0,
	0x83, 0x29, 0x90, 0x12, 0x4c, 0xcc, 0xcd, 0xcc, 0xcb, 0xd7, 0x07, 0x93, 0x10, 0x25, 0x52, 0x92,
	0xc9, 0xf9, 0xc5, 0xb9, 0xf9, 0xc5, 0xf1, 0x60, 0x9e, 0x3e, 0x84, 0x03, 0x95, 0x12, 0x49, 0xcf,
	0x4f, 0xcf, 0x87, 0x88, 0x83, 0x58, 0x10, 0x51, 0xa5, 0x5c, 0x2e, 0x2e, 0x17, 0xb8, 0x3d, 0x42,
	0x52, 0x5c, 0x1c, 0x30, 0xd3, 0x25, 0x18, 0x15, 0x18, 0x35, 0x78, 0x82, 0xe0, 0x7c, 0x21, 0x67,
	0x2e, 0xb6, 0xc4, 0xdc, 0xfc, 0xd2, 0xbc, 0x12, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x4e, 0x27, 0xed,
	0x13, 0xf7, 0xe4, 0x19, 0x6e, 0xdd, 0x93, 0x17, 0x85, 0xd8, 0x52, 0x9c, 0x92, 0xad, 0x97, 0x99,
	0xaf, 0x9f, 0x9b, 0x58, 0x92, 0xa1, 0xe7, 0x99, 0x57, 0x72, 0x69, 0x8b, 0x2e, 0x17, 0xd4, 0x7a,
	0xcf, 0xbc, 0x92, 0x20, 0xa8, 0x56, 0x27, 0xd7, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63,
	0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96,
	0x63, 0x88, 0xd2, 0x4e, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0x49,
	0xcd, 0xc9, 0xc9, 0x2f, 0xd2, 0xcd, 0xcc, 0xd7, 0x87, 0x04, 0x49, 0x05, 0x22, 0x50, 0x4a, 0x2a,
	0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x8e, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x67, 0x38,
	0x08, 0x27, 0x33, 0x01, 0x00, 0x00,
}

func (m *Delegation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Delegation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Delegation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDelegation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Reporter) > 0 {
		i -= len(m.Reporter)
		copy(dAtA[i:], m.Reporter)
		i = encodeVarintDelegation(dAtA, i, uint64(len(m.Reporter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDelegation(dAtA []byte, offset int, v uint64) int {
	offset -= sovDelegation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Delegation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Reporter)
	if l > 0 {
		n += 1 + l + sovDelegation(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovDelegation(uint64(l))
	return n
}

func sovDelegation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDelegation(x uint64) (n int) {
	return sovDelegation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Delegation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDelegation
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
			return fmt.Errorf("proto: Delegation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Delegation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reporter", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDelegation
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDelegation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reporter = append(m.Reporter[:0], dAtA[iNdEx:postIndex]...)
			if m.Reporter == nil {
				m.Reporter = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegation
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
				return ErrInvalidLengthDelegation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelegation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDelegation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDelegation
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
func skipDelegation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDelegation
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
					return 0, ErrIntOverflowDelegation
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
					return 0, ErrIntOverflowDelegation
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
				return 0, ErrInvalidLengthDelegation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDelegation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDelegation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDelegation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDelegation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDelegation = fmt.Errorf("proto: unexpected end of group")
)
