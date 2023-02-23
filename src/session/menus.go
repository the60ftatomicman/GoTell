package session

import (
	"example/gotell/src/screen/region"
	"example/gotell/src/tile"
	"strconv"
)

func MENU_SPELL(hasMana bool,attacked []string) []string {
	menu := []string{"Currently Casting [Spell]"}
	if(hasMana){
		menu[0] = menu[0]+": OOPS! Not enough mana! "
	}
	if(attacked != nil){
		menu = append(menu, "You ["+attacked[0]+"] ["+attacked[1]+"].")
	}
	menu = append(menu, "Press [wasd] to cast in that direction")
	menu = append(menu, "Press [x] to cast on SELF")
	return menu
}

func MENU_ITEM(item *tile.Item)[]string{
	menu := []string{
		"["+item.Name+"] selected.",
		"Press (u) to USE",
		"Press (c) to CONVERT",
		"Press (d) to DROP",
	}
	return menu
}

func MENU_MOVING(pickup string)[]string{
	menu := []string{
		"Currently [MOVING]: Press [wasd] to move",
	}
	if pickup != "" {menu = append(menu,"Press [p] to pickup ["+pickup+"]")}
	return menu
}

func MENU_INVENTORY(p *region.Profile,badInput string)[]string {
	menu := []string{
		"Currently viewing [Invetory]",
		"No items currently Selectable",
	}
	if(badInput == "" || badInput == "i"){
		itemCount := len(p.Player.Items)
		if(itemCount > 0){
			end := ""
			if(itemCount > 1){
				end = "-"+strconv.Itoa(itemCount)
			}
			menu[1] = "Press [1"+end+"] to select an item"
		}
	}else {
		menu[1] = "OOPS! ["+badInput+"] is just not valid input!"
		menu[2] = "No items currently Selectable"
	}
	return menu
}