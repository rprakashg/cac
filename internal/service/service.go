package service

import (
	"context"

	"github.com/rprakashg/cac/internal/domain"
)

type Service interface {

	// Policy
	CreatePolicy(ctx context.Context, policy domain.Policy)
	ListPolicies(ctx context.Context, params domain.ListPoliciesParam)
}
