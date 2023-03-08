package overrides

// Attributes
// These are used to put properties on a tile object.
// *Most* logic checks against attributes!
// TODO -- honestly this ought to just be a TILE thing, not so much a CORE thing. Same for ICONS

const ATTR_SOLID      = "SOLID;"      // Whether or not a player can walk through
const ATTR_FOREGROUND = "FORGROUND;"  // If this tile should "mask" other tiles. Used mostly for FOG
const ATTR_CLIMBABLE  = "CLIMABLE;"   // UNUSED - will be for ladders

// ---- Item Attribubtes

const ATTR_EQUIPTABLE = "EQUIPTABLE;" // Used to determine if an ITEM is an equiptable object and it's interation should apply on pickup
const ATTR_USABLE     = "USEABLE;"    // Used to determine if an ITEM needs to be picked and used
const ATTR_SPELL      = "SPELL;"      // Used to determine if an ITEM is a spell

// ---- Enemy Attribbutes

const ATTR_FIGHTABLE  = "FIGHTABLE;"  // Whether or not a player can engage in combat
const ATTR_POISONOUS  = "POISONOUS;"  // Will add poison to stats effects. Poison prevents the entity from gaining health from FOG
const ATTR_MANABURN   = "MANABURN;"   // Will add mana burn to stats effects. Mana burn prevents the entity from gaining mana from FOG
const ATTR_BOSS       = "BOSS;"       // Whether or not an enemy is a boss (killing boss should end level!)