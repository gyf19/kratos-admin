package auth

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"

	authnEngine "github.com/tx7do/kratos-authn/engine"
	authn "github.com/tx7do/kratos-authn/middleware"

	authzEngine "github.com/tx7do/kratos-authz/engine"
	authz "github.com/tx7do/kratos-authz/middleware"

	"github.com/tx7do/go-utils/trans"

	"kratos-admin/pkg/entgo/viewer"
)

var action = authzEngine.Action("ANY")

// Server 衔接认证和权鉴
func Server(opts ...Option) middleware.Middleware {
	op := options{
		setTenantId:   false,
		setOperatorId: true,
	}
	for _, o := range opts {
		o(&op)
	}

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				return nil, ErrWrongContext
			}

			authnClaims, ok := authn.FromContext(ctx)
			if !ok {
				return nil, ErrWrongContext
			}

			tokenPayload, err := NewUserTokenPayloadWithClaims(authnClaims)
			if err != nil {
				return nil, ErrExtractUserInfoFailed
			}

			// 校验访问令牌是否存在
			if op.isExistAccessToken != nil {
				if !op.isExistAccessToken(ctx, tokenPayload.UserId) {
					return nil, ErrAccessTokenExpired
				}
			}

			ctx = viewer.NewContext(ctx, viewer.UserViewer{
				Role:     tokenPayload.GetAuthority(),
				TenantId: trans.Ptr(tokenPayload.TenantId),
			})

			sub, err := authnClaims.GetSubject()
			if err != nil {
				return nil, ErrExtractSubjectFailed
			}

			path := authzEngine.Resource(tr.Operation())

			authzClaims := authzEngine.AuthClaims{
				Subject:  (*authzEngine.Subject)(&sub),
				Action:   &action,
				Resource: &path,
			}

			ctx = authz.NewContext(ctx, &authzClaims)

			return handler(ctx, req)
		}
	}
}

func FromContext(ctx context.Context) (*UserTokenPayload, error) {
	claims, ok := authnEngine.AuthClaimsFromContext(ctx)
	if !ok {
		return nil, ErrMissingJwtToken
	}

	return NewUserTokenPayloadWithClaims(claims)
}
