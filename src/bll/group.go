package bll

import (
	"context"

	"github.com/teambition/urbs-setting/src/model"
	"github.com/teambition/urbs-setting/src/tpl"
)

// Group ...
type Group struct {
	ms *model.Models
}

// List 返回群组列表
func (b *Group) List(ctx context.Context, kind string, pg tpl.Pagination) (*tpl.GroupsRes, error) {
	groups, total, err := b.ms.Group.Find(context.WithValue(ctx, model.ReadDB, true), kind, pg)
	if err != nil {
		return nil, err
	}
	res := &tpl.GroupsRes{Result: groups}
	res.TotalSize = total
	if len(res.Result) > pg.PageSize {
		res.NextPageToken = tpl.IDToPageToken(res.Result[pg.PageSize].ID)
		res.Result = res.Result[:pg.PageSize]
	}
	return res, nil
}

// ListLabels ...
func (b *Group) ListLabels(ctx context.Context, kind, uid string, pg tpl.Pagination) (*tpl.MyLabelsRes, error) {
	group, err := b.ms.Group.Acquire(context.WithValue(ctx, model.ReadDB, true), kind, uid)
	if err != nil {
		return nil, err
	}

	labels, total, err := b.ms.Group.FindLabels(ctx, group.ID, pg)
	if err != nil {
		return nil, err
	}

	res := &tpl.MyLabelsRes{Result: labels}
	res.TotalSize = total
	if len(res.Result) > pg.PageSize {
		res.NextPageToken = tpl.IDToPageToken(res.Result[pg.PageSize].ID)
		res.Result = res.Result[:pg.PageSize]
	}
	return res, nil
}

// ListMembers ...
func (b *Group) ListMembers(ctx context.Context, kind, uid string, pg tpl.Pagination) (*tpl.GroupMembersRes, error) {
	group, err := b.ms.Group.Acquire(context.WithValue(ctx, model.ReadDB, true), kind, uid)
	if err != nil {
		return nil, err
	}

	members, total, err := b.ms.Group.FindMembers(ctx, group.ID, pg)
	if err != nil {
		return nil, err
	}

	res := &tpl.GroupMembersRes{Result: members}
	res.TotalSize = total
	if len(res.Result) > pg.PageSize {
		res.NextPageToken = tpl.IDToPageToken(res.Result[pg.PageSize].ID)
		res.Result = res.Result[:pg.PageSize]
	}
	return res, nil
}

// ListSettings ...
func (b *Group) ListSettings(ctx context.Context, req tpl.MySettingsQueryURL) (*tpl.MySettingsRes, error) {
	group, err := b.ms.Group.Acquire(ctx, req.Kind, req.UID)
	if err != nil {
		return nil, err
	}

	readCtx := context.WithValue(ctx, model.ReadDB, true)
	var productID int64
	var moduleID int64
	var settingID int64

	if req.Product != "" {
		productID, err = b.ms.Product.AcquireID(readCtx, req.Product)
		if err != nil {
			return nil, err
		}
	}
	if productID > 0 && req.Module != "" {
		moduleID, err = b.ms.Module.AcquireID(readCtx, productID, req.Module)
		if err != nil {
			return nil, err
		}
	}

	if moduleID > 0 && req.Setting != "" {
		settingID, err = b.ms.Setting.AcquireID(readCtx, moduleID, req.Setting)
		if err != nil {
			return nil, err
		}
	}

	pg := req.Pagination
	settings, total, err := b.ms.Group.FindSettings(readCtx, group.ID, productID, moduleID, settingID, pg, req.Channel, req.Client)
	if err != nil {
		return nil, err
	}

	res := &tpl.MySettingsRes{Result: settings}
	res.TotalSize = total
	if len(res.Result) > pg.PageSize {
		res.NextPageToken = tpl.IDToPageToken(res.Result[pg.PageSize].ID)
		res.Result = res.Result[:pg.PageSize]
	}
	return res, nil
}

// CheckExists ...
func (b *Group) CheckExists(ctx context.Context, kind, uid string) bool {
	group, _ := b.ms.Group.FindByUID(ctx, kind, uid, "id")
	return group != nil
}

// BatchAdd ...
func (b *Group) BatchAdd(ctx context.Context, groups []tpl.GroupBody) error {
	return b.ms.Group.BatchAdd(ctx, groups)
}

// BatchAddMembers 批量给群组添加成员，如果用户未加入系统，则会自动加入
func (b *Group) BatchAddMembers(ctx context.Context, kind, uid string, users []string) error {
	group, err := b.ms.Group.Acquire(ctx, kind, uid)
	if err != nil {
		return err
	}

	if err = b.ms.User.BatchAdd(ctx, users); err != nil {
		return err
	}

	return b.ms.Group.BatchAddMembers(ctx, group, users)
}

// RemoveMembers ...
func (b *Group) RemoveMembers(ctx context.Context, kind, uid, userUID string, syncLt int64) error {
	group, err := b.ms.Group.Acquire(ctx, kind, uid)
	if err != nil {
		return err
	}

	var userID int64
	if userUID != "" {
		if user, _ := b.ms.User.FindByUID(ctx, userUID, "id"); user != nil {
			userID = user.ID
		}
	}

	return b.ms.Group.RemoveMembers(ctx, group.ID, userID, syncLt)
}

// Update ...
func (b *Group) Update(ctx context.Context, kind, uid string, body tpl.GroupUpdateBody) (*tpl.GroupRes, error) {
	group, err := b.ms.Group.Acquire(ctx, kind, uid)
	if err != nil {
		return nil, err
	}
	group, err = b.ms.Group.Update(ctx, group.ID, body.ToMap())
	if err != nil {
		return nil, err
	}
	return &tpl.GroupRes{Result: *group}, nil
}

// Delete ...
func (b *Group) Delete(ctx context.Context, kind, uid string) error {
	group, _ := b.ms.Group.FindByUID(ctx, kind, uid, "id")
	if group == nil {
		return nil
	}
	return b.ms.Group.Delete(ctx, group.ID)
}
