package models

import "context"

type Participantrecord struct {
	APIVersion *string     `json:"apiVersion,omitempty"`
	Kind       *string     `json:"kind,omitempty"`
	Spec       interface{} `json:"spec,omitempty"` // твій конкретний `Entry` буде десеріалізовано окремо
}

func (p *Participantrecord) ContextValidate(ctx context.Context) error {
	return nil
}
