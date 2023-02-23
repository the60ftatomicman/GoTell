package region

import (
	"example/gotell/src/tile"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const MAP_LEFT = 0
const MAP_TOP = 3
const MAP_LINES = 20
const MAP_COLUMNS = 80


// Level
// Where the action takes place.
// This is where we display players, enemies, items, etc.
// You know, the gameplay.
type Level struct {
	Name        string        `default:"Training"`
	Filename    string        `default:"map.txt"`
	Buffer      [][]tile.Tile
	Enemies     []tile.Enemy
	enemySpawns [][]int
	Items       []tile.Item
	itemSpawns  [][]int
}

func (m *Level) Initialize(b [][]tile.Tile) {
	m.Buffer  = initializeBuffer(MAP_LINES, MAP_COLUMNS, b,tile.BLANK)
	m.AssignEnemies(tile.GenerateEnemiesFromFile())
	m.AssignItems(tile.GenerateItemsFromFile())
	//Add fog AFTERWARDS
	for rIdx,row := range m.Buffer {
		for cIdx,column := range row {
			if(column == tile.BLANK){
				m.Buffer[rIdx][cIdx] = tile.FOG
			}
		}
	}
}

func (m *Level) Get() (int, int, int, int, [][]tile.Tile) {
	return MAP_LEFT, MAP_TOP, MAP_LINES, MAP_COLUMNS, m.Buffer
}

func (m *Level) Refresh(){}

func (m *Level) ReadDataFromFile() [][]tile.Tile {
	tiles := [][]tile.Tile{}
	// READ MAP DATA
	fileData := []string{
		"79w",
		"w",
		"w",
		"w",
		"w",
		"w,10b,5si",
		"w",
		"w,2se",
		"5w,2se",
		"5w",
		"5w",
		"10w,20b,20w,20b,10w",
		"10w,20b,20w,20b,10w",
		"10w,20b,20w,20b,10w",
	}
	for r,row := range fileData {
		var nextRow []tile.Tile = fileParser(row)
		for c,nextCell := range nextRow{
			if(nextCell.Name == tile.ENEMY_SPAWN.Name){
				m.enemySpawns = append(m.enemySpawns, []int{r+MAP_TOP,c+MAP_LEFT})
			}
			if(nextCell.Name == tile.ITEM_SPAWN.Name){
				m.itemSpawns = append(m.itemSpawns, []int{r+MAP_TOP,c+MAP_LEFT})
			}
		}
		tiles = append(tiles,nextRow)
	}
	return tiles
}
//
//
//
//
//
var dataConverter = map[string]tile.Tile{
     "w": tile.WALL,
     "b": tile.BLANK,
     "l": tile.LADDER,
	"se": tile.ENEMY_SPAWN,
	"si": tile.ITEM_SPAWN,
}

func fileParser(tileColVals string) []tile.Tile{
	tileStrings := strings.Split(tileColVals, ",")
	tiles := []tile.Tile{}
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
			val, keyExist := dataConverter[strings.ReplaceAll(strTile,strconv.Itoa(numTile),"")]
			if(keyExist){
				tiles = append(tiles,val)
			}else{
				tiles = append(tiles,tile.NULL)
			}
		}
	}
	return tiles
}
//
//
//
//
//TODO simplify this a bit.
func (m *Level) AssignEnemies(enemyList [10][]tile.Enemy) {
	//Always assign a boss
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(m.enemySpawns), func(i, j int) { m.enemySpawns[i], m.enemySpawns[j] = m.enemySpawns[j], m.enemySpawns[i] })
	placedEnemies := []tile.Enemy{}
	//First set BOSS
	enemyList[0][0].X           = m.enemySpawns[0][1]
	enemyList[0][0].Y           = m.enemySpawns[0][0]
	enemyList[0][0].Stats.Level = 10
	placedEnemies               = append(placedEnemies, enemyList[0][0])
	//Now place other enemies
	currentLevelPool  := 1
	enemiesPerlevel   := 1
	spawnsLeft        := len(m.enemySpawns)
	if (spawnsLeft > 10){ 
		enemiesPerlevel = int(math.Round(float64(spawnsLeft) / 10))
	}
	currentCountAtLevel := 0
	for sIdx,spawn := range m.enemySpawns {
		if(sIdx > 0){
			enemy := enemyList[currentLevelPool][rand.Intn(len(enemyList[currentLevelPool]))]
			//Assign enemy XY based on the enemy spawn
			//m.Buffer[spawn[1]][spawn[0]] = tile.BLANK // this is SADLY not working!
			enemy.X = spawn[1]
			enemy.Y = spawn[0]
			enemy.Stats.Level = currentLevelPool
			//Clear both our arrays of the offending value
			placedEnemies = append(placedEnemies, enemy)
			currentCountAtLevel += 1
			if(currentCountAtLevel >= enemiesPerlevel){
				currentCountAtLevel = 0
				currentLevelPool++
			}
		}
	}
	m.Enemies = placedEnemies
}

func (m *Level) AssignItems(itemList []tile.Item) {
	//Always assign a boss
	placedItems := []tile.Item{}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(m.itemSpawns), func(i, j int) { m.itemSpawns[i], m.itemSpawns[j] = m.itemSpawns[j], m.itemSpawns[i] })
	for sIdx,spawn := range m.itemSpawns {
		item := itemList[sIdx]
		item.X = spawn[1]
		item.Y = spawn[0]
		placedItems= append(placedItems, item)
	}
	m.Items = placedItems
}