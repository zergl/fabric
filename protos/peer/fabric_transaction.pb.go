// Code generated by protoc-gen-go.
// source: peer/fabric_transaction.proto
// DO NOT EDIT!

package peer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type InvalidTransaction_Cause int32

const (
	InvalidTransaction_TxIdAlreadyExists      InvalidTransaction_Cause = 0
	InvalidTransaction_RWConflictDuringCommit InvalidTransaction_Cause = 1
)

var InvalidTransaction_Cause_name = map[int32]string{
	0: "TxIdAlreadyExists",
	1: "RWConflictDuringCommit",
}
var InvalidTransaction_Cause_value = map[string]int32{
	"TxIdAlreadyExists":      0,
	"RWConflictDuringCommit": 1,
}

func (x InvalidTransaction_Cause) String() string {
	return proto.EnumName(InvalidTransaction_Cause_name, int32(x))
}
func (InvalidTransaction_Cause) EnumDescriptor() ([]byte, []int) { return fileDescriptor11, []int{1, 0} }

// This message is necessary to facilitate the verification of the signature
// (in the signature field) over the bytes of the transaction (in the
// transactionBytes field).
type SignedTransaction struct {
	// The bytes of the Transaction. NDD
	TransactionBytes []byte `protobuf:"bytes,1,opt,name=transactionBytes,proto3" json:"transactionBytes,omitempty"`
	// Signature of the transactionBytes The public key of the signature is in
	// the header field of TransactionAction There might be multiple
	// TransactionAction, so multiple headers, but there should be same
	// transactor identity (cert) in all headers
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *SignedTransaction) Reset()                    { *m = SignedTransaction{} }
func (m *SignedTransaction) String() string            { return proto.CompactTextString(m) }
func (*SignedTransaction) ProtoMessage()               {}
func (*SignedTransaction) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

// This is used to wrap an invalid Transaction with the cause
type InvalidTransaction struct {
	Transaction *Transaction             `protobuf:"bytes,1,opt,name=transaction" json:"transaction,omitempty"`
	Cause       InvalidTransaction_Cause `protobuf:"varint,2,opt,name=cause,enum=protos.InvalidTransaction_Cause" json:"cause,omitempty"`
}

func (m *InvalidTransaction) Reset()                    { *m = InvalidTransaction{} }
func (m *InvalidTransaction) String() string            { return proto.CompactTextString(m) }
func (*InvalidTransaction) ProtoMessage()               {}
func (*InvalidTransaction) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{1} }

func (m *InvalidTransaction) GetTransaction() *Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

// The transaction to be sent to the ordering service. A transaction contains
// one or more TransactionAction. Each TransactionAction binds a proposal to
// potentially multiple actions. The transaction is atomic meaning that either
// all actions in the transaction will be committed or none will.  Note that
// while a Transaction might include more than one Header, the Header.creator
// field must be the same in each.
// A single client is free to issue a number of independent Proposal, each with
// their header (Header) and request payload (ChaincodeProposalPayload).  Each
// proposal is independently endorsed generating an action
// (ProposalResponsePayload) with one signature per Endorser. Any number of
// independent proposals (and their action) might be included in a transaction
// to ensure that they are treated atomically.
type Transaction struct {
	// Version indicates message protocol version.
	Version int32 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	// Timestamp is the local time that the
	// message was created by the sender
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	// The payload is an array of TransactionAction. An array is necessary to
	// accommodate multiple actions per transaction
	Actions []*TransactionAction `protobuf:"bytes,3,rep,name=actions" json:"actions,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{2} }

func (m *Transaction) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Transaction) GetActions() []*TransactionAction {
	if m != nil {
		return m.Actions
	}
	return nil
}

// TransactionAction binds a proposal to its action.  The type field in the
// header dictates the type of action to be applied to the ledger.
type TransactionAction struct {
	// The header of the proposal action, which is the proposal header
	Header []byte `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// The payload of the action as defined by the type in the header For
	// chaincode, it's the bytes of ChaincodeActionPayload
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *TransactionAction) Reset()                    { *m = TransactionAction{} }
func (m *TransactionAction) String() string            { return proto.CompactTextString(m) }
func (*TransactionAction) ProtoMessage()               {}
func (*TransactionAction) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{3} }

func init() {
	proto.RegisterType((*SignedTransaction)(nil), "protos.SignedTransaction")
	proto.RegisterType((*InvalidTransaction)(nil), "protos.InvalidTransaction")
	proto.RegisterType((*Transaction)(nil), "protos.Transaction")
	proto.RegisterType((*TransactionAction)(nil), "protos.TransactionAction")
	proto.RegisterEnum("protos.InvalidTransaction_Cause", InvalidTransaction_Cause_name, InvalidTransaction_Cause_value)
}

