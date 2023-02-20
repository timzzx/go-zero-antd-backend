package logic

import (
	"context"

	"tapi/internal/svc"
	"tapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserListLogic) UserList(req *types.UserListRequest) (resp *types.UserListResponse, err error) {
	u := l.svcCtx.BkModel.User
	r := l.svcCtx.BkModel.Role

	var data []types.User

	if req.Name != "" {

		err = u.WithContext(l.ctx).Where(u.Status.Eq(1)).Where(u.Name.Like(req.Name+"%")).LeftJoin(r, u.RoleID.EqCol(r.ID)).Select(u.ID.As("Id"), u.Name, r.ID.As("RoleId"), r.Name.As("RoleName"), u.Utime, u.Ctime).Debug().Scan(&data)
	} else {
		err = u.WithContext(l.ctx).Where(u.Status.Eq(1)).LeftJoin(r, u.RoleID.EqCol(r.ID)).Select(u.ID.As("Id"), u.Name, r.ID.As("RoleId"), r.Name.As("RoleName"), u.Utime, u.Ctime).Scan(&data)
	}

	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return &types.UserListResponse{
				Code: 500,
				Msg:  err.Error(),
			}, nil
		}
	}
	return &types.UserListResponse{
		Code: 200,
		Msg:  "用户列表",
		Data: data,
	}, nil
}
