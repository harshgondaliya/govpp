// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              20.09-release
// source: /usr/share/vpp/api/core/tcp.api.json

// Package tcp contains generated bindings for API file tcp.api.
//
// Contents:
//   2 messages
//
package tcp

import (
	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
	ip_types "github.com/edwarnicke/govpp/pkg/v2009/binapi/ip_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "tcp"
	APIVersion = "2.0.0"
	VersionCrc = 0x8264ed4a
)

// TCPConfigureSrcAddresses defines message 'tcp_configure_src_addresses'.
type TCPConfigureSrcAddresses struct {
	VrfID        uint32           `binapi:"u32,name=vrf_id" json:"vrf_id,omitempty"`
	FirstAddress ip_types.Address `binapi:"address,name=first_address" json:"first_address,omitempty"`
	LastAddress  ip_types.Address `binapi:"address,name=last_address" json:"last_address,omitempty"`
}

func (m *TCPConfigureSrcAddresses) Reset()               { *m = TCPConfigureSrcAddresses{} }
func (*TCPConfigureSrcAddresses) GetMessageName() string { return "tcp_configure_src_addresses" }
func (*TCPConfigureSrcAddresses) GetCrcString() string   { return "4b02b946" }
func (*TCPConfigureSrcAddresses) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *TCPConfigureSrcAddresses) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.VrfID
	size += 1      // m.FirstAddress.Af
	size += 1 * 16 // m.FirstAddress.Un
	size += 1      // m.LastAddress.Af
	size += 1 * 16 // m.LastAddress.Un
	return size
}
func (m *TCPConfigureSrcAddresses) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.VrfID)
	buf.EncodeUint8(uint8(m.FirstAddress.Af))
	buf.EncodeBytes(m.FirstAddress.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(uint8(m.LastAddress.Af))
	buf.EncodeBytes(m.LastAddress.Un.XXX_UnionData[:], 16)
	return buf.Bytes(), nil
}
func (m *TCPConfigureSrcAddresses) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.VrfID = buf.DecodeUint32()
	m.FirstAddress.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.FirstAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.LastAddress.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.LastAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	return nil
}

// TCPConfigureSrcAddressesReply defines message 'tcp_configure_src_addresses_reply'.
type TCPConfigureSrcAddressesReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *TCPConfigureSrcAddressesReply) Reset() { *m = TCPConfigureSrcAddressesReply{} }
func (*TCPConfigureSrcAddressesReply) GetMessageName() string {
	return "tcp_configure_src_addresses_reply"
}
func (*TCPConfigureSrcAddressesReply) GetCrcString() string { return "e8d4e804" }
func (*TCPConfigureSrcAddressesReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *TCPConfigureSrcAddressesReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *TCPConfigureSrcAddressesReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *TCPConfigureSrcAddressesReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_tcp_binapi_init() }
func file_tcp_binapi_init() {
	api.RegisterMessage((*TCPConfigureSrcAddresses)(nil), "tcp_configure_src_addresses_4b02b946")
	api.RegisterMessage((*TCPConfigureSrcAddressesReply)(nil), "tcp_configure_src_addresses_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*TCPConfigureSrcAddresses)(nil),
		(*TCPConfigureSrcAddressesReply)(nil),
	}
}