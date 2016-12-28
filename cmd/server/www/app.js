var N8 = N8 || {};  // bign8 global namespace

N8.games = (function(w, d) {
  "use strict";

  var output, game, moves, newSocket, move_set = [];

  var Writer = function(name, cls) {
    this.send = function(m) {
      var p = document.createElement('p');
      p.classList.add('list-group-item');
      if (cls == undefined) {
        p.innerHTML = '<b>' + name + ': </b>';
      } else {
        p.classList.add(cls);
      }
      p.innerHTML += m;
      output.appendChild(p);
      output.scrollTop = output.scrollHeight;
    };
  };

  var systemMessage = new Writer('System', 'list-group-item-info').send;
  var userMessage = new Writer('Me').send;

  function chooseMove(move) {
    newSocket.send('g', move.Name + '\n');
    moves.innerHTML = '';
    game.innerHTML = move.SVG;
  }

  function buildMoveButton(move) {
    var li = document.createElement('button');
    li.className = 'list-group-item';
    li.innerHTML = move.Name;
    li.addEventListener('mouseover', function() {
      game.innerHTML = move.SVG;
    }, false);
    li.addEventListener('click', function() {
      chooseMove(move);
    });
    moves.appendChild(li);
  }

  function gameMessage(m) {
    var obj = JSON.parse(m);
    console.log(obj);
    game.innerHTML = obj.SVG;
    moves.innerHTML = '';
    move_set = obj.Moves;
    for (var i = 0; i < obj.Moves.length; i++) {
      buildMoveButton(obj.Moves[i], obj.SVG)
    }
    moves.addEventListener('mouseout', function() {
      game.innerHTML = obj.SVG;
    }, false);
  }

  // Window on-load event
  w.addEventListener('load', function() {
    var input = document.getElementById("input");
    var loc = document.location.toString().replace("http://", "ws://") + '/socket';
    newSocket = new RoomSocket(loc);
    newSocket.listen('s', systemMessage);
    newSocket.listen('u', new Writer('Opponent').send);
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
  }, false);

  return {
    chooseMove : function(move_string) {
      for (var i = 0; i < move_set.length; i++) {
        if (move_set[i].Name == move_string) {
          return chooseMove(move_set[i]);
        }
      }
      console.log('Move not found:', move_string);
    },
  };
})(window, document);
