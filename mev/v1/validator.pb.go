// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: v1/validator.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type GetBundlesRequest struct {
	MinBlockHeight uint64 `protobuf:"varint,1,opt,name=min_block_height,json=minBlockHeight,proto3" json:"min_block_height,omitempty"`
}

func (m *GetBundlesRequest) Reset()         { *m = GetBundlesRequest{} }
func (m *GetBundlesRequest) String() string { return proto.CompactTextString(m) }
func (*GetBundlesRequest) ProtoMessage()    {}
func (*GetBundlesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b975b097e577fdff, []int{0}
}
func (m *GetBundlesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetBundlesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetBundlesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetBundlesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBundlesRequest.Merge(m, src)
}
func (m *GetBundlesRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetBundlesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBundlesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBundlesRequest proto.InternalMessageInfo

func (m *GetBundlesRequest) GetMinBlockHeight() uint64 {
	if m != nil {
		return m.MinBlockHeight
	}
	return 0
}

type GetBundlesResponse struct {
	Bundles map[uint64]*Bundles `protobuf:"bytes,1,rep,name=bundles,proto3" json:"bundles,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *GetBundlesResponse) Reset()         { *m = GetBundlesResponse{} }
func (m *GetBundlesResponse) String() string { return proto.CompactTextString(m) }
func (*GetBundlesResponse) ProtoMessage()    {}
func (*GetBundlesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b975b097e577fdff, []int{1}
}
func (m *GetBundlesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetBundlesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetBundlesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetBundlesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBundlesResponse.Merge(m, src)
}
func (m *GetBundlesResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetBundlesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBundlesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBundlesResponse proto.InternalMessageInfo

func (m *GetBundlesResponse) GetBundles() map[uint64]*Bundles {
	if m != nil {
		return m.Bundles
	}
	return nil
}

type Bundles struct {
	Bundles []*Bundle `protobuf:"bytes,1,rep,name=bundles,proto3" json:"bundles,omitempty"`
}

func (m *Bundles) Reset()         { *m = Bundles{} }
func (m *Bundles) String() string { return proto.CompactTextString(m) }
func (*Bundles) ProtoMessage()    {}
func (*Bundles) Descriptor() ([]byte, []int) {
	return fileDescriptor_b975b097e577fdff, []int{2}
}
func (m *Bundles) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Bundles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Bundles.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Bundles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bundles.Merge(m, src)
}
func (m *Bundles) XXX_Size() int {
	return m.Size()
}
func (m *Bundles) XXX_DiscardUnknown() {
	xxx_messageInfo_Bundles.DiscardUnknown(m)
}

var xxx_messageInfo_Bundles proto.InternalMessageInfo

func (m *Bundles) GetBundles() []*Bundle {
	if m != nil {
		return m.Bundles
	}
	return nil
}

type Bundle struct {
	BlockHeight  uint64   `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	Transactions [][]byte `protobuf:"bytes,2,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (m *Bundle) Reset()         { *m = Bundle{} }
func (m *Bundle) String() string { return proto.CompactTextString(m) }
func (*Bundle) ProtoMessage()    {}
func (*Bundle) Descriptor() ([]byte, []int) {
	return fileDescriptor_b975b097e577fdff, []int{3}
}
func (m *Bundle) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Bundle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Bundle.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Bundle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bundle.Merge(m, src)
}
func (m *Bundle) XXX_Size() int {
	return m.Size()
}
func (m *Bundle) XXX_DiscardUnknown() {
	xxx_messageInfo_Bundle.DiscardUnknown(m)
}

var xxx_messageInfo_Bundle proto.InternalMessageInfo

func (m *Bundle) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *Bundle) GetTransactions() [][]byte {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func init() {
	proto.RegisterType((*GetBundlesRequest)(nil), "seiprotocol.seichain.mev.GetBundlesRequest")
	proto.RegisterType((*GetBundlesResponse)(nil), "seiprotocol.seichain.mev.GetBundlesResponse")
	proto.RegisterMapType((map[uint64]*Bundles)(nil), "seiprotocol.seichain.mev.GetBundlesResponse.BundlesEntry")
	proto.RegisterType((*Bundles)(nil), "seiprotocol.seichain.mev.Bundles")
	proto.RegisterType((*Bundle)(nil), "seiprotocol.seichain.mev.Bundle")
}

func init() { proto.RegisterFile("v1/validator.proto", fileDescriptor_b975b097e577fdff) }

var fileDescriptor_b975b097e577fdff = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xcf, 0x6b, 0xe2, 0x40,
	0x14, 0xc7, 0x1d, 0xdd, 0x55, 0x78, 0x06, 0x71, 0x87, 0x3d, 0x04, 0x0f, 0x21, 0xe6, 0x14, 0x58,
	0x4d, 0x58, 0x77, 0x61, 0x77, 0x85, 0xbd, 0x08, 0xb6, 0xbd, 0xb5, 0xa4, 0xb7, 0x42, 0x91, 0x24,
	0x4e, 0x93, 0xc1, 0x64, 0xc6, 0x66, 0x26, 0x81, 0xfc, 0x17, 0xfd, 0xb3, 0x3c, 0x7a, 0xec, 0xb1,
	0xe8, 0x3f, 0x52, 0x4c, 0xb4, 0xda, 0x1f, 0xd2, 0xf6, 0xf6, 0xe6, 0x33, 0xef, 0x7d, 0xdf, 0x97,
	0xef, 0x03, 0x9c, 0xfd, 0xb4, 0x33, 0x37, 0xa2, 0x53, 0x57, 0xf2, 0xc4, 0x9a, 0x27, 0x5c, 0x72,
	0xac, 0x0a, 0x42, 0x8b, 0xca, 0xe7, 0x91, 0x25, 0x08, 0xf5, 0x43, 0x97, 0x32, 0x2b, 0x26, 0x59,
	0xe7, 0x7b, 0xc0, 0x03, 0x5e, 0x7c, 0xd9, 0x9b, 0xaa, 0xec, 0x37, 0xfe, 0xc3, 0xb7, 0x53, 0x22,
	0x47, 0x29, 0x9b, 0x46, 0x44, 0x38, 0xe4, 0x36, 0x25, 0x42, 0x62, 0x13, 0xda, 0x31, 0x65, 0x13,
	0x2f, 0xe2, 0xfe, 0x6c, 0x12, 0x12, 0x1a, 0x84, 0x52, 0x45, 0x3a, 0x32, 0xbf, 0x38, 0xad, 0x98,
	0xb2, 0xd1, 0x06, 0x9f, 0x15, 0xd4, 0x58, 0x20, 0xc0, 0x87, 0xf3, 0x62, 0xce, 0x99, 0x20, 0xf8,
	0x12, 0x1a, 0x5e, 0x89, 0x54, 0xa4, 0xd7, 0xcc, 0xe6, 0xe0, 0x9f, 0x75, 0xcc, 0x97, 0xf5, 0x7a,
	0xdc, 0xda, 0xbe, 0xc7, 0x4c, 0x26, 0xb9, 0xb3, 0x53, 0xea, 0x5c, 0x83, 0x72, 0xf8, 0x81, 0xdb,
	0x50, 0x9b, 0x91, 0x7c, 0x6b, 0x6c, 0x53, 0xe2, 0x3f, 0xf0, 0x35, 0x73, 0xa3, 0x94, 0xa8, 0x55,
	0x1d, 0x99, 0xcd, 0x41, 0xf7, 0xf8, 0xd2, 0xdd, 0xc6, 0xb2, 0x7f, 0x58, 0xfd, 0x8b, 0x8c, 0x31,
	0x34, 0xb6, 0x14, 0x0f, 0x5f, 0xda, 0xd7, 0xdf, 0x53, 0x7a, 0x72, 0x69, 0x9c, 0x43, 0xbd, 0x44,
	0xb8, 0x0b, 0xca, 0x1b, 0x09, 0x36, 0xbd, 0x7d, 0x7c, 0xd8, 0x00, 0x45, 0x26, 0x2e, 0x13, 0xae,
	0x2f, 0x29, 0x67, 0x42, 0xad, 0xea, 0x35, 0x53, 0x71, 0x9e, 0xb1, 0x41, 0x0e, 0xad, 0x52, 0xf0,
	0x22, 0xe1, 0x19, 0x9d, 0x92, 0x04, 0x07, 0x00, 0xfb, 0xd0, 0xf0, 0x8f, 0x8f, 0x45, 0x5b, 0x5c,
	0xb6, 0xd3, 0xfb, 0xcc, 0x1d, 0x46, 0x27, 0x8b, 0x95, 0x86, 0x96, 0x2b, 0x0d, 0x3d, 0xac, 0x34,
	0x74, 0xb7, 0xd6, 0x2a, 0xcb, 0xb5, 0x56, 0xb9, 0x5f, 0x6b, 0x95, 0xab, 0x5e, 0x40, 0x65, 0x98,
	0x7a, 0x96, 0xcf, 0x63, 0x3b, 0xfe, 0x3d, 0x13, 0x94, 0xdb, 0x82, 0x46, 0xbc, 0x1f, 0x93, 0xac,
	0x5f, 0xc8, 0x7b, 0xe9, 0x4d, 0x3f, 0xe0, 0xb6, 0xcc, 0xe7, 0x44, 0x78, 0xf5, 0x02, 0xfd, 0x7a,
	0x0c, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x27, 0xf1, 0x07, 0xb1, 0x02, 0x00, 0x00,
}

func (m *GetBundlesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetBundlesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetBundlesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MinBlockHeight != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.MinBlockHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetBundlesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetBundlesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetBundlesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Bundles) > 0 {
		for k := range m.Bundles {
			v := m.Bundles[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintValidator(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i = encodeVarintValidator(dAtA, i, uint64(k))
			i--
			dAtA[i] = 0x8
			i = encodeVarintValidator(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Bundles) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Bundles) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Bundles) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Bundles) > 0 {
		for iNdEx := len(m.Bundles) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Bundles[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintValidator(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Bundle) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Bundle) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Bundle) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Transactions) > 0 {
		for iNdEx := len(m.Transactions) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Transactions[iNdEx])
			copy(dAtA[i:], m.Transactions[iNdEx])
			i = encodeVarintValidator(dAtA, i, uint64(len(m.Transactions[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.BlockHeight != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintValidator(dAtA []byte, offset int, v uint64) int {
	offset -= sovValidator(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetBundlesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MinBlockHeight != 0 {
		n += 1 + sovValidator(uint64(m.MinBlockHeight))
	}
	return n
}

func (m *GetBundlesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Bundles) > 0 {
		for k, v := range m.Bundles {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovValidator(uint64(l))
			}
			mapEntrySize := 1 + sovValidator(uint64(k)) + l
			n += mapEntrySize + 1 + sovValidator(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *Bundles) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Bundles) > 0 {
		for _, e := range m.Bundles {
			l = e.Size()
			n += 1 + l + sovValidator(uint64(l))
		}
	}
	return n
}

func (m *Bundle) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockHeight != 0 {
		n += 1 + sovValidator(uint64(m.BlockHeight))
	}
	if len(m.Transactions) > 0 {
		for _, b := range m.Transactions {
			l = len(b)
			n += 1 + l + sovValidator(uint64(l))
		}
	}
	return n
}

func sovValidator(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozValidator(x uint64) (n int) {
	return sovValidator(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetBundlesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValidator
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
			return fmt.Errorf("proto: GetBundlesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetBundlesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinBlockHeight", wireType)
			}
			m.MinBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinBlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthValidator
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
func (m *GetBundlesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValidator
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
			return fmt.Errorf("proto: GetBundlesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetBundlesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bundles", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Bundles == nil {
				m.Bundles = make(map[uint64]*Bundles)
			}
			var mapkey uint64
			var mapvalue *Bundles
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowValidator
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
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowValidator
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowValidator
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthValidator
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthValidator
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &Bundles{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipValidator(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthValidator
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Bundles[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthValidator
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
func (m *Bundles) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValidator
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
			return fmt.Errorf("proto: Bundles: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Bundles: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bundles", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bundles = append(m.Bundles, &Bundle{})
			if err := m.Bundles[len(m.Bundles)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthValidator
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
func (m *Bundle) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValidator
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
			return fmt.Errorf("proto: Bundle: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Bundle: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transactions", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Transactions = append(m.Transactions, make([]byte, postIndex-iNdEx))
			copy(m.Transactions[len(m.Transactions)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthValidator
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
func skipValidator(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowValidator
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
					return 0, ErrIntOverflowValidator
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
					return 0, ErrIntOverflowValidator
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
				return 0, ErrInvalidLengthValidator
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupValidator
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthValidator
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthValidator        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowValidator          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupValidator = fmt.Errorf("proto: unexpected end of group")
)
