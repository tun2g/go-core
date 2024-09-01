package constants

type ValidationType int

const (
	BODY ValidationType = iota
	QUERY
	PARAM
)