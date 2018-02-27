// Code generated by govpp binapi-generator DO NOT EDIT.
// Package acl represents the VPP binary API of the 'acl' VPP module.
// Generated from '/usr/share/vpp/api/acl.api.json'
package acl

import "git.fd.io/govpp.git/api"

// ACLRule represents the VPP binary API data type 'acl_rule'.
// Generated from '/usr/share/vpp/api/acl.api.json', line 863:
//
//            "acl_rule",
//            [
//                "u8",
//                "is_permit"
//            ],
//            [
//                "u8",
//                "is_ipv6"
//            ],
//            [
//                "u8",
//                "src_ip_addr",
//                16
//            ],
//            [
//                "u8",
//                "src_ip_prefix_len"
//            ],
//            [
//                "u8",
//                "dst_ip_addr",
//                16
//            ],
//            [
//                "u8",
//                "dst_ip_prefix_len"
//            ],
//            [
//                "u8",
//                "proto"
//            ],
//            [
//                "u16",
//                "srcport_or_icmptype_first"
//            ],
//            [
//                "u16",
//                "srcport_or_icmptype_last"
//            ],
//            [
//                "u16",
//                "dstport_or_icmpcode_first"
//            ],
//            [
//                "u16",
//                "dstport_or_icmpcode_last"
//            ],
//            [
//                "u8",
//                "tcp_flags_mask"
//            ],
//            [
//                "u8",
//                "tcp_flags_value"
//            ],
//            {
//                "crc": "0x6f99bf4d"
//            }
//
type ACLRule struct {
	IsPermit               uint8
	IsIpv6                 uint8
	SrcIPAddr              []byte `struc:"[16]byte"`
	SrcIPPrefixLen         uint8
	DstIPAddr              []byte `struc:"[16]byte"`
	DstIPPrefixLen         uint8
	Proto                  uint8
	SrcportOrIcmptypeFirst uint16
	SrcportOrIcmptypeLast  uint16
	DstportOrIcmpcodeFirst uint16
	DstportOrIcmpcodeLast  uint16
	TCPFlagsMask           uint8
	TCPFlagsValue          uint8
}

func (*ACLRule) GetTypeName() string {
	return "acl_rule"
}
func (*ACLRule) GetCrcString() string {
	return "6f99bf4d"
}

// MacipACLRule represents the VPP binary API data type 'macip_acl_rule'.
// Generated from '/usr/share/vpp/api/acl.api.json', line 923:
//
//            "macip_acl_rule",
//            [
//                "u8",
//                "is_permit"
//            ],
//            [
//                "u8",
//                "is_ipv6"
//            ],
//            [
//                "u8",
//                "src_mac",
//                6
//            ],
//            [
//                "u8",
//                "src_mac_mask",
//                6
//            ],
//            [
//                "u8",
//                "src_ip_addr",
//                16
//            ],
//            [
//                "u8",
//                "src_ip_prefix_len"
//            ],
//            {
//                "crc": "0x70589f1e"
//            }
//
type MacipACLRule struct {
	IsPermit       uint8
	IsIpv6         uint8
	SrcMac         []byte `struc:"[6]byte"`
	SrcMacMask     []byte `struc:"[6]byte"`
	SrcIPAddr      []byte `struc:"[16]byte"`
	SrcIPPrefixLen uint8
}

func (*MacipACLRule) GetTypeName() string {
	return "macip_acl_rule"
}
func (*MacipACLRule) GetCrcString() string {
	return "70589f1e"
}

// ACLPluginGetVersion represents the VPP binary API message 'acl_plugin_get_version'.
// Generated from '/usr/share/vpp/api/acl.api.json', line 60:
//
//            "acl_plugin_get_version",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            {
//                "crc": "0x51077d14"
//            }
//
type ACLPluginGetVersion struct {
}

func (*ACLPluginGetVersion) GetMessageName() string {
	return "acl_plugin_get_version"
}
func (*ACLPluginGetVersion) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ACLPluginGetVersion) GetCrcString() string {
	return "51077d14"
}
func NewACLPluginGetVersion() api.Message {
	return &ACLPluginGetVersion{}
}

