package robocode

type BattleSpec struct {
	season string
	mode string
}

type Battle interface {
	RunBattle(BattleSpec) (BattleResult, error)
}

