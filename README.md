### OVERVIEW ###
This is a game engine which generates ANSI code for a TCP connection over a terminal.
Using "Cells" "Tiles" and "Regions" you can build a screen.

### HOW TO RUN THIS ###
1) go run .
- This starts our server

2) On Mac:
- run stty -icanon && nc localhost 9002

3) On Windows:
- Download PuTTy
- ![image description](docs/imgs/putty_one.png)
- As seen above: Insert these settings (in red); click SAVE
- ![image description](docs/imgs/putty_two.png)
- As seen above: Click Terminal and set "Local line editing" and to "Force Off". Save again.
- Now click OPEN on the main screen again. 
- From now on if you SAVE these settings you'll just have to "Load" than "Open" in the future

### MAIN DEMO ###
- v0.0.1 Initial Commit
- ![image description](docs/imgs/v0.0.1.gif)
- v0.0.3 Better coloring, Now have most "in game" regions defined.
- ![image description](docs/imgs/v0.0.3.gif)
- v0.0.4 We have COOOOMBAT now.
- v0.0.5 Fog exist! Health returned on fog.
-- Stats for players and enemies exist as well
-- examples of how to update screens now implemented
-- switched to CELLS which are just stacks of tiles
- ![image description](docs/imgs/v0.0.5.gif)
- v0.0.6 Items-a-palooza
-- Items are now a thing to pickup and uses
-- this includes potions, equiptment AND spells
-- One spell implemented
-- ton of refactor work
- ![image description](docs/imgs/v0.0.6.gif)
### MAIN STRUCTS ###
- Screen
Think old computers with all their talk of "column count" and you have a good idea on how this will work.
A screen is a collection of CELLS and a string which represents what is sent to the users terminal
- Cells
An array of "TILES". These help determine Zindex on the screen.
We the programmer push "tiles" to represent characters we want to display in a column/line and than compile it into the "RAW" string which is the ANSI code needed to be sent to the user terminal
- Tile
A tile is a representationo of a single Column/Line cooordinate.
#Icon# represents the actual "character" you'll see.
#Color# stores the ANSI code to change the characters color (just foreground for now)
Why?
Because its hard to calculate WHERE on our screen a character will appear when just directly using strings. ANSI codes are just ugly strings which become invisible when submitted to the clientt terminal.
This makes it easier for you; the programmer.
- iRegion
We'll want different sections of the screen; and an easy way to determine what to draw.
Regions are just screen sections.
Thus far I have:
level (aka map)
profile (aka user info area)

### What is DONE ###
1) Can draw to screen
 - with color (foreground background)
 - icons are in solo file
2) Can do input
- nice
3) detect wall collision
4) detect enemy collision
5) Fog exist
6) Code is generic RPG in terms of text but the theme has been decided:
Taking place during the golden age of invention (this is just a fancy way to say steampunk)
a railroad mogul buys up a town with a mine in hopes to uncover it's treasures.
You select a class to take on a mission; attempt to conquer the map and return with riches.
7) Magic and items integrated
8) Map from file parser COMPLETED


### TODO ###
1) Make input it's own class thingy
2) Read from file for player,items, and enemies (Maps done)
3) Do some actual enemy stuff
--- want drops
--- want spawns
4) Intro screen and map selection
5) Add FOG wall detection (cannot see through walls)
7) Start thinking about town hub interactions.