// ACLPluginGetVersionReply represents the VPP binary API message 'acl_plugin_get_version_reply'.
// Generated from '/usr/share/vpp/api/acl.api.json', line 78:
//
//            "acl_plugin_get_version_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u32",
//                "major"
//            ],
//            [
//                "u32",
//                "minor"
//            ],
//            {
//                "crc": "0x9b32cf86"
//            }
//
type ACLPluginGetVersionReply struct {
	Major uint32
	Minor uint32
}

func (*ACLPluginGetVersionReply) GetMessageName() string {
	return "acl_plugin_get_version_reply"
}
func (*ACLPluginGetVersionReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ACLPluginGetVersionReply) GetCrcString() string {
	return "9b32cf86"
}
func NewACLPluginGetVersionReply() api.Message {
	return &ACLPluginGetVersionReply{}
}

// ACLPluginControlPing represents the VPP binary API message 'acl_plugin_control_ping'.
// Generated from '/usr/share/vpp/api/acl.api.json', line 100:
//
//            "acl_plugin_control_ping",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            {
//                "crc": "0x51077d14"
//            }
//
type ACLPluginControlPing struct {
}

func (*ACLPluginControlPing) GetMessageName() string {
	return "acl_plugin_control_ping"
}
func (*ACLPluginControlPing) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ACLPluginControlPing) GetCrcString() string {
	return "51077d14"
}
func NewACLPluginControlPing() api.Message {
	return &ACLPluginControlPing{}
}

// ACLPluginControlPingReply represents the VPP binary API message 'acl_plugin_control_ping_reply'.
// Generated from '/usr/share/vpp/api/acl.api.json', line 118:
//
//            "acl_plugin_control_ping_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "vpe_pid"
//            ],
//            {
//                "crc": "0xf6b0b8ca"
//            }
//
type ACLPluginControlPingReply struct {
	Retval      int32
	ClientIndex uint32
	VpePid      uint32
}

func (*ACLPluginControlPingReply) GetMessageName() string {
	return "acl_plugin_control_ping_reply"
}
func (*ACLPluginControlPingReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ACLPluginControlPingReply) GetCrcString() string {
	return "f6b0b8ca"
}
func NewACLPluginControlPingReply() api.Message {
	return &ACLPluginControlPingReply{}
}

// ACLAddReplace represents the VPP binary API message 'acl_add_replace'.
// Generated from '/usr/share/vpp/api/acl.api.json', line 144:
//
//            "acl_add_replace",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u32",
//                "acl_index"
//            ],
//            [
//                "u8",
//                "tag",
//                64
//            ],
//            [
//                "u32",
//                "count"
//            ],
//            [
//                "vl_api_acl_rule_t",
//                "r",
//                0,
//                "count"
//            ],
//            {
//                "crc": "0xe839997e"
//            }
//
type ACLAddReplace struct {
	ACLIndex uint32
	Tag      []byte `struc:"[64]byte"`
	Count    uint32 `struc:"sizeof=R"`
	R        []ACLRule
}

func (*ACLAddReplace) GetMessageName() string {
	return "acl_add_replace"
}
func (*ACLAddReplace) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ACLAddReplace) GetCrcString() string {
	return "e839997e"
}
func NewACLAddReplace() api.Message {
	return &ACLAddReplace{}
}

// ACLAddReplaceReply represents the VPP binary API message 'acl_add_replace_reply'.
// Generated from '/usr/share/vpp/api/acl.api.json', line 181:
//
//            "acl_add_replace_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u32",
//                "acl_index"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            {
//                "crc": "0xac407b0c"
//            }
//
type ACLAddReplaceReply struct {
	ACLIndex uint32
	Retval   int32
}

func (*ACLAddReplaceReply) GetMessageName() string {
	return "acl_add_replace_reply"
}
func (*ACLAddReplaceReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ACLAddReplaceReply) GetCrcString() string {
	return "ac407b0c"
}
func NewACLAddReplaceReply() api.Message {
	return &ACLAddReplaceReply{}
}

// ACLDel represents the VPP binary API message 'acl_del'.
// Generated from '/usr/share/vpp/api/acl.api.json', line 203:
//
//            "acl_del",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u32",
//                "acl_index"
//            ],
//            {
//                "crc": "0xef34fea4"
//            }
//
type ACLDel struct {
	ACLIndex uint32
}

