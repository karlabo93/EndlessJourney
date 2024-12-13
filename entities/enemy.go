package entities

import "EndlessJourney/components"

type Enemy struct {
	*Sprite
	FollowsPlayer bool
	CombatComp    *components.EnemyCombat
}
