package logic

import (
	"context"
	"database/sql"
	"github.com/j128919965/gopkg/errors"
	"github.com/j128919965/gopkg/stringx"
	"queoj/service/user/internal/model"

	"queoj/service/user/internal/svc"
	"queoj/service/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type ChangePasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	return &ChangePasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChangePasswordLogic) ChangePassword(in *user.ChangePasswordReq) (*user.Result, error) {
	if err := stringx.ValidPassword(in.New)  ;err!=nil {
		return nil , err
	}

	u := model.User{}
	l.svcCtx.Db.Where(&model.User{ID: in.Id}).First(&u)

	code, err := l.svcCtx.Redis.Get(codeEmailPrefix + u.Email)
	if err != nil {
		return nil, err
	}

	if code == "" || code != in.Code {
		return nil, errors.New("验证码无效或已过期", 400)
	}

	newPassword := stringx.Encrypt(l.svcCtx.Config.Salt,in.New)
	u.Password = sql.NullString{String: newPassword,Valid: true}
	l.svcCtx.Db.Save(u)
	return &user.Result{Success: true}, nil
}