func (*ACLDel) GetMessageName() string {
	return "acl_del"
}
func (*ACLDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ACLDel) GetCrcString() string {
	return "ef34fea4"
}
func NewACLDel() api.Message {
	return &ACLDel{}
}

// ACLDelReply represents the VPP binary API message 'acl_del_reply'.
// Generated from '/usr/share/vpp/api/acl.api.json', line 225:
//
//            "acl_del_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            {
//                "crc": "0xe8d4e804"
//            }
//
type ACLDelReply struct {
	Retval int32
}

func (*ACLDelReply) GetMessageName() string {
	return "acl_del_reply"
}
func (*ACLDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ACLDelReply) GetCrcString() string {
	return "e8d4e804"
}
func NewACLDelReply() api.Message {
	return &ACLDelReply{}
}

// ACLInterfaceAddDel represents the VPP binary API message 'acl_interface_add_del'.
// Generated from '/usr/share/vpp/api/acl.api.json', line 243:
//
//            "acl_interface_add_del",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u8",
//                "is_add"
//            ],
//            [
//                "u8",
//                "is_input"
//            ],
//            [
//                "u32",
//                "sw_if_index"
//            ],
//            [
//                "u32",
//                "acl_index"
//            ],
//            {
//                "crc": "0x0b2aedd1"
//            }
//
type ACLInterfaceAddDel struct {
	IsAdd     uint8
	IsInput   uint8
	SwIfIndex uint32
	ACLIndex  uint32
}

func (*ACLInterfaceAddDel) GetMessageName() string {
	return "acl_interface_add_del"
}
func (*ACLInterfaceAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ACLInterfaceAddDel) GetCrcString() string {
	return "0b2aedd1"
}
func NewACLInterfaceAddDel() api.Message {
	return &ACLInterfaceAddDel{}
}

// ACLInterfaceAddDelReply represents the VPP binary API message 'acl_interface_add_del_reply'.
// Generated from '/usr/share/vpp/api/acl.api.json', line 277:
//
//            "acl_interface_add_del_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            {
//                "crc": "0xe8d4e804"
//            }
//
type ACLInterfaceAddDelReply struct {
	Retval int32
}

func (*ACLInterfaceAddDelReply) GetMessageName() string {
	return "acl_interface_add_del_reply"
}
func (*ACLInterfaceAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ACLInterfaceAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func NewACLInterfaceAddDelReply() api.Message {
	return &ACLInterfaceAddDelReply{}
}

// ACLInterfaceSetACLList represents the VPP binary API message 'acl_interface_set_acl_list'.
// Generated from '/usr/share/vpp/api/acl.api.json', line 295:
//
//            "acl_interface_set_acl_list",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u32",
//                "sw_if_index"
//            ],
//            [
//                "u8",
//                "count"
//            ],
//            [
//                "u8",
//                "n_input"
//            ],
//            [
//                "u32",
//                "acls",
//                0,
//                "count"
//            ],
//            {
//                "crc": "0x8baece38"
//            }
//
type ACLInterfaceSetACLList struct {
	SwIfIndex uint32
	Count     uint8 `struc:"sizeof=Acls"`
	NInput    uint8
	Acls      []uint32
}

func (*ACLInterfaceSetACLList) GetMessageName() string {
	return "acl_interface_set_acl_list"
}
func (*ACLInterfaceSetACLList) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ACLInterfaceSetACLList) GetCrcString() string {
	return "8baece38"
}
func NewACLInterfaceSetACLList() api.Message {
	return &ACLInterfaceSetACLList{}
}

// ACLInterfaceSetACLListReply represents the VPP binary API message 'acl_interface_set_acl_list_reply'.
// Generated from '/usr/share/vpp/api/acl.api.json', line 331:
//
//            "acl_interface_set_acl_list_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            {
//                "crc": "0xe8d4e804"
//            }
//
type ACLInterfaceSetACLListReply struct {
	Retval int32
}

func (*ACLInterfaceSetACLListReply) GetMessageName() string {
	return "acl_interface_set_acl_list_reply"
}
func (*ACLInterfaceSetACLListReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ACLInterfaceSetACLListReply) GetCrcString() string {
	return "e8d4e804"
}
func NewACLInterfaceSetACLListReply() api.Message {
	return &ACLInterfaceSetACLListReply{}
}

