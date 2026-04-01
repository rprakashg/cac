package models

import (
	"encoding/json"

	"github.com/rprakashg/cac/internal/domain"
)

type Policy struct {

	// The desired state, stored as opaque JSON object.
	Spec *JSONField[domain.PolicySpec] `gorm:"type:jsonb"`

	// The last reported state, stored as opaque JSON object.
	Status *JSONField[domain.PolicyStatus] `gorm:"type:jsonb"`
}

func (d Policy) String() string {
	val, _ := json.Marshal(d)
	return string(val)
}
