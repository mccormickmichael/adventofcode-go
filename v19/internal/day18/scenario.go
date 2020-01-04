package day18

import "fmt"

type scenarioHandler interface {
	push(*scenario)
	pop() *scenario
	minSteps() int
	setPathSteps(string, int)
}

type scenario struct {
	vault   *Vault
	elf     *Elf
	path    string
	steps   int
}

func (s *scenario) resolve(handler scenarioHandler) {
	if len(s.vault.uncollectedKeys) == 0 {
		handler.setPathSteps(s.path, s.steps)
		return
	}

	s.vault.closeDoors()
	for _, key := range s.elf.keys {
		s.vault.openDoor(key.door)
	}

	goals := newGoals(s.vault.maze, s)
	goals.find()

	if len(goals.found) == 0 {
		// TODO: log something! Can't get to all they keys via this path!

		return
	}

	// edge cases: len(goals) == 0 but len(uncollectedKeys) > 0 :
	//    can't collect any more keys by this scenario chain. Log & return

	for _, goal := range goals.found {

		// don't bother with this scenario if it's already more steps than the best
		if goal.dist + s.steps > handler.minSteps() {
			// TODO: log something
			continue
		}
		scenarioPath := fmt.Sprintf("%s %s", s.path, goal.key.name)
		scenarioElf := &Elf{goal.key.loc, append(s.elf.keys, goal.key)}
		scenarioVault := &Vault{
			maze:            s.vault.maze,
			uncollectedKeys: removeKey(goal.key, s.vault.uncollectedKeys),
			doors:           s.vault.doors,
		}

		handler.push(&scenario{scenarioVault, scenarioElf, scenarioPath, goal.dist + s.steps})
	}


}

func removeKey(key Key, uncollectedKeys []Key) []Key{
	for i, k := range uncollectedKeys {
		if k.name == key.name {
			return append(uncollectedKeys[:i], uncollectedKeys[i+1:]...)
		}
	}
	return uncollectedKeys
}
