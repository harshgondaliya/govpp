// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              21.06-release
// source: /usr/share/vpp/api/core/qos.api.json

// Package qos contains generated bindings for API file qos.api.
//
// Contents:
//   1 enum
//   5 structs
//  19 messages
//
package qos

import (
	"strconv"

	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
	interface_types "github.com/edwarnicke/govpp/binapi/interface_types"
	_ "github.com/edwarnicke/govpp/binapi/ip_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "qos"
	APIVersion = "1.1.1"
	VersionCrc = 0x7b7b5955
)

// QosSource defines enum 'qos_source'.
type QosSource uint8

const (
	QOS_API_SOURCE_EXT  QosSource = 0
	QOS_API_SOURCE_VLAN QosSource = 1
	QOS_API_SOURCE_MPLS QosSource = 2
	QOS_API_SOURCE_IP   QosSource = 3
)

var (
	QosSource_name = map[uint8]string{
		0: "QOS_API_SOURCE_EXT",
		1: "QOS_API_SOURCE_VLAN",
		2: "QOS_API_SOURCE_MPLS",
		3: "QOS_API_SOURCE_IP",
	}
	QosSource_value = map[string]uint8{
		"QOS_API_SOURCE_EXT":  0,
		"QOS_API_SOURCE_VLAN": 1,
		"QOS_API_SOURCE_MPLS": 2,
		"QOS_API_SOURCE_IP":   3,
	}
)

func (x QosSource) String() string {
	s, ok := QosSource_name[uint8(x)]
	if ok {
		return s
	}
	return "QosSource(" + strconv.Itoa(int(x)) + ")"
}

// QosEgressMap defines type 'qos_egress_map'.
type QosEgressMap struct {
	ID   uint32             `binapi:"u32,name=id" json:"id,omitempty"`
	Rows [4]QosEgressMapRow `binapi:"qos_egress_map_row[4],name=rows" json:"rows,omitempty"`
}

// QosEgressMapRow defines type 'qos_egress_map_row'.
type QosEgressMapRow struct {
	Outputs []byte `binapi:"u8[256],name=outputs" json:"outputs,omitempty"`
}

// QosMark defines type 'qos_mark'.
type QosMark struct {
	SwIfIndex    uint32    `binapi:"u32,name=sw_if_index" json:"sw_if_index,omitempty"`
	MapID        uint32    `binapi:"u32,name=map_id" json:"map_id,omitempty"`
	OutputSource QosSource `binapi:"qos_source,name=output_source" json:"output_source,omitempty"`
}

// QosRecord defines type 'qos_record'.
type QosRecord struct {
	SwIfIndex   interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	InputSource QosSource                      `binapi:"qos_source,name=input_source" json:"input_source,omitempty"`
}

// QosStore defines type 'qos_store'.
type QosStore struct {
	SwIfIndex   interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	InputSource QosSource                      `binapi:"qos_source,name=input_source" json:"input_source,omitempty"`
	Value       uint8                          `binapi:"u8,name=value" json:"value,omitempty"`
}

// QosEgressMapDelete defines message 'qos_egress_map_delete'.
type QosEgressMapDelete struct {
	ID uint32 `binapi:"u32,name=id" json:"id,omitempty"`
}

func (m *QosEgressMapDelete) Reset()               { *m = QosEgressMapDelete{} }
func (*QosEgressMapDelete) GetMessageName() string { return "qos_egress_map_delete" }
func (*QosEgressMapDelete) GetCrcString() string   { return "3a91bde5" }
func (*QosEgressMapDelete) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *QosEgressMapDelete) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.ID
	return size
}
func (m *QosEgressMapDelete) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.ID)
	return buf.Bytes(), nil
}
func (m *QosEgressMapDelete) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.ID = buf.DecodeUint32()
	return nil
}

