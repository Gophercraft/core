package models

type VisualClass uint8

const (
	VisualPrecast VisualClass = iota
	VisualCast
	VisualImpact
	VisualState
	VisualStateDone
	VisualChannel
)
