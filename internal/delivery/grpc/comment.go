package grpcHandler

import (
	"context"
	"fmt"

	"github.com/kodinggo/gb-2-api-comment-service/internal/model"
	pb "github.com/kodinggo/gb-2-api-comment-service/pb/comment_service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CommentgRPCHandler struct {
	pb.UnimplementedCommentServiceServer
	commentUsecase model.CommentUseCase
}

func InitgRPCHanlder(commentUsecase model.CommentUseCase) pb.CommentServiceServer {
	return &CommentgRPCHandler{commentUsecase: commentUsecase}
}
func (h *CommentgRPCHandler) FindAllByStoryID(ctx context.Context, req *pb.FindAllByStoryIDRequest) (*pb.Comments, error) {
	if req.StoryId < 1 {
		return nil, fmt.Errorf("story id can't below one")
	}

	comments, err := h.commentUsecase.FindByStoryId(ctx, req.StoryId)
	if err != nil {
		return nil, fmt.Errorf("failed to get data %w", err)
	}
	data := ConvertModeltoProto(comments)
	response := pb.Comments{
		Comments: data,
	}
	return &response, nil
}
func ConvertModeltoProto(data []*model.Comment) []*pb.Comment {
	var protoComments []*pb.Comment
	for _, comment := range data {
		protoComments = append(protoComments, &pb.Comment{
			Id:        comment.ID,
			Comment:   comment.Comment,
			CreatedAt: timestamppb.New(comment.Created_at),
			UpdatedAt: timestamppb.New(*comment.Updated_at),
		})
	}

	return protoComments
}
