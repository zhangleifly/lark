// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// SetSheetValueImage 该接口用于根据 spreadsheetToken 和 range 向单个格子写入图片。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDNxYjL1QTM24SN0EjN
func (r *DriveService) SetSheetValueImage(ctx context.Context, request *SetSheetValueImageReq, options ...MethodOptionFunc) (*SetSheetValueImageResp, *Response, error) {
	if r.cli.mock.mockDriveSetSheetValueImage != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#SetSheetValueImage mock enable")
		return r.cli.mock.mockDriveSetSheetValueImage(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "SetSheetValueImage",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values_image",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,

		NeedUserAccessToken: true,
	}
	resp := new(setSheetValueImageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockDriveSetSheetValueImage(f func(ctx context.Context, request *SetSheetValueImageReq, options ...MethodOptionFunc) (*SetSheetValueImageResp, *Response, error)) {
	r.mockDriveSetSheetValueImage = f
}

func (r *Mock) UnMockDriveSetSheetValueImage() {
	r.mockDriveSetSheetValueImage = nil
}

type SetSheetValueImageReq struct {
	SpreadSheetToken string `path:"spreadsheetToken" json:"-"` // spreadsheet的token，获取方式见 [对接前说明](/ssl:ttdoc/ukTMukTMukTM/uczNzUjL3czM14yN3MTN) 的第 4 项
	Range            string `json:"range,omitempty"`           // 查询范围  range=<sheetId>!<开始格子>:<结束格子> 如：xxxx!A1:D5，详见 [对接前说明](/ssl:ttdoc/ukTMukTMukTM/uczNzUjL3czM14yN3MTN) 的第 5 项。此处限定为一个格子，如: xxxx!A1:A1
	Image            []byte `json:"image,omitempty"`           // 需要写入的图片二进制流，支持  "PNG", "JPEG", "JPG", "GIF", "BMP", "JFIF", "EXIF", "TIFF", "BPG", "WEBP", "HEIC" 等图片格式
	Name             string `json:"name,omitempty"`            // 写入的图片名字
}

type setSheetValueImageResp struct {
	Code int64                   `json:"code,omitempty"`
	Msg  string                  `json:"msg,omitempty"`
	Data *SetSheetValueImageResp `json:"data,omitempty"`
}

type SetSheetValueImageResp struct {
	SpreadSheetToken string `json:"spreadsheetToken,omitempty"` // spreadsheet 的 token
}
