// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.3
// Source: meta_land_worker.proto

package metalandworker

import (
	"context"

	"metaLand/app/worker/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	TriggerUCETaskRequest  = pb.TriggerUCETaskRequest
	TriggerUCETaskResponse = pb.TriggerUCETaskResponse

	MetaLandWorker interface {
		TriggerUCETask(ctx context.Context, in *TriggerUCETaskRequest, opts ...grpc.CallOption) (*TriggerUCETaskResponse, error)
	}

	defaultMetaLandWorker struct {
		cli zrpc.Client
	}
)

func NewMetaLandWorker(cli zrpc.Client) MetaLandWorker {
	return &defaultMetaLandWorker{
		cli: cli,
	}
}

func (m *defaultMetaLandWorker) TriggerUCETask(ctx context.Context, in *TriggerUCETaskRequest, opts ...grpc.CallOption) (*TriggerUCETaskResponse, error) {
	client := pb.NewMetaLandWorkerClient(m.cli.Conn())
	return client.TriggerUCETask(ctx, in, opts...)
}
