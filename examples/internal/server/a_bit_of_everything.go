package server

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	examples "github.com/davidhong1/grpc-gateway/v2/examples/internal/proto/examplepb"
	"github.com/davidhong1/grpc-gateway/v2/examples/internal/proto/oneofenum"
	"github.com/davidhong1/grpc-gateway/v2/examples/internal/proto/pathenum"
	"github.com/davidhong1/grpc-gateway/v2/examples/internal/proto/sub"
	"github.com/davidhong1/grpc-gateway/v2/examples/internal/proto/sub2"
	"github.com/rogpeppe/fastuuid"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	statuspb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// Implements of ABitOfEverythingServiceServer

var uuidgen = fastuuid.MustNewGenerator()

type _ABitOfEverythingServer struct {
	v map[string]*examples.ABitOfEverything
	m sync.Mutex
}

type ABitOfEverythingServer interface {
	examples.ABitOfEverythingServiceServer
	examples.StreamServiceServer
}

func newABitOfEverythingServer() ABitOfEverythingServer {
	return &_ABitOfEverythingServer{
		v: make(map[string]*examples.ABitOfEverything),
	}
}

func (s *_ABitOfEverythingServer) Create(ctx context.Context, msg *examples.ABitOfEverything) (*examples.ABitOfEverything, error) {
	s.m.Lock()
	defer s.m.Unlock()

	grpclog.Info(msg)
	var uuid string
	for {
		uuid = fmt.Sprintf("%x", uuidgen.Next())
		if _, ok := s.v[uuid]; !ok {
			break
		}
	}
	s.v[uuid] = msg
	s.v[uuid].Uuid = uuid
	grpclog.Infof("%v", s.v[uuid])
	return s.v[uuid], nil
}

func (s *_ABitOfEverythingServer) CreateBody(ctx context.Context, msg *examples.ABitOfEverything) (*examples.ABitOfEverything, error) {
	return s.Create(ctx, msg)
}

func (s *_ABitOfEverythingServer) CreateBook(ctx context.Context, req *examples.CreateBookRequest) (*examples.Book, error) {
	return &examples.Book{}, nil
}

func (s *_ABitOfEverythingServer) UpdateBook(ctx context.Context, req *examples.UpdateBookRequest) (*examples.Book, error) {
	return &examples.Book{}, nil
}

func (s *_ABitOfEverythingServer) BulkCreate(stream examples.StreamService_BulkCreateServer) error {
	ctx := stream.Context()

	if header, ok := metadata.FromIncomingContext(ctx); ok {
		if v, ok := header["error"]; ok {
			return status.Errorf(codes.InvalidArgument, "error metadata: %v", v)
		}
	}

	count := 0
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		count++
		grpclog.Info(msg)
		if _, err = s.Create(ctx, msg); err != nil {
			return err
		}
	}

	err := stream.SendHeader(metadata.New(map[string]string{
		"count": fmt.Sprintf("%d", count),
	}))
	if err != nil {
		return nil
	}

	stream.SetTrailer(metadata.New(map[string]string{
		"foo": "foo2",
		"bar": "bar2",
	}))
	return stream.SendAndClose(new(emptypb.Empty))
}

func (s *_ABitOfEverythingServer) Lookup(ctx context.Context, msg *sub2.IdMessage) (*examples.ABitOfEverything, error) {
	s.m.Lock()
	defer s.m.Unlock()
	grpclog.Info(msg)

	err := grpc.SendHeader(ctx, metadata.New(map[string]string{
		"uuid": msg.Uuid,
	}))
	if err != nil {
		return nil, err
	}

	if a, ok := s.v[msg.Uuid]; ok {
		return a, nil
	}

	grpc.SetTrailer(ctx, metadata.New(map[string]string{
		"foo": "foo2",
		"bar": "bar2",
	}))
	return nil, status.Errorf(codes.NotFound, "not found")
}

func (s *_ABitOfEverythingServer) List(opt *examples.Options, stream examples.StreamService_ListServer) error {
	s.m.Lock()
	defer s.m.Unlock()

	err := stream.SendHeader(metadata.New(map[string]string{
		"count": fmt.Sprintf("%d", len(s.v)),
	}))
	if err != nil {
		return nil
	}

	for _, msg := range s.v {
		if err := stream.Send(msg); err != nil {
			return err
		}
	}

	if opt.Error {
		stream.SetTrailer(metadata.New(map[string]string{
			"foo": "foo2",
			"bar": "bar2",
		}))
		return status.Error(codes.InvalidArgument, "error")
	}
	return nil
}

