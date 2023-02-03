package session

import (
	"example/gotell/src/screen/region"
	"example/gotell/src/tile"
	"math"
	"strconv"
)

func handleInputMoving(input string, p *tile.Player, s *Session) {
	p.PrvY = p.Y
	p.PrvX = p.X
	switch input {
	case "w":
		{
			if p.Y > 1 {
				p.Y -= 1
				p.DirY = -1
				p.DirX = 0
			}
		}
	case "d":
		{
			if p.X < region.MAP_COLUMNS-1 {
				p.X += 1
				p.DirY = 0
				p.DirX = 1
			}
		}
	case "s":
		{
			if p.Y < region.MAP_LINES-1 {
				p.Y += 1
				p.DirY = 1
				p.DirX = 0
			}
		}
	case "a":
		{
			if p.X > 1 {
				p.X -= 1
				p.DirY = 0
				p.DirX = -1
			}
		}
	}
	s.Info.Set("Currently [MOVING]: WASD (moves), switch to (i)nventory, (Q)uit")
	// Test if we are fighting
	enemy_msgs    := []string{""}
	for idx,enemy := range s.Enemies {
		enemyXdelta  := enemy.X - p.X
		enemyYdelta  := enemy.Y - p.Y
		delta        := math.Abs(float64(enemyXdelta)) + math.Abs(float64(enemyYdelta))
		base_enemy_msg := "Enemy ["+enemy.Name+"] Level ["+strconv.Itoa(enemy.Stats.Level)+"] Health ["+strconv.Itoa(enemy.Stats.Health)+"]"
		if (delta < 2) {
			if (delta == 0) {
				//FIGHTING!
				removeEnemy := s.Enemies[idx].Interaction(p)
				s.Profile.Health = strconv.Itoa(p.Stats.Health)
				s.Profile.Level = strconv.Itoa(p.Stats.Level)
				s.Profile.XP = strconv.Itoa(p.Stats.XP)
				enemy_msgs[0] = "ATTACKED: "+base_enemy_msg
				if(removeEnemy){
					enemy_msgs[0] = "DEFEATED ["+enemy.Name+"]"
					s.Screen.Buffer[p.Y][p.X].Pop()
					//TODO -- is this right?
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
				s.Info.Set("Currently [MOVING]: WASD (moves), switch to (i)nventory, (Q)uit")
			}
		}
	}
	// -- check our HEALTH of all things!
	if(p.Stats.Health <= 0){
		s.Info.Set(p.Name+" has died. Try (r)eviving")
		s.State = STATE_DEAD
		p.X = p.PrvX
		p.Y = p.PrvY
	}else{
		// -- We are not dead nor did we fight. movement time
		nextTile := s.Screen.Buffer[p.Y][p.X].Get()
		if preventMovement(&p.Tile, &nextTile) {
			p.X = p.PrvX
			p.Y = p.PrvY
		}
		// -- now do the player placement
		s.Screen.Buffer[p.PrvY][p.PrvX].Pop()
		s.Screen.Buffer[p.Y][p.X].Set(p.Tile);
		// -- remove any fog, loop to see who is nearby
		//TODO -- block based on who is nearby
		xStart,xEnd,xInc,yStart,yEnd,yInc := p.GetViewRanges()
		for c := xStart; c != xEnd; c += xInc {
			for r := yStart; r != yEnd; r+= yInc  {
				removeFog(s,c,r)
			}
		}
	}
	s.Profile.Refresh()
	s.Info.Refresh()
}