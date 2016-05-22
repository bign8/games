var output, game, moves, newSocket;

function userMessage(m) {
  var p = document.createElement('p');
  p.innerHTML = m;
  output.appendChild(p);
}
systemMessage = userMessage; // TODO: update in future

function buildMoveButton(move) {
  var li = document.createElement('button');
  li.className = 'list-group-item';
  li.innerHTML = move.Name;
  li.addEventListener('mouseover', function() {
    game.innerHTML = move.SVG;
  }, false);
  li.addEventListener('click', function() {
    console.log('clicking');
    newSocket.send('g', move.Name + '\n');
    moves.innerHTML = '';
    game.innerHTML = move.SVG;
  });
  moves.appendChild(li);
}

function gameMessage(m) {
  var obj = JSON.parse(m);
  console.log(obj);
  game.innerHTML = obj.SVG;
  moves.innerHTML = '';
  for (var i = 0; i < obj.Moves.length; i++) {
    buildMoveButton(obj.Moves[i], obj.SVG)
  }
  moves.addEventListener('mouseout', function() {
    game.innerHTML = obj.SVG;
  }, false);
}

function init() {
  var input = document.getElementById("input");
  var loc = document.location.toString().replace("http://", "ws://") + '/socket';
  newSocket = new RoomSocket(loc);
  newSocket.listen('s', systemMessage);
  newSocket.listen('u', userMessage);
  newSocket.listen('g', gameMessage);
  input.addEventListener("keyup", function(e) {
    if (e.keyCode == 13) {
      var m = input.value;
      input.value = "";
      newSocket.send('u', m + '\n');
      userMessage(m);
    }
  }, false);
  output = document.getElementById("output");
  game = document.getElementById("game");
  moves = document.getElementById("moves");
  newSocket.onclose = systemMessage.bind(this, "Connection Closed.");
}

window.addEventListener("load", init, false);