// QosEgressMapDeleteReply defines message 'qos_egress_map_delete_reply'.
type QosEgressMapDeleteReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *QosEgressMapDeleteReply) Reset()               { *m = QosEgressMapDeleteReply{} }
func (*QosEgressMapDeleteReply) GetMessageName() string { return "qos_egress_map_delete_reply" }
func (*QosEgressMapDeleteReply) GetCrcString() string   { return "e8d4e804" }
func (*QosEgressMapDeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *QosEgressMapDeleteReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *QosEgressMapDeleteReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *QosEgressMapDeleteReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// QosEgressMapDetails defines message 'qos_egress_map_details'.
type QosEgressMapDetails struct {
	Map QosEgressMap `binapi:"qos_egress_map,name=map" json:"map,omitempty"`
}

func (m *QosEgressMapDetails) Reset()               { *m = QosEgressMapDetails{} }
func (*QosEgressMapDetails) GetMessageName() string { return "qos_egress_map_details" }
func (*QosEgressMapDetails) GetCrcString() string   { return "46c5653c" }
func (*QosEgressMapDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *QosEgressMapDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Map.ID
	for j2 := 0; j2 < 4; j2++ {
		size += 1 * 256 // m.Map.Rows[j2].Outputs
	}
	return size
}
func (m *QosEgressMapDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Map.ID)
	for j1 := 0; j1 < 4; j1++ {
		buf.EncodeBytes(m.Map.Rows[j1].Outputs, 256)
	}
	return buf.Bytes(), nil
}
func (m *QosEgressMapDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Map.ID = buf.DecodeUint32()
	for j1 := 0; j1 < 4; j1++ {
		m.Map.Rows[j1].Outputs = make([]byte, 256)
		copy(m.Map.Rows[j1].Outputs, buf.DecodeBytes(len(m.Map.Rows[j1].Outputs)))
	}
	return nil
}

// QosEgressMapDump defines message 'qos_egress_map_dump'.
type QosEgressMapDump struct{}

func (m *QosEgressMapDump) Reset()               { *m = QosEgressMapDump{} }
func (*QosEgressMapDump) GetMessageName() string { return "qos_egress_map_dump" }
func (*QosEgressMapDump) GetCrcString() string   { return "51077d14" }
func (*QosEgressMapDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *QosEgressMapDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *QosEgressMapDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *QosEgressMapDump) Unmarshal(b []byte) error {
	return nil
}

// QosEgressMapUpdate defines message 'qos_egress_map_update'.
type QosEgressMapUpdate struct {
	Map QosEgressMap `binapi:"qos_egress_map,name=map" json:"map,omitempty"`
}

func (m *QosEgressMapUpdate) Reset()               { *m = QosEgressMapUpdate{} }
func (*QosEgressMapUpdate) GetMessageName() string { return "qos_egress_map_update" }
func (*QosEgressMapUpdate) GetCrcString() string   { return "6d1c065f" }
func (*QosEgressMapUpdate) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *QosEgressMapUpdate) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Map.ID
	for j2 := 0; j2 < 4; j2++ {
		size += 1 * 256 // m.Map.Rows[j2].Outputs
	}
	return size
}
func (m *QosEgressMapUpdate) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Map.ID)
	for j1 := 0; j1 < 4; j1++ {
		buf.EncodeBytes(m.Map.Rows[j1].Outputs, 256)
	}
	return buf.Bytes(), nil
}
func (m *QosEgressMapUpdate) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Map.ID = buf.DecodeUint32()
	for j1 := 0; j1 < 4; j1++ {
		m.Map.Rows[j1].Outputs = make([]byte, 256)
		copy(m.Map.Rows[j1].Outputs, buf.DecodeBytes(len(m.Map.Rows[j1].Outputs)))
	}
	return nil
}

// QosEgressMapUpdateReply defines message 'qos_egress_map_update_reply'.
type QosEgressMapUpdateReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *QosEgressMapUpdateReply) Reset()               { *m = QosEgressMapUpdateReply{} }
func (*QosEgressMapUpdateReply) GetMessageName() string { return "qos_egress_map_update_reply" }
func (*QosEgressMapUpdateReply) GetCrcString() string   { return "e8d4e804" }
func (*QosEgressMapUpdateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *QosEgressMapUpdateReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *QosEgressMapUpdateReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *QosEgressMapUpdateReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// QosMarkDetails defines message 'qos_mark_details'.
type QosMarkDetails struct {
	Mark QosMark `binapi:"qos_mark,name=mark" json:"mark,omitempty"`
}

func (m *QosMarkDetails) Reset()               { *m = QosMarkDetails{} }
func (*QosMarkDetails) GetMessageName() string { return "qos_mark_details" }
func (*QosMarkDetails) GetCrcString() string   { return "89fe81a9" }
func (*QosMarkDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *QosMarkDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Mark.SwIfIndex
	size += 4 // m.Mark.MapID
	size += 1 // m.Mark.OutputSource
	return size
}
func (m *QosMarkDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Mark.SwIfIndex)
	buf.EncodeUint32(m.Mark.MapID)
	buf.EncodeUint8(uint8(m.Mark.OutputSource))
	return buf.Bytes(), nil
}
func (m *QosMarkDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Mark.SwIfIndex = buf.DecodeUint32()
	m.Mark.MapID = buf.DecodeUint32()
	m.Mark.OutputSource = QosSource(buf.DecodeUint8())
	return nil
}

