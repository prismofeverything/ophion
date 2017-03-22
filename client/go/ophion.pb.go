// Code generated by protoc-gen-go.
// source: ophion.proto
// DO NOT EDIT!

/*
Package ophion is a generated protocol buffer package.

It is generated from these files:
	ophion.proto

It has these top-level messages:
	GraphQuery
	GraphStatement
	HasStatement
	SelectStatement
	Vertex
	Edge
	QueryResult
*/
package ophion

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/struct"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GraphQuery struct {
	Query []*GraphStatement `protobuf:"bytes,1,rep,name=query" json:"query,omitempty"`
}

func (m *GraphQuery) Reset()                    { *m = GraphQuery{} }
func (m *GraphQuery) String() string            { return proto.CompactTextString(m) }
func (*GraphQuery) ProtoMessage()               {}
func (*GraphQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GraphQuery) GetQuery() []*GraphStatement {
	if m != nil {
		return m.Query
	}
	return nil
}

type GraphStatement struct {
	// Types that are valid to be assigned to Statement:
	//	*GraphStatement_V
	//	*GraphStatement_E
	//	*GraphStatement_Label
	//	*GraphStatement_Has
	//	*GraphStatement_As
	//	*GraphStatement_In
	//	*GraphStatement_Out
	//	*GraphStatement_InEdge
	//	*GraphStatement_OutEdge
	//	*GraphStatement_InVertex
	//	*GraphStatement_OutVertex
	//	*GraphStatement_Select
	//	*GraphStatement_Limit
	//	*GraphStatement_Count
	//	*GraphStatement_AddV
	//	*GraphStatement_Property
	//	*GraphStatement_AddE
	//	*GraphStatement_To
	//	*GraphStatement_Drop
	Statement isGraphStatement_Statement `protobuf_oneof:"statement"`
}

func (m *GraphStatement) Reset()                    { *m = GraphStatement{} }
func (m *GraphStatement) String() string            { return proto.CompactTextString(m) }
func (*GraphStatement) ProtoMessage()               {}
func (*GraphStatement) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isGraphStatement_Statement interface {
	isGraphStatement_Statement()
}

type GraphStatement_V struct {
	V string `protobuf:"bytes,1,opt,name=V,oneof"`
}
type GraphStatement_E struct {
	E string `protobuf:"bytes,2,opt,name=E,oneof"`
}
type GraphStatement_Label struct {
	Label string `protobuf:"bytes,3,opt,name=label,oneof"`
}
type GraphStatement_Has struct {
	Has *HasStatement `protobuf:"bytes,4,opt,name=has,oneof"`
}
type GraphStatement_As struct {
	As string `protobuf:"bytes,5,opt,name=as,oneof"`
}
type GraphStatement_In struct {
	In string `protobuf:"bytes,6,opt,name=in,oneof"`
}
type GraphStatement_Out struct {
	Out string `protobuf:"bytes,7,opt,name=out,oneof"`
}
type GraphStatement_InEdge struct {
	InEdge string `protobuf:"bytes,8,opt,name=inEdge,oneof"`
}
type GraphStatement_OutEdge struct {
	OutEdge string `protobuf:"bytes,9,opt,name=outEdge,oneof"`
}
type GraphStatement_InVertex struct {
	InVertex string `protobuf:"bytes,10,opt,name=inVertex,oneof"`
}
type GraphStatement_OutVertex struct {
	OutVertex string `protobuf:"bytes,11,opt,name=outVertex,oneof"`
}
type GraphStatement_Select struct {
	Select *SelectStatement `protobuf:"bytes,12,opt,name=select,oneof"`
}
type GraphStatement_Limit struct {
	Limit int64 `protobuf:"varint,13,opt,name=limit,oneof"`
}
type GraphStatement_Count struct {
	Count string `protobuf:"bytes,14,opt,name=count,oneof"`
}
type GraphStatement_AddV struct {
	AddV string `protobuf:"bytes,15,opt,name=addV,oneof"`
}
type GraphStatement_Property struct {
	Property *google_protobuf1.Struct `protobuf:"bytes,16,opt,name=property,oneof"`
}
type GraphStatement_AddE struct {
	AddE string `protobuf:"bytes,17,opt,name=addE,oneof"`
}
type GraphStatement_To struct {
	To string `protobuf:"bytes,18,opt,name=to,oneof"`
}
type GraphStatement_Drop struct {
	Drop string `protobuf:"bytes,19,opt,name=drop,oneof"`
}

