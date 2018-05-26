package backgammon

import "github.com/bign8/games"

// Game is the fully described version of TTT
var Game = games.Game{
	Name:  "Backgammon",
	Slug:  "backgammon",
	Board: board,
	Start: nil,
	AI:    nil,
}

// Source: https://commons.wikimedia.org/wiki/File:BackgammonBoard.svg
// TODO: slim down the logic/representation verbosity
var board = `
<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewbox="-1 -1 282 242" version="1.1">
  <g>
    <rect id="felt" style="fill:#164623" width="260" height="220" x="10" y="10" />
    <path id="board" style="fill:#906739" d="M 0,0 0,240 280,240 280,0 0,0 z m 10,10 120,0 0,220 -120,0 0,-220 z m 140,0 120,0 0,220 -120,0 0,-220 z" />
    <g transform="matrix(1.0577079,0,0,1.0577079,-3.6212082,-1.5144618)" id="g6785">
      <path style="fill:#dedcae;fill-opacity:1;stroke:#000000;stroke-width:0.2275199" d="M 33.240105,11 42.340901,102.00796 51.441697,11 H 33.240105 z" id="path6787" />
      <path style="fill:#dedcae;fill-opacity:1;stroke:#000000;stroke-width:0.2275199" d="M 69.64329,11 78.744086,102.00796 87.844882,11 H 69.64329 z" id="path6789" />
      <path style="fill:#dedcae;fill-opacity:1;stroke:#000000;stroke-width:0.2275199" d="m 106.04647,11 9.1008,91.00796 L 124.24807,11 h -18.2016 z" id="path6791" />
      <path style="fill:#da251d;fill-opacity:1;stroke:#000000;stroke-width:0.2275199" d="M 15.038512,11 24.139309,102.00796 33.240105,11 H 15.038512 z" id="path6793" />
      <path style="fill:#da251d;fill-opacity:1;stroke:#000000;stroke-width:0.2275199" d="M 51.441697,11 60.542493,102.00796 69.64329,11 H 51.441697 z" id="path6795" />
      <path style="fill:#da251d;fill-opacity:1;stroke:#000000;stroke-width:0.2275199" d="M 87.844882,11 96.945678,102.00796 106.04647,11 H 87.844882 z" id="path6797" />
    </g>
    <g id="g6799" transform="matrix(-1.0577079,0,0,-1.0577079,143.70331,241.51446)">
      <path id="path6801" d="M 33.240105,11 42.340901,102.00796 51.441697,11 H 33.240105 z" style="fill:#dedcae;fill-opacity:1;stroke:#000000;stroke-width:0.2275199" />
      <path id="path6803" d="M 69.64329,11 78.744086,102.00796 87.844882,11 H 69.64329 z" style="fill:#dedcae;fill-opacity:1;stroke:#000000;stroke-width:0.2275199" />
      <path id="path6805" d="m 106.04647,11 9.1008,91.00796 L 124.24807,11 h -18.2016 z" style="fill:#dedcae;fill-opacity:1;stroke:#000000;stroke-width:0.2275199" />
      <path id="path6807" d="M 15.038512,11 24.139309,102.00796 33.240105,11 H 15.038512 z" style="fill:#da251d;fill-opacity:1;stroke:#000000;stroke-width:0.2275199" />
      <path id="path6809" d="M 51.441697,11 60.542493,102.00796 69.64329,11 H 51.441697 z" style="fill:#da251d;fill-opacity:1;stroke:#000000;stroke-width:0.2275199" />
      <path id="path6811" d="M 87.844882,11 96.945678,102.00796 106.04647,11 H 87.844882 z" style="fill:#da251d;fill-opacity:1;stroke:#000000;stroke-width:0.2275199" />
    </g>
    <g id="g6679" transform="matrix(1.0577079,0,0,1.0577079,136.37879,-1.5144618)">
      <path id="path6627" d="M 33.240105,11 42.340901,102.00796 51.441697,11 H 33.240105 z" style="fill:#dedcae;fill-opacity:1;stroke:#000000;stroke-width:0.2275199" />
      <path id="path6625" d="M 69.64329,11 78.744086,102.00796 87.844882,11 H 69.64329 z" style="fill:#dedcae;fill-opacity:1;stroke:#000000;stroke-width:0.2275199" />
      <path id="path6623" d="m 106.04647,11 9.1008,91.00796 L 124.24807,11 h -18.2016 z" style="fill:#dedcae;fill-opacity:1;stroke:#000000;stroke-width:0.2275199" />
      <path id="path6595" d="M 15.038512,11 24.139309,102.00796 33.240105,11 H 15.038512 z" style="fill:#da251d;fill-opacity:1;stroke:#000000;stroke-width:0.2275199" />
      <path id="path6593" d="M 51.441697,11 60.542493,102.00796 69.64329,11 H 51.441697 z" style="fill:#da251d;fill-opacity:1;stroke:#000000;stroke-width:0.2275199" />
      <path id="path6591" d="M 87.844882,11 96.945678,102.00796 106.04647,11 H 87.844882 z" style="fill:#da251d;fill-opacity:1;stroke:#000000;stroke-width:0.2275199" />
    </g>
    <g transform="matrix(-1.0577079,0,0,-1.0577079,283.70331,241.51446)" id="g6720">
      <path style="fill:#dedcae;fill-opacity:1;stroke:#000000;stroke-width:0.2275199" d="M 33.240105,11 42.340901,102.00796 51.441697,11 H 33.240105 z" id="path6722" />
      <path style="fill:#dedcae;fill-opacity:1;stroke:#000000;stroke-width:0.2275199" d="M 69.64329,11 78.744086,102.00796 87.844882,11 H 69.64329 z" id="path6724" />
      <path style="fill:#dedcae;fill-opacity:1;stroke:#000000;stroke-width:0.2275199" d="m 106.04647,11 9.1008,91.00796 L 124.24807,11 h -18.2016 z" id="path6726" />
      <path style="fill:#da251d;fill-opacity:1;stroke:#000000;stroke-width:0.2275199" d="M 15.038512,11 24.139309,102.00796 33.240105,11 H 15.038512 z" id="path6728" />
      <path style="fill:#da251d;fill-opacity:1;stroke:#000000;stroke-width:0.2275199" d="M 51.441697,11 60.542493,102.00796 69.64329,11 H 51.441697 z" id="path6730" />
      <path style="fill:#da251d;fill-opacity:1;stroke:#000000;stroke-width:0.2275199" d="M 87.844882,11 96.945678,102.00796 106.04647,11 H 87.844882 z" id="path6732" />
    </g>
  </g>
</svg>
`
