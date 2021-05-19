// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// BatchCreateBitableTable 新增多个数据表
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table/batch_create
func (r *BitableService) BatchCreateBitableTable(ctx context.Context, request *BatchCreateBitableTableReq, options ...MethodOptionFunc) (*BatchCreateBitableTableResp, *Response, error) {
	if r.cli.mock.mockBitableBatchCreateBitableTable != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#BatchCreateBitableTable mock enable")
		return r.cli.mock.mockBitableBatchCreateBitableTable(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:        "Bitable",
		API:          "BatchCreateBitableTable",
		Method:       "POST",
		URL:          "https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/batch_create",
		Body:         request,
		MethodOption: newMethodOption(options),

		NeedUserAccessToken: true,
	}
	resp := new(batchCreateBitableTableResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockBitableBatchCreateBitableTable(f func(ctx context.Context, request *BatchCreateBitableTableReq, options ...MethodOptionFunc) (*BatchCreateBitableTableResp, *Response, error)) {
	r.mockBitableBatchCreateBitableTable = f
}

func (r *Mock) UnMockBitableBatchCreateBitableTable() {
	r.mockBitableBatchCreateBitableTable = nil
}

type BatchCreateBitableTableReq struct {
	UserIDType *IDType                            `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 userid
	AppToken   string                             `path:"app_token" json:"-"`     // bitable app token, 示例值："appbcbWCzen6D8dezhoCH2RpMAh"
	Tables     []*BatchCreateBitableTableReqTable `json:"tables,omitempty"`       // tables
}

type BatchCreateBitableTableReqTable struct {
	Name *string `json:"name,omitempty"` // 数据表 名字, 示例值："table1"
}

type batchCreateBitableTableResp struct {
	Code int64                        `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                       `json:"msg,omitempty"`  // 错误描述
	Data *BatchCreateBitableTableResp `json:"data,omitempty"` //
}

type BatchCreateBitableTableResp struct {
	TableIDs []string `json:"table_ids,omitempty"` // table ids
}
