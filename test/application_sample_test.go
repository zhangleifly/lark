// Code generated by lark_sdk_gen. DO NOT EDIT.

package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Application_Sample_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppALLPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		cli.Mock().MockGetAppAccessToken(mockGetTenantAccessTokenFailed)
		moduleCli := cli.Application

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.IsApplicationUserAdmin(ctx, &lark.IsApplicationUserAdminReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetApplicationUserAdminScope(ctx, &lark.GetApplicationUserAdminScopeReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetApplicationAppVisibility(ctx, &lark.GetApplicationAppVisibilityReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetApplicationUserVisibleApp(ctx, &lark.GetApplicationUserVisibleAppReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetApplicationAppList(ctx, &lark.GetApplicationAppListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateApplicationAppVisibility(ctx, &lark.UpdateApplicationAppVisibilityReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetApplicationAppAdminUserList(ctx, &lark.GetApplicationAppAdminUserListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CheckUserIsInApplicationPaidScope(ctx, &lark.CheckUserIsInApplicationPaidScopeReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetApplicationOrderList(ctx, &lark.GetApplicationOrderListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetApplicationOrder(ctx, &lark.GetApplicationOrderReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})
	})

	t.Run("request mock failed", func(t *testing.T) {
		cli := AppALLPermission.Ins()
		moduleCli := cli.Application

		t.Run("", func(t *testing.T) {
			cli.Mock().MockApplicationIsApplicationUserAdmin(func(ctx context.Context, request *lark.IsApplicationUserAdminReq, options ...lark.MethodOptionFunc) (*lark.IsApplicationUserAdminResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApplicationIsApplicationUserAdmin()

			_, _, err := moduleCli.IsApplicationUserAdmin(ctx, &lark.IsApplicationUserAdminReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockApplicationGetApplicationUserAdminScope(func(ctx context.Context, request *lark.GetApplicationUserAdminScopeReq, options ...lark.MethodOptionFunc) (*lark.GetApplicationUserAdminScopeResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApplicationGetApplicationUserAdminScope()

			_, _, err := moduleCli.GetApplicationUserAdminScope(ctx, &lark.GetApplicationUserAdminScopeReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockApplicationGetApplicationAppVisibility(func(ctx context.Context, request *lark.GetApplicationAppVisibilityReq, options ...lark.MethodOptionFunc) (*lark.GetApplicationAppVisibilityResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApplicationGetApplicationAppVisibility()

			_, _, err := moduleCli.GetApplicationAppVisibility(ctx, &lark.GetApplicationAppVisibilityReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockApplicationGetApplicationUserVisibleApp(func(ctx context.Context, request *lark.GetApplicationUserVisibleAppReq, options ...lark.MethodOptionFunc) (*lark.GetApplicationUserVisibleAppResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApplicationGetApplicationUserVisibleApp()

			_, _, err := moduleCli.GetApplicationUserVisibleApp(ctx, &lark.GetApplicationUserVisibleAppReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockApplicationGetApplicationAppList(func(ctx context.Context, request *lark.GetApplicationAppListReq, options ...lark.MethodOptionFunc) (*lark.GetApplicationAppListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApplicationGetApplicationAppList()

			_, _, err := moduleCli.GetApplicationAppList(ctx, &lark.GetApplicationAppListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockApplicationUpdateApplicationAppVisibility(func(ctx context.Context, request *lark.UpdateApplicationAppVisibilityReq, options ...lark.MethodOptionFunc) (*lark.UpdateApplicationAppVisibilityResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApplicationUpdateApplicationAppVisibility()

			_, _, err := moduleCli.UpdateApplicationAppVisibility(ctx, &lark.UpdateApplicationAppVisibilityReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockApplicationGetApplicationAppAdminUserList(func(ctx context.Context, request *lark.GetApplicationAppAdminUserListReq, options ...lark.MethodOptionFunc) (*lark.GetApplicationAppAdminUserListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApplicationGetApplicationAppAdminUserList()

			_, _, err := moduleCli.GetApplicationAppAdminUserList(ctx, &lark.GetApplicationAppAdminUserListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockApplicationCheckUserIsInApplicationPaidScope(func(ctx context.Context, request *lark.CheckUserIsInApplicationPaidScopeReq, options ...lark.MethodOptionFunc) (*lark.CheckUserIsInApplicationPaidScopeResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApplicationCheckUserIsInApplicationPaidScope()

			_, _, err := moduleCli.CheckUserIsInApplicationPaidScope(ctx, &lark.CheckUserIsInApplicationPaidScopeReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockApplicationGetApplicationOrderList(func(ctx context.Context, request *lark.GetApplicationOrderListReq, options ...lark.MethodOptionFunc) (*lark.GetApplicationOrderListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApplicationGetApplicationOrderList()

			_, _, err := moduleCli.GetApplicationOrderList(ctx, &lark.GetApplicationOrderListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockApplicationGetApplicationOrder(func(ctx context.Context, request *lark.GetApplicationOrderReq, options ...lark.MethodOptionFunc) (*lark.GetApplicationOrderResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApplicationGetApplicationOrder()

			_, _, err := moduleCli.GetApplicationOrder(ctx, &lark.GetApplicationOrderReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})
	})

	t.Run("response is failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.Application

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.IsApplicationUserAdmin(ctx, &lark.IsApplicationUserAdminReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetApplicationUserAdminScope(ctx, &lark.GetApplicationUserAdminScopeReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetApplicationAppVisibility(ctx, &lark.GetApplicationAppVisibilityReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetApplicationUserVisibleApp(ctx, &lark.GetApplicationUserVisibleAppReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetApplicationAppList(ctx, &lark.GetApplicationAppListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateApplicationAppVisibility(ctx, &lark.UpdateApplicationAppVisibilityReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetApplicationAppAdminUserList(ctx, &lark.GetApplicationAppAdminUserListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CheckUserIsInApplicationPaidScope(ctx, &lark.CheckUserIsInApplicationPaidScopeReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetApplicationOrderList(ctx, &lark.GetApplicationOrderListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetApplicationOrder(ctx, &lark.GetApplicationOrderReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})
	})

	t.Run("fake request is failed", func(t *testing.T) {
		cli := AppALLPermission.Ins()
		moduleCli := cli.Application
		cli.Mock().MockRawRequest(func(ctx context.Context, req *lark.RawRequestReq, resp interface{}) (response *lark.Response, err error) {
			return nil, fmt.Errorf("fake raw request")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.IsApplicationUserAdmin(ctx, &lark.IsApplicationUserAdminReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetApplicationUserAdminScope(ctx, &lark.GetApplicationUserAdminScopeReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetApplicationAppVisibility(ctx, &lark.GetApplicationAppVisibilityReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetApplicationUserVisibleApp(ctx, &lark.GetApplicationUserVisibleAppReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetApplicationAppList(ctx, &lark.GetApplicationAppListReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateApplicationAppVisibility(ctx, &lark.UpdateApplicationAppVisibilityReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetApplicationAppAdminUserList(ctx, &lark.GetApplicationAppAdminUserListReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CheckUserIsInApplicationPaidScope(ctx, &lark.CheckUserIsInApplicationPaidScopeReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetApplicationOrderList(ctx, &lark.GetApplicationOrderListReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetApplicationOrder(ctx, &lark.GetApplicationOrderReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})
	})
}
