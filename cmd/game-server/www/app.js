var output;

function showMessage(m) {
  var p = document.createElement("p");
  p.innerHTML = m.data ? m.data : m;
  output.appendChild(p);
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
  websocket.onmessage = showMessage;
  websocket.onclose = showMessage.bind(this, "Connection Closed.");
}

window.addEventListener("load", init, false);