func (s *_ABitOfEverythingServer) Download(opt *examples.Options, stream examples.StreamService_DownloadServer) error {
	msgs := []*httpbody.HttpBody{{
		ContentType: "text/html",
		Data:        []byte("Hello 1"),
	}, {
		ContentType: "text/html",
		Data:        []byte("Hello 2"),
	}}

	for _, msg := range msgs {
		if err := stream.Send(msg); err != nil {
			return err
		}

		time.Sleep(5 * time.Millisecond)
	}

	if opt.Error {
		stream.SetTrailer(metadata.New(map[string]string{
			"foo": "foo2",
			"bar": "bar2",
		}))
		return status.Error(codes.InvalidArgument, "error")
	}
	return nil
}

func (s *_ABitOfEverythingServer) Custom(ctx context.Context, msg *examples.ABitOfEverything) (*examples.ABitOfEverything, error) {
	s.m.Lock()
	defer s.m.Unlock()

	grpclog.Info(msg)
	if _, ok := s.v[msg.Uuid]; ok {
		s.v[msg.Uuid] = msg
	} else {
		return nil, status.Errorf(codes.NotFound, "not found")
	}
	return msg, nil
}

func (s *_ABitOfEverythingServer) DoubleColon(ctx context.Context, msg *examples.ABitOfEverything) (*examples.ABitOfEverything, error) {
	s.m.Lock()
	defer s.m.Unlock()

	grpclog.Info(msg)
	if _, ok := s.v[msg.Uuid]; ok {
		s.v[msg.Uuid] = msg
	} else {
		return nil, status.Errorf(codes.NotFound, "not found")
	}
	return msg, nil
}

func (s *_ABitOfEverythingServer) Update(ctx context.Context, msg *examples.ABitOfEverything) (*emptypb.Empty, error) {
	s.m.Lock()
	defer s.m.Unlock()

	grpclog.Info(msg)
	if _, ok := s.v[msg.Uuid]; ok {
		s.v[msg.Uuid] = msg
	} else {
		return nil, status.Errorf(codes.NotFound, "not found")
	}
	return new(emptypb.Empty), nil
}

func (s *_ABitOfEverythingServer) UpdateV2(ctx context.Context, msg *examples.UpdateV2Request) (*emptypb.Empty, error) {
	grpclog.Info(msg)
	// If there is no update mask do a regular update
	if msg.UpdateMask == nil || len(msg.UpdateMask.GetPaths()) == 0 {
		return s.Update(ctx, msg.Abe)
	}

	s.m.Lock()
	defer s.m.Unlock()
	if a, ok := s.v[msg.Abe.Uuid]; ok {
		applyFieldMask(a, msg.Abe, msg.UpdateMask)
	} else {
		return nil, status.Errorf(codes.NotFound, "not found")
	}
	return new(emptypb.Empty), nil
}

func (s *_ABitOfEverythingServer) Delete(ctx context.Context, msg *sub2.IdMessage) (*emptypb.Empty, error) {
	s.m.Lock()
	defer s.m.Unlock()

	grpclog.Info(msg)
	if _, ok := s.v[msg.Uuid]; ok {
		delete(s.v, msg.Uuid)
	} else {
		return nil, status.Errorf(codes.NotFound, "not found")
	}
	return new(emptypb.Empty), nil
}

func (s *_ABitOfEverythingServer) GetQuery(ctx context.Context, msg *examples.ABitOfEverything) (*emptypb.Empty, error) {
	s.m.Lock()
	defer s.m.Unlock()

	grpclog.Info(msg)
	if _, ok := s.v[msg.Uuid]; ok {
		s.v[msg.Uuid] = msg
	} else {
		return nil, status.Errorf(codes.NotFound, "not found")
	}
	return new(emptypb.Empty), nil
}

func (s *_ABitOfEverythingServer) GetRepeatedQuery(ctx context.Context, msg *examples.ABitOfEverythingRepeated) (*examples.ABitOfEverythingRepeated, error) {
	s.m.Lock()
	defer s.m.Unlock()

	grpclog.Info(msg)
	return msg, nil
}

func (s *_ABitOfEverythingServer) Echo(ctx context.Context, msg *sub.StringMessage) (*sub.StringMessage, error) {
	s.m.Lock()
	defer s.m.Unlock()

	grpclog.Info(msg)
	return msg, nil
}

