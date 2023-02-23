package region

import (
	"bufio"
	"example/gotell/src/tile"
	"math"
	"math/rand"
	"os"
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
	maxEnemies  int
	Items       []tile.Item
	itemSpawns  [][]int
	maxItems    int
}

func (m *Level) Initialize(b [][]tile.Tile) {
	m.Buffer  = initializeBuffer(MAP_LINES, MAP_COLUMNS, b,tile.BLANK)

	//remove spawns add fog
	for rIdx,row := range m.Buffer {
		for cIdx,column := range row {
			if(column == tile.ENEMY_SPAWN || column == tile.ITEM_SPAWN){
				m.Buffer[rIdx][cIdx] = tile.BLANK
				column = tile.BLANK
			}
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

//TODO -- make this pull ALL 3 things
func (m *Level) getFileRegions()[][]string {
	fileData      := [][]string{}
	currentRegion := -1
	skipLine      := false
	readFile,err := os.Open("./utilities/data/demolevel.txt")
	if(err != nil){
		panic(err)
	}
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
	
    for fileScanner.Scan() {
		fileLine := fileScanner.Text()
		
		//TODO make this a regex
		if(fileLine == "#### METADATA ####" || fileLine == "#### LEVEL ####" || fileLine == "#### ENEMY ####" || fileLine == "#### ITEM ####"){
			skipLine = true
			fileData = append(fileData, []string{})
		}
		if(fileLine == "#### NOTES ####"){
			currentRegion = -1
		}
		if(!skipLine && currentRegion > -1){
			fileData[currentRegion] = append(fileData[currentRegion], fileLine)
		}
		if(skipLine){
			skipLine = false
			currentRegion += 1
		}
		
    }
    readFile.Close()
	return fileData
}

func (m *Level) ReadDataFromFile() [][]tile.Tile {
	tiles := [][]tile.Tile{}
	//Open that data file
	fileData := m.getFileRegions()
	//Assign levels
	for r,row := range fileData[1] {
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
	//TODO -- write a parser class to make this cleaner. Think, we have to also parse player data
	m.parseMetadata(fileData[0])
	m.assignEnemies(tile.GenerateEnemiesFromFile(fileData[2]))
	m.assignItems(tile.GenerateItemsFromFile(fileData[3]))
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
func (m *Level) parseMetadata(metaData []string) {
	for _, md := range metaData {
		keyVal := strings.Split(md, ":")
		switch(keyVal[0]){
			case "max_enemies":{
				val,_ := strconv.Atoi(keyVal[1])
				m.maxEnemies = val
			}
			case "max_items":{
				val,_ := strconv.Atoi(keyVal[1])
				m.maxItems = val
			}
		}
	}
}
//TODO simplify this a bit.
func (m *Level) assignEnemies(enemyList [10][]tile.Enemy) {
	//Always assign a boss
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(m.enemySpawns), func(i, j int) { m.enemySpawns[i], m.enemySpawns[j] = m.enemySpawns[j], m.enemySpawns[i] })
	//Now Truncate Spawns
	if(len(m.enemySpawns) > m.maxEnemies){
		m.enemySpawns = m.enemySpawns[:m.maxEnemies]
	}
	placedEnemies := []tile.Enemy{}
	//First set BOSS
	enemyList[0][0].X           = m.enemySpawns[0][1]
	enemyList[0][0].Y           = m.enemySpawns[0][0]
	enemyList[0][0].Stats.Level = 10
	placedEnemies               = append(placedEnemies, enemyList[0][0])
	//Now place other enemies
	currentLevelPool  := 1
	enemiesPerLevel   := 1
	spawnsLeft        := len(m.enemySpawns)
	if (spawnsLeft > 10){ 
		enemiesPerLevel = int(math.Round(float64(spawnsLeft) / 10))
		if enemiesPerLevel == 1 {
			enemiesPerLevel = 2
		}
	}
	currentCountAtLevel := 0
	for sIdx,spawn := range m.enemySpawns {
		if(sIdx > 0){
			enemy := enemyList[currentLevelPool][rand.Intn(len(enemyList[currentLevelPool]))]
			//Assign enemy XY based on the enemy spawn
			enemy.X = spawn[1]
			enemy.Y = spawn[0]
			enemy.Stats.Level = currentLevelPool
			//Clear both our arrays of the offending value
			placedEnemies = append(placedEnemies, enemy)
			currentCountAtLevel += 1
			if(currentCountAtLevel >= enemiesPerLevel){
				currentCountAtLevel = 0
				currentLevelPool++
			}
		}
	}
	m.Enemies = placedEnemies
}

func (m *Level) assignItems(itemList []tile.Item) {
	//Always assign a boss
	placedItems := []tile.Item{}
	rand.Seed(time.Now().UnixNano())
	//Shuffle and Truncate Spawns
	rand.Shuffle(len(m.itemSpawns), func(i, j int) { m.itemSpawns[i], m.itemSpawns[j] = m.itemSpawns[j], m.itemSpawns[i] })
	if(len(m.itemSpawns) > m.maxItems){
		m.itemSpawns = m.itemSpawns[:m.maxItems]
	}
	//Shuffle Items
	rand.Shuffle(len(itemList), func(i, j int) { itemList[i], itemList[j] = itemList[j], itemList[i] })
	for sIdx,spawn := range m.itemSpawns {
		item := itemList[sIdx]
		item.X = spawn[1]
		item.Y = spawn[0]
		placedItems= append(placedItems, item)
	}
	m.Items = placedItems
}