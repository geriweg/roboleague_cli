package robocode

import "fmt"

var result BattleResult

func RunBattle(season string, mode string, bots []Bot) (BattleResult, error) {
	fmt.Printf("func is called with: %s\n", season)
	err := DirectoryExists(season)
	if err != nil {
		return result, err
	}

	return result, nil
}
