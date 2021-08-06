// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
	"io"
)

// GetACSAccessRecordPhoto
//
// 用户在门禁考勤机上成功开门或打卡后，智能门禁应用都会生成一条门禁记录，对于使用人脸识别方式进行开门的识别记录，还会有抓拍图。
// 可以用该接口下载开门时的人脸识别照片
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/acs-v1/access_record-access_photo/get
func (r *ACSService) GetACSAccessRecordPhoto(ctx context.Context, request *GetACSAccessRecordPhotoReq, options ...MethodOptionFunc) (*GetACSAccessRecordPhotoResp, *Response, error) {
	if r.cli.mock.mockACSGetACSAccessRecordPhoto != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] ACS#GetACSAccessRecordPhoto mock enable")
		return r.cli.mock.mockACSGetACSAccessRecordPhoto(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "ACS",
		API:                   "GetACSAccessRecordPhoto",
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/acs/v1/access_records/:access_record_id/access_photo",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getACSAccessRecordPhotoResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockACSGetACSAccessRecordPhoto(f func(ctx context.Context, request *GetACSAccessRecordPhotoReq, options ...MethodOptionFunc) (*GetACSAccessRecordPhotoResp, *Response, error)) {
	r.mockACSGetACSAccessRecordPhoto = f
}

func (r *Mock) UnMockACSGetACSAccessRecordPhoto() {
	r.mockACSGetACSAccessRecordPhoto = nil
}

type GetACSAccessRecordPhotoReq struct {
	AccessRecordID string `path:"access_record_id" json:"-"` // 门禁访问记录 ID, 示例值："6939433228970082591"
}

type getACSAccessRecordPhotoResp struct {
	IsFile bool                         `json:"is_file,omitempty"`
	Code   int64                        `json:"code,omitempty"`
	Msg    string                       `json:"msg,omitempty"`
	Data   *GetACSAccessRecordPhotoResp `json:"data,omitempty"`
}

func (r *getACSAccessRecordPhotoResp) SetReader(file io.Reader) {
	if r.Data == nil {
		r.Data = &GetACSAccessRecordPhotoResp{}
	}
	r.Data.File = file
}

type GetACSAccessRecordPhotoResp struct {
	File io.Reader `json:"file,omitempty"`
}
