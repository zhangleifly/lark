// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// UpdateDrivePublicPermission 该接口用于根据 filetoken 更新文档的公共设置。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/ukTM3UjL5EzN14SOxcTN
func (r *DriveService) UpdateDrivePublicPermission(ctx context.Context, request *UpdateDrivePublicPermissionReq, options ...MethodOptionFunc) (*UpdateDrivePublicPermissionResp, *Response, error) {
	if r.cli.mock.mockDriveUpdateDrivePublicPermission != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#UpdateDrivePublicPermission mock enable")
		return r.cli.mock.mockDriveUpdateDrivePublicPermission(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "UpdateDrivePublicPermission",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/drive/permission/public/update",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,

		NeedUserAccessToken: true,
	}
	resp := new(updateDrivePublicPermissionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockDriveUpdateDrivePublicPermission(f func(ctx context.Context, request *UpdateDrivePublicPermissionReq, options ...MethodOptionFunc) (*UpdateDrivePublicPermissionResp, *Response, error)) {
	r.mockDriveUpdateDrivePublicPermission = f
}

func (r *Mock) UnMockDriveUpdateDrivePublicPermission() {
	r.mockDriveUpdateDrivePublicPermission = nil
}

type UpdateDrivePublicPermissionReq struct {
	Token                 string  `json:"token,omitempty"`                    // 文件的 token，获取方式见 [对接前说明](/ssl:ttdoc/ukTMukTMukTM/uczNzUjL3czM14yN3MTN)的第 4 项
	Type                  string  `json:"type,omitempty"`                     // 文档类型  "doc"  or  "sheet"
	CopyPrintExportStatus *bool   `json:"copy_print_export_status,omitempty"` // 可创建副本/打印/导出/复制设置（不传则保持原值）：<br>true - 所有可访问此文档的用户<br>false - 有编辑权限的用户
	Comment               *bool   `json:"comment,omitempty"`                  // 可评论设置（不传则保持原值）：<br>true - 所有可访问此文档的用户<br>false - 有编辑权限的用户
	TenantShareable       *bool   `json:"tenant_shareable,omitempty"`         // 租户内用户是否有共享权限（不传则保持原值）
	LinkShareEntity       *string `json:"link_share_entity,omitempty"`        // 链接共享（不传则保持原值）：<br>"tenant_readable" - 组织内获得链接的人可阅读<br>"tenant_editable" - 组织内获得链接的人可编辑<br>"anyone_readable" - 获得链接的任何人可阅读<br>"anyone_editable" - 获得链接的任何人可编辑
	ExternalAccess        *bool   `json:"external_access,omitempty"`          // 是否允许分享到租户外开关（不传则保持原值）
	InviteExternal        *bool   `json:"invite_external,omitempty"`          // 非owner是否允许邀请外部人（不传则保持原值）
}

type updateDrivePublicPermissionResp struct {
	Code int64                            `json:"code,omitempty"`
	Msg  string                           `json:"msg,omitempty"`
	Data *UpdateDrivePublicPermissionResp `json:"data,omitempty"`
}

type UpdateDrivePublicPermissionResp struct {
	IsSuccess bool `json:"is_success,omitempty"` // 是否成功
}