func (*GraphStatement_V) isGraphStatement_Statement()         {}
func (*GraphStatement_E) isGraphStatement_Statement()         {}
func (*GraphStatement_Label) isGraphStatement_Statement()     {}
func (*GraphStatement_Has) isGraphStatement_Statement()       {}
func (*GraphStatement_As) isGraphStatement_Statement()        {}
func (*GraphStatement_In) isGraphStatement_Statement()        {}
func (*GraphStatement_Out) isGraphStatement_Statement()       {}
func (*GraphStatement_InEdge) isGraphStatement_Statement()    {}
func (*GraphStatement_OutEdge) isGraphStatement_Statement()   {}
func (*GraphStatement_InVertex) isGraphStatement_Statement()  {}
func (*GraphStatement_OutVertex) isGraphStatement_Statement() {}
func (*GraphStatement_Select) isGraphStatement_Statement()    {}
func (*GraphStatement_Limit) isGraphStatement_Statement()     {}
func (*GraphStatement_Count) isGraphStatement_Statement()     {}
func (*GraphStatement_AddV) isGraphStatement_Statement()      {}
func (*GraphStatement_Property) isGraphStatement_Statement()  {}
func (*GraphStatement_AddE) isGraphStatement_Statement()      {}
func (*GraphStatement_To) isGraphStatement_Statement()        {}
func (*GraphStatement_Drop) isGraphStatement_Statement()      {}

func (m *GraphStatement) GetStatement() isGraphStatement_Statement {
	if m != nil {
		return m.Statement
	}
	return nil
}

func (m *GraphStatement) GetV() string {
	if x, ok := m.GetStatement().(*GraphStatement_V); ok {
		return x.V
	}
	return ""
}

func (m *GraphStatement) GetE() string {
	if x, ok := m.GetStatement().(*GraphStatement_E); ok {
		return x.E
	}
	return ""
}

func (m *GraphStatement) GetLabel() string {
	if x, ok := m.GetStatement().(*GraphStatement_Label); ok {
		return x.Label
	}
	return ""
}

func (m *GraphStatement) GetHas() *HasStatement {
	if x, ok := m.GetStatement().(*GraphStatement_Has); ok {
		return x.Has
	}
	return nil
}

func (m *GraphStatement) GetAs() string {
	if x, ok := m.GetStatement().(*GraphStatement_As); ok {
		return x.As
	}
	return ""
}

func (m *GraphStatement) GetIn() string {
	if x, ok := m.GetStatement().(*GraphStatement_In); ok {
		return x.In
	}
	return ""
}

func (m *GraphStatement) GetOut() string {
	if x, ok := m.GetStatement().(*GraphStatement_Out); ok {
		return x.Out
	}
	return ""
}

func (m *GraphStatement) GetInEdge() string {
	if x, ok := m.GetStatement().(*GraphStatement_InEdge); ok {
		return x.InEdge
	}
	return ""
}

func (m *GraphStatement) GetOutEdge() string {
	if x, ok := m.GetStatement().(*GraphStatement_OutEdge); ok {
		return x.OutEdge
	}
	return ""
}

func (m *GraphStatement) GetInVertex() string {
	if x, ok := m.GetStatement().(*GraphStatement_InVertex); ok {
		return x.InVertex
	}
	return ""
}

func (m *GraphStatement) GetOutVertex() string {
	if x, ok := m.GetStatement().(*GraphStatement_OutVertex); ok {
		return x.OutVertex
	}
	return ""
}

func (m *GraphStatement) GetSelect() *SelectStatement {
	if x, ok := m.GetStatement().(*GraphStatement_Select); ok {
		return x.Select
	}
	return nil
}

func (m *GraphStatement) GetLimit() int64 {
	if x, ok := m.GetStatement().(*GraphStatement_Limit); ok {
		return x.Limit
	}
	return 0
}

func (m *GraphStatement) GetCount() string {
	if x, ok := m.GetStatement().(*GraphStatement_Count); ok {
		return x.Count
	}
	return ""
}

func (m *GraphStatement) GetAddV() string {
	if x, ok := m.GetStatement().(*GraphStatement_AddV); ok {
		return x.AddV
	}
	return ""
}

