package logic

import (
	"context"

	"github.com/Corrots/gozero/apps/app/api/internal/svc"
	"github.com/Corrots/gozero/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeBannerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeBannerLogic {
	return &HomeBannerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeBannerLogic) HomeBanner() (*types.HomeBannerResponse, error) {
	// todo: add your logic here and delete this line
	banner := &types.Banner{
		ID:   100,
		Name: "ABC",
		URL:  "https://www.baidu.com",
	}
	var banners []*types.Banner
	banners = append(banners, banner)

	resp := &types.HomeBannerResponse{
		Banners: banners,
	}

	return resp, nil
}
