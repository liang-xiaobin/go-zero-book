package logic

import (
	"book/service/search/api/internal/svc"
	"book/service/search/api/internal/types"
	"book/service/user/rpc/types/user"
	"context"
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchReq) (*types.SearchReply, error) {
	userIdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId")))
	logx.Infof("userId: %s \n", userIdNumber)
	userId, err := userIdNumber.Int64()
	if err != nil {
		return nil, err
	}

	// 使用user rpc
	reply, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdReq{
		Id: userId,
	})
	if err != nil {
		return nil, err
	}
	logx.Infof("userName: %s \n", reply.Name)
	return &types.SearchReply{}, nil
}