func (m *GraphStatement) GetProperty() *google_protobuf1.Struct {
	if x, ok := m.GetStatement().(*GraphStatement_Property); ok {
		return x.Property
	}
	return nil
}

func (m *GraphStatement) GetAddE() string {
	if x, ok := m.GetStatement().(*GraphStatement_AddE); ok {
		return x.AddE
	}
	return ""
}

func (m *GraphStatement) GetTo() string {
	if x, ok := m.GetStatement().(*GraphStatement_To); ok {
		return x.To
	}
	return ""
}

func (m *GraphStatement) GetDrop() string {
	if x, ok := m.GetStatement().(*GraphStatement_Drop); ok {
		return x.Drop
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*GraphStatement) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _GraphStatement_OneofMarshaler, _GraphStatement_OneofUnmarshaler, _GraphStatement_OneofSizer, []interface{}{
		(*GraphStatement_V)(nil),
		(*GraphStatement_E)(nil),
		(*GraphStatement_Label)(nil),
		(*GraphStatement_Has)(nil),
		(*GraphStatement_As)(nil),
		(*GraphStatement_In)(nil),
		(*GraphStatement_Out)(nil),
		(*GraphStatement_InEdge)(nil),
		(*GraphStatement_OutEdge)(nil),
		(*GraphStatement_InVertex)(nil),
		(*GraphStatement_OutVertex)(nil),
		(*GraphStatement_Select)(nil),
		(*GraphStatement_Limit)(nil),
		(*GraphStatement_Count)(nil),
		(*GraphStatement_AddV)(nil),
		(*GraphStatement_Property)(nil),
		(*GraphStatement_AddE)(nil),
		(*GraphStatement_To)(nil),
		(*GraphStatement_Drop)(nil),
	}
}

func _GraphStatement_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*GraphStatement)
	// statement
	switch x := m.Statement.(type) {
	case *GraphStatement_V:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.V)
	case *GraphStatement_E:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.E)
	case *GraphStatement_Label:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Label)
	case *GraphStatement_Has:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Has); err != nil {
			return err
		}
	case *GraphStatement_As:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.As)
	case *GraphStatement_In:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.In)
	case *GraphStatement_Out:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Out)
	case *GraphStatement_InEdge:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.InEdge)
	case *GraphStatement_OutEdge:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.OutEdge)
	case *GraphStatement_InVertex:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.InVertex)
	case *GraphStatement_OutVertex:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.OutVertex)
	case *GraphStatement_Select:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Select); err != nil {
			return err
		}
	case *GraphStatement_Limit:
		b.EncodeVarint(13<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Limit))
	case *GraphStatement_Count:
		b.EncodeVarint(14<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Count)
	case *GraphStatement_AddV:
		b.EncodeVarint(15<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.AddV)
	case *GraphStatement_Property:
		b.EncodeVarint(16<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Property); err != nil {
			return err
		}
	case *GraphStatement_AddE:
		b.EncodeVarint(17<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.AddE)
	case *GraphStatement_To:
		b.EncodeVarint(18<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.To)
	case *GraphStatement_Drop:
		b.EncodeVarint(19<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Drop)
	case nil:
	default:
		return fmt.Errorf("GraphStatement.Statement has unexpected type %T", x)
	}
	return nil
}