// ACLDump represents the VPP binary API message 'acl_dump'.
// Generated from '/usr/share/vpp/api/acl.api.json', line 349:
//
//            "acl_dump",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u32",
//                "acl_index"
//            ],
//            {
//                "crc": "0xef34fea4"
//            }
//
type ACLDump struct {
	ACLIndex uint32
}

func (*ACLDump) GetMessageName() string {
	return "acl_dump"
}
func (*ACLDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ACLDump) GetCrcString() string {
	return "ef34fea4"
}
func NewACLDump() api.Message {
	return &ACLDump{}
}

// ACLDetails represents the VPP binary API message 'acl_details'.
// Generated from '/usr/share/vpp/api/acl.api.json', line 371:
//
//            "acl_details",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u32",
//                "acl_index"
//            ],
//            [
//                "u8",
//                "tag",
//                64
//            ],
//            [
//                "u32",
//                "count"
//            ],
//            [
//                "vl_api_acl_rule_t",
//                "r",
//                0,
//                "count"
//            ],
//            {
//                "crc": "0x5bd895be"
//            }
//
type ACLDetails struct {
	ACLIndex uint32
	Tag      []byte `struc:"[64]byte"`
	Count    uint32 `struc:"sizeof=R"`
	R        []ACLRule
}

func (*ACLDetails) GetMessageName() string {
	return "acl_details"
}
func (*ACLDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ACLDetails) GetCrcString() string {
	return "5bd895be"
}
func NewACLDetails() api.Message {
	return &ACLDetails{}
}

// ACLInterfaceListDump represents the VPP binary API message 'acl_interface_list_dump'.
// Generated from '/usr/share/vpp/api/acl.api.json', line 404:
//
//            "acl_interface_list_dump",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u32",
//                "sw_if_index"
//            ],
//            {
//                "crc": "0x529cb13f"
//            }
//
type ACLInterfaceListDump struct {
	SwIfIndex uint32
}

func (*ACLInterfaceListDump) GetMessageName() string {
	return "acl_interface_list_dump"
}
func (*ACLInterfaceListDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ACLInterfaceListDump) GetCrcString() string {
	return "529cb13f"
}
func NewACLInterfaceListDump() api.Message {
	return &ACLInterfaceListDump{}
}

// ACLInterfaceListDetails represents the VPP binary API message 'acl_interface_list_details'.
// Generated from '/usr/share/vpp/api/acl.api.json', line 426:
//
//            "acl_interface_list_details",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u32",
//                "sw_if_index"
//            ],
//            [
//                "u8",
//                "count"
//            ],
//            [
//                "u8",
//                "n_input"
//            ],
//            [
//                "u32",
//                "acls",
//                0,
//                "count"
//            ],
//            {
//                "crc": "0xd5e80809"
//            }
//
type ACLInterfaceListDetails struct {
	SwIfIndex uint32
	Count     uint8 `struc:"sizeof=Acls"`
	NInput    uint8
	Acls      []uint32
}

func (*ACLInterfaceListDetails) GetMessageName() string {
	return "acl_interface_list_details"
}
func (*ACLInterfaceListDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ACLInterfaceListDetails) GetCrcString() string {
	return "d5e80809"
}
func NewACLInterfaceListDetails() api.Message {
	return &ACLInterfaceListDetails{}
}

// MacipACLAdd represents the VPP binary API message 'macip_acl_add'.
// Generated from '/usr/share/vpp/api/acl.api.json', line 458:
//
//            "macip_acl_add",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u8",
//                "tag",
//                64
//            ],
//            [
//                "u32",
//                "count"
//            ],
//            [
//                "vl_api_macip_acl_rule_t",
//                "r",
//                0,
//                "count"
//            ],
//            {
//                "crc": "0xb3d3d65a"
//            }
//
type MacipACLAdd struct {
	Tag   []byte `struc:"[64]byte"`
	Count uint32 `struc:"sizeof=R"`
	R     []MacipACLRule
}

func (*MacipACLAdd) GetMessageName() string {
	return "macip_acl_add"
}
func (*MacipACLAdd) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*MacipACLAdd) GetCrcString() string {
	return "b3d3d65a"
}
func NewMacipACLAdd() api.Message {
	return &MacipACLAdd{}
}

