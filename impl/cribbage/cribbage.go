package cribbage

import "github.com/bign8/games"

// Game is the fully described version of TTT
var Game = games.Game{
	Name:   "Cribbage",
	Slug:   "cribbage",
	Board:  board,
	Counts: []uint8{2, 3},
	Start:  nil,
	AI:     nil,
}

// https://commons.wikimedia.org/wiki/File:Cribbage_Board.svg
// Path Tutorial: https://developer.mozilla.org/en-US/docs/Web/SVG/Tutorial/Paths
// TODO: slim down the logic/representation verbosity
var board = `
<svg version="1.1" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 680 225">
  <rect id="sr" x="600" y="10" width="40" height="20" style="fill: rgb(100%,0%,0%);stroke: black;stroke-width:0.5px;"></rect>
  <rect id="sg" x="600" y="30" width="40" height="20" style="fill: rgb(0%,100%,0%);stroke: black;stroke-width:0.5px;"></rect>
  <rect id="sb" x="600" y="50" width="40" height="20" style="fill: rgb(0%,0%,100%);stroke: black;stroke-width:0.5px;"></rect>

  <path id="br" d="M595 10
                   h-455
                   C  10  10,   5 215, 140 215
                   h 455
                   C 675 215, 675  80, 595  80
                   h-455
                   v  20
                   h 455
                   C 650 100, 650 195, 595 195
                   h-455
                   C  30 195,  30  30, 140  30
                   h 455
                   Z" style="fill: rgb(100%,0%,0%);stroke: black;stroke-width:0.5px;"></path>
  <path id="bg" d="M595 30
                   h-455
                   C  30  30,  30 195, 140 195
                   h 455
                   C 650 195, 650 100, 595 100
                   h-455
                   v  20
                   h 455
                   C 625 120, 625 175, 595 175
                   h-455
                   C  60 175,  60  50, 140  50
                   h 455
                   Z" style="fill: rgb(0%,100%,0%);stroke: black;stroke-width:0.5px;"></path>
  <path id="bb" d="M595 50
                   h-455
                   C  60  50,  60 175, 140 175
                   h 455
                   C 625 175, 625 120, 595 120
                   h-455
                   v  20
                   h 455
                   C 605 140, 605 155, 595 155
                   h-455
                   C  90 155,  90  70, 140 70
                   h 455
                   Z" style="fill: rgb(0%,0%,100%);stroke: black;stroke-width:0.5px;"></path>


  <path id="dividers" d="M595  80 v60
                         M595 155 v60
                         M530  10 v60
                         M530  80 v60
                         M530 155 v60
                         M465  10 v60
                         M465  80 v60
                         M465 155 v60
                         M400  10 v60
                         M400  80 v60
                         M400 155 v60
                         M335  10 v60
                         M335  80 v60
                         M335 155 v60
                         M270  10 v60
                         M270  80 v60
                         M270 155 v60
                         M205  10 v60
                         M205  80 v60
                         M205 155 v60
                         M140  10 v60
                         M140 155 v60
                         M40 113.3858 h60" style="stroke: black; stroke-width: 0.5;"></path>
</svg>
`
