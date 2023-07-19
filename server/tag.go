package server

import (
	"context"
	"encoding/json"

	"github.com/lc-1010/OneBlogService/pkg/blog_api"
	"github.com/lc-1010/OneBlogService/pkg/errcode"
	pb "github.com/lc-1010/OneBlogService/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TagService struct {
	pb.UnimplementedTagServiceServer
}

func NewTagServer() *TagService {
	return &TagService{}
}

// GetTagList(context.Context, *GetTagListRequst) (*GetTagListReply, error)
func (t *TagService) GetTagList(c context.Context, r *pb.GetTagListRequst) (*pb.GetTagListReply, error) {

	api := blog_api.NewAPI("http://127.0.0.1:8000")
	body, err := api.GetTagList(c, r.GetName())
	if err != nil {
		return nil, errcode.TogRPCError(errcode.ErrorGetTagListFail)
	}
	tagList := pb.GetTagListReply{}
	err = json.Unmarshal(body, &tagList)
	if err != nil {
		return nil, errcode.TogRPCError(errcode.ServerError)
	}
	return &tagList, nil
}

func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return grpc.DialContext(ctx, target, opts...)
}
