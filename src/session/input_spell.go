package session

import (
	"strconv"
)

func handleInputSpell(input string, s *Session) bool{
	s.Info.Set(MENU_SPELL(false,nil)...)
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
		//TODO -- most likely a new attribute
		for idx,enemy := range s.Level[s.currLevel].Enemies{
			//Loop enemies
			if(enemy.X == targetX && enemy.Y == targetY){
				enemyStatus := "Damaged"
				item.Interaction(&s.Level[s.currLevel].Enemies[idx].Stats)
				s.Player.Stats.UpdateMana(item.Cost)
				if s.Level[s.currLevel].Enemies[idx].Stats.Health <=0{
					s.Level[s.currLevel].Buffer[enemy.Y][enemy.X].Pop()
					s.Level[s.currLevel].Enemies = append(s.Level[s.currLevel].Enemies[:idx], s.Level[s.currLevel].Enemies[idx+1:]...)
					enemyStatus = "Killed"
				}
				s.Info.Set(MENU_SPELL(true,[]string{enemyStatus,enemy.Name})...)
				
			}
		}
	}else{
		if(remainingMana < 0){
			s.Info.Set(MENU_SPELL(true,nil)...)
		}
	}
	return false
}