// QosMarkDetailsReply defines message 'qos_mark_details_reply'.
type QosMarkDetailsReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *QosMarkDetailsReply) Reset()               { *m = QosMarkDetailsReply{} }
func (*QosMarkDetailsReply) GetMessageName() string { return "qos_mark_details_reply" }
func (*QosMarkDetailsReply) GetCrcString() string   { return "e8d4e804" }
func (*QosMarkDetailsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *QosMarkDetailsReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *QosMarkDetailsReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *QosMarkDetailsReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// QosMarkDump defines message 'qos_mark_dump'.
type QosMarkDump struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *QosMarkDump) Reset()               { *m = QosMarkDump{} }
func (*QosMarkDump) GetMessageName() string { return "qos_mark_dump" }
func (*QosMarkDump) GetCrcString() string   { return "f9e6675e" }
func (*QosMarkDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *QosMarkDump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *QosMarkDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *QosMarkDump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// QosMarkEnableDisable defines message 'qos_mark_enable_disable'.
type QosMarkEnableDisable struct {
	Enable bool    `binapi:"bool,name=enable,default=true" json:"enable,omitempty"`
	Mark   QosMark `binapi:"qos_mark,name=mark" json:"mark,omitempty"`
}

func (m *QosMarkEnableDisable) Reset()               { *m = QosMarkEnableDisable{} }
func (*QosMarkEnableDisable) GetMessageName() string { return "qos_mark_enable_disable" }
func (*QosMarkEnableDisable) GetCrcString() string   { return "1a010f74" }
func (*QosMarkEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *QosMarkEnableDisable) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.Enable
	size += 4 // m.Mark.SwIfIndex
	size += 4 // m.Mark.MapID
	size += 1 // m.Mark.OutputSource
	return size
}
func (m *QosMarkEnableDisable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.Enable)
	buf.EncodeUint32(m.Mark.SwIfIndex)
	buf.EncodeUint32(m.Mark.MapID)
	buf.EncodeUint8(uint8(m.Mark.OutputSource))
	return buf.Bytes(), nil
}
func (m *QosMarkEnableDisable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Enable = buf.DecodeBool()
	m.Mark.SwIfIndex = buf.DecodeUint32()
	m.Mark.MapID = buf.DecodeUint32()
	m.Mark.OutputSource = QosSource(buf.DecodeUint8())
	return nil
}

// QosMarkEnableDisableReply defines message 'qos_mark_enable_disable_reply'.
type QosMarkEnableDisableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *QosMarkEnableDisableReply) Reset()               { *m = QosMarkEnableDisableReply{} }
func (*QosMarkEnableDisableReply) GetMessageName() string { return "qos_mark_enable_disable_reply" }
func (*QosMarkEnableDisableReply) GetCrcString() string   { return "e8d4e804" }
func (*QosMarkEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *QosMarkEnableDisableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *QosMarkEnableDisableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *QosMarkEnableDisableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// QosRecordDetails defines message 'qos_record_details'.
type QosRecordDetails struct {
	Record QosRecord `binapi:"qos_record,name=record" json:"record,omitempty"`
}

func (m *QosRecordDetails) Reset()               { *m = QosRecordDetails{} }
func (*QosRecordDetails) GetMessageName() string { return "qos_record_details" }
func (*QosRecordDetails) GetCrcString() string   { return "a425d4d3" }
func (*QosRecordDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *QosRecordDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Record.SwIfIndex
	size += 1 // m.Record.InputSource
	return size
}
func (m *QosRecordDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.Record.SwIfIndex))
	buf.EncodeUint8(uint8(m.Record.InputSource))
	return buf.Bytes(), nil
}
func (m *QosRecordDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Record.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Record.InputSource = QosSource(buf.DecodeUint8())
	return nil
}

// QosRecordDump defines message 'qos_record_dump'.
type QosRecordDump struct{}

