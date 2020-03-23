package api

import (
	"github.com/teambition/gear"
	"github.com/teambition/urbs-setting/src/bll"
	"github.com/teambition/urbs-setting/src/tpl"
)

// Label ..
type Label struct {
	blls *bll.Blls
}

// List ..
func (a *Label) List(ctx *gear.Context) error {
	req := tpl.ProductPaginationURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	res, err := a.blls.Label.List(ctx, req.Product, req.Pagination)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// Create ..
func (a *Label) Create(ctx *gear.Context) error {
	req := tpl.ProductURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}

	body := tpl.LabelBody{}
	if err := ctx.ParseBody(&body); err != nil {
		return err
	}

	res, err := a.blls.Label.Create(ctx, req.Product, body.Name, body.Desc)
	if err != nil {
		return err
	}

	return ctx.OkJSON(res)
}

// Update ..
func (a *Label) Update(ctx *gear.Context) error {
	req := tpl.ProductLabelURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}

	body := tpl.LabelUpdateBody{}
	if err := ctx.ParseBody(&body); err != nil {
		return err
	}

	res, err := a.blls.Label.Update(ctx, req.Product, req.Label, body)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// Offline ..
func (a *Label) Offline(ctx *gear.Context) error {
	req := tpl.ProductLabelURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	res, err := a.blls.Label.Offline(ctx, req.Product, req.Label)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}

// Assign ..
func (a *Label) Assign(ctx *gear.Context) error {
	req := tpl.ProductLabelURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}

	body := tpl.UsersGroupsBody{}
	if err := ctx.ParseBody(&body); err != nil {
		return err
	}

	if err := a.blls.Label.Assign(ctx, req.Product, req.Label, body.Users, body.Groups); err != nil {
		return err
	}
	return ctx.OkJSON(tpl.BoolRes{Result: true})
}

// Delete ..
func (a *Label) Delete(ctx *gear.Context) error {
	req := tpl.ProductLabelURL{}
	if err := ctx.ParseURL(&req); err != nil {
		return err
	}
	res, err := a.blls.Label.Delete(ctx, req.Product, req.Label)
	if err != nil {
		return err
	}
	return ctx.OkJSON(res)
}
