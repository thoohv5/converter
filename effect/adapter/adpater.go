package adapter

import (
	"github.com/thoohv5/converter/effect"
	"github.com/thoohv5/converter/effect/ent"
	"github.com/thoohv5/converter/effect/md"
)

type Type int

const (
	ModelEffect Type = iota
	SchemaEffect
	MdEffect
)

func GetAdapter(t Type) effect.IEffect {
	var iEffect effect.IEffect
	switch t {
	case MdEffect:
		iEffect = md.New()
	case SchemaEffect:
		iEffect = ent.New()
	default:
		iEffect = effect.New()
	}
	return iEffect
}
