// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              21.06-release
// source: /usr/share/vpp/api/core/vxlan_gbp.api.json

// Package vxlan_gbp contains generated bindings for API file vxlan_gbp.api.
//
// Contents:
//   1 enum
//   1 struct
//   6 messages
//
package vxlan_gbp

import (
	"strconv"

	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
	interface_types "github.com/edwarnicke/govpp/binapi/interface_types"
	ip_types "github.com/edwarnicke/govpp/binapi/ip_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "vxlan_gbp"
	APIVersion = "1.1.1"
	VersionCrc = 0xb40203a1
)

// VxlanGbpAPITunnelMode defines enum 'vxlan_gbp_api_tunnel_mode'.
type VxlanGbpAPITunnelMode uint32

const (
	VXLAN_GBP_API_TUNNEL_MODE_L2 VxlanGbpAPITunnelMode = 1
	VXLAN_GBP_API_TUNNEL_MODE_L3 VxlanGbpAPITunnelMode = 2
)

var (
	VxlanGbpAPITunnelMode_name = map[uint32]string{
		1: "VXLAN_GBP_API_TUNNEL_MODE_L2",
		2: "VXLAN_GBP_API_TUNNEL_MODE_L3",
	}
	VxlanGbpAPITunnelMode_value = map[string]uint32{
		"VXLAN_GBP_API_TUNNEL_MODE_L2": 1,
		"VXLAN_GBP_API_TUNNEL_MODE_L3": 2,
	}
)

func (x VxlanGbpAPITunnelMode) String() string {
	s, ok := VxlanGbpAPITunnelMode_name[uint32(x)]
	if ok {
		return s
	}
	return "VxlanGbpAPITunnelMode(" + strconv.Itoa(int(x)) + ")"
}

// VxlanGbpTunnel defines type 'vxlan_gbp_tunnel'.
type VxlanGbpTunnel struct {
	Instance       uint32                         `binapi:"u32,name=instance" json:"instance,omitempty"`
	Src            ip_types.Address               `binapi:"address,name=src" json:"src,omitempty"`
	Dst            ip_types.Address               `binapi:"address,name=dst" json:"dst,omitempty"`
	McastSwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=mcast_sw_if_index" json:"mcast_sw_if_index,omitempty"`
	EncapTableID   uint32                         `binapi:"u32,name=encap_table_id" json:"encap_table_id,omitempty"`
	Vni            uint32                         `binapi:"u32,name=vni" json:"vni,omitempty"`
	SwIfIndex      interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	Mode           VxlanGbpAPITunnelMode          `binapi:"vxlan_gbp_api_tunnel_mode,name=mode" json:"mode,omitempty"`
}

// SwInterfaceSetVxlanGbpBypass defines message 'sw_interface_set_vxlan_gbp_bypass'.
// InProgress: the message form may change in the future versions
type SwInterfaceSetVxlanGbpBypass struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	IsIPv6    bool                           `binapi:"bool,name=is_ipv6" json:"is_ipv6,omitempty"`
	Enable    bool                           `binapi:"bool,name=enable,default=true" json:"enable,omitempty"`
}

func (m *SwInterfaceSetVxlanGbpBypass) Reset() { *m = SwInterfaceSetVxlanGbpBypass{} }
func (*SwInterfaceSetVxlanGbpBypass) GetMessageName() string {
	return "sw_interface_set_vxlan_gbp_bypass"
}
func (*SwInterfaceSetVxlanGbpBypass) GetCrcString() string { return "65247409" }
func (*SwInterfaceSetVxlanGbpBypass) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SwInterfaceSetVxlanGbpBypass) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	size += 1 // m.IsIPv6
	size += 1 // m.Enable
	return size
}
func (m *SwInterfaceSetVxlanGbpBypass) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeBool(m.IsIPv6)
	buf.EncodeBool(m.Enable)
	return buf.Bytes(), nil
}
func (m *SwInterfaceSetVxlanGbpBypass) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.IsIPv6 = buf.DecodeBool()
	m.Enable = buf.DecodeBool()
	return nil
}

