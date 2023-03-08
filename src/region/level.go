package region

import (
	"bufio"
	"example/gotell/src/core/screen"
	"example/gotell/src/core/tile"
	overrides "example/gotell/src/core_overrides"
	"example/gotell/src/object"
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
	Player		*object.Player   // This is part of our remove from session refactor
	Buffer      [][]tile.Cell // Welcome to PAIN COUNTRY! you are trying to port in the foreground logic and object logic (rightfully so) here.
	Enemies     []object.Enemy
	enemySpawns [][]int
	maxEnemies  int
	Items       []object.Item
	itemSpawns  [][]int
	maxItems    int
}

func (m *Level) Initialize(b [][]tile.Tile) {
	initalTiles := screen.InitializeBuffer(MAP_LINES, MAP_COLUMNS, b,tile.BLANK)
	//-- Set all non objects, clear spawns as well
	for rIdx,r := range initalTiles {
		m.Buffer = append(m.Buffer, []tile.Cell{})
		for cIdx,c := range r {
			m.Buffer[rIdx] = append(m.Buffer[rIdx],tile.GenerateNewCell())
			if(c.Name == overrides.ENEMY_SPAWN.Name || c.Name == overrides.ITEM_SPAWN.Name){
				m.Buffer[rIdx][cIdx].Pop()
			}else{
				m.Buffer[rIdx][cIdx].Set(c)
			}
		}
	}
	//-- Place Enemies
	for _,enemy := range m.Enemies {
		m.Buffer[enemy.Y][enemy.X].Set(enemy.Tile)
	}
	//-- Place Items
	for _,item := range m.Items {
		m.Buffer[item.Y][item.X].Set(item.Tile)
	}
	//-- Now place player and the ladder underneath them!
	m.Buffer[m.Player.Y][m.Player.X].Clear()
	m.Buffer[m.Player.Y][m.Player.X].Set(overrides.LADDER)
	m.Buffer[m.Player.Y][m.Player.X].Set(m.Player.Tile)
	//-- NOW fog.
	vision := m.Player.Stats.Vision * 2
	left  := m.Player.X-vision
	if(left < 0){left = 0}
	right := m.Player.X+vision
	if(right > MAP_COLUMNS){right = MAP_COLUMNS}
	up    := m.Player.Y-vision
	if(up < 0){up = 0}
	down  := m.Player.Y+vision
	if(down > MAP_LINES){down = MAP_LINES}

	for rIdx,r := range m.Buffer{
		for cIdx,_ := range r {
			inPlayerSpace := (cIdx > left && cIdx < right && rIdx > up && rIdx < down)
			if(!inPlayerSpace){
					m.Buffer[rIdx][cIdx].Set(overrides.FOG)
			}
		} 
	}
}

func (m *Level) Get() (int, int, int, int, [][]tile.Tile) {
	bufferTop := [][]tile.Tile{}
	for rIdx,r := range m.Buffer{
		bufferTop = append(bufferTop,[]tile.Tile{})
		for _,c := range r {
			bufferTop[rIdx] = append(bufferTop[rIdx],c.Get())
		} 
	}
	return MAP_LEFT, MAP_TOP, MAP_LINES, MAP_COLUMNS, bufferTop
}

func (m *Level) Refresh(){}
//
//
//
//
//
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
			if(nextCell.Name == overrides.ENEMY_SPAWN.Name){
				m.enemySpawns = append(m.enemySpawns, []int{r,c})
			}
			if(nextCell.Name == overrides.ITEM_SPAWN.Name){
				m.itemSpawns = append(m.itemSpawns, []int{r,c})
			}
		}
		tiles = append(tiles,nextRow)
	}
	//TODO -- write a parser class to make this cleaner. Think, we have to also parse player data
	m.parseMetadata(fileData[0])
	m.assignEnemies(object.GenerateEnemiesFromFile(fileData[2]))
	m.assignItems(object.GenerateItemsFromFile(fileData[3]))
	return tiles
}
//
//
//
//
//
var dataConverter = map[string]tile.Tile{
     "w": overrides.WALL,
     "b": tile.BLANK,
     "l": overrides.LADDER,
	"se": overrides.ENEMY_SPAWN,
	"si": overrides.ITEM_SPAWN,
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
func (m *Level) assignEnemies(enemyList [10][]object.Enemy) {
	//Always assign a boss
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(m.enemySpawns), func(i, j int) { m.enemySpawns[i], m.enemySpawns[j] = m.enemySpawns[j], m.enemySpawns[i] })
	//Now Truncate Spawns
	if(len(m.enemySpawns) > m.maxEnemies){
		m.enemySpawns = m.enemySpawns[:m.maxEnemies]
	}
	placedEnemies := []object.Enemy{}
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

func (m *Level) assignItems(itemList []object.Item) {
	//Always assign a boss
	placedItems := []object.Item{}
	rand.Seed(time.Now().UnixNano())
	//Shuffle and Truncate Spawns
	rand.Shuffle(len(m.itemSpawns), func(i, j int) { m.itemSpawns[i], m.itemSpawns[j] = m.itemSpawns[j], m.itemSpawns[i] })
	if(len(m.itemSpawns) > m.maxItems){
		m.itemSpawns = m.itemSpawns[:m.maxItems]
	}
	//Shuffle Items
	rand.Shuffle(len(itemList), func(i, j int) { itemList[i], itemList[j] = itemList[j], itemList[i] })
	for sIdx,spawn := range m.itemSpawns {
		if(sIdx < len(itemList)){
			item := itemList[sIdx]
			item.X = spawn[1]
			item.Y = spawn[0]
			placedItems= append(placedItems, item)
		}
	}
	m.Items = placedItems
}