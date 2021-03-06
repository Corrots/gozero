package logic

import (
	"context"

	"github.com/Corrots/gozero/apps/order/admin/internal/svc"
	"github.com/Corrots/gozero/apps/order/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLogic {
	return &AdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminLogic) Admin(req *types.Request) (*types.Response, error) {
	// todo: add your logic here and delete this line
	resp := &types.Response{Message: req.Name}

	return resp, nil
}