func _GraphStatement_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*GraphStatement)
	switch tag {
	case 1: // statement.V
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Statement = &GraphStatement_V{x}
		return true, err
	case 2: // statement.E
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Statement = &GraphStatement_E{x}
		return true, err
	case 3: // statement.label
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Statement = &GraphStatement_Label{x}
		return true, err
	case 4: // statement.has
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HasStatement)
		err := b.DecodeMessage(msg)
		m.Statement = &GraphStatement_Has{msg}
		return true, err
	case 5: // statement.as
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Statement = &GraphStatement_As{x}
		return true, err
	case 6: // statement.in
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Statement = &GraphStatement_In{x}
		return true, err
	case 7: // statement.out
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Statement = &GraphStatement_Out{x}
		return true, err
	case 8: // statement.inEdge
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Statement = &GraphStatement_InEdge{x}
		return true, err
	case 9: // statement.outEdge
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Statement = &GraphStatement_OutEdge{x}
		return true, err
	case 10: // statement.inVertex
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Statement = &GraphStatement_InVertex{x}
		return true, err
	case 11: // statement.outVertex
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Statement = &GraphStatement_OutVertex{x}
		return true, err
	case 12: // statement.select
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SelectStatement)
		err := b.DecodeMessage(msg)
		m.Statement = &GraphStatement_Select{msg}
		return true, err
	case 13: // statement.limit
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Statement = &GraphStatement_Limit{int64(x)}
		return true, err
	case 14: // statement.count
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Statement = &GraphStatement_Count{x}
		return true, err
	case 15: // statement.addV
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Statement = &GraphStatement_AddV{x}
		return true, err
	case 16: // statement.property
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf1.Struct)
		err := b.DecodeMessage(msg)
		m.Statement = &GraphStatement_Property{msg}
		return true, err
	case 17: // statement.addE
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Statement = &GraphStatement_AddE{x}
		return true, err
	case 18: // statement.to
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Statement = &GraphStatement_To{x}
		return true, err
	case 19: // statement.drop
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Statement = &GraphStatement_Drop{x}
		return true, err
	default:
		return false, nil
	}
}

