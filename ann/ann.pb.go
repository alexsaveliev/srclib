// Code generated by protoc-gen-gogo.
// source: ann.proto
// DO NOT EDIT!

/*
Package ann is a generated protocol buffer package.

It is generated from these files:
	ann.proto

It has these top-level messages:
	Ann
*/
package ann

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto"

import sourcegraph_com_sqs_pbtypes "sourcegraph.com/sqs/pbtypes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// An Ann is a source code annotation.
//
// Annotations are unique on (Repo, CommitID, UnitType, Unit, File,
// StartLine, EndLine, Type).
type Ann struct {
	// Repo is the VCS repository in which this ann exists.
	Repo string `protobuf:"bytes,1,opt,name=Repo,proto3" json:"Repo,omitempty"`
	// CommitID is the ID of the VCS commit that this ann exists
	// in. The CommitID is always a full commit ID (40 hexadecimal
	// characters for git and hg), never a branch or tag name.
	CommitID string `protobuf:"bytes,2,opt,name=CommitID,proto3" json:"CommitID,omitempty"`
	// UnitType is the source unit type that the annotation exists
	// on. It is either the source unit type during whose processing
	// the annotation was detected/created. Multiple annotations may
	// exist on the same file from different source unit types if a
	// file is contained in multiple source units.
	UnitType string `protobuf:"bytes,3,opt,name=UnitType,proto3" json:"UnitType,omitempty"`
	// Unit is the name of the source unit that this ann exists in.
	Unit string `protobuf:"bytes,4,opt,name=Unit,proto3" json:"Unit,omitempty"`
	// File is the filename in which this Ann exists.
	File string `protobuf:"bytes,5,opt,name=File,proto3" json:"File,omitempty"`
	// StartLine is the line number (inclusive, 1-indexed) of the
	// beginning of the annotation.
	StartLine uint32 `protobuf:"varint,6,opt,name=StartLine,proto3" json:"StartLine"`
	// EndLine is the line number (inclusive, 1-indexed) of the end of
	// the annotation.
	EndLine uint32 `protobuf:"varint,7,opt,name=EndLine,proto3" json:"EndLine"`
	// Type is the type of the annotation. See this package's type
	// constants for a list of possible types.
	Type string `protobuf:"bytes,8,opt,name=Type,proto3" json:"Type"`
	// Data contains arbitrary JSON data that is specific to this
	// annotation type (e.g., the link URL for Link annotations).
	Data sourcegraph_com_sqs_pbtypes.RawMessage `protobuf:"bytes,9,opt,name=Data,proto3,customtype=sourcegraph.com/sqs/pbtypes.RawMessage" json:"Data,omitempty"`
}

func (m *Ann) Reset()         { *m = Ann{} }
func (m *Ann) String() string { return proto.CompactTextString(m) }
func (*Ann) ProtoMessage()    {}
