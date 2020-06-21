# Dominos

CLI for playing a games of Dominos 

Specify number of players as cmd argument, e.g.
./dominos --players=4

or use default of 2 players

First enter name of each player
Pick first 7 dominos anonymously from a table of 28

Each player takes turn to either play one of their dominos or pick up another domino from the table.
When playing a domino, the player must first pick the domino, using the ID and then specify which end of the board to play it.
The current board is displayed as a list of paired numbers, see example.

    Current board:
    -------------
    <-[0|0]-[0|4]-[4|4]-[4|6]-[6|6]-[6|3]-[3|3]-[3|2]-[2|2]-[2|4]->
  
    You can play either: 0 or 4
  
    Player Bob
    Dominos: 1: [5|5] 2: [2|6] 3: [0|5] 
  
    Pick a option
    1: Play
    2: Pick up
  
    Player Bob
    Play a domino
    -> 3
    Play at either the left(1) or right(2) or 0 to return
    -> 1
