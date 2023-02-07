package session

import (
	"strconv"
)

func handleInputSpell(input string, s *Session){
	s.Info.Set(
		"Currently Casting [Spell]",
		"Press [wasd] to cast in that direction [x] to cast on SELF",
		"Press [i] to return to inventory",
		"Press [m] to return to moving",
	)
	idx,_         := strconv.Atoi(s.Profile.SelectedItem)
	//TODO -- error handling
	idx            = idx - 1
	item          := s.Player.Items[idx]
	targetX       := s.Player.X
	targetY       := s.Player.Y
	castPressed   := false
	remainingMana := s.Player.Stats.Mana + item.Cost
	switch input {
		case "x":{
			idx,_ := strconv.Atoi(s.Profile.SelectedItem)
			idx = idx - 1
			item := s.Player.Items[idx]
			item.Interaction(&s.Player.Stats)
		}
		case "w":{
			targetY -= 1
			castPressed = true	
		}
		case "s":{
			targetY += 1
			castPressed = true	
		}
		case "a":{
			targetX -= 1
			castPressed = true	
		}
		case "d":{
			targetX += 1
			castPressed = true	
		}
	}
	if(castPressed && remainingMana >= 0){
		//TODO -- to make more interesting spells we'll have to update this :/
		for idx,enemy := range s.Enemies{
			//Loop enemies
			if(enemy.X == targetX && enemy.Y == targetY){
				enemyStatus := "Damaged"
				item.Interaction(&s.Enemies[idx].Stats)
				s.Player.Stats.UpdateMana(item.Cost)
				if s.Enemies[idx].Stats.Health <=0{
					s.Screen.Buffer[enemy.Y][enemy.X].Pop()
					s.Enemies = append(s.Enemies[:idx], s.Enemies[idx+1:]...)
					enemyStatus = "Killed"
				}
				s.Info.Set(
					"You ["+enemyStatus+"] enemy ["+enemy.Name+"]. Currently Casting [Spell]",
					"Press [wasd] to cast in that direction [x] to cast on SELF",
					"Press [i] to return to inventory",
					"Press [m] to return to moving",
				)
			}
		}
	}else{
		if(remainingMana < 0){
			s.Info.Set(
				"OOPS! Not enough mana! Currently Casting [Spell]",
				"Press [wasd] to cast in that direction [x] to cast on SELF",
				"Press [i] to return to inventory",
				"Press [m] to return to moving",
			)
		}
	}
	s.Info.Refresh()
}