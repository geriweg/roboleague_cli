package robocode

import "os/user"

type Battle struct {
	season string
	mode string
	bots []Bot
}

func PrepareBattleFile(fs fileSystem, battle *Battle) error {
	_, err := fs.Create(battle.season)
	if err != nil {
		return err
	}

	return nil
}


