package tpl

import (
	"github.com/teambition/gear"
	"github.com/teambition/urbs-setting/src/schema"
)

// ProductUpdateBody ...
type ProductUpdateBody struct {
	Desc *string `json:"desc"`
}

// Validate 实现 gear.BodyTemplate。
func (t *ProductUpdateBody) Validate() error {
	if t.Desc == nil {
		return gear.ErrBadRequest.WithMsgf("desc required")
	}

	if len(*t.Desc) > 1022 {
		return gear.ErrBadRequest.WithMsgf("desc too long: %d", len(*t.Desc))
	}
	return nil
}

// ToMap ...
func (t *ProductUpdateBody) ToMap() map[string]interface{} {
	changed := make(map[string]interface{})
	if t.Desc != nil {
		changed["description"] = *t.Desc
	}
	return changed
}

// ProductURL ...
type ProductURL struct {
	Product string `json:"product" param:"product"`
}

// Validate 实现 gear.BodyTemplate。
func (t *ProductURL) Validate() error {
	if !validNameReg.MatchString(t.Product) {
		return gear.ErrBadRequest.WithMsgf("invalid product name: %s", t.Product)
	}
	return nil
}

// UIDProductURL ...
type UIDProductURL struct {
	Pagination
	UID     string `json:"uid" param:"uid"`
	Product string `json:"product" query:"product"`
}

// Validate 实现 gear.BodyTemplate。
func (t *UIDProductURL) Validate() error {
	if !validIDReg.MatchString(t.UID) {
		return gear.ErrBadRequest.WithMsgf("invalid uid: %s", t.UID)
	}
	if !validNameReg.MatchString(t.Product) {
		return gear.ErrBadRequest.WithMsgf("invalid product name: %s", t.Product)
	}

	if err := t.Pagination.Validate(); err != nil {
		return err
	}
	return nil
}

// ProductPaginationURL ...
type ProductPaginationURL struct {
	Pagination
	Product string `json:"product" param:"product"`
}

// Validate 实现 gear.BodyTemplate。
func (t *ProductPaginationURL) Validate() error {
	if !validNameReg.MatchString(t.Product) {
		return gear.ErrBadRequest.WithMsgf("invalid product name: %s", t.Product)
	}
	if err := t.Pagination.Validate(); err != nil {
		return err
	}
	return nil
}

// ProductLabelURL ...
type ProductLabelURL struct {
	ProductPaginationURL
	Label string `json:"label" param:"label"`
}

// Validate 实现 gear.BodyTemplate。
func (t *ProductLabelURL) Validate() error {
	if !validLabelReg.MatchString(t.Label) {
		return gear.ErrBadRequest.WithMsgf("invalid label: %s", t.Label)
	}
	if err := t.ProductPaginationURL.Validate(); err != nil {
		return err
	}
	return nil
}

// ProductModuleURL ...
type ProductModuleURL struct {
	ProductPaginationURL
	Module string `json:"module" param:"module"`
}

// Validate 实现 gear.BodyTemplate。
func (t *ProductModuleURL) Validate() error {
	if !validNameReg.MatchString(t.Module) {
		return gear.ErrBadRequest.WithMsgf("invalid module name: %s", t.Module)
	}
	if err := t.ProductPaginationURL.Validate(); err != nil {
		return err
	}
	return nil
}

// ProductModuleSettingURL ...
type ProductModuleSettingURL struct {
	ProductModuleURL
	Setting string `json:"setting" param:"setting"`
}

// Validate 实现 gear.BodyTemplate。
func (t *ProductModuleSettingURL) Validate() error {
	if !validNameReg.MatchString(t.Setting) {
		return gear.ErrBadRequest.WithMsgf("invalid setting name: %s", t.Setting)
	}
	if err := t.ProductModuleURL.Validate(); err != nil {
		return err
	}
	return nil
}

// ProductsRes ...
type ProductsRes struct {
	SuccessResponseType
	Result []schema.Product `json:"result"`
}

// ProductRes ...
type ProductRes struct {
	SuccessResponseType
	Result schema.Product `json:"result"`
}

// ProductStatistics ...
type ProductStatistics struct {
	Labels   int64 `json:"labels"`
	Modules  int64 `json:"modules"`
	Settings int64 `json:"settings"`
	Release  int64 `json:"release"`
	Status   int64 `json:"status"`
}
