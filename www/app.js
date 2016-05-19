var output;

function showMessage(m) {
  var p = document.createElement("p");
  p.innerHTML = m.data ? m.data : m;
  output.appendChild(p);
}

function init() {
  var input = document.getElementById("input");
  var websocket = new WebSocket('ws://' + document.location.host + '/api/v0.0.0/socket');
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
