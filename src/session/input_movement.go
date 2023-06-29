package session

import (
	"example/gotell/src/core/tile"
	overrides "example/gotell/src/core_overrides"
	"example/gotell/src/region"
	"math"
	"strconv"
)

func handleInputMoving(input string,s *Session) bool{
	s.Info.Set(MENU_MOVING("")...)
	PrvY := s.Player.Y
	PrvX := s.Player.X
	switch input {
		case "w":
			{
				if  s.Player.Y > 1 {
					 s.Player.Y -= 1
					 s.Player.DirY = -1
					 s.Player.DirX = 0
				}
			}
		case "d":
			{
				if  s.Player.X < region.MAP_COLUMNS-1 {
					 s.Player.X += 1
					 s.Player.DirY = 0
					 s.Player.DirX = 1
				}
			}
		case "s":
			{
				if  s.Player.Y < region.MAP_LINES-1 {
					 s.Player.Y += 1
					 s.Player.DirY = 1
					 s.Player.DirX = 0
				}
			}
		case "a":
			{
				if  s.Player.X > 1 {
					 s.Player.X -= 1
					 s.Player.DirY = 0
					 s.Player.DirX = -1
				}
			}
	}
	//New potential coordinates chosen.
	// Check Items
	for _,item := range s.Level[s.currLevel].Items {
		if (item.X == s.Player.X && item.Y == s.Player.Y){
			s.Info.Set(MENU_MOVING(item.Name)...)
		}
	}
	// Test if we are fighting
	prev_status   := s.Info.Message[0];
	enemy_msgs    := []string{prev_status}
	for idx,enemy := range s.Level[s.currLevel].Enemies {
		enemyXdelta    := enemy.X - s.Player.X
		enemyYdelta    := enemy.Y - s.Player.Y
		delta          := math.Abs(float64(enemyXdelta)) + math.Abs(float64(enemyYdelta))

		if (delta < 2) {
			pDmg,eDmg := enemy.CalcDefeat(&s.Player.Stats)
			base_enemy_msg := "Enemy ["+enemy.Name+"] Lvl ["+strconv.Itoa(enemy.Stats.Level)+"] " 
			base_enemy_msg += "Hp ["+strconv.Itoa(enemy.Stats.Health)+"/"+strconv.Itoa(enemy.Stats.GetHealthWithMod())+"] "
			base_enemy_msg +="Dmg you ["+strconv.Itoa(eDmg)+"] "
			base_enemy_msg +="Dmg them ["+strconv.Itoa(pDmg)+"]"
			if (delta == 0) {
				//FIGHTING!
				removeEnemy     := s.Level[s.currLevel].Enemies[idx].Interaction(&s.Player.Stats)
				enemy_msgs[0]    = "ATTACKED ["+s.Player.GetDirString()+"]: "+base_enemy_msg
				if(removeEnemy){
					enemy_msgs[0] = "DEFEATED ["+enemy.Name+"]"
					s.Level[s.currLevel].Buffer[s.Player.Y][s.Player.X].Pop()
					s.Level[s.currLevel].Enemies = append(s.Level[s.currLevel].Enemies[:idx], s.Level[s.currLevel].Enemies[idx+1:]...)
				}

			}
			if (delta == 1){
				//We are near an enemy in the cardinal dir
				if(enemyXdelta < 0){enemy_msgs = append(enemy_msgs,"WEST: "+base_enemy_msg)}
				if(enemyXdelta > 0){enemy_msgs = append(enemy_msgs,"EAST: "+base_enemy_msg)}
				if(enemyYdelta < 0){enemy_msgs = append(enemy_msgs,"NORTH: "+base_enemy_msg)}
				if(enemyYdelta > 0){enemy_msgs = append(enemy_msgs,"SOUTH: "+base_enemy_msg)}
				if(tile.CheckAttributes(enemy.Tile,overrides.ATTR_BOSS)){
					if(s.Level[s.currLevel].BossMessage != ""){
						s.Popup.Set(s.Level[s.currLevel].BossMessage)
						s.Level[s.currLevel].BossMessage = ""
					}
				}
			}
			if(enemy_msgs[0] != "" || len(enemy_msgs) > 1){
				s.Info.Set(enemy_msgs...)
			}else{
				s.Info.Set(prev_status)
			}
		}
	}
	// -- check our HEALTH of all things!
	if(s.Player.Stats.Health <= 0){
		//TODO -- make this proper!
		s.Info.Set(s.Player.Name+" has died. Try (r)eviving")
		s.State = STATE_DEAD
		s.Player.X = PrvX
		s.Player.Y = PrvY
	}else{
		// -- We are not dead nor did we fight. movement time
		nextTile := s.Level[s.currLevel].Buffer[s.Player.Y][s.Player.X].Get()
		if preventMovement(&s.Player.Tile, &nextTile) {
			s.Player.X = PrvX
			s.Player.Y = PrvY
		}
		// -- now do the player placement
		s.Level[s.currLevel].Buffer[PrvY][PrvX].Pop()
		s.Level[s.currLevel].Buffer[s.Player.Y][s.Player.X].Set(s.Player.Tile);
		// -- remove any fog, loop to see who is nearby
		//TODO -- block based on who is nearby
		xStart,xEnd,xInc,yStart,yEnd,yInc := s.Player.GetViewRanges()
		for c := xStart; c != xEnd; c += xInc {
			for r := yStart; r != yEnd; r+= yInc  {
				removeFog(s,c,r)
			}
		}
	}

	return false
}
//
//
// Helper functions only useful in this state
//
//
func preventMovement(tA *tile.Tile, tB *tile.Tile) bool {
	var prevent bool = false
	if(tile.CheckAttributes(*tA,overrides.ATTR_SOLID) && tile.CheckAttributes(*tB,overrides.ATTR_SOLID)) {
		prevent = true
	}

	return prevent
}