func (m *QosRecordDump) Reset()               { *m = QosRecordDump{} }
func (*QosRecordDump) GetMessageName() string { return "qos_record_dump" }
func (*QosRecordDump) GetCrcString() string   { return "51077d14" }
func (*QosRecordDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *QosRecordDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *QosRecordDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *QosRecordDump) Unmarshal(b []byte) error {
	return nil
}

// QosRecordEnableDisable defines message 'qos_record_enable_disable'.
type QosRecordEnableDisable struct {
	Enable bool      `binapi:"bool,name=enable,default=true" json:"enable,omitempty"`
	Record QosRecord `binapi:"qos_record,name=record" json:"record,omitempty"`
}

func (m *QosRecordEnableDisable) Reset()               { *m = QosRecordEnableDisable{} }
func (*QosRecordEnableDisable) GetMessageName() string { return "qos_record_enable_disable" }
func (*QosRecordEnableDisable) GetCrcString() string   { return "2f1a4a38" }
func (*QosRecordEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *QosRecordEnableDisable) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.Enable
	size += 4 // m.Record.SwIfIndex
	size += 1 // m.Record.InputSource
	return size
}
func (m *QosRecordEnableDisable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.Enable)
	buf.EncodeUint32(uint32(m.Record.SwIfIndex))
	buf.EncodeUint8(uint8(m.Record.InputSource))
	return buf.Bytes(), nil
}
func (m *QosRecordEnableDisable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Enable = buf.DecodeBool()
	m.Record.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Record.InputSource = QosSource(buf.DecodeUint8())
	return nil
}

// QosRecordEnableDisableReply defines message 'qos_record_enable_disable_reply'.
type QosRecordEnableDisableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *QosRecordEnableDisableReply) Reset()               { *m = QosRecordEnableDisableReply{} }
func (*QosRecordEnableDisableReply) GetMessageName() string { return "qos_record_enable_disable_reply" }
func (*QosRecordEnableDisableReply) GetCrcString() string   { return "e8d4e804" }
func (*QosRecordEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *QosRecordEnableDisableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *QosRecordEnableDisableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *QosRecordEnableDisableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// QosStoreDetails defines message 'qos_store_details'.
type QosStoreDetails struct {
	Store QosStore `binapi:"qos_store,name=store" json:"store,omitempty"`
}

func (m *QosStoreDetails) Reset()               { *m = QosStoreDetails{} }
func (*QosStoreDetails) GetMessageName() string { return "qos_store_details" }
func (*QosStoreDetails) GetCrcString() string   { return "3ee0aad7" }
func (*QosStoreDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *QosStoreDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Store.SwIfIndex
	size += 1 // m.Store.InputSource
	size += 1 // m.Store.Value
	return size
}
func (m *QosStoreDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.Store.SwIfIndex))
	buf.EncodeUint8(uint8(m.Store.InputSource))
	buf.EncodeUint8(m.Store.Value)
	return buf.Bytes(), nil
}
func (m *QosStoreDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Store.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Store.InputSource = QosSource(buf.DecodeUint8())
	m.Store.Value = buf.DecodeUint8()
	return nil
}

// QosStoreDump defines message 'qos_store_dump'.
type QosStoreDump struct{}

func (m *QosStoreDump) Reset()               { *m = QosStoreDump{} }
func (*QosStoreDump) GetMessageName() string { return "qos_store_dump" }
func (*QosStoreDump) GetCrcString() string   { return "51077d14" }
func (*QosStoreDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *QosStoreDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *QosStoreDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *QosStoreDump) Unmarshal(b []byte) error {
	return nil
}

// QosStoreEnableDisable defines message 'qos_store_enable_disable'.
type QosStoreEnableDisable struct {
	Enable bool     `binapi:"bool,name=enable,default=true" json:"enable,omitempty"`
	Store  QosStore `binapi:"qos_store,name=store" json:"store,omitempty"`
}

