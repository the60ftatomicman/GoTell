package overrides

import "example/gotell/src/core/tile"

const (
	// ---- Generic Tile Attribubtes
	ATTR_SOLID  tile.Attributes = "SOLID;"     // Whether or not a player can walk through
	ATTR_FOREGROUND             = "FORGROUND;" // If this tile should "mask" other tiles. Used mostly for FOG
	ATTR_CLIMBABLE              = "CLIMABLE;"  // UNUSED - will be for ladders
	// ---- Item Attribubtes
	ATTR_EQUIPTABLE = "EQUIPTABLE;" // Used to determine if an ITEM is an equiptable object and it's interation should apply on pickup
	ATTR_USABLE     = "USEABLE;"    // Used to determine if an ITEM needs to be picked and used
	ATTR_SPELL      = "SPELL;"      // Used to determine if an ITEM is a spell
	// ---- Enemy Attribbutes
	ATTR_FIGHTABLE = "FIGHTABLE;" // Whether or not a player can engage in combat
	ATTR_POISONOUS = "POISONOUS;" // Will add poison to stats effects. Poison prevents the entity from gaining health from FOG
	ATTR_MANABURN  = "MANABURN;"  // Will add mana burn to stats effects. Mana burn prevents the entity from gaining mana from FOG
	ATTR_BOSS      = "BOSS;"      // Whether or not an enemy is a boss (killing boss should end level!)
)