func (s *_ABitOfEverythingServer) BulkEcho(stream examples.StreamService_BulkEchoServer) error {
	var msgs []*sub.StringMessage
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		msgs = append(msgs, msg)
	}

	hmd := metadata.New(map[string]string{
		"foo": "foo1",
		"bar": "bar1",
	})
	if err := stream.SendHeader(hmd); err != nil {
		return err
	}

	for _, msg := range msgs {
		grpclog.Info(msg)
		if err := stream.Send(msg); err != nil {
			return err
		}
	}

	stream.SetTrailer(metadata.New(map[string]string{
		"foo": "foo2",
		"bar": "bar2",
	}))
	return nil
}

func (s *_ABitOfEverythingServer) DeepPathEcho(ctx context.Context, msg *examples.ABitOfEverything) (*examples.ABitOfEverything, error) {
	s.m.Lock()
	defer s.m.Unlock()

	grpclog.Info(msg)
	return msg, nil
}

func (s *_ABitOfEverythingServer) NoBindings(ctx context.Context, msg *durationpb.Duration) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *_ABitOfEverythingServer) Timeout(ctx context.Context, msg *emptypb.Empty) (*emptypb.Empty, error) {
	<-ctx.Done()
	return nil, status.FromContextError(ctx.Err()).Err()
}

func (s *_ABitOfEverythingServer) ErrorWithDetails(ctx context.Context, msg *emptypb.Empty) (*emptypb.Empty, error) {
	stat, err := status.New(codes.Unknown, "with details").
		WithDetails(&errdetails.DebugInfo{
			StackEntries: []string{"foo:1"},
			Detail:       "error debug details",
		})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unexpected error adding details: %s", err)
	}
	return nil, stat.Err()
}

func (s *_ABitOfEverythingServer) GetMessageWithBody(ctx context.Context, msg *examples.MessageWithBody) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *_ABitOfEverythingServer) PostWithEmptyBody(ctx context.Context, msg *examples.Body) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *_ABitOfEverythingServer) CheckGetQueryParams(ctx context.Context, msg *examples.ABitOfEverything) (*examples.ABitOfEverything, error) {
	return msg, nil
}

func (s *_ABitOfEverythingServer) CheckNestedEnumGetQueryParams(ctx context.Context, msg *examples.ABitOfEverything) (*examples.ABitOfEverything, error) {
	return msg, nil
}

func (s *_ABitOfEverythingServer) CheckPostQueryParams(ctx context.Context, msg *examples.ABitOfEverything) (*examples.ABitOfEverything, error) {
	return msg, nil
}

func (s *_ABitOfEverythingServer) OverwriteRequestContentType(ctx context.Context, msg *examples.Body) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *_ABitOfEverythingServer) OverwriteResponseContentType(ctx context.Context, msg *emptypb.Empty) (*wrapperspb.StringValue, error) {
	return &wrapperspb.StringValue{}, nil
}

func (s *_ABitOfEverythingServer) CheckExternalPathEnum(ctx context.Context, msg *pathenum.MessageWithPathEnum) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *_ABitOfEverythingServer) CheckExternalNestedPathEnum(ctx context.Context, msg *pathenum.MessageWithNestedPathEnum) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *_ABitOfEverythingServer) CheckStatus(ctx context.Context, empty *emptypb.Empty) (*examples.CheckStatusResponse, error) {
	return &examples.CheckStatusResponse{Status: &statuspb.Status{}}, nil
}

func (s *_ABitOfEverythingServer) Exists(ctx context.Context, msg *examples.ABitOfEverything) (*emptypb.Empty, error) {
	if _, ok := s.v[msg.Uuid]; ok {
		return new(emptypb.Empty), nil
	}

	return nil, status.Errorf(codes.NotFound, "not found")
}

func (s *_ABitOfEverythingServer) CustomOptionsRequest(ctx context.Context, msg *examples.ABitOfEverything) (*emptypb.Empty, error) {
	err := grpc.SendHeader(ctx, metadata.New(map[string]string{
		"Allow": "OPTIONS, GET, HEAD, POST, PUT, TRACE",
	}))
	if err != nil {
		return nil, err
	}
	return new(emptypb.Empty), nil
}

func (s *_ABitOfEverythingServer) TraceRequest(ctx context.Context, msg *examples.ABitOfEverything) (*examples.ABitOfEverything, error) {
	return msg, nil
}

func (s *_ABitOfEverythingServer) PostOneofEnum(ctx context.Context, msg *oneofenum.OneofEnumMessage) (*emptypb.Empty, error) {
	return new(emptypb.Empty), nil
}

func (s *_ABitOfEverythingServer) PostRequiredMessageType(ctx context.Context, req *examples.RequiredMessageTypeRequest) (*emptypb.Empty, error) {
	return new(emptypb.Empty), nil
}