func (m *QosStoreEnableDisable) Reset()               { *m = QosStoreEnableDisable{} }
func (*QosStoreEnableDisable) GetMessageName() string { return "qos_store_enable_disable" }
func (*QosStoreEnableDisable) GetCrcString() string   { return "f3abcc8b" }
func (*QosStoreEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *QosStoreEnableDisable) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.Enable
	size += 4 // m.Store.SwIfIndex
	size += 1 // m.Store.InputSource
	size += 1 // m.Store.Value
	return size
}
func (m *QosStoreEnableDisable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.Enable)
	buf.EncodeUint32(uint32(m.Store.SwIfIndex))
	buf.EncodeUint8(uint8(m.Store.InputSource))
	buf.EncodeUint8(m.Store.Value)
	return buf.Bytes(), nil
}
func (m *QosStoreEnableDisable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Enable = buf.DecodeBool()
	m.Store.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Store.InputSource = QosSource(buf.DecodeUint8())
	m.Store.Value = buf.DecodeUint8()
	return nil
}

// QosStoreEnableDisableReply defines message 'qos_store_enable_disable_reply'.
type QosStoreEnableDisableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *QosStoreEnableDisableReply) Reset()               { *m = QosStoreEnableDisableReply{} }
func (*QosStoreEnableDisableReply) GetMessageName() string { return "qos_store_enable_disable_reply" }
func (*QosStoreEnableDisableReply) GetCrcString() string   { return "e8d4e804" }
func (*QosStoreEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *QosStoreEnableDisableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *QosStoreEnableDisableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *QosStoreEnableDisableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_qos_binapi_init() }
func file_qos_binapi_init() {
	api.RegisterMessage((*QosEgressMapDelete)(nil), "qos_egress_map_delete_3a91bde5")
	api.RegisterMessage((*QosEgressMapDeleteReply)(nil), "qos_egress_map_delete_reply_e8d4e804")
	api.RegisterMessage((*QosEgressMapDetails)(nil), "qos_egress_map_details_46c5653c")
	api.RegisterMessage((*QosEgressMapDump)(nil), "qos_egress_map_dump_51077d14")
	api.RegisterMessage((*QosEgressMapUpdate)(nil), "qos_egress_map_update_6d1c065f")
	api.RegisterMessage((*QosEgressMapUpdateReply)(nil), "qos_egress_map_update_reply_e8d4e804")
	api.RegisterMessage((*QosMarkDetails)(nil), "qos_mark_details_89fe81a9")
	api.RegisterMessage((*QosMarkDetailsReply)(nil), "qos_mark_details_reply_e8d4e804")
	api.RegisterMessage((*QosMarkDump)(nil), "qos_mark_dump_f9e6675e")
	api.RegisterMessage((*QosMarkEnableDisable)(nil), "qos_mark_enable_disable_1a010f74")
	api.RegisterMessage((*QosMarkEnableDisableReply)(nil), "qos_mark_enable_disable_reply_e8d4e804")
	api.RegisterMessage((*QosRecordDetails)(nil), "qos_record_details_a425d4d3")
	api.RegisterMessage((*QosRecordDump)(nil), "qos_record_dump_51077d14")
	api.RegisterMessage((*QosRecordEnableDisable)(nil), "qos_record_enable_disable_2f1a4a38")
	api.RegisterMessage((*QosRecordEnableDisableReply)(nil), "qos_record_enable_disable_reply_e8d4e804")
	api.RegisterMessage((*QosStoreDetails)(nil), "qos_store_details_3ee0aad7")
	api.RegisterMessage((*QosStoreDump)(nil), "qos_store_dump_51077d14")
	api.RegisterMessage((*QosStoreEnableDisable)(nil), "qos_store_enable_disable_f3abcc8b")
	api.RegisterMessage((*QosStoreEnableDisableReply)(nil), "qos_store_enable_disable_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*QosEgressMapDelete)(nil),
		(*QosEgressMapDeleteReply)(nil),
		(*QosEgressMapDetails)(nil),
		(*QosEgressMapDump)(nil),
		(*QosEgressMapUpdate)(nil),
		(*QosEgressMapUpdateReply)(nil),
		(*QosMarkDetails)(nil),
		(*QosMarkDetailsReply)(nil),
		(*QosMarkDump)(nil),
		(*QosMarkEnableDisable)(nil),
		(*QosMarkEnableDisableReply)(nil),
		(*QosRecordDetails)(nil),
		(*QosRecordDump)(nil),
		(*QosRecordEnableDisable)(nil),
		(*QosRecordEnableDisableReply)(nil),
		(*QosStoreDetails)(nil),
		(*QosStoreDump)(nil),
		(*QosStoreEnableDisable)(nil),
		(*QosStoreEnableDisableReply)(nil),
	}
}
