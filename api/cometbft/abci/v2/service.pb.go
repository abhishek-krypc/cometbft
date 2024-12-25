// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cometbft/abci/v2/service.proto

package v2

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("cometbft/abci/v2/service.proto", fileDescriptor_9a03478482bd42d5) }

var fileDescriptor_9a03478482bd42d5 = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0xdd, 0x6e, 0xd3, 0x30,
	0x18, 0x86, 0x5b, 0x69, 0x0c, 0xe1, 0x09, 0x01, 0xe6, 0x6c, 0x82, 0x30, 0x7e, 0xc4, 0x3f, 0x89,
	0x28, 0x57, 0xb0, 0x46, 0x9d, 0xa8, 0x36, 0xc1, 0x60, 0xd3, 0x90, 0x38, 0x22, 0xcd, 0xbe, 0x10,
	0xab, 0xa9, 0x6d, 0x6c, 0xa7, 0x5a, 0xb9, 0x0a, 0x2e, 0x8a, 0x03, 0x0e, 0x7b, 0xc8, 0x21, 0x6a,
	0x6f, 0x04, 0xa5, 0xb1, 0x09, 0x49, 0x5c, 0xaf, 0xa7, 0x7e, 0x9f, 0x3c, 0xaf, 0x1d, 0x7d, 0xd2,
	0x87, 0xbc, 0x98, 0x4d, 0x40, 0x8d, 0x12, 0x15, 0x44, 0xa3, 0x98, 0x04, 0xd3, 0x5e, 0x20, 0x41,
	0x4c, 0x49, 0x0c, 0x3e, 0x17, 0x4c, 0x31, 0x7c, 0xd3, 0xe4, 0x7e, 0x91, 0xfb, 0xd3, 0xde, 0xee,
	0x9d, 0xd6, 0x17, 0x6a, 0xc6, 0x41, 0x96, 0x7c, 0xef, 0xe7, 0x0e, 0xda, 0xd9, 0xef, 0x87, 0xc3,
	0x93, 0xd2, 0x82, 0x07, 0x68, 0x6b, 0x10, 0xa7, 0x0c, 0xdf, 0xf5, 0x9b, 0x22, 0xbf, 0x38, 0xff,
	0x08, 0xdf, 0x72, 0x90, 0x6a, 0xd7, 0x5b, 0x17, 0x4b, 0xce, 0xa8, 0x04, 0xfc, 0x16, 0x5d, 0x39,
	0xc8, 0x72, 0x99, 0x62, 0x0b, 0xb8, 0x0a, 0x8c, 0xe8, 0xde, 0xda, 0x5c, 0x9b, 0x06, 0x68, 0x6b,
	0x48, 0x13, 0xeb, 0x85, 0x8a, 0x73, 0xc7, 0x85, 0xca, 0x58, 0x6b, 0xde, 0xa1, 0xab, 0x61, 0x0a,
	0xf1, 0xf8, 0xf4, 0x02, 0xef, 0xb5, 0x51, 0x1d, 0x19, 0xd9, 0x7d, 0x07, 0x51, 0x3d, 0xf0, 0x43,
	0x0e, 0x62, 0x66, 0x7b, 0xe0, 0x2a, 0x70, 0x3c, 0x50, 0xe7, 0xda, 0x74, 0x88, 0xb6, 0x43, 0x36,
	0x99, 0x10, 0x85, 0x2d, 0x68, 0x99, 0x18, 0xd7, 0xde, 0x7a, 0x40, 0xcb, 0x4e, 0xd1, 0xb5, 0x21,
	0x25, 0x2a, 0x4c, 0x23, 0x42, 0xf1, 0x03, 0xdb, 0x3f, 0xd1, 0xa1, 0x51, 0x3e, 0x74, 0x32, 0xda,
	0xfa, 0x05, 0x5d, 0x3f, 0x22, 0x52, 0x9d, 0xd0, 0x88, 0xcb, 0x94, 0x29, 0x89, 0x1f, 0xb7, 0xbf,
	0xaa, 0x01, 0xc6, 0xfe, 0xe4, 0x52, 0xae, 0x6a, 0x78, 0x9f, 0x24, 0x20, 0x4c, 0x62, 0x6b, 0xa8,
	0x01, 0x8e, 0x86, 0x06, 0xa7, 0x1b, 0x32, 0x74, 0xeb, 0x88, 0x45, 0xe7, 0xe6, 0x3c, 0x4c, 0x73,
	0x3a, 0xc6, 0xcf, 0x2d, 0xf7, 0x6b, 0x42, 0xa6, 0xe9, 0xc5, 0x46, 0xac, 0x6e, 0x63, 0x08, 0xef,
	0x73, 0x9e, 0xcd, 0xea, 0x75, 0x16, 0x45, 0x9b, 0x32, 0x7d, 0x2f, 0x37, 0x83, 0x75, 0x61, 0x82,
	0x6e, 0x1c, 0x0b, 0xe0, 0x91, 0x80, 0x63, 0xc1, 0x38, 0x93, 0x51, 0x86, 0x9f, 0xb6, 0x05, 0x0d,
	0xc4, 0x54, 0x3d, 0xdb, 0x80, 0xfc, 0xbf, 0x87, 0xc5, 0x20, 0xa5, 0xbb, 0xa7, 0x86, 0x38, 0x7b,
	0x1a, 0xa4, 0xee, 0xf9, 0x84, 0xd0, 0xe0, 0x42, 0x01, 0x3d, 0x3f, 0x63, 0x0a, 0xb0, 0x65, 0x4a,
	0xab, 0xd4, 0xd8, 0x1f, 0xb9, 0x21, 0x2d, 0x16, 0xe8, 0xf6, 0x19, 0x08, 0x92, 0xcc, 0x8a, 0xd3,
	0x55, 0x2e, 0x09, 0xa3, 0xd8, 0xf2, 0xb7, 0x2d, 0x98, 0xa9, 0x7a, 0xb5, 0x21, 0x5d, 0x4d, 0xf7,
	0x01, 0xa1, 0x51, 0x46, 0xbe, 0x43, 0x3f, 0x63, 0xf1, 0xd8, 0x36, 0xdd, 0x35, 0xc0, 0x31, 0xdd,
	0x0d, 0xae, 0x6c, 0xe8, 0x1f, 0xfe, 0x5a, 0x78, 0xdd, 0xf9, 0xc2, 0xeb, 0xfe, 0x59, 0x78, 0xdd,
	0x1f, 0x4b, 0xaf, 0x33, 0x5f, 0x7a, 0x9d, 0xdf, 0x4b, 0xaf, 0xf3, 0xf9, 0xf5, 0x57, 0xa2, 0xd2,
	0x7c, 0x54, 0x88, 0x82, 0x7f, 0x9b, 0xa0, 0x5a, 0x09, 0x9c, 0x04, 0xcd, 0xfd, 0x30, 0xda, 0x5e,
	0xad, 0x86, 0x37, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x61, 0xbd, 0x0a, 0x94, 0x6c, 0x06, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ABCIServiceClient is the client API for ABCIService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ABCIServiceClient interface {
	// Echo returns back the same message it is sent.
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
	// Flush flushes the write buffer.
	Flush(ctx context.Context, in *FlushRequest, opts ...grpc.CallOption) (*FlushResponse, error)
	// Info returns information about the application state.
	Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error)
	// CheckTx validates a transaction.
	CheckTx(ctx context.Context, in *CheckTxRequest, opts ...grpc.CallOption) (*CheckTxResponse, error)
	// Query queries the application state.
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
	// Commit commits a block of transactions.
	Commit(ctx context.Context, in *CommitRequest, opts ...grpc.CallOption) (*CommitResponse, error)
	// InitChain initializes the blockchain.
	InitChain(ctx context.Context, in *InitChainRequest, opts ...grpc.CallOption) (*InitChainResponse, error)
	// ListSnapshots lists all the available snapshots.
	ListSnapshots(ctx context.Context, in *ListSnapshotsRequest, opts ...grpc.CallOption) (*ListSnapshotsResponse, error)
	// OfferSnapshot sends a snapshot offer.
	OfferSnapshot(ctx context.Context, in *OfferSnapshotRequest, opts ...grpc.CallOption) (*OfferSnapshotResponse, error)
	// LoadSnapshotChunk returns a chunk of snapshot.
	LoadSnapshotChunk(ctx context.Context, in *LoadSnapshotChunkRequest, opts ...grpc.CallOption) (*LoadSnapshotChunkResponse, error)
	// ApplySnapshotChunk applies a chunk of snapshot.
	ApplySnapshotChunk(ctx context.Context, in *ApplySnapshotChunkRequest, opts ...grpc.CallOption) (*ApplySnapshotChunkResponse, error)
	// PrepareProposal returns a proposal for the next block.
	PrepareProposal(ctx context.Context, in *PrepareProposalRequest, opts ...grpc.CallOption) (*PrepareProposalResponse, error)
	// ProcessProposal validates a proposal.
	ProcessProposal(ctx context.Context, in *ProcessProposalRequest, opts ...grpc.CallOption) (*ProcessProposalResponse, error)
	// ExtendVote extends a vote with application-injected data (vote extensions).
	ExtendVote(ctx context.Context, in *ExtendVoteRequest, opts ...grpc.CallOption) (*ExtendVoteResponse, error)
	// VerifyVoteExtension verifies a vote extension.
	VerifyVoteExtension(ctx context.Context, in *VerifyVoteExtensionRequest, opts ...grpc.CallOption) (*VerifyVoteExtensionResponse, error)
	// FinalizeBlock finalizes a block.
	FinalizeBlock(ctx context.Context, in *FinalizeBlockRequest, opts ...grpc.CallOption) (*FinalizeBlockResponse, error)
}

