// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// CreateDriveImportTask 创建导入任务。支持导入为 doc、sheet、bitable，参考[导入用户指南](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/import_task/import-user-guide)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/import_task/create
func (r *DriveService) CreateDriveImportTask(ctx context.Context, request *CreateDriveImportTaskReq, options ...MethodOptionFunc) (*CreateDriveImportTaskResp, *Response, error) {
	if r.cli.mock.mockDriveCreateDriveImportTask != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#CreateDriveImportTask mock enable")
		return r.cli.mock.mockDriveCreateDriveImportTask(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "CreateDriveImportTask",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/drive/v1/import_tasks",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createDriveImportTaskResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockDriveCreateDriveImportTask(f func(ctx context.Context, request *CreateDriveImportTaskReq, options ...MethodOptionFunc) (*CreateDriveImportTaskResp, *Response, error)) {
	r.mockDriveCreateDriveImportTask = f
}

func (r *Mock) UnMockDriveCreateDriveImportTask() {
	r.mockDriveCreateDriveImportTask = nil
}

type CreateDriveImportTaskReq struct {
	FileExtension string                         `json:"file_extension,omitempty"` // 导入文件格式后缀, 示例值："xlsx"
	FileToken     string                         `json:"file_token,omitempty"`     // 导入文件Drive FileToken, 示例值："boxcnxe5OxxxxxxxSNdsJviENsk"
	Type          string                         `json:"type,omitempty"`           // 导入目标云文档格式, 示例值："sheet"
	FileName      *string                        `json:"file_name,omitempty"`      // 导入目标云文档文件名 ，若为空使用Drive文件名, 示例值："test"
	Point         *CreateDriveImportTaskReqPoint `json:"point,omitempty"`          // 挂载点
}

type CreateDriveImportTaskReqPoint struct {
	MountType int64  `json:"mount_type,omitempty"` // 挂载类型, 示例值：1, 可选值有: `1`：挂载到云空间
	MountKey  string `json:"mount_key,omitempty"`  // 挂载位置,对于mount_type=1, 云空间目录token，空表示根目录, 示例值：""fldxxxxxxxx""
}

type createDriveImportTaskResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *CreateDriveImportTaskResp `json:"data,omitempty"`
}

type CreateDriveImportTaskResp struct {
	Ticket string `json:"ticket,omitempty"` // 导入任务ID
}
