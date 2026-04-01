package models

import (
	"encoding/json"
	"strconv"

	"github.com/rprakashg/cac/internal/domain"
	"github.com/rprakashg/cac/internal/errors"
	"github.com/samber/lo"
)

type Policy struct {
	Resource

	// The desired state, stored as opaque JSON object.
	Spec *JSONField[domain.PolicySpec] `gorm:"type:jsonb"`

	// The last reported state, stored as opaque JSON object.
	Status *JSONField[domain.PolicyStatus] `gorm:"type:jsonb"`
}

func (d Policy) String() string {
	val, _ := json.Marshal(d)
	return string(val)
}

func NewPolicyFromApiResource(resource *domain.Policy) (*Policy, error) {
	if resource == nil || resource.Metadata.Name == nil {
		return &Policy{}, nil
	}

	status := domain.PolicyStatus{Conditions: []domain.Condition{}}

	if resource.Status != nil {
		status = *resource.Status
	}
	var resourceVersion *int64
	if resource.Metadata.ResourceVersion != nil {
		i, err := strconv.ParseInt(lo.FromPtr(resource.Metadata.ResourceVersion), 10, 64)
		if err != nil {
			return nil, errors.ErrIllegalResourceVersionFormat
		}
		resourceVersion = &i
	}
	return &Policy{
		Resource: Resource{
			Name:            *resource.Metadata.Name,
			Labels:          lo.FromPtrOr(resource.Metadata.Labels, make(map[string]string)),
			Annotations:     lo.FromPtrOr(resource.Metadata.Annotations, make(map[string]string)),
			Generation:      resource.Metadata.Generation,
			Owner:           resource.Metadata.Owner,
			ResourceVersion: resourceVersion,
		},
		Spec:   MakeJSONField(resource.Spec),
		Status: MakeJSONField(status),
	}, nil
}
