var output, game;

function userMessage(m) {
  var p = document.createElement('p');
  p.innerHTML = m;
  output.appendChild(p);
}

systemMessage = userMessage; // TODO: update in future

function gameMessage(m) {
  game.innerHTML = m;
}

function init() {
  var input = document.getElementById("input");
  var loc = document.location.toString().replace("http://", "ws://") + '/socket';
  var newSocket = new RoomSocket(loc);
  newSocket.listen('s', systemMessage);
  newSocket.listen('u', userMessage);
  newSocket.listen('g', gameMessage);
  input.addEventListener("keyup", function(e) {
    if (e.keyCode == 13) {
      var m = input.value;
      input.value = "";
      newSocket.send('u', m);
      userMessage(m);
    }
  }, false);
  output = document.getElementById("output");
  game = document.getElementById("game");
  newSocket.onclose = systemMessage.bind(this, "Connection Closed.");
}

window.addEventListener("load", init, false);