func init() { proto.RegisterFile("peer/fabric_transaction.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x92, 0xdf, 0x4a, 0xc3, 0x30,
	0x18, 0xc5, 0xad, 0x63, 0x1b, 0xfb, 0x2a, 0xb2, 0x45, 0x1c, 0x75, 0x28, 0x8e, 0x5e, 0xf9, 0x07,
	0x5a, 0xd8, 0x50, 0xc4, 0xbb, 0x6d, 0xee, 0x62, 0xb7, 0x75, 0x20, 0x08, 0x22, 0x69, 0x9b, 0x75,
	0x81, 0xb6, 0x29, 0x49, 0x3a, 0xd6, 0x17, 0xf1, 0x75, 0x7c, 0x35, 0xbb, 0x66, 0xdd, 0xaa, 0xf3,
	0xaa, 0x9c, 0xf4, 0x97, 0x93, 0x93, 0x93, 0x0f, 0xae, 0x12, 0x42, 0xb8, 0xbd, 0xc0, 0x2e, 0xa7,
	0xde, 0xa7, 0xe4, 0x38, 0x16, 0xd8, 0x93, 0x94, 0xc5, 0x56, 0xc2, 0x99, 0x64, 0xa8, 0x51, 0x7c,
	0x44, 0xef, 0x3a, 0x60, 0x2c, 0x08, 0x89, 0x5d, 0x48, 0x37, 0x5d, 0xd8, 0x92, 0x46, 0x44, 0x48,
	0x1c, 0x25, 0x0a, 0x34, 0x3f, 0xa0, 0xf3, 0x4a, 0x83, 0x98, 0xf8, 0xf3, 0xbd, 0x07, 0xba, 0x83,
	0x76, 0xc5, 0x72, 0x9c, 0x49, 0x22, 0x0c, 0xad, 0xaf, 0xdd, 0x9c, 0x38, 0x07, 0xeb, 0xe8, 0x12,
	0x5a, 0x22, 0x37, 0xc0, 0x32, 0xe5, 0xc4, 0x38, 0x2e, 0xa0, 0xfd, 0x82, 0xf9, 0xad, 0x01, 0x9a,
	0xc5, 0x2b, 0x1c, 0xd2, 0x5f, 0x07, 0x3c, 0x80, 0x5e, 0x31, 0x2a, 0xbc, 0xf5, 0xc1, 0x99, 0x8a,
	0x24, 0xac, 0x0a, 0xe9, 0x54, 0x39, 0xf4, 0x08, 0x75, 0x0f, 0xa7, 0x42, 0x9d, 0x73, 0x3a, 0xe8,
	0x97, 0x1b, 0x0e, 0x4f, 0xb0, 0x26, 0x1b, 0xce, 0x51, 0xb8, 0xf9, 0x0c, 0xf5, 0x42, 0xa3, 0x73,
	0xe8, 0xcc, 0xd7, 0x33, 0x7f, 0x14, 0x72, 0x82, 0xfd, 0x6c, 0xba, 0xa6, 0x42, 0x8a, 0xf6, 0x11,
	0xea, 0x41, 0xd7, 0x79, 0x9b, 0xb0, 0x78, 0x11, 0x52, 0x4f, 0xbe, 0xa4, 0x9c, 0xc6, 0xc1, 0x84,
	0x45, 0x11, 0x95, 0x6d, 0xcd, 0xfc, 0xd2, 0x40, 0xaf, 0x46, 0x37, 0xa0, 0xb9, 0x22, 0x5c, 0x94,
	0xb1, 0xeb, 0x4e, 0x29, 0xd1, 0x13, 0xb4, 0x76, 0xed, 0x16, 0x09, 0xf5, 0x41, 0xcf, 0x52, 0xfd,
	0x5b, 0x65, 0xff, 0xd6, 0xbc, 0x24, 0x9c, 0x3d, 0x8c, 0x86, 0xd0, 0x54, 0xee, 0xc2, 0xa8, 0xf5,
	0x6b, 0xf9, 0xbe, 0x8b, 0x7f, 0xaa, 0x18, 0xa9, 0x42, 0x4a, 0xd2, 0x9c, 0xe6, 0x77, 0xf9, 0xfb,
	0x17, 0x75, 0xa1, 0xb1, 0xcc, 0xaf, 0x46, 0xf8, 0xf6, 0xbd, 0xb6, 0x6a, 0x93, 0x3a, 0xc1, 0x59,
	0xc8, 0xb0, 0xbf, 0x7d, 0xa3, 0x52, 0x8e, 0xef, 0xdf, 0x6f, 0x03, 0x2a, 0x97, 0xa9, 0x6b, 0x79,
	0x2c, 0xb2, 0x97, 0x59, 0x42, 0x78, 0x48, 0xfc, 0x60, 0x37, 0x5c, 0x6a, 0x74, 0x84, 0xbd, 0x99,
	0x37, 0x57, 0x8d, 0xd5, 0xf0, 0x27, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x48, 0x72, 0x86, 0x7e, 0x02,
	0x00, 0x00,
}
