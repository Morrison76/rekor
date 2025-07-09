package models

type Participantrecord struct {
	APIVersion *string     `json:"apiVersion,omitempty"`
	Kind       *string     `json:"kind,omitempty"`
	Spec       interface{} `json:"spec,omitempty"` // твій конкретний `Entry` буде десеріалізовано окремо
}
