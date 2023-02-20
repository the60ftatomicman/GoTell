package session

import (
	"example/gotell/src/screen/region"
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
	for _,item := range s.Items {
		if (item.X == s.Player.X && item.Y == s.Player.Y){
			s.Info.Set(MENU_MOVING(item.Name)...)
		}
	}
	// Test if we are fighting
	prev_status   := s.Info.Message[0];
	enemy_msgs    := []string{prev_status}
	for idx,enemy := range s.Enemies {
		enemyXdelta    := enemy.X - s.Player.X
		enemyYdelta    := enemy.Y - s.Player.Y
		delta          := math.Abs(float64(enemyXdelta)) + math.Abs(float64(enemyYdelta))
		base_enemy_msg := "Enemy ["+enemy.Name+"] Level ["+strconv.Itoa(enemy.Stats.Level)+"] Health ["+strconv.Itoa(enemy.Stats.Health)+"] Hits ["+strconv.Itoa(enemy.CalcDefeat(&s.Player.Stats))+"]"

		if (delta < 2) {
			if (delta == 0) {
				//FIGHTING!
				removeEnemy     := s.Enemies[idx].Interaction(&s.Player.Stats)
				enemy_msgs[0]    = "ATTACKED ["+s.Player.GetDirString()+"]: "+base_enemy_msg
				if(removeEnemy){
					enemy_msgs[0] = "DEFEATED ["+enemy.Name+"]"
					s.Screen.Buffer[s.Player.Y][s.Player.X].Pop()
					s.Enemies = append(s.Enemies[:idx], s.Enemies[idx+1:]...)
				}

			}
			if (delta == 1){
				//We are near an enemy in the cardinal dir
				if(enemyXdelta < 0){enemy_msgs = append(enemy_msgs,"WEST: "+base_enemy_msg)}
				if(enemyXdelta > 0){enemy_msgs = append(enemy_msgs,"EAST: "+base_enemy_msg)}
				if(enemyYdelta < 0){enemy_msgs = append(enemy_msgs,"NORTH: "+base_enemy_msg)}
				if(enemyYdelta > 0){enemy_msgs = append(enemy_msgs,"SOUTH: "+base_enemy_msg)}
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
		nextTile := s.Screen.Buffer[s.Player.Y][s.Player.X].Get()
		if preventMovement(&s.Player.Tile, &nextTile) {
			s.Player.X = PrvX
			s.Player.Y = PrvY
		}
		// -- now do the player placement
		s.Screen.Buffer[PrvY][PrvX].Pop()
		s.Screen.Buffer[s.Player.Y][s.Player.X].Set(s.Player.Tile);
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