// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rpc.proto

/*
	Package rpc is a generated protocol buffer package.

	It is generated from these files:
		rpc.proto

	It has these top-level messages:
		ResultStatus
		SyncInfo
*/
package rpc

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import tendermint "github.com/hyperledger/burrow/consensus/tendermint"
import validator "github.com/hyperledger/burrow/acm/validator"
import _ "github.com/golang/protobuf/ptypes/timestamp"

import github_com_hyperledger_burrow_binary "github.com/hyperledger/burrow/binary"
import time "time"

import types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ResultStatus struct {
	ChainID       string                                        `protobuf:"bytes,1,opt,name=ChainID,proto3" json:"ChainID,omitempty"`
	RunID         string                                        `protobuf:"bytes,2,opt,name=RunID,proto3" json:"RunID,omitempty"`
	BurrowVersion string                                        `protobuf:"bytes,3,opt,name=BurrowVersion,proto3" json:"BurrowVersion,omitempty"`
	GenesisHash   github_com_hyperledger_burrow_binary.HexBytes `protobuf:"bytes,4,opt,name=GenesisHash,proto3,customtype=github.com/hyperledger/burrow/binary.HexBytes" json:"GenesisHash"`
	NodeInfo      *tendermint.NodeInfo                          `protobuf:"bytes,5,opt,name=NodeInfo" json:"NodeInfo,omitempty"`
	SyncInfo      *SyncInfo                                     `protobuf:"bytes,6,opt,name=SyncInfo" json:"SyncInfo,omitempty"`
	ValidatorInfo *validator.Validator                          `protobuf:"bytes,7,opt,name=ValidatorInfo" json:"ValidatorInfo,omitempty"`
}

func (m *ResultStatus) Reset()                    { *m = ResultStatus{} }
func (m *ResultStatus) String() string            { return proto.CompactTextString(m) }
func (*ResultStatus) ProtoMessage()               {}
func (*ResultStatus) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{0} }

func (m *ResultStatus) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

func (m *ResultStatus) GetRunID() string {
	if m != nil {
		return m.RunID
	}
	return ""
}

func (m *ResultStatus) GetBurrowVersion() string {
	if m != nil {
		return m.BurrowVersion
	}
	return ""
}

func (m *ResultStatus) GetNodeInfo() *tendermint.NodeInfo {
	if m != nil {
		return m.NodeInfo
	}
	return nil
}

func (m *ResultStatus) GetSyncInfo() *SyncInfo {
	if m != nil {
		return m.SyncInfo
	}
	return nil
}

func (m *ResultStatus) GetValidatorInfo() *validator.Validator {
	if m != nil {
		return m.ValidatorInfo
	}
	return nil
}

func (*ResultStatus) XXX_MessageName() string {
	return "rpc.ResultStatus"
}

type SyncInfo struct {
	LatestBlockHeight uint64                                        `protobuf:"varint,1,opt,name=LatestBlockHeight,proto3" json:""`
	LatestBlockHash   github_com_hyperledger_burrow_binary.HexBytes `protobuf:"bytes,2,opt,name=LatestBlockHash,proto3,customtype=github.com/hyperledger/burrow/binary.HexBytes" json:"LatestBlockHash"`
	LatestAppHash     github_com_hyperledger_burrow_binary.HexBytes `protobuf:"bytes,3,opt,name=LatestAppHash,proto3,customtype=github.com/hyperledger/burrow/binary.HexBytes" json:"LatestAppHash"`
	// Timestamp of block as set by the block proposer
	LatestBlockTime time.Time `protobuf:"bytes,4,opt,name=LatestBlockTime,stdtime" json:"LatestBlockTime"`
	// Time at which we committed the last block
	LatestBlockSeenTime time.Time `protobuf:"bytes,5,opt,name=LatestBlockSeenTime,stdtime" json:"LatestBlockSeenTime"`
	// When catching up in fast sync
	CatchingUp bool `protobuf:"varint,6,opt,name=CatchingUp,proto3" json:"CatchingUp,omitempty"`
}

