package tile

import (
	"example/gotell/src/core"
	"regexp"
	"strconv"
	"strings"
)

type Tile struct {
	Name              string              `default:"UNASSIGNED"`
	Icon              core.Icons          `default:Icons(ICON_NULL)`
	Color             core.TermCodes      `default:""`
	BGColor           core.TermCodes      `default:""`
	Status            string              `default:"OK"`
	Attribute         string              `default:""`
}

var BLANK = Tile{
	Name: "BLANK",
	Icon: core.Icons(core.ICON_BLANK),
	Color:     core.TermCodes(core.BgBlack),
	BGColor:   core.TermCodes(core.BgBlack),
}
var NULL = Tile{
	Name: "NULL",
	Icon: core.Icons(core.ICON_NULL),
	Color:     core.TermCodes(core.FgYellow),
	BGColor:   core.TermCodes(core.BgMagenta),
}
var WALL = Tile{
	Name:      "WALL",
	Icon:      core.Icons(core.ICON_WALL),
	Color:     core.TermCodes(core.FgBlue),
	BGColor:   core.TermCodes(core.BgBlue),
	Attribute: core.ATTR_SOLID,
}
var PROFILE_V = Tile{
	Name:      "PROFILE_V",
	Icon:      core.Icons(core.ICON_PROFIlE_V),
	Color:     core.TermCodes(core.FgCyan),
	BGColor:   core.TermCodes(core.BgBlack),
	Attribute: core.ATTR_SOLID,
}
var PROFILE_H = Tile{
	Name:      "PROFILE_H",
	Icon:      core.Icons(core.ICON_PROFIlE_H),
	Color:     core.TermCodes(core.FgCyan),
	BGColor:   core.TermCodes(core.BgBlack),
	Attribute: core.ATTR_SOLID,
}
var INFO_V = Tile{
	Name:      "INFO_V",
	Icon:      core.Icons(core.ICON_PROFIlE_V),
	Color:     core.TermCodes(core.FgCyan),
	BGColor:   core.TermCodes(core.BgBlack),
	Attribute: core.ATTR_SOLID,
}
var INFO_H = Tile{
	Name:      "INFO_H",
	Icon:      core.Icons(core.ICON_PROFIlE_H),
	Color:     core.TermCodes(core.FgCyan),
	BGColor:   core.TermCodes(core.BgBlack),
	Attribute: core.ATTR_SOLID,
}
//
var LADDER = Tile{
	Name:  "LADDER",
	Icon:  core.Icons(core.ICON_LADDER),
	Color: core.TermCodes(core.FgYellow),
	Attribute: core.ATTR_CLIMBABLE,
}
var EMPTY = Tile{
	Name:  "EMPTY",
	Icon:  core.Icons(core.ICON_EMPTY),
	Color: core.TermCodes(core.FgWhite),
}
var FOG = Tile{
	Name:  "FOG",
	Icon:  core.Icons(core.ICON_FOG),
	Color: core.TermCodes(core.FgGrey),
	BGColor: core.TermCodes(core.BgGrey),
	Attribute: core.ATTR_FOREGROUND,
}
//
var ENEMY_BASIC = Tile{
	Name:      "ENEMY",
	Icon:      core.ICON_ENEMY,
	Color:     core.TermCodes(core.FgWhite),
	BGColor:   core.TermCodes(core.BgRed),
	Attribute: core.ATTR_FIGHTABLE + core.ATTR_SOLID,
}
var ENEMY_BOSS = Tile{
	Name:      "ENEMY",
	Icon:      core.ICON_BOSS,
	Color:     core.TermCodes(core.FgWhite),
	BGColor:   core.TermCodes(core.BgRed),
	Attribute: core.ATTR_FIGHTABLE + core.ATTR_SOLID + core.ATTR_BOSS,
}

var POTION_HEALTH = Tile{
	Name:      "ITEM",
	Icon:      core.ICON_HEALTH,
	Color:     core.TermCodes(core.FgRed),
	BGColor:   core.TermCodes(core.BgBlack),
	Attribute: core.ATTR_USABLE,
}
var POTION_MANA = Tile{
	Name:      "ITEM",
	Icon:      core.ICON_MANA,
	Color:     core.TermCodes(core.FgCyan),
	BGColor:   core.TermCodes(core.BgBlack),
	Attribute: core.ATTR_USABLE,
}
var EQUIPTMENT= Tile{
	Name:      "ITEM",
	Icon:      core.ICON_EQUIPTMENT,
	Color:     core.TermCodes(core.FgYellow),
	BGColor:   core.TermCodes(core.BgBlack),
	Attribute: core.ATTR_EQUIPTABLE,
}
var SPELL = Tile{
	Name:      "ITEM",
	Icon:      core.ICON_SPELL,
	Color:     core.TermCodes(core.FgMagenta),
	BGColor:   core.TermCodes(core.BgBlack),
	Attribute: core.ATTR_USABLE+core.ATTR_SPELL,
}


func GENERIC_TEXT(character string, colors ...core.TermCodes) Tile{
	bgColor := core.TermCodes(core.BgBlack)
	if(len(colors) > 1){
		bgColor = core.TermCodes(colors[1])
	}
	return Tile{
		Name: "PROFILE_TEXT",
		Icon: core.StringToIcon(character),
		Color: core.TermCodes(colors[0]),
		BGColor: bgColor,
	}
}

func GenerateHorizontalDivider(length int,bookend Tile,fill Tile) []Tile {
	t := []Tile{bookend}
	for i := 0; i < length; i++ {
		t = append(t, fill)
	}
	t = append(t, bookend)
	return t
}

func CheckAttributes(t Tile, attr string) bool{
	return strings.Contains(t.Attribute, attr)
}

var tileConverter = map[string]Tile{
    "w": WALL,
    "b": BLANK,
    "l": LADDER,
}

func FileParser(tileColVals string) []Tile{
	tileStrings := strings.Split(tileColVals, ",")
	tiles := []Tile{}
	for _, strTile := range tileStrings {
		//see if we have a # count
		numTile := 1
		re,regErr := regexp.Compile(`\d{1,}`)
		if(regErr == nil){
			matches := re.FindStringSubmatch(strTile)
			if(len(matches) > 0){
				nt,_ := strconv.Atoi(matches[0])
				numTile = nt
			}
		}
		
		for i:= 0; i < numTile ; i++ {
			val, keyExist := tileConverter[strings.ReplaceAll(strTile,strconv.Itoa(numTile),"")]
			if(keyExist){
				tiles = append(tiles,val)
			}else{
				tiles = append(tiles,NULL)
			}
		}
	}
	return tiles
}