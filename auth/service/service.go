package service

import (
	"context"
	"time"

	"github.com/micro/go-micro/v2/auth"
	pb "github.com/micro/go-micro/v2/auth/service/proto"
	"github.com/micro/go-micro/v2/client"
)

// NewAuth returns a new instance of the Auth service
func NewAuth(opts ...auth.Option) auth.Auth {
	svc := new(svc)
	svc.Init(opts...)
	return svc
}

// svc is the service implementation of the Auth interface
type svc struct {
	options auth.Options
	auth    pb.AuthService
}

func (s *svc) String() string {
	return "service"
}

func (s *svc) Init(opts ...auth.Option) {
	for _, o := range opts {
		o(&s.options)
	}

	dc := client.DefaultClient
	s.auth = pb.NewAuthService("go.micro.auth", dc)
}

func (s *svc) Options() auth.Options {
	return s.options
}

// Generate a new account
func (s *svc) Generate(id string, opts ...auth.GenerateOption) (*auth.Account, error) {
	options := auth.NewGenerateOptions(opts...)

	rsp, err := s.auth.Generate(context.TODO(), &pb.GenerateRequest{
		Id:           id,
		Roles:        options.Roles,
		Metadata:     options.Metadata,
		SecretExpiry: int64(options.SecretExpiry.Seconds()),
	})
	if err != nil {
		return nil, err
	}

	return serializeAccount(rsp.Account), nil
}

// Grant access to a resource
func (s *svc) Grant(role string, res *auth.Resource) error {
	_, err := s.auth.Grant(context.TODO(), &pb.GrantRequest{
		Role: role,
		Resource: &pb.Resource{
			Type:     res.Type,
			Name:     res.Name,
			Endpoint: res.Endpoint,
		},
	})
	return err
}

// Revoke access to a resource
func (s *svc) Revoke(role string, res *auth.Resource) error {
	_, err := s.auth.Revoke(context.TODO(), &pb.RevokeRequest{
		Role: role,
		Resource: &pb.Resource{
			Type:     res.Type,
			Name:     res.Name,
			Endpoint: res.Endpoint,
		},
	})
	return err
}

// Verify an account has access to a resource
func (s *svc) Verify(acc *auth.Account, res *auth.Resource) error {
	_, err := s.auth.Verify(context.TODO(), &pb.VerifyRequest{
		Account: &pb.Account{
			Id:    acc.ID,
			Roles: acc.Roles,
		},
		Resource: &pb.Resource{
			Type:     res.Type,
			Name:     res.Name,
			Endpoint: res.Endpoint,
		},
	})
	return err
}

// Inspect a token
func (s *svc) Inspect(token string) (*auth.Account, error) {
	rsp, err := s.auth.Inspect(context.TODO(), &pb.InspectRequest{
		Token: token,
	})
	if err != nil {
		return nil, err
	}

	return serializeAccount(rsp.Account), nil
}

// Refresh an account using a secret
func (s *svc) Refresh(secret string, opts ...auth.RefreshOption) (*auth.Token, error) {
	options := auth.NewRefreshOptions(opts...)

	rsp, err := s.auth.Refresh(context.Background(), &pb.RefreshRequest{
		Secret:      secret,
		TokenExpiry: int64(options.TokenExpiry.Seconds()),
	})
	if err != nil {
		return nil, err
	}

	return serializeToken(rsp.Token), nil
}

func serializeToken(t *pb.Token) *auth.Token {
	return &auth.Token{
		Token:    t.Token,
		Type:     t.Type,
		Created:  time.Unix(t.Created, 0),
		Expiry:   time.Unix(t.Expiry, 0),
		Subject:  t.Subject,
		Roles:    t.Roles,
		Metadata: t.Metadata,
	}
}

func serializeAccount(a *pb.Account) *auth.Account {
	var secret *auth.Token
	if a.Secret != nil {
		secret = serializeToken(a.Secret)
	}

	return &auth.Account{
		ID:       a.Id,
		Roles:    a.Roles,
		Metadata: a.Metadata,
		Secret:   secret,
	}
}
