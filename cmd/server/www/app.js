(function(w, d) {
  "use strict";

  var output, game, moves, newSocket;

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

  function chooseMove(slug) {
    newSocket.send('g', slug + '\n');
    moves.innerHTML = '';
    // TODO: hide all other moves
    // TODO: show this move specifically
  }

  function buildMoveButton(move, cls) {
    var li = document.createElement('button');
    li.className = cls;
    li.innerHTML = move.name;
    li.addEventListener('mouseover', shower(move.slug), false);
    li.addEventListener('mouseout', hideer(move.slug), false);
    li.addEventListener('click', function() {
      chooseMove(move.slug);
    });
    return li;
  }

  function buildMoveGroup(title, moves) {
    var group = document.createElement('div'),
      title = document.createElement('h4'),
      text = document.createElement('div');
    group.className = "list-group-item";
    title.className = "list-group-item-heading";
    title.innerHTML = title;
    text.className = "list-group-item-text";
    for (var move of moves) {
      text.appendChild(buildMoveButton(move, 'btn btn-default'));
      text.appendChild(document.createTextNode(' '));
    }
    group.appendChild(title);
    group.appendChild(text);
    return group
  }

  function shower(name) {
    return function() {
      game.querySelector("[data-slug="+name+"]").style.opacity = '1';
    }
  }

  function hideer(name) {
    return function() {
      game.querySelector("[data-slug="+name+"]").style.opacity = '0';
    };
  }

  function setupSVG() {
    var slugs = game.querySelectorAll("[data-slug]"),
      shows = game.querySelectorAll("[data-show]");

    for (var slug of slugs) slug.style.opacity = "0";
    for (var show of shows) {
      show.style.opacity = "0";
      show.style.stroke = "none";
      show.addEventListener("mouseover", shower(show.dataset.show), false);
      show.addEventListener("mouseout", hideer(show.dataset.show), false);
      show.addEventListener("click", function(e) {
        chooseMove(e.target.dataset.show);
      }, false);
    }
  }

  function gameMessage(m) {
    var obj = JSON.parse(m);
    console.log(obj);
    game.innerHTML = obj.svg;
    moves.innerHTML = '';

    // pre-process the svg and add move selection handlers
    setupSVG();

    // moves is unset
    if (!obj.hasOwnProperty('moves')) {
      // TODO: disable moves panel
      return
    }

    // Generate map of moves grouped by move type
    var byType = {};
    for (var m of obj.moves) {
      if (!byType.hasOwnProperty(m.type)) byType[m.type] = [];
      byType[m.type].push(m);
    }

    // State 1: moves are all of the same type
    if (Object.keys(byType).length == 1)
      for (var m of obj.moves) moves.appendChild(buildMoveButton(m, 'list-group-item'));

    // State 2: moves all have various types
    else {
      var keys = Object.keys(byType);
      keys.sort();
      for (var key of keys) moves.appendChild(buildMoveGroup(key, byType[key]));
    }
  }

  // Window on-load event
  w.addEventListener('load', function() {
    var input = document.getElementById("input");
    var loc = document.location.toString().replace("http://", "ws://").replace("https://", "wss://") + '/socket';
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
})(window, document);
