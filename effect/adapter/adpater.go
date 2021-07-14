package adapter

import (
	"github.com/thoohv5/converter/effect"
	"github.com/thoohv5/converter/effect/ent"
)

type Type int

const (
	ModelEffect Type = iota
	SchemaEffect
)

func GetAdapter(t Type) effect.IEffect {
	var iEffect effect.IEffect
	switch t {
	case SchemaEffect:
		iEffect = ent.New()
	default:
		iEffect = effect.New()
	}
	return iEffect
}