func _GraphStatement_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*GraphStatement)
	// statement
	switch x := m.Statement.(type) {
	case *GraphStatement_V:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.V)))
		n += len(x.V)
	case *GraphStatement_E:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.E)))
		n += len(x.E)
	case *GraphStatement_Label:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Label)))
		n += len(x.Label)
	case *GraphStatement_Has:
		s := proto.Size(x.Has)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *GraphStatement_As:
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.As)))
		n += len(x.As)
	case *GraphStatement_In:
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.In)))
		n += len(x.In)
	case *GraphStatement_Out:
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Out)))
		n += len(x.Out)
	case *GraphStatement_InEdge:
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.InEdge)))
		n += len(x.InEdge)
	case *GraphStatement_OutEdge:
		n += proto.SizeVarint(9<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.OutEdge)))
		n += len(x.OutEdge)
	case *GraphStatement_InVertex:
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.InVertex)))
		n += len(x.InVertex)
	case *GraphStatement_OutVertex:
		n += proto.SizeVarint(11<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.OutVertex)))
		n += len(x.OutVertex)
	case *GraphStatement_Select:
		s := proto.Size(x.Select)
		n += proto.SizeVarint(12<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *GraphStatement_Limit:
		n += proto.SizeVarint(13<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Limit))
	case *GraphStatement_Count:
		n += proto.SizeVarint(14<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Count)))
		n += len(x.Count)
	case *GraphStatement_AddV:
		n += proto.SizeVarint(15<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.AddV)))
		n += len(x.AddV)
	case *GraphStatement_Property:
		s := proto.Size(x.Property)
		n += proto.SizeVarint(16<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *GraphStatement_AddE:
		n += proto.SizeVarint(17<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.AddE)))
		n += len(x.AddE)
	case *GraphStatement_To:
		n += proto.SizeVarint(18<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.To)))
		n += len(x.To)
	case *GraphStatement_Drop:
		n += proto.SizeVarint(19<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Drop)))
		n += len(x.Drop)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type HasStatement struct {
	Key    string   `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Within []string `protobuf:"bytes,2,rep,name=within" json:"within,omitempty"`
}

func (m *HasStatement) Reset()                    { *m = HasStatement{} }
func (m *HasStatement) String() string            { return proto.CompactTextString(m) }
func (*HasStatement) ProtoMessage()               {}
func (*HasStatement) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *HasStatement) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *HasStatement) GetWithin() []string {
	if m != nil {
		return m.Within
	}
	return nil
}

type SelectStatement struct {
	Steps []string `protobuf:"bytes,1,rep,name=steps" json:"steps,omitempty"`
}

func (m *SelectStatement) Reset()                    { *m = SelectStatement{} }
func (m *SelectStatement) String() string            { return proto.CompactTextString(m) }
func (*SelectStatement) ProtoMessage()               {}
func (*SelectStatement) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SelectStatement) GetSteps() []string {
	if m != nil {
		return m.Steps
	}
	return nil
}

type Vertex struct {
	Gid   string `protobuf:"bytes,1,opt,name=gid" json:"gid,omitempty"`
	Label string `protobuf:"bytes,2,opt,name=label" json:"label,omitempty"`
	// map<string, string> properties = 3;
	Properties *google_protobuf1.Struct `protobuf:"bytes,3,opt,name=properties" json:"properties,omitempty"`
}

func (m *Vertex) Reset()                    { *m = Vertex{} }
func (m *Vertex) String() string            { return proto.CompactTextString(m) }
func (*Vertex) ProtoMessage()               {}
func (*Vertex) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Vertex) GetGid() string {
	if m != nil {
		return m.Gid
	}
	return ""
}

func (m *Vertex) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Vertex) GetProperties() *google_protobuf1.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

type Edge struct {
	Gid   string `protobuf:"bytes,1,opt,name=gid" json:"gid,omitempty"`
	Label string `protobuf:"bytes,2,opt,name=label" json:"label,omitempty"`
	In    string `protobuf:"bytes,3,opt,name=in" json:"in,omitempty"`
	Out   string `protobuf:"bytes,4,opt,name=out" json:"out,omitempty"`
	// map<string, string> properties = 5;
	Properties *google_protobuf1.Struct `protobuf:"bytes,5,opt,name=properties" json:"properties,omitempty"`
}

func (m *Edge) Reset()                    { *m = Edge{} }
func (m *Edge) String() string            { return proto.CompactTextString(m) }
func (*Edge) ProtoMessage()               {}
func (*Edge) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Edge) GetGid() string {
	if m != nil {
		return m.Gid
	}
	return ""
}

func (m *Edge) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Edge) GetIn() string {
	if m != nil {
		return m.In
	}
	return ""
}

func (m *Edge) GetOut() string {
	if m != nil {
		return m.Out
	}
	return ""
}

func (m *Edge) GetProperties() *google_protobuf1.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

type QueryResult struct {
	// Types that are valid to be assigned to Result:
	//	*QueryResult_Struct
	//	*QueryResult_Vertex
	//	*QueryResult_Edge
	//	*QueryResult_IntValue
	//	*QueryResult_FloatValue
	//	*QueryResult_StrValue
	Result isQueryResult_Result `protobuf_oneof:"result"`
}

func (m *QueryResult) Reset()                    { *m = QueryResult{} }
func (m *QueryResult) String() string            { return proto.CompactTextString(m) }
func (*QueryResult) ProtoMessage()               {}
func (*QueryResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type isQueryResult_Result interface {
	isQueryResult_Result()
}

type QueryResult_Struct struct {
	Struct *google_protobuf1.Struct `protobuf:"bytes,1,opt,name=struct,oneof"`
}
type QueryResult_Vertex struct {
	Vertex *Vertex `protobuf:"bytes,2,opt,name=vertex,oneof"`
}
type QueryResult_Edge struct {
	Edge *Edge `protobuf:"bytes,3,opt,name=edge,oneof"`
}
type QueryResult_IntValue struct {
	IntValue int32 `protobuf:"varint,4,opt,name=int_value,json=intValue,oneof"`
}
type QueryResult_FloatValue struct {
	FloatValue float64 `protobuf:"fixed64,5,opt,name=float_value,json=floatValue,oneof"`
}
type QueryResult_StrValue struct {
	StrValue string `protobuf:"bytes,6,opt,name=str_value,json=strValue,oneof"`
}

func (*QueryResult_Struct) isQueryResult_Result()     {}
func (*QueryResult_Vertex) isQueryResult_Result()     {}
func (*QueryResult_Edge) isQueryResult_Result()       {}
func (*QueryResult_IntValue) isQueryResult_Result()   {}
func (*QueryResult_FloatValue) isQueryResult_Result() {}
func (*QueryResult_StrValue) isQueryResult_Result()   {}

func (m *QueryResult) GetResult() isQueryResult_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *QueryResult) GetStruct() *google_protobuf1.Struct {
	if x, ok := m.GetResult().(*QueryResult_Struct); ok {
		return x.Struct
	}
	return nil
}

func (m *QueryResult) GetVertex() *Vertex {
	if x, ok := m.GetResult().(*QueryResult_Vertex); ok {
		return x.Vertex
	}
	return nil
}

func (m *QueryResult) GetEdge() *Edge {
	if x, ok := m.GetResult().(*QueryResult_Edge); ok {
		return x.Edge
	}
	return nil
}

func (m *QueryResult) GetIntValue() int32 {
	if x, ok := m.GetResult().(*QueryResult_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (m *QueryResult) GetFloatValue() float64 {
	if x, ok := m.GetResult().(*QueryResult_FloatValue); ok {
		return x.FloatValue
	}
	return 0
}

func (m *QueryResult) GetStrValue() string {
	if x, ok := m.GetResult().(*QueryResult_StrValue); ok {
		return x.StrValue
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*QueryResult) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _QueryResult_OneofMarshaler, _QueryResult_OneofUnmarshaler, _QueryResult_OneofSizer, []interface{}{
		(*QueryResult_Struct)(nil),
		(*QueryResult_Vertex)(nil),
		(*QueryResult_Edge)(nil),
		(*QueryResult_IntValue)(nil),
		(*QueryResult_FloatValue)(nil),
		(*QueryResult_StrValue)(nil),
	}
}

func _QueryResult_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*QueryResult)
	// result
	switch x := m.Result.(type) {
	case *QueryResult_Struct:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Struct); err != nil {
			return err
		}
	case *QueryResult_Vertex:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Vertex); err != nil {
			return err
		}
	case *QueryResult_Edge:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Edge); err != nil {
			return err
		}
	case *QueryResult_IntValue:
		b.EncodeVarint(4<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.IntValue))
	case *QueryResult_FloatValue:
		b.EncodeVarint(5<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.FloatValue))
	case *QueryResult_StrValue:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.StrValue)
	case nil:
	default:
		return fmt.Errorf("QueryResult.Result has unexpected type %T", x)
	}
	return nil
}

func _QueryResult_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*QueryResult)
	switch tag {
	case 1: // result.struct
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf1.Struct)
		err := b.DecodeMessage(msg)
		m.Result = &QueryResult_Struct{msg}
		return true, err
	case 2: // result.vertex
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Vertex)
		err := b.DecodeMessage(msg)
		m.Result = &QueryResult_Vertex{msg}
		return true, err
	case 3: // result.edge
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Edge)
		err := b.DecodeMessage(msg)
		m.Result = &QueryResult_Edge{msg}
		return true, err
	case 4: // result.int_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Result = &QueryResult_IntValue{int32(x)}
		return true, err
	case 5: // result.float_value
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Result = &QueryResult_FloatValue{math.Float64frombits(x)}
		return true, err
	case 6: // result.str_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Result = &QueryResult_StrValue{x}
		return true, err
	default:
		return false, nil
	}
}

func _QueryResult_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*QueryResult)
	// result
	switch x := m.Result.(type) {
	case *QueryResult_Struct:
		s := proto.Size(x.Struct)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *QueryResult_Vertex:
		s := proto.Size(x.Vertex)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *QueryResult_Edge:
		s := proto.Size(x.Edge)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *QueryResult_IntValue:
		n += proto.SizeVarint(4<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.IntValue))
	case *QueryResult_FloatValue:
		n += proto.SizeVarint(5<<3 | proto.WireFixed64)
		n += 8
	case *QueryResult_StrValue:
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.StrValue)))
		n += len(x.StrValue)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*GraphQuery)(nil), "ophion.GraphQuery")
	proto.RegisterType((*GraphStatement)(nil), "ophion.GraphStatement")
	proto.RegisterType((*HasStatement)(nil), "ophion.HasStatement")
	proto.RegisterType((*SelectStatement)(nil), "ophion.SelectStatement")
	proto.RegisterType((*Vertex)(nil), "ophion.Vertex")
	proto.RegisterType((*Edge)(nil), "ophion.Edge")
	proto.RegisterType((*QueryResult)(nil), "ophion.QueryResult")
}

func init() { proto.RegisterFile("ophion.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 674 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcf, 0x6f, 0xd3, 0x3a,
	0x1c, 0x5f, 0x9a, 0x26, 0x6b, 0xbe, 0xed, 0xeb, 0xfa, 0xbc, 0x6a, 0xcf, 0xaa, 0xb6, 0xa7, 0xbe,
	0x5c, 0x5e, 0x84, 0xa0, 0x65, 0x43, 0x08, 0xb4, 0xe3, 0xa4, 0x8a, 0x5e, 0xc9, 0x50, 0x2f, 0x1c,
	0x90, 0xb7, 0x7a, 0xad, 0x45, 0x16, 0x87, 0xd8, 0x29, 0xec, 0xca, 0x8d, 0x33, 0x27, 0xfe, 0x2e,
	0xfe, 0x05, 0xfe, 0x09, 0x6e, 0xc8, 0x5f, 0xdb, 0xa3, 0x43, 0x02, 0xed, 0x96, 0xcf, 0x0f, 0xdb,
	0xf9, 0x7e, 0xf2, 0x71, 0xa0, 0x27, 0xab, 0xb5, 0x90, 0xe5, 0xa4, 0xaa, 0xa5, 0x96, 0x24, 0xb6,
	0x68, 0x74, 0xb8, 0x92, 0x72, 0x55, 0xf0, 0x29, 0xab, 0xc4, 0x94, 0x95, 0xa5, 0xd4, 0x4c, 0x0b,
	0x59, 0x2a, 0xeb, 0xba, 0x55, 0x11, 0x5d, 0x34, 0x57, 0x53, 0xa5, 0xeb, 0xe6, 0x52, 0x5b, 0x35,
	0x3d, 0x05, 0x78, 0x51, 0xb3, 0x6a, 0xfd, 0xb2, 0xe1, 0xf5, 0x0d, 0x79, 0x08, 0xd1, 0x3b, 0xf3,
	0x40, 0x83, 0x71, 0x98, 0x75, 0x4f, 0x0e, 0x26, 0xee, 0x3c, 0xb4, 0x9c, 0x6b, 0xa6, 0xf9, 0x35,
	0x2f, 0x75, 0x6e, 0x4d, 0xe9, 0x97, 0x36, 0xf4, 0xef, 0x2a, 0xa4, 0x0f, 0xc1, 0x82, 0x06, 0xe3,
	0x20, 0x4b, 0xe6, 0x3b, 0x79, 0xb0, 0x30, 0x78, 0x46, 0x5b, 0x1e, 0xcf, 0xc8, 0x01, 0x44, 0x05,
	0xbb, 0xe0, 0x05, 0x0d, 0x1d, 0x67, 0x21, 0xc9, 0x20, 0x5c, 0x33, 0x45, 0xdb, 0xe3, 0x20, 0xeb,
	0x9e, 0x0c, 0xfd, 0xb1, 0x73, 0xa6, 0x6e, 0xb7, 0x9e, 0xef, 0xe4, 0xc6, 0x42, 0x06, 0xd0, 0x62,
	0x8a, 0x46, 0x6e, 0x79, 0xcb, 0x32, 0xa2, 0xa4, 0xb1, 0x67, 0x44, 0x49, 0x08, 0x84, 0xb2, 0xd1,
	0x74, 0xd7, 0x51, 0x06, 0x10, 0x0a, 0xb1, 0x28, 0x67, 0xcb, 0x15, 0xa7, 0x1d, 0x47, 0x3b, 0x4c,
	0x46, 0xb0, 0x2b, 0x1b, 0x8d, 0x52, 0xe2, 0x24, 0x4f, 0x90, 0x43, 0xe8, 0x88, 0x72, 0xc1, 0x6b,
	0xcd, 0x3f, 0x50, 0x70, 0xe2, 0x2d, 0x43, 0xfe, 0x85, 0x44, 0x36, 0xda, 0xc9, 0x5d, 0x27, 0xff,
	0xa4, 0xc8, 0x31, 0xc4, 0x8a, 0x17, 0xfc, 0x52, 0xd3, 0x1e, 0x0e, 0xf6, 0x8f, 0x1f, 0xec, 0x1c,
	0xd9, 0xed, 0xd9, 0x9c, 0x11, 0x03, 0x12, 0xd7, 0x42, 0xd3, 0xbf, 0xc6, 0x41, 0x16, 0x62, 0x40,
	0x06, 0x1a, 0xfe, 0x52, 0x36, 0xa5, 0xa6, 0x7d, 0x1f, 0x1c, 0x42, 0x32, 0x84, 0x36, 0x5b, 0x2e,
	0x17, 0x74, 0xcf, 0xd1, 0x88, 0xc8, 0x53, 0xe8, 0x54, 0xb5, 0xac, 0x78, 0xad, 0x6f, 0xe8, 0xc0,
	0x1d, 0x6d, 0x6b, 0x30, 0xf1, 0x35, 0x98, 0x9c, 0x63, 0x0d, 0xcc, 0x3c, 0xde, 0xea, 0x36, 0x9b,
	0xd1, 0xbf, 0xb7, 0x36, 0x9b, 0x99, 0x7c, 0xb5, 0xa4, 0xc4, 0xe7, 0xab, 0xa5, 0xf1, 0x2d, 0x6b,
	0x59, 0xd1, 0x7d, 0xef, 0x33, 0xe8, 0xac, 0x0b, 0x89, 0xf2, 0x13, 0xa5, 0xcf, 0xa1, 0xb7, 0xfd,
	0xf5, 0xc8, 0x00, 0xc2, 0xb7, 0xfc, 0xc6, 0x56, 0x23, 0x37, 0x8f, 0xe4, 0x00, 0xe2, 0xf7, 0x42,
	0xaf, 0x45, 0x49, 0x5b, 0xe3, 0x30, 0x4b, 0x72, 0x87, 0xd2, 0xff, 0x61, 0xef, 0x97, 0x78, 0xc8,
	0x10, 0x22, 0xa5, 0x79, 0xa5, 0xb0, 0x96, 0x49, 0x6e, 0x41, 0x2a, 0x20, 0x76, 0x39, 0x0f, 0x20,
	0x5c, 0x89, 0xa5, 0xdf, 0x7c, 0x25, 0x96, 0x66, 0x85, 0xed, 0x19, 0x76, 0xcf, 0xb7, 0xec, 0x19,
	0x80, 0x9b, 0x55, 0x70, 0x85, 0x15, 0xfc, 0x7d, 0x30, 0xf9, 0x96, 0x35, 0xfd, 0x14, 0x40, 0x1b,
	0xfb, 0x70, 0xdf, 0x93, 0xfa, 0xd8, 0x49, 0x2c, 0x39, 0x36, 0x72, 0x60, 0x1b, 0xd9, 0xb6, 0xeb,
	0x4c, 0x1f, 0xef, 0xbe, 0x4b, 0x74, 0xff, 0x77, 0xf9, 0x1e, 0x40, 0x17, 0x6f, 0x6b, 0xce, 0x55,
	0x53, 0x68, 0x2c, 0x19, 0xba, 0xf0, 0xad, 0xfe, 0xf8, 0xa5, 0x9d, 0x91, 0x64, 0x10, 0x6f, 0x6c,
	0x69, 0x5b, 0xb8, 0xa4, 0xef, 0x7b, 0x69, 0xf3, 0x34, 0x4e, 0xab, 0x93, 0x14, 0xda, 0xdc, 0x5c,
	0x0c, 0x9b, 0x55, 0xcf, 0xfb, 0x4c, 0x16, 0xe6, 0xbb, 0x1b, 0x8d, 0x1c, 0x41, 0x22, 0x4a, 0xfd,
	0x66, 0xc3, 0x8a, 0x86, 0xe3, 0x84, 0x91, 0xbd, 0x24, 0x7a, 0x61, 0x18, 0xf2, 0x1f, 0x74, 0xaf,
	0x0a, 0xc9, 0xbc, 0xc1, 0x4c, 0x1a, 0xcc, 0x77, 0x72, 0x40, 0xd2, 0x5a, 0x8e, 0x4c, 0x73, 0x6a,
	0x67, 0xf0, 0x17, 0xb9, 0xa3, 0x74, 0x8d, 0xf2, 0x59, 0x07, 0xe2, 0x1a, 0x67, 0x3d, 0x79, 0x0d,
	0x91, 0xfd, 0x51, 0xe5, 0x90, 0xbc, 0xaa, 0xd9, 0x86, 0xd7, 0x8a, 0x15, 0x84, 0xdc, 0xf9, 0x4d,
	0xa1, 0x61, 0xb4, 0xef, 0xb9, 0xad, 0xa8, 0xd2, 0xd1, 0xc7, 0xaf, 0xdf, 0x3e, 0xb7, 0x86, 0xe9,
	0xde, 0x74, 0x73, 0x3c, 0x5d, 0x19, 0xf3, 0x23, 0xfc, 0x93, 0x9d, 0x06, 0x0f, 0x1e, 0x07, 0x17,
	0x31, 0x06, 0xf6, 0xe4, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0x94, 0x30, 0x62, 0x45, 0x65, 0x05,
	0x00, 0x00,
}
