package hrpc

import (
	"context"
	"github.com/zhanchangbao/hbasegosdk/pb"
	"google.golang.org/protobuf/proto"
)

// ListProcedure to represent a ListProcedure request
type ListProcedure struct {
	base
}

// NewListProcedure creates a new ListProcedureStruct with default fields
func NewListProcedure() *ListProcedure {
	return &ListProcedure{
		base{
			ctx:      context.Background(),
			table:    []byte{},
			resultch: make(chan RPCResult, 1),
		},
	}
}

// Name returns the name of the rpc function
func (p *ListProcedure) Name() string {
	return "ListProcedures"
}

// Description returns the description of this RPC call.
func (p *ListProcedure) Description() string {
	return p.Name()
}

// ToProto returns the Protobuf message to be sent
func (p *ListProcedure) ToProto() proto.Message {
	return &pb.ListProceduresRequest{}
}

// NewResponse returns the empty protobuf response
func (p *ListProcedure) NewResponse() proto.Message {
	return &pb.ListProceduresResponse{}
}
