// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetTicketCustomizedField
//
// 该接口用于获取工单自定义字段详情。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket_customized_field/get-ticket-customized-field
func (r *HelpdeskService) GetTicketCustomizedField(ctx context.Context, request *GetTicketCustomizedFieldReq, options ...MethodOptionFunc) (*GetTicketCustomizedFieldResp, *Response, error) {
	if r.cli.mock.mockHelpdeskGetTicketCustomizedField != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#GetTicketCustomizedField mock enable")
		return r.cli.mock.mockHelpdeskGetTicketCustomizedField(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Helpdesk",
		API:                   "GetTicketCustomizedField",
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/helpdesk/v1/ticket_customized_fields/:ticket_customized_field_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,

		NeedHelpdeskAuth: true,
	}
	resp := new(getTicketCustomizedFieldResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockHelpdeskGetTicketCustomizedField(f func(ctx context.Context, request *GetTicketCustomizedFieldReq, options ...MethodOptionFunc) (*GetTicketCustomizedFieldResp, *Response, error)) {
	r.mockHelpdeskGetTicketCustomizedField = f
}

func (r *Mock) UnMockHelpdeskGetTicketCustomizedField() {
	r.mockHelpdeskGetTicketCustomizedField = nil
}

type GetTicketCustomizedFieldReq struct {
	TicketCustomizedFieldID string `path:"ticket_customized_field_id" json:"-"` // 工单自定义字段ID, 示例值："6948728206392295444"
}

type getTicketCustomizedFieldResp struct {
	Code int64                         `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                        `json:"msg,omitempty"`  // 错误描述
	Data *GetTicketCustomizedFieldResp `json:"data,omitempty"` //
}

type GetTicketCustomizedFieldResp struct {
	TicketCustomizedFieldID string                                 `json:"ticket_customized_field_id,omitempty"` // 工单自定义字段ID
	HelpdeskID              string                                 `json:"helpdesk_id,omitempty"`                // 服务台ID
	KeyName                 string                                 `json:"key_name,omitempty"`                   // 键名
	DisplayName             string                                 `json:"display_name,omitempty"`               // 名称
	Position                string                                 `json:"position,omitempty"`                   // 字段在列表后台管理列表中的位置
	FieldType               string                                 `json:"field_type,omitempty"`                 // 类型
	Description             string                                 `json:"description,omitempty"`                // 描述
	Visible                 bool                                   `json:"visible,omitempty"`                    // 是否可见
	Editable                bool                                   `json:"editable,omitempty"`                   // 是否可以修改
	Required                bool                                   `json:"required,omitempty"`                   // 是否必填
	CreatedAt               string                                 `json:"created_at,omitempty"`                 // 创建时间
	UpdatedAt               string                                 `json:"updated_at,omitempty"`                 // 更新时间
	CreatedBy               *GetTicketCustomizedFieldRespCreatedBy `json:"created_by,omitempty"`                 // 创建用户
	UpdatedBy               *GetTicketCustomizedFieldRespUpdatedBy `json:"updated_by,omitempty"`                 // 更新用户
	DropdownAllowMultiple   bool                                   `json:"dropdown_allow_multiple,omitempty"`    // 是否支持多选，仅在字段类型是dropdown的时候有效
}

type GetTicketCustomizedFieldRespCreatedBy struct {
	ID        string `json:"id,omitempty"`         // 用户ID
	AvatarURL string `json:"avatar_url,omitempty"` // 用户头像url
	Name      string `json:"name,omitempty"`       // 用户名
	Email     string `json:"email,omitempty"`      // 用户邮箱
}

type GetTicketCustomizedFieldRespUpdatedBy struct {
	ID              string                  `json:"id,omitempty"`               // 用户ID
	AvatarURL       string                  `json:"avatar_url,omitempty"`       // 用户头像url
	Name            string                  `json:"name,omitempty"`             // 用户名
	Email           string                  `json:"email,omitempty"`            // 用户邮箱
	DropdownOptions *HelpdeskDropdownOption `json:"dropdown_options,omitempty"` // 下拉列表选项
}
