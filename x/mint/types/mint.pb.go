// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stargaze/mint/v1beta1/mint.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Minter represents the minting state.
type Minter struct {
	// current annual expected provisions
	AnnualProvisions github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=annual_provisions,json=annualProvisions,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"annual_provisions" yaml:"annual_provisions"`
}

func (m *Minter) Reset()         { *m = Minter{} }
func (m *Minter) String() string { return proto.CompactTextString(m) }
func (*Minter) ProtoMessage()    {}
func (*Minter) Descriptor() ([]byte, []int) {
	return fileDescriptor_736e1519c3888655, []int{0}
}
func (m *Minter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Minter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Minter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Minter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Minter.Merge(m, src)
}
func (m *Minter) XXX_Size() int {
	return m.Size()
}
func (m *Minter) XXX_DiscardUnknown() {
	xxx_messageInfo_Minter.DiscardUnknown(m)
}

var xxx_messageInfo_Minter proto.InternalMessageInfo

// Params holds parameters for the mint module.
type Params struct {
	// type of coin to mint
	MintDenom string `protobuf:"bytes,1,opt,name=mint_denom,json=mintDenom,proto3" json:"mint_denom,omitempty"`
	// the time the chain starts
	StartTime time.Time `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time" yaml:"start_time"`
	// initial annual provisions
	InitialAnnualProvisions github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=initial_annual_provisions,json=initialAnnualProvisions,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"initial_annual_provisions" yaml:"initial_annual_provisions"`
	// factor to reduce inflation by each year
	ReductionFactor github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=reduction_factor,json=reductionFactor,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reduction_factor" yaml:"reduction_factor"`
	// expected blocks per year
	BlocksPerYear uint64 `protobuf:"varint,5,opt,name=blocks_per_year,json=blocksPerYear,proto3" json:"blocks_per_year,omitempty" yaml:"blocks_per_year"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_736e1519c3888655, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMintDenom() string {
	if m != nil {
		return m.MintDenom
	}
	return ""
}

func (m *Params) GetStartTime() time.Time {
	if m != nil {
		return m.StartTime
	}
	return time.Time{}
}

func (m *Params) GetBlocksPerYear() uint64 {
	if m != nil {
		return m.BlocksPerYear
	}
	return 0
}

func init() {
	proto.RegisterType((*Minter)(nil), "deposnetwork.mint.v1beta1.Minter")
	proto.RegisterType((*Params)(nil), "deposnetwork.mint.v1beta1.Params")
}

func init() {
	proto.RegisterFile("deposnetwork/mint/v1beta1/mint.proto", fileDescriptor_736e1519c3888655)
}

var fileDescriptor_736e1519c3888655 = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xc1, 0x6e, 0xd3, 0x30,
	0x1c, 0xc6, 0x63, 0x56, 0x2a, 0xd5, 0x08, 0x6d, 0x8b, 0x80, 0x85, 0x48, 0x4b, 0xa2, 0x1c, 0x50,
	0x2f, 0x8b, 0xb5, 0x21, 0x2e, 0xbb, 0x11, 0x4d, 0x48, 0x20, 0x90, 0xaa, 0x88, 0x03, 0x70, 0x89,
	0x9c, 0xd4, 0x0b, 0xd6, 0xe2, 0x38, 0xb2, 0x9d, 0x8e, 0x72, 0xe3, 0x05, 0xd0, 0x8e, 0x1c, 0x79,
	0x9c, 0x1d, 0x27, 0x4e, 0x88, 0x43, 0x40, 0xed, 0x1b, 0xf4, 0x09, 0x50, 0xec, 0x76, 0x43, 0x9d,
	0x38, 0xf4, 0x94, 0xf8, 0x97, 0x2f, 0x7f, 0x7f, 0xdf, 0x97, 0x18, 0x06, 0x52, 0x61, 0x51, 0xe0,
	0xcf, 0x04, 0x31, 0x5a, 0x29, 0x34, 0x39, 0xcc, 0x88, 0xc2, 0x87, 0x7a, 0x11, 0xd5, 0x82, 0x2b,
	0x6e, 0x3f, 0x5c, 0x29, 0x22, 0x0d, 0x97, 0x0a, 0xf7, 0x41, 0xc1, 0x0b, 0xae, 0x15, 0xa8, 0xbb,
	0x33, 0x62, 0xd7, 0x2f, 0x38, 0x2f, 0x4a, 0x82, 0xf4, 0x2a, 0x6b, 0x4e, 0x91, 0xa2, 0x8c, 0x48,
	0x85, 0x59, 0x6d, 0x04, 0xe1, 0x17, 0x00, 0xfb, 0x6f, 0x68, 0xa5, 0x88, 0xb0, 0xcf, 0xe1, 0x2e,
	0xae, 0xaa, 0x06, 0x97, 0x69, 0x2d, 0xf8, 0x84, 0x4a, 0xca, 0x2b, 0xe9, 0x80, 0x00, 0x0c, 0x07,
	0xf1, 0xab, 0xcb, 0xd6, 0xb7, 0x7e, 0xb5, 0xfe, 0x93, 0x82, 0xaa, 0x8f, 0x4d, 0x16, 0xe5, 0x9c,
	0xa1, 0x9c, 0x4b, 0xc6, 0xe5, 0xf2, 0x72, 0x20, 0xc7, 0x67, 0x48, 0x4d, 0x6b, 0x22, 0xa3, 0x13,
	0x92, 0x2f, 0x5a, 0xdf, 0x99, 0x62, 0x56, 0x1e, 0x87, 0xb7, 0x06, 0x86, 0xc9, 0x8e, 0x61, 0xa3,
	0x1b, 0xf4, 0x63, 0x0b, 0xf6, 0x47, 0x58, 0x60, 0x26, 0xed, 0x7d, 0x08, 0xbb, 0x54, 0xe9, 0x98,
	0x54, 0x9c, 0x99, 0xcd, 0x93, 0x41, 0x47, 0x4e, 0x3a, 0x60, 0xbf, 0x83, 0xb0, 0x4b, 0xaf, 0xd2,
	0x2e, 0x86, 0x73, 0x27, 0x00, 0xc3, 0x7b, 0x47, 0x6e, 0x64, 0x32, 0x46, 0xab, 0x8c, 0xd1, 0xdb,
	0x55, 0xc6, 0x78, 0xbf, 0xf3, 0xbd, 0x68, 0xfd, 0x5d, 0xe3, 0xe6, 0xe6, 0xdd, 0xf0, 0xe2, 0xb7,
	0x0f, 0x92, 0x81, 0x06, 0x9d, 0xdc, 0xfe, 0x0a, 0xe0, 0x63, 0x5a, 0x51, 0x45, 0x71, 0x99, 0xde,
	0x6e, 0x61, 0x4b, 0xb7, 0x90, 0x6c, 0xdc, 0x42, 0x60, 0xf6, 0xfd, 0xef, 0xe0, 0x30, 0xd9, 0x5b,
	0x3e, 0x7b, 0xbe, 0x56, 0x8a, 0xad, 0xe0, 0x8e, 0x20, 0xe3, 0x26, 0x57, 0x94, 0x57, 0xe9, 0x29,
	0xce, 0x15, 0x17, 0x4e, 0x4f, 0xdb, 0x78, 0xb9, 0xb1, 0x8d, 0x3d, 0x63, 0x63, 0x7d, 0x5e, 0x98,
	0x6c, 0x5f, 0xa3, 0x17, 0x9a, 0xd8, 0x31, 0xdc, 0xce, 0x4a, 0x9e, 0x9f, 0xc9, 0xb4, 0x26, 0x22,
	0x9d, 0x12, 0x2c, 0x9c, 0xbb, 0x01, 0x18, 0xf6, 0x62, 0x77, 0xd1, 0xfa, 0x8f, 0xcc, 0x98, 0x35,
	0x41, 0x98, 0xdc, 0x37, 0x64, 0x44, 0xc4, 0x7b, 0x82, 0xc5, 0x71, 0xef, 0xdb, 0x77, 0xdf, 0x8a,
	0x5f, 0x5f, 0xce, 0x3c, 0x70, 0x35, 0xf3, 0xc0, 0x9f, 0x99, 0x07, 0x2e, 0xe6, 0x9e, 0x75, 0x35,
	0xf7, 0xac, 0x9f, 0x73, 0xcf, 0xfa, 0x70, 0xf4, 0x8f, 0xef, 0xba, 0xc9, 0x4a, 0x9a, 0x1f, 0xe0,
	0x73, 0x22, 0x39, 0x23, 0xe8, 0xfa, 0xe7, 0x9f, 0x3c, 0x43, 0x9f, 0xcc, 0x09, 0xd0, 0x39, 0xb2,
	0xbe, 0xfe, 0xb8, 0x4f, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xb6, 0xd1, 0x32, 0x8f, 0x1f, 0x03,
	0x00, 0x00,
}

func (m *Minter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Minter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Minter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.AnnualProvisions.Size()
		i -= size
		if _, err := m.AnnualProvisions.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlocksPerYear != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.BlocksPerYear))
		i--
		dAtA[i] = 0x28
	}
	{
		size := m.ReductionFactor.Size()
		i -= size
		if _, err := m.ReductionFactor.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.InitialAnnualProvisions.Size()
		i -= size
		if _, err := m.InitialAnnualProvisions.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintMint(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	if len(m.MintDenom) > 0 {
		i -= len(m.MintDenom)
		copy(dAtA[i:], m.MintDenom)
		i = encodeVarintMint(dAtA, i, uint64(len(m.MintDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMint(dAtA []byte, offset int, v uint64) int {
	offset -= sovMint(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Minter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.AnnualProvisions.Size()
	n += 1 + l + sovMint(uint64(l))
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MintDenom)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovMint(uint64(l))
	l = m.InitialAnnualProvisions.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.ReductionFactor.Size()
	n += 1 + l + sovMint(uint64(l))
	if m.BlocksPerYear != 0 {
		n += 1 + sovMint(uint64(m.BlocksPerYear))
	}
	return n
}

func sovMint(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMint(x uint64) (n int) {
	return sovMint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Minter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: Minter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Minter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnnualProvisions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AnnualProvisions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialAnnualProvisions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InitialAnnualProvisions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReductionFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReductionFactor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlocksPerYear", wireType)
			}
			m.BlocksPerYear = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlocksPerYear |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func skipMint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
				return 0, ErrInvalidLengthMint
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMint
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMint
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMint        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMint          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMint = fmt.Errorf("proto: unexpected end of group")
)