// MacipACLAddReply represents the VPP binary API message 'macip_acl_add_reply'.
// Generated from '/usr/share/vpp/api/acl.api.json', line 491:
//
//            "macip_acl_add_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u32",
//                "acl_index"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            {
//                "crc": "0xac407b0c"
//            }
//
type MacipACLAddReply struct {
	ACLIndex uint32
	Retval   int32
}

func (*MacipACLAddReply) GetMessageName() string {
	return "macip_acl_add_reply"
}
func (*MacipACLAddReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*MacipACLAddReply) GetCrcString() string {
	return "ac407b0c"
}
func NewMacipACLAddReply() api.Message {
	return &MacipACLAddReply{}
}

// MacipACLAddReplace represents the VPP binary API message 'macip_acl_add_replace'.
// Generated from '/usr/share/vpp/api/acl.api.json', line 513:
//
//            "macip_acl_add_replace",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u32",
//                "acl_index"
//            ],
//            [
//                "u8",
//                "tag",
//                64
//            ],
//            [
//                "u32",
//                "count"
//            ],
//            [
//                "vl_api_macip_acl_rule_t",
//                "r",
//                0,
//                "count"
//            ],
//            {
//                "crc": "0xa0e8c01b"
//            }
//
type MacipACLAddReplace struct {
	ACLIndex uint32
	Tag      []byte `struc:"[64]byte"`
	Count    uint32 `struc:"sizeof=R"`
	R        []MacipACLRule
}

func (*MacipACLAddReplace) GetMessageName() string {
	return "macip_acl_add_replace"
}
func (*MacipACLAddReplace) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*MacipACLAddReplace) GetCrcString() string {
	return "a0e8c01b"
}
func NewMacipACLAddReplace() api.Message {
	return &MacipACLAddReplace{}
}

// MacipACLAddReplaceReply represents the VPP binary API message 'macip_acl_add_replace_reply'.
// Generated from '/usr/share/vpp/api/acl.api.json', line 550:
//
//            "macip_acl_add_replace_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u32",
//                "acl_index"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            {
//                "crc": "0xac407b0c"
//            }
//
type MacipACLAddReplaceReply struct {
	ACLIndex uint32
	Retval   int32
}

func (*MacipACLAddReplaceReply) GetMessageName() string {
	return "macip_acl_add_replace_reply"
}
func (*MacipACLAddReplaceReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*MacipACLAddReplaceReply) GetCrcString() string {
	return "ac407b0c"
}
func NewMacipACLAddReplaceReply() api.Message {
	return &MacipACLAddReplaceReply{}
}

// MacipACLDel represents the VPP binary API message 'macip_acl_del'.
// Generated from '/usr/share/vpp/api/acl.api.json', line 572:
//
//            "macip_acl_del",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u32",
//                "acl_index"
//            ],
//            {
//                "crc": "0xef34fea4"
//            }
//
type MacipACLDel struct {
	ACLIndex uint32
}

func (*MacipACLDel) GetMessageName() string {
	return "macip_acl_del"
}
func (*MacipACLDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*MacipACLDel) GetCrcString() string {
	return "ef34fea4"
}
func NewMacipACLDel() api.Message {
	return &MacipACLDel{}
}

// MacipACLDelReply represents the VPP binary API message 'macip_acl_del_reply'.
// Generated from '/usr/share/vpp/api/acl.api.json', line 594:
//
//            "macip_acl_del_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            {
//                "crc": "0xe8d4e804"
//            }
//
type MacipACLDelReply struct {
	Retval int32
}

func (*MacipACLDelReply) GetMessageName() string {
	return "macip_acl_del_reply"
}
func (*MacipACLDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*MacipACLDelReply) GetCrcString() string {
	return "e8d4e804"
}
func NewMacipACLDelReply() api.Message {
	return &MacipACLDelReply{}
}

// MacipACLInterfaceAddDel represents the VPP binary API message 'macip_acl_interface_add_del'.
// Generated from '/usr/share/vpp/api/acl.api.json', line 612:
//
//            "macip_acl_interface_add_del",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u8",
//                "is_add"
//            ],
//            [
//                "u32",
//                "sw_if_index"
//            ],
//            [
//                "u32",
//                "acl_index"
//            ],
//            {
//                "crc": "0x6a6be97c"
//            }
//
type MacipACLInterfaceAddDel struct {
	IsAdd     uint8
	SwIfIndex uint32
	ACLIndex  uint32
}

