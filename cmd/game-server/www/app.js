var output, game;

function showMessage(m) {
  var load = m.data ? m.data : m;
  if (load.trimLeft().slice(0,4) == '<svg') {
    game.innerHTML = load;
  } else {
    var p = document.createElement("p");
    p.innerHTML = load;
    output.appendChild(p);
  }
}

function init() {
  var input = document.getElementById("input");
  var loc = document.location.toString().replace("http://", "ws://") + '/socket';
  var websocket = new WebSocket(loc);
  input.addEventListener("keyup", function(e) {
    if (e.keyCode == 13) {
      var m = input.value;
      input.value = "";
      websocket.send(m);
      showMessage(m);
    }
  }, false);
  output = document.getElementById("output");
  game = document.getElementById("game");
  websocket.onmessage = showMessage;
  websocket.onclose = showMessage.bind(this, "Connection Closed.");
}

window.addEventListener("load", init, false);
