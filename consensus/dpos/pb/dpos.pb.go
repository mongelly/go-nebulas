// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dpos.proto

/*
Package dpospb is a generated protocol buffer package.

It is generated from these files:
	dpos.proto

It has these top-level messages:
	DposState
*/
package dpospb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type DposState struct {
	DynastyRoot     []byte `protobuf:"bytes,1,opt,name=dynasty_root,json=dynastyRoot,proto3" json:"dynasty_root,omitempty"`
	NextDynastyRoot []byte `protobuf:"bytes,2,opt,name=next_dynasty_root,json=nextDynastyRoot,proto3" json:"next_dynasty_root,omitempty"`
	DelegateRoot    []byte `protobuf:"bytes,3,opt,name=delegate_root,json=delegateRoot,proto3" json:"delegate_root,omitempty"`
	CandidateRoot   []byte `protobuf:"bytes,4,opt,name=candidate_root,json=candidateRoot,proto3" json:"candidate_root,omitempty"`
	VoteRoot        []byte `protobuf:"bytes,5,opt,name=vote_root,json=voteRoot,proto3" json:"vote_root,omitempty"`
	MintCntRoot     []byte `protobuf:"bytes,6,opt,name=mint_cnt_root,json=mintCntRoot,proto3" json:"mint_cnt_root,omitempty"`
}

func (m *DposState) Reset()                    { *m = DposState{} }
func (m *DposState) String() string            { return proto.CompactTextString(m) }
func (*DposState) ProtoMessage()               {}
func (*DposState) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{0} }

func (m *DposState) GetDynastyRoot() []byte {
	if m != nil {
		return m.DynastyRoot
	}
	return nil
}

func (m *DposState) GetNextDynastyRoot() []byte {
	if m != nil {
		return m.NextDynastyRoot
	}
	return nil
}

func (m *DposState) GetDelegateRoot() []byte {
	if m != nil {
		return m.DelegateRoot
	}
	return nil
}

func (m *DposState) GetCandidateRoot() []byte {
	if m != nil {
		return m.CandidateRoot
	}
	return nil
}

func (m *DposState) GetVoteRoot() []byte {
	if m != nil {
		return m.VoteRoot
	}
	return nil
}

func (m *DposState) GetMintCntRoot() []byte {
	if m != nil {
		return m.MintCntRoot
	}
	return nil
}

func init() {
	proto.RegisterType((*DposState)(nil), "dpospb.DposState")
}

func init() { proto.RegisterFile("dpos.proto", fileDescriptorDpos) }

var fileDescriptorDpos = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0xcf, 0xc1, 0xea, 0x82, 0x40,
	0x10, 0xc7, 0x71, 0xfc, 0xff, 0x4b, 0x72, 0xd2, 0x22, 0x4f, 0x41, 0x97, 0x32, 0x82, 0xe8, 0xd0,
	0xa5, 0x47, 0xc8, 0x27, 0xb0, 0x07, 0x90, 0xd5, 0x5d, 0x42, 0xa8, 0x99, 0x45, 0x87, 0xc8, 0x27,
	0xee, 0x35, 0x62, 0xc7, 0x94, 0xba, 0x2d, 0xdf, 0xdf, 0xe7, 0xb0, 0x03, 0xa0, 0x2d, 0x35, 0x47,
	0x5b, 0x13, 0x53, 0xec, 0xbb, 0xb7, 0x2d, 0x92, 0x97, 0x07, 0x41, 0x6a, 0xa9, 0xb9, 0xb0, 0x62,
	0x13, 0x6f, 0x20, 0xd4, 0x2d, 0xaa, 0x86, 0xdb, 0xbc, 0x26, 0xe2, 0xa5, 0xb7, 0xf6, 0xf6, 0x61,
	0x36, 0xfd, 0xb4, 0x8c, 0x88, 0xe3, 0x03, 0x2c, 0xd0, 0x3c, 0x39, 0xff, 0x71, 0x7f, 0xe2, 0xe6,
	0x6e, 0x48, 0xbf, 0xec, 0x16, 0x22, 0x6d, 0x6e, 0xe6, 0xaa, 0xd8, 0x74, 0xee, 0x5f, 0x5c, 0xd8,
	0x47, 0x41, 0x3b, 0x98, 0x95, 0x0a, 0x75, 0xa5, 0x07, 0x35, 0x12, 0x15, 0x0d, 0x55, 0xd8, 0x0a,
	0x82, 0x07, 0xf5, 0x62, 0x2c, 0x62, 0xe2, 0x82, 0x8c, 0x09, 0x44, 0xf7, 0x0a, 0x39, 0x2f, 0x91,
	0x3b, 0xe0, 0x77, 0x1f, 0x77, 0xf1, 0x8c, 0xec, 0x4c, 0xe1, 0xcb, 0xe1, 0xa7, 0x77, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xad, 0xd1, 0x4b, 0x9d, 0x06, 0x01, 0x00, 0x00,
}