func (*MacipACLInterfaceAddDel) GetMessageName() string {
	return "macip_acl_interface_add_del"
}
func (*MacipACLInterfaceAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*MacipACLInterfaceAddDel) GetCrcString() string {
	return "6a6be97c"
}
func NewMacipACLInterfaceAddDel() api.Message {
	return &MacipACLInterfaceAddDel{}
}

// MacipACLInterfaceAddDelReply represents the VPP binary API message 'macip_acl_interface_add_del_reply'.
// Generated from '/usr/share/vpp/api/acl.api.json', line 642:
//
//            "macip_acl_interface_add_del_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            {
//                "crc": "0xe8d4e804"
//            }
//
type MacipACLInterfaceAddDelReply struct {
	Retval int32
}

func (*MacipACLInterfaceAddDelReply) GetMessageName() string {
	return "macip_acl_interface_add_del_reply"
}
func (*MacipACLInterfaceAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*MacipACLInterfaceAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func NewMacipACLInterfaceAddDelReply() api.Message {
	return &MacipACLInterfaceAddDelReply{}
}

// MacipACLDump represents the VPP binary API message 'macip_acl_dump'.
// Generated from '/usr/share/vpp/api/acl.api.json', line 660:
//
//            "macip_acl_dump",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u32",
//                "acl_index"
//            ],
//            {
//                "crc": "0xef34fea4"
//            }
//
type MacipACLDump struct {
	ACLIndex uint32
}

func (*MacipACLDump) GetMessageName() string {
	return "macip_acl_dump"
}
func (*MacipACLDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*MacipACLDump) GetCrcString() string {
	return "ef34fea4"
}
func NewMacipACLDump() api.Message {
	return &MacipACLDump{}
}

// MacipACLDetails represents the VPP binary API message 'macip_acl_details'.
// Generated from '/usr/share/vpp/api/acl.api.json', line 682:
//
//            "macip_acl_details",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u32",
//                "acl_index"
//            ],
//            [
//                "u8",
//                "tag",
//                64
//            ],
//            [
//                "u32",
//                "count"
//            ],
//            [
//                "vl_api_macip_acl_rule_t",
//                "r",
//                0,
//                "count"
//            ],
//            {
//                "crc": "0xdd2b55ba"
//            }
//
type MacipACLDetails struct {
	ACLIndex uint32
	Tag      []byte `struc:"[64]byte"`
	Count    uint32 `struc:"sizeof=R"`
	R        []MacipACLRule
}

func (*MacipACLDetails) GetMessageName() string {
	return "macip_acl_details"
}
func (*MacipACLDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*MacipACLDetails) GetCrcString() string {
	return "dd2b55ba"
}
func NewMacipACLDetails() api.Message {
	return &MacipACLDetails{}
}

// MacipACLInterfaceGet represents the VPP binary API message 'macip_acl_interface_get'.
// Generated from '/usr/share/vpp/api/acl.api.json', line 715:
//
//            "macip_acl_interface_get",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            {
//                "crc": "0x51077d14"
//            }
//
type MacipACLInterfaceGet struct {
}

func (*MacipACLInterfaceGet) GetMessageName() string {
	return "macip_acl_interface_get"
}
func (*MacipACLInterfaceGet) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*MacipACLInterfaceGet) GetCrcString() string {
	return "51077d14"
}
func NewMacipACLInterfaceGet() api.Message {
	return &MacipACLInterfaceGet{}
}

// MacipACLInterfaceGetReply represents the VPP binary API message 'macip_acl_interface_get_reply'.
// Generated from '/usr/share/vpp/api/acl.api.json', line 733:
//
//            "macip_acl_interface_get_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u32",
//                "count"
//            ],
//            [
//                "u32",
//                "acls",
//                0,
//                "count"
//            ],
//            {
//                "crc": "0xaccf9b05"
//            }
//
type MacipACLInterfaceGetReply struct {
	Count uint32 `struc:"sizeof=Acls"`
	Acls  []uint32
}