func (m *SyncInfo) Reset()                    { *m = SyncInfo{} }
func (m *SyncInfo) String() string            { return proto.CompactTextString(m) }
func (*SyncInfo) ProtoMessage()               {}
func (*SyncInfo) Descriptor() ([]byte, []int) { return fileDescriptorRpc, []int{1} }

func (m *SyncInfo) GetLatestBlockHeight() uint64 {
	if m != nil {
		return m.LatestBlockHeight
	}
	return 0
}

func (m *SyncInfo) GetLatestBlockTime() time.Time {
	if m != nil {
		return m.LatestBlockTime
	}
	return time.Time{}
}

func (m *SyncInfo) GetLatestBlockSeenTime() time.Time {
	if m != nil {
		return m.LatestBlockSeenTime
	}
	return time.Time{}
}

func (m *SyncInfo) GetCatchingUp() bool {
	if m != nil {
		return m.CatchingUp
	}
	return false
}

func (*SyncInfo) XXX_MessageName() string {
	return "rpc.SyncInfo"
}
func init() {
	proto.RegisterType((*ResultStatus)(nil), "rpc.ResultStatus")
	golang_proto.RegisterType((*ResultStatus)(nil), "rpc.ResultStatus")
	proto.RegisterType((*SyncInfo)(nil), "rpc.SyncInfo")
	golang_proto.RegisterType((*SyncInfo)(nil), "rpc.SyncInfo")
}
func (m *ResultStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResultStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ChainID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.ChainID)))
		i += copy(dAtA[i:], m.ChainID)
	}
	if len(m.RunID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.RunID)))
		i += copy(dAtA[i:], m.RunID)
	}
	if len(m.BurrowVersion) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.BurrowVersion)))
		i += copy(dAtA[i:], m.BurrowVersion)
	}
	dAtA[i] = 0x22
	i++
	i = encodeVarintRpc(dAtA, i, uint64(m.GenesisHash.Size()))
	n1, err := m.GenesisHash.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.NodeInfo != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.NodeInfo.Size()))
		n2, err := m.NodeInfo.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.SyncInfo != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.SyncInfo.Size()))
		n3, err := m.SyncInfo.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.ValidatorInfo != nil {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.ValidatorInfo.Size()))
		n4, err := m.ValidatorInfo.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func (m *SyncInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SyncInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.LatestBlockHeight != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.LatestBlockHeight))
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintRpc(dAtA, i, uint64(m.LatestBlockHash.Size()))
	n5, err := m.LatestBlockHash.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	dAtA[i] = 0x1a
	i++
	i = encodeVarintRpc(dAtA, i, uint64(m.LatestAppHash.Size()))
	n6, err := m.LatestAppHash.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	dAtA[i] = 0x22
	i++
	i = encodeVarintRpc(dAtA, i, uint64(types.SizeOfStdTime(m.LatestBlockTime)))
	n7, err := types.StdTimeMarshalTo(m.LatestBlockTime, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n7
	dAtA[i] = 0x2a
	i++
	i = encodeVarintRpc(dAtA, i, uint64(types.SizeOfStdTime(m.LatestBlockSeenTime)))
	n8, err := types.StdTimeMarshalTo(m.LatestBlockSeenTime, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n8
	if m.CatchingUp {
		dAtA[i] = 0x30
		i++
		if m.CatchingUp {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeVarintRpc(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ResultStatus) Size() (n int) {
	var l int
	_ = l
	l = len(m.ChainID)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.RunID)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.BurrowVersion)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = m.GenesisHash.Size()
	n += 1 + l + sovRpc(uint64(l))
	if m.NodeInfo != nil {
		l = m.NodeInfo.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.SyncInfo != nil {
		l = m.SyncInfo.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.ValidatorInfo != nil {
		l = m.ValidatorInfo.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}

func (m *SyncInfo) Size() (n int) {
	var l int
	_ = l
	if m.LatestBlockHeight != 0 {
		n += 1 + sovRpc(uint64(m.LatestBlockHeight))
	}
	l = m.LatestBlockHash.Size()
	n += 1 + l + sovRpc(uint64(l))
	l = m.LatestAppHash.Size()
	n += 1 + l + sovRpc(uint64(l))
	l = types.SizeOfStdTime(m.LatestBlockTime)
	n += 1 + l + sovRpc(uint64(l))
	l = types.SizeOfStdTime(m.LatestBlockSeenTime)
	n += 1 + l + sovRpc(uint64(l))
	if m.CatchingUp {
		n += 2
	}
	return n
}

func sovRpc(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRpc(x uint64) (n int) {
	return sovRpc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ResultStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ResultStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResultStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RunID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RunID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurrowVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BurrowVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GenesisHash.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NodeInfo == nil {
				m.NodeInfo = &tendermint.NodeInfo{}
			}
			if err := m.NodeInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SyncInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SyncInfo == nil {
				m.SyncInfo = &SyncInfo{}
			}
			if err := m.SyncInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ValidatorInfo == nil {
				m.ValidatorInfo = &validator.Validator{}
			}
			if err := m.ValidatorInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
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
func (m *SyncInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SyncInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SyncInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestBlockHeight", wireType)
			}
			m.LatestBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LatestBlockHeight |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestBlockHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LatestBlockHash.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestAppHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LatestAppHash.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestBlockTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := types.StdTimeUnmarshal(&m.LatestBlockTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestBlockSeenTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := types.StdTimeUnmarshal(&m.LatestBlockSeenTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CatchingUp", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CatchingUp = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
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
func skipRpc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRpc
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
					return 0, ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRpc
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRpc
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRpc
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRpc(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRpc = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRpc   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("rpc.proto", fileDescriptorRpc) }
func init() { golang_proto.RegisterFile("rpc.proto", fileDescriptorRpc) }

var fileDescriptorRpc = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x3f, 0x6f, 0xd3, 0x40,
	0x18, 0xc6, 0x7b, 0xf9, 0xd3, 0xa6, 0x97, 0x46, 0x85, 0xa3, 0x43, 0x94, 0xc1, 0x0e, 0x15, 0x43,
	0x18, 0x70, 0x50, 0x10, 0x42, 0x62, 0xc3, 0x45, 0x22, 0x95, 0x50, 0x87, 0x4b, 0x09, 0x12, 0x0c,
	0xc8, 0x76, 0xde, 0xda, 0x27, 0xec, 0x3b, 0xeb, 0xee, 0x0c, 0xe4, 0x5b, 0xf0, 0x91, 0x18, 0x18,
	0x32, 0x32, 0x31, 0x30, 0x04, 0x94, 0x6e, 0x7c, 0x0a, 0x94, 0x73, 0x9c, 0x38, 0x05, 0x21, 0x55,
	0xdd, 0xfc, 0x3e, 0xcf, 0xfb, 0xfe, 0xee, 0xfc, 0xdc, 0x8b, 0xf7, 0x65, 0x1a, 0x38, 0xa9, 0x14,
	0x5a, 0x90, 0xaa, 0x4c, 0x83, 0xce, 0x83, 0x90, 0xe9, 0x28, 0xf3, 0x9d, 0x40, 0x24, 0xfd, 0x50,
	0x84, 0xa2, 0x6f, 0x3c, 0x3f, 0xbb, 0x30, 0x95, 0x29, 0xcc, 0x57, 0x3e, 0xd3, 0xb9, 0xa5, 0x81,
	0x4f, 0x40, 0x26, 0x8c, 0xeb, 0x95, 0x72, 0xf8, 0xc1, 0x8b, 0xd9, 0xc4, 0xd3, 0x42, 0xae, 0x04,
	0x3b, 0x14, 0x22, 0x8c, 0x61, 0x03, 0xd2, 0x2c, 0x01, 0xa5, 0xbd, 0x24, 0xcd, 0x1b, 0x8e, 0xbf,
	0x57, 0xf0, 0x01, 0x05, 0x95, 0xc5, 0x7a, 0xa4, 0x3d, 0x9d, 0x29, 0xd2, 0xc6, 0x7b, 0x27, 0x91,
	0xc7, 0xf8, 0xe9, 0xf3, 0x36, 0xea, 0xa2, 0xde, 0x3e, 0x2d, 0x4a, 0x72, 0x84, 0xeb, 0x34, 0x5b,
	0xea, 0x15, 0xa3, 0xe7, 0x05, 0xb9, 0x87, 0x5b, 0x6e, 0x26, 0xa5, 0xf8, 0x38, 0x06, 0xa9, 0x98,
	0xe0, 0xed, 0xaa, 0x71, 0xb7, 0x45, 0xf2, 0x1a, 0x37, 0x5f, 0x00, 0x07, 0xc5, 0xd4, 0xd0, 0x53,
	0x51, 0xbb, 0xd6, 0x45, 0xbd, 0x03, 0xf7, 0xf1, 0x6c, 0x6e, 0xef, 0xfc, 0x98, 0xdb, 0xe5, 0xdf,
	0x8e, 0xa6, 0x29, 0xc8, 0x18, 0x26, 0x21, 0xc8, 0xbe, 0x6f, 0x10, 0x7d, 0x9f, 0x71, 0x4f, 0x4e,
	0x9d, 0x21, 0x7c, 0x72, 0xa7, 0x1a, 0x14, 0x2d, 0x93, 0xc8, 0x43, 0xdc, 0x38, 0x13, 0x13, 0x38,
	0xe5, 0x17, 0xa2, 0x5d, 0xef, 0xa2, 0x5e, 0x73, 0x70, 0xe4, 0x94, 0x62, 0x29, 0x3c, 0xba, 0xee,
	0x22, 0xf7, 0x71, 0x63, 0x34, 0xe5, 0x81, 0x99, 0xd8, 0x35, 0x13, 0x2d, 0x67, 0xf9, 0x0e, 0x85,
	0x48, 0xd7, 0x36, 0x79, 0x8a, 0x5b, 0xe3, 0x22, 0x50, 0xd3, 0xbf, 0xb7, 0x3a, 0x61, 0x13, 0xf3,
	0xda, 0xa7, 0xdb, 0xad, 0xc7, 0x5f, 0xab, 0x9b, 0x73, 0xc8, 0x00, 0xdf, 0x7e, 0xe9, 0x69, 0x50,
	0xda, 0x8d, 0x45, 0xf0, 0x7e, 0x08, 0x2c, 0x8c, 0xb4, 0x89, 0xb7, 0xe6, 0xd6, 0x7e, 0xcf, 0xed,
	0x1d, 0xfa, 0xb7, 0x4d, 0xde, 0xe1, 0xc3, 0xb2, 0xb8, 0x8c, 0xad, 0x72, 0x93, 0xd8, 0xae, 0xd2,
	0xc8, 0x5b, 0xdc, 0xca, 0xa5, 0x67, 0x69, 0x6a, 0xf0, 0xd5, 0x9b, 0xe0, 0xb7, 0x59, 0xe4, 0x6c,
	0xeb, 0xf6, 0xe7, 0x2c, 0x01, 0xf3, 0xe8, 0xcd, 0x41, 0xc7, 0xc9, 0x57, 0xd2, 0x29, 0x56, 0xd2,
	0x39, 0x2f, 0x56, 0xd2, 0x6d, 0x2c, 0x8f, 0xfe, 0xfc, 0xd3, 0x46, 0xf4, 0xea, 0x30, 0x19, 0xe3,
	0x3b, 0x25, 0x69, 0x04, 0xc0, 0x0d, 0xb3, 0x7e, 0x0d, 0xe6, 0xbf, 0x00, 0xc4, 0xc2, 0xf8, 0xc4,
	0xd3, 0x41, 0xc4, 0x78, 0xf8, 0x2a, 0x35, 0xfb, 0xd0, 0xa0, 0x25, 0xc5, 0x7d, 0x32, 0x5b, 0x58,
	0xe8, 0xdb, 0xc2, 0x42, 0xbf, 0x16, 0x16, 0xfa, 0x72, 0x69, 0xa1, 0xd9, 0xa5, 0x85, 0xde, 0xdc,
	0xfd, 0x7f, 0x36, 0x32, 0x0d, 0xfc, 0x5d, 0x73, 0x97, 0x47, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x7a, 0x31, 0xa5, 0x7d, 0xe4, 0x03, 0x00, 0x00,
}