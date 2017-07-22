package robocode

type BattleRunner interface {
	Run() (BattleResult, error)
}

type Battle struct {
	season string
	mode string
	bots []Bot
}

func (b *Battle)Run() (BattleResult, error) {
	var result BattleResult
	// TODO prepare battle file


	// TODO execute battle
	return result, nil
}

func (b *Battle)PrepareBattleFile(fs fileSystem) error {
	_, err := fs.Create(b.season)
	if err != nil {
		return err
	}

	return nil
}