type aBCIServiceClient struct {
	cc grpc1.ClientConn
}

func NewABCIServiceClient(cc grpc1.ClientConn) ABCIServiceClient {
	return &aBCIServiceClient{cc}
}

func (c *aBCIServiceClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/cometbft.abci.v2.ABCIService/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIServiceClient) Flush(ctx context.Context, in *FlushRequest, opts ...grpc.CallOption) (*FlushResponse, error) {
	out := new(FlushResponse)
	err := c.cc.Invoke(ctx, "/cometbft.abci.v2.ABCIService/Flush", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIServiceClient) Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := c.cc.Invoke(ctx, "/cometbft.abci.v2.ABCIService/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIServiceClient) CheckTx(ctx context.Context, in *CheckTxRequest, opts ...grpc.CallOption) (*CheckTxResponse, error) {
	out := new(CheckTxResponse)
	err := c.cc.Invoke(ctx, "/cometbft.abci.v2.ABCIService/CheckTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIServiceClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/cometbft.abci.v2.ABCIService/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIServiceClient) Commit(ctx context.Context, in *CommitRequest, opts ...grpc.CallOption) (*CommitResponse, error) {
	out := new(CommitResponse)
	err := c.cc.Invoke(ctx, "/cometbft.abci.v2.ABCIService/Commit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIServiceClient) InitChain(ctx context.Context, in *InitChainRequest, opts ...grpc.CallOption) (*InitChainResponse, error) {
	out := new(InitChainResponse)
	err := c.cc.Invoke(ctx, "/cometbft.abci.v2.ABCIService/InitChain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIServiceClient) ListSnapshots(ctx context.Context, in *ListSnapshotsRequest, opts ...grpc.CallOption) (*ListSnapshotsResponse, error) {
	out := new(ListSnapshotsResponse)
	err := c.cc.Invoke(ctx, "/cometbft.abci.v2.ABCIService/ListSnapshots", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIServiceClient) OfferSnapshot(ctx context.Context, in *OfferSnapshotRequest, opts ...grpc.CallOption) (*OfferSnapshotResponse, error) {
	out := new(OfferSnapshotResponse)
	err := c.cc.Invoke(ctx, "/cometbft.abci.v2.ABCIService/OfferSnapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIServiceClient) LoadSnapshotChunk(ctx context.Context, in *LoadSnapshotChunkRequest, opts ...grpc.CallOption) (*LoadSnapshotChunkResponse, error) {
	out := new(LoadSnapshotChunkResponse)
	err := c.cc.Invoke(ctx, "/cometbft.abci.v2.ABCIService/LoadSnapshotChunk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIServiceClient) ApplySnapshotChunk(ctx context.Context, in *ApplySnapshotChunkRequest, opts ...grpc.CallOption) (*ApplySnapshotChunkResponse, error) {
	out := new(ApplySnapshotChunkResponse)
	err := c.cc.Invoke(ctx, "/cometbft.abci.v2.ABCIService/ApplySnapshotChunk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIServiceClient) PrepareProposal(ctx context.Context, in *PrepareProposalRequest, opts ...grpc.CallOption) (*PrepareProposalResponse, error) {
	out := new(PrepareProposalResponse)
	err := c.cc.Invoke(ctx, "/cometbft.abci.v2.ABCIService/PrepareProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIServiceClient) ProcessProposal(ctx context.Context, in *ProcessProposalRequest, opts ...grpc.CallOption) (*ProcessProposalResponse, error) {
	out := new(ProcessProposalResponse)
	err := c.cc.Invoke(ctx, "/cometbft.abci.v2.ABCIService/ProcessProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIServiceClient) ExtendVote(ctx context.Context, in *ExtendVoteRequest, opts ...grpc.CallOption) (*ExtendVoteResponse, error) {
	out := new(ExtendVoteResponse)
	err := c.cc.Invoke(ctx, "/cometbft.abci.v2.ABCIService/ExtendVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIServiceClient) VerifyVoteExtension(ctx context.Context, in *VerifyVoteExtensionRequest, opts ...grpc.CallOption) (*VerifyVoteExtensionResponse, error) {
	out := new(VerifyVoteExtensionResponse)
	err := c.cc.Invoke(ctx, "/cometbft.abci.v2.ABCIService/VerifyVoteExtension", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIServiceClient) FinalizeBlock(ctx context.Context, in *FinalizeBlockRequest, opts ...grpc.CallOption) (*FinalizeBlockResponse, error) {
	out := new(FinalizeBlockResponse)
	err := c.cc.Invoke(ctx, "/cometbft.abci.v2.ABCIService/FinalizeBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ABCIServiceServer is the server API for ABCIService service.
type ABCIServiceServer interface {
	// Echo returns back the same message it is sent.
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
	// Flush flushes the write buffer.
	Flush(context.Context, *FlushRequest) (*FlushResponse, error)
	// Info returns information about the application state.
	Info(context.Context, *InfoRequest) (*InfoResponse, error)
	// CheckTx validates a transaction.
	CheckTx(context.Context, *CheckTxRequest) (*CheckTxResponse, error)
	// Query queries the application state.
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
	// Commit commits a block of transactions.
	Commit(context.Context, *CommitRequest) (*CommitResponse, error)
	// InitChain initializes the blockchain.
	InitChain(context.Context, *InitChainRequest) (*InitChainResponse, error)
	// ListSnapshots lists all the available snapshots.
	ListSnapshots(context.Context, *ListSnapshotsRequest) (*ListSnapshotsResponse, error)
	// OfferSnapshot sends a snapshot offer.
	OfferSnapshot(context.Context, *OfferSnapshotRequest) (*OfferSnapshotResponse, error)
	// LoadSnapshotChunk returns a chunk of snapshot.
	LoadSnapshotChunk(context.Context, *LoadSnapshotChunkRequest) (*LoadSnapshotChunkResponse, error)
	// ApplySnapshotChunk applies a chunk of snapshot.
	ApplySnapshotChunk(context.Context, *ApplySnapshotChunkRequest) (*ApplySnapshotChunkResponse, error)
	// PrepareProposal returns a proposal for the next block.
	PrepareProposal(context.Context, *PrepareProposalRequest) (*PrepareProposalResponse, error)
	// ProcessProposal validates a proposal.
	ProcessProposal(context.Context, *ProcessProposalRequest) (*ProcessProposalResponse, error)
	// ExtendVote extends a vote with application-injected data (vote extensions).
	ExtendVote(context.Context, *ExtendVoteRequest) (*ExtendVoteResponse, error)
	// VerifyVoteExtension verifies a vote extension.
	VerifyVoteExtension(context.Context, *VerifyVoteExtensionRequest) (*VerifyVoteExtensionResponse, error)
	// FinalizeBlock finalizes a block.
	FinalizeBlock(context.Context, *FinalizeBlockRequest) (*FinalizeBlockResponse, error)
}

// UnimplementedABCIServiceServer can be embedded to have forward compatible implementations.
type UnimplementedABCIServiceServer struct {
}

func (*UnimplementedABCIServiceServer) Echo(ctx context.Context, req *EchoRequest) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (*UnimplementedABCIServiceServer) Flush(ctx context.Context, req *FlushRequest) (*FlushResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Flush not implemented")
}
func (*UnimplementedABCIServiceServer) Info(ctx context.Context, req *InfoRequest) (*InfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (*UnimplementedABCIServiceServer) CheckTx(ctx context.Context, req *CheckTxRequest) (*CheckTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckTx not implemented")
}
func (*UnimplementedABCIServiceServer) Query(ctx context.Context, req *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (*UnimplementedABCIServiceServer) Commit(ctx context.Context, req *CommitRequest) (*CommitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Commit not implemented")
}
func (*UnimplementedABCIServiceServer) InitChain(ctx context.Context, req *InitChainRequest) (*InitChainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitChain not implemented")
}
func (*UnimplementedABCIServiceServer) ListSnapshots(ctx context.Context, req *ListSnapshotsRequest) (*ListSnapshotsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSnapshots not implemented")
}
func (*UnimplementedABCIServiceServer) OfferSnapshot(ctx context.Context, req *OfferSnapshotRequest) (*OfferSnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OfferSnapshot not implemented")
}
func (*UnimplementedABCIServiceServer) LoadSnapshotChunk(ctx context.Context, req *LoadSnapshotChunkRequest) (*LoadSnapshotChunkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadSnapshotChunk not implemented")
}
func (*UnimplementedABCIServiceServer) ApplySnapshotChunk(ctx context.Context, req *ApplySnapshotChunkRequest) (*ApplySnapshotChunkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplySnapshotChunk not implemented")
}
func (*UnimplementedABCIServiceServer) PrepareProposal(ctx context.Context, req *PrepareProposalRequest) (*PrepareProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrepareProposal not implemented")
}
func (*UnimplementedABCIServiceServer) ProcessProposal(ctx context.Context, req *ProcessProposalRequest) (*ProcessProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessProposal not implemented")
}
func (*UnimplementedABCIServiceServer) ExtendVote(ctx context.Context, req *ExtendVoteRequest) (*ExtendVoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExtendVote not implemented")
}
func (*UnimplementedABCIServiceServer) VerifyVoteExtension(ctx context.Context, req *VerifyVoteExtensionRequest) (*VerifyVoteExtensionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyVoteExtension not implemented")
}
func (*UnimplementedABCIServiceServer) FinalizeBlock(ctx context.Context, req *FinalizeBlockRequest) (*FinalizeBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinalizeBlock not implemented")
}

func RegisterABCIServiceServer(s grpc1.Server, srv ABCIServiceServer) {
	s.RegisterService(&_ABCIService_serviceDesc, srv)
}

func _ABCIService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cometbft.abci.v2.ABCIService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIServiceServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIService_Flush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIServiceServer).Flush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cometbft.abci.v2.ABCIService/Flush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIServiceServer).Flush(ctx, req.(*FlushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIService_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIServiceServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cometbft.abci.v2.ABCIService/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIServiceServer).Info(ctx, req.(*InfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIService_CheckTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIServiceServer).CheckTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cometbft.abci.v2.ABCIService/CheckTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIServiceServer).CheckTx(ctx, req.(*CheckTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIService_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIServiceServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cometbft.abci.v2.ABCIService/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIServiceServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIService_Commit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIServiceServer).Commit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cometbft.abci.v2.ABCIService/Commit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIServiceServer).Commit(ctx, req.(*CommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIService_InitChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitChainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIServiceServer).InitChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cometbft.abci.v2.ABCIService/InitChain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIServiceServer).InitChain(ctx, req.(*InitChainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIService_ListSnapshots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSnapshotsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIServiceServer).ListSnapshots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cometbft.abci.v2.ABCIService/ListSnapshots",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIServiceServer).ListSnapshots(ctx, req.(*ListSnapshotsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIService_OfferSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OfferSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIServiceServer).OfferSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cometbft.abci.v2.ABCIService/OfferSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIServiceServer).OfferSnapshot(ctx, req.(*OfferSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIService_LoadSnapshotChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadSnapshotChunkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIServiceServer).LoadSnapshotChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cometbft.abci.v2.ABCIService/LoadSnapshotChunk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIServiceServer).LoadSnapshotChunk(ctx, req.(*LoadSnapshotChunkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIService_ApplySnapshotChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplySnapshotChunkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIServiceServer).ApplySnapshotChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cometbft.abci.v2.ABCIService/ApplySnapshotChunk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIServiceServer).ApplySnapshotChunk(ctx, req.(*ApplySnapshotChunkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIService_PrepareProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareProposalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIServiceServer).PrepareProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cometbft.abci.v2.ABCIService/PrepareProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIServiceServer).PrepareProposal(ctx, req.(*PrepareProposalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIService_ProcessProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessProposalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIServiceServer).ProcessProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cometbft.abci.v2.ABCIService/ProcessProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIServiceServer).ProcessProposal(ctx, req.(*ProcessProposalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIService_ExtendVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExtendVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIServiceServer).ExtendVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cometbft.abci.v2.ABCIService/ExtendVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIServiceServer).ExtendVote(ctx, req.(*ExtendVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIService_VerifyVoteExtension_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyVoteExtensionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIServiceServer).VerifyVoteExtension(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cometbft.abci.v2.ABCIService/VerifyVoteExtension",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIServiceServer).VerifyVoteExtension(ctx, req.(*VerifyVoteExtensionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIService_FinalizeBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinalizeBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIServiceServer).FinalizeBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cometbft.abci.v2.ABCIService/FinalizeBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIServiceServer).FinalizeBlock(ctx, req.(*FinalizeBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var ABCIService_serviceDesc = _ABCIService_serviceDesc
var _ABCIService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cometbft.abci.v2.ABCIService",
	HandlerType: (*ABCIServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _ABCIService_Echo_Handler,
		},
		{
			MethodName: "Flush",
			Handler:    _ABCIService_Flush_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _ABCIService_Info_Handler,
		},
		{
			MethodName: "CheckTx",
			Handler:    _ABCIService_CheckTx_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _ABCIService_Query_Handler,
		},
		{
			MethodName: "Commit",
			Handler:    _ABCIService_Commit_Handler,
		},
		{
			MethodName: "InitChain",
			Handler:    _ABCIService_InitChain_Handler,
		},
		{
			MethodName: "ListSnapshots",
			Handler:    _ABCIService_ListSnapshots_Handler,
		},
		{
			MethodName: "OfferSnapshot",
			Handler:    _ABCIService_OfferSnapshot_Handler,
		},
		{
			MethodName: "LoadSnapshotChunk",
			Handler:    _ABCIService_LoadSnapshotChunk_Handler,
		},
		{
			MethodName: "ApplySnapshotChunk",
			Handler:    _ABCIService_ApplySnapshotChunk_Handler,
		},
		{
			MethodName: "PrepareProposal",
			Handler:    _ABCIService_PrepareProposal_Handler,
		},
		{
			MethodName: "ProcessProposal",
			Handler:    _ABCIService_ProcessProposal_Handler,
		},
		{
			MethodName: "ExtendVote",
			Handler:    _ABCIService_ExtendVote_Handler,
		},
		{
			MethodName: "VerifyVoteExtension",
			Handler:    _ABCIService_VerifyVoteExtension_Handler,
		},
		{
			MethodName: "FinalizeBlock",
			Handler:    _ABCIService_FinalizeBlock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cometbft/abci/v2/service.proto",
}