// SwInterfaceSetVxlanGbpBypassReply defines message 'sw_interface_set_vxlan_gbp_bypass_reply'.
// InProgress: the message form may change in the future versions
type SwInterfaceSetVxlanGbpBypassReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SwInterfaceSetVxlanGbpBypassReply) Reset() { *m = SwInterfaceSetVxlanGbpBypassReply{} }
func (*SwInterfaceSetVxlanGbpBypassReply) GetMessageName() string {
	return "sw_interface_set_vxlan_gbp_bypass_reply"
}
func (*SwInterfaceSetVxlanGbpBypassReply) GetCrcString() string { return "e8d4e804" }
func (*SwInterfaceSetVxlanGbpBypassReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SwInterfaceSetVxlanGbpBypassReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SwInterfaceSetVxlanGbpBypassReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SwInterfaceSetVxlanGbpBypassReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// VxlanGbpTunnelAddDel defines message 'vxlan_gbp_tunnel_add_del'.
// InProgress: the message form may change in the future versions
type VxlanGbpTunnelAddDel struct {
	IsAdd  bool           `binapi:"bool,name=is_add,default=true" json:"is_add,omitempty"`
	Tunnel VxlanGbpTunnel `binapi:"vxlan_gbp_tunnel,name=tunnel" json:"tunnel,omitempty"`
}

func (m *VxlanGbpTunnelAddDel) Reset()               { *m = VxlanGbpTunnelAddDel{} }
func (*VxlanGbpTunnelAddDel) GetMessageName() string { return "vxlan_gbp_tunnel_add_del" }
func (*VxlanGbpTunnelAddDel) GetCrcString() string   { return "6c743427" }
func (*VxlanGbpTunnelAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *VxlanGbpTunnelAddDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.IsAdd
	size += 4      // m.Tunnel.Instance
	size += 1      // m.Tunnel.Src.Af
	size += 1 * 16 // m.Tunnel.Src.Un
	size += 1      // m.Tunnel.Dst.Af
	size += 1 * 16 // m.Tunnel.Dst.Un
	size += 4      // m.Tunnel.McastSwIfIndex
	size += 4      // m.Tunnel.EncapTableID
	size += 4      // m.Tunnel.Vni
	size += 4      // m.Tunnel.SwIfIndex
	size += 4      // m.Tunnel.Mode
	return size
}
func (m *VxlanGbpTunnelAddDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsAdd)
	buf.EncodeUint32(m.Tunnel.Instance)
	buf.EncodeUint8(uint8(m.Tunnel.Src.Af))
	buf.EncodeBytes(m.Tunnel.Src.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(uint8(m.Tunnel.Dst.Af))
	buf.EncodeBytes(m.Tunnel.Dst.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(uint32(m.Tunnel.McastSwIfIndex))
	buf.EncodeUint32(m.Tunnel.EncapTableID)
	buf.EncodeUint32(m.Tunnel.Vni)
	buf.EncodeUint32(uint32(m.Tunnel.SwIfIndex))
	buf.EncodeUint32(uint32(m.Tunnel.Mode))
	return buf.Bytes(), nil
}
func (m *VxlanGbpTunnelAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeBool()
	m.Tunnel.Instance = buf.DecodeUint32()
	m.Tunnel.Src.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Tunnel.Src.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Tunnel.Dst.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Tunnel.Dst.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Tunnel.McastSwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Tunnel.EncapTableID = buf.DecodeUint32()
	m.Tunnel.Vni = buf.DecodeUint32()
	m.Tunnel.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Tunnel.Mode = VxlanGbpAPITunnelMode(buf.DecodeUint32())
	return nil
}

// VxlanGbpTunnelAddDelReply defines message 'vxlan_gbp_tunnel_add_del_reply'.
// InProgress: the message form may change in the future versions
type VxlanGbpTunnelAddDelReply struct {
	Retval    int32                          `binapi:"i32,name=retval" json:"retval,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *VxlanGbpTunnelAddDelReply) Reset()               { *m = VxlanGbpTunnelAddDelReply{} }
func (*VxlanGbpTunnelAddDelReply) GetMessageName() string { return "vxlan_gbp_tunnel_add_del_reply" }
func (*VxlanGbpTunnelAddDelReply) GetCrcString() string   { return "5383d31f" }
func (*VxlanGbpTunnelAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *VxlanGbpTunnelAddDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.SwIfIndex
	return size
}
func (m *VxlanGbpTunnelAddDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *VxlanGbpTunnelAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// VxlanGbpTunnelDetails defines message 'vxlan_gbp_tunnel_details'.
// InProgress: the message form may change in the future versions
type VxlanGbpTunnelDetails struct {
	Tunnel VxlanGbpTunnel `binapi:"vxlan_gbp_tunnel,name=tunnel" json:"tunnel,omitempty"`
}

func (m *VxlanGbpTunnelDetails) Reset()               { *m = VxlanGbpTunnelDetails{} }
func (*VxlanGbpTunnelDetails) GetMessageName() string { return "vxlan_gbp_tunnel_details" }
func (*VxlanGbpTunnelDetails) GetCrcString() string   { return "66e94a89" }
func (*VxlanGbpTunnelDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *VxlanGbpTunnelDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.Tunnel.Instance
	size += 1      // m.Tunnel.Src.Af
	size += 1 * 16 // m.Tunnel.Src.Un
	size += 1      // m.Tunnel.Dst.Af
	size += 1 * 16 // m.Tunnel.Dst.Un
	size += 4      // m.Tunnel.McastSwIfIndex
	size += 4      // m.Tunnel.EncapTableID
	size += 4      // m.Tunnel.Vni
	size += 4      // m.Tunnel.SwIfIndex
	size += 4      // m.Tunnel.Mode
	return size
}
func (m *VxlanGbpTunnelDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Tunnel.Instance)
	buf.EncodeUint8(uint8(m.Tunnel.Src.Af))
	buf.EncodeBytes(m.Tunnel.Src.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(uint8(m.Tunnel.Dst.Af))
	buf.EncodeBytes(m.Tunnel.Dst.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(uint32(m.Tunnel.McastSwIfIndex))
	buf.EncodeUint32(m.Tunnel.EncapTableID)
	buf.EncodeUint32(m.Tunnel.Vni)
	buf.EncodeUint32(uint32(m.Tunnel.SwIfIndex))
	buf.EncodeUint32(uint32(m.Tunnel.Mode))
	return buf.Bytes(), nil
}
func (m *VxlanGbpTunnelDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Tunnel.Instance = buf.DecodeUint32()
	m.Tunnel.Src.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Tunnel.Src.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Tunnel.Dst.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Tunnel.Dst.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Tunnel.McastSwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Tunnel.EncapTableID = buf.DecodeUint32()
	m.Tunnel.Vni = buf.DecodeUint32()
	m.Tunnel.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Tunnel.Mode = VxlanGbpAPITunnelMode(buf.DecodeUint32())
	return nil
}

// VxlanGbpTunnelDump defines message 'vxlan_gbp_tunnel_dump'.
// InProgress: the message form may change in the future versions
type VxlanGbpTunnelDump struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index,default=4294967295" json:"sw_if_index,omitempty"`
}

func (m *VxlanGbpTunnelDump) Reset()               { *m = VxlanGbpTunnelDump{} }
func (*VxlanGbpTunnelDump) GetMessageName() string { return "vxlan_gbp_tunnel_dump" }
func (*VxlanGbpTunnelDump) GetCrcString() string   { return "f9e6675e" }
func (*VxlanGbpTunnelDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *VxlanGbpTunnelDump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *VxlanGbpTunnelDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *VxlanGbpTunnelDump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

func init() { file_vxlan_gbp_binapi_init() }
func file_vxlan_gbp_binapi_init() {
	api.RegisterMessage((*SwInterfaceSetVxlanGbpBypass)(nil), "sw_interface_set_vxlan_gbp_bypass_65247409")
	api.RegisterMessage((*SwInterfaceSetVxlanGbpBypassReply)(nil), "sw_interface_set_vxlan_gbp_bypass_reply_e8d4e804")
	api.RegisterMessage((*VxlanGbpTunnelAddDel)(nil), "vxlan_gbp_tunnel_add_del_6c743427")
	api.RegisterMessage((*VxlanGbpTunnelAddDelReply)(nil), "vxlan_gbp_tunnel_add_del_reply_5383d31f")
	api.RegisterMessage((*VxlanGbpTunnelDetails)(nil), "vxlan_gbp_tunnel_details_66e94a89")
	api.RegisterMessage((*VxlanGbpTunnelDump)(nil), "vxlan_gbp_tunnel_dump_f9e6675e")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*SwInterfaceSetVxlanGbpBypass)(nil),
		(*SwInterfaceSetVxlanGbpBypassReply)(nil),
		(*VxlanGbpTunnelAddDel)(nil),
		(*VxlanGbpTunnelAddDelReply)(nil),
		(*VxlanGbpTunnelDetails)(nil),
		(*VxlanGbpTunnelDump)(nil),
	}
}