func (*MacipACLInterfaceGetReply) GetMessageName() string {
	return "macip_acl_interface_get_reply"
}
func (*MacipACLInterfaceGetReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*MacipACLInterfaceGetReply) GetCrcString() string {
	return "accf9b05"
}
func NewMacipACLInterfaceGetReply() api.Message {
	return &MacipACLInterfaceGetReply{}
}

// MacipACLInterfaceListDump represents the VPP binary API message 'macip_acl_interface_list_dump'.
// Generated from '/usr/share/vpp/api/acl.api.json', line 757:
//
//            "macip_acl_interface_list_dump",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u32",
//                "sw_if_index"
//            ],
//            {
//                "crc": "0x529cb13f"
//            }
//
type MacipACLInterfaceListDump struct {
	SwIfIndex uint32
}

func (*MacipACLInterfaceListDump) GetMessageName() string {
	return "macip_acl_interface_list_dump"
}
func (*MacipACLInterfaceListDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*MacipACLInterfaceListDump) GetCrcString() string {
	return "529cb13f"
}
func NewMacipACLInterfaceListDump() api.Message {
	return &MacipACLInterfaceListDump{}
}

// MacipACLInterfaceListDetails represents the VPP binary API message 'macip_acl_interface_list_details'.
// Generated from '/usr/share/vpp/api/acl.api.json', line 779:
//
//            "macip_acl_interface_list_details",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u32",
//                "sw_if_index"
//            ],
//            [
//                "u8",
//                "count"
//            ],
//            [
//                "u32",
//                "acls",
//                0,
//                "count"
//            ],
//            {
//                "crc": "0x29783fa0"
//            }
//
type MacipACLInterfaceListDetails struct {
	SwIfIndex uint32
	Count     uint8 `struc:"sizeof=Acls"`
	Acls      []uint32
}

func (*MacipACLInterfaceListDetails) GetMessageName() string {
	return "macip_acl_interface_list_details"
}
func (*MacipACLInterfaceListDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*MacipACLInterfaceListDetails) GetCrcString() string {
	return "29783fa0"
}
func NewMacipACLInterfaceListDetails() api.Message {
	return &MacipACLInterfaceListDetails{}
}

// ACLInterfaceSetEtypeWhitelist represents the VPP binary API message 'acl_interface_set_etype_whitelist'.
// Generated from '/usr/share/vpp/api/acl.api.json', line 807:
//
//            "acl_interface_set_etype_whitelist",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u32",
//                "sw_if_index"
//            ],
//            [
//                "u8",
//                "count"
//            ],
//            [
//                "u8",
//                "n_input"
//            ],
//            [
//                "u16",
//                "whitelist",
//                0,
//                "count"
//            ],
//            {
//                "crc": "0xf515efc5"
//            }
//
type ACLInterfaceSetEtypeWhitelist struct {
	SwIfIndex uint32
	Count     uint8 `struc:"sizeof=Whitelist"`
	NInput    uint8
	Whitelist []uint16
}

func (*ACLInterfaceSetEtypeWhitelist) GetMessageName() string {
	return "acl_interface_set_etype_whitelist"
}
func (*ACLInterfaceSetEtypeWhitelist) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*ACLInterfaceSetEtypeWhitelist) GetCrcString() string {
	return "f515efc5"
}
func NewACLInterfaceSetEtypeWhitelist() api.Message {
	return &ACLInterfaceSetEtypeWhitelist{}
}

// ACLInterfaceSetEtypeWhitelistReply represents the VPP binary API message 'acl_interface_set_etype_whitelist_reply'.
// Generated from '/usr/share/vpp/api/acl.api.json', line 843:
//
//            "acl_interface_set_etype_whitelist_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            {
//                "crc": "0xe8d4e804"
//            }
//
type ACLInterfaceSetEtypeWhitelistReply struct {
	Retval int32
}

func (*ACLInterfaceSetEtypeWhitelistReply) GetMessageName() string {
	return "acl_interface_set_etype_whitelist_reply"
}
func (*ACLInterfaceSetEtypeWhitelistReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*ACLInterfaceSetEtypeWhitelistReply) GetCrcString() string {
	return "e8d4e804"
}
func NewACLInterfaceSetEtypeWhitelistReply() api.Message {
	return &ACLInterfaceSetEtypeWhitelistReply{}
}
