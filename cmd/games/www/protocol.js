function RoomSocket(loc) {
  var socket = new WebSocket(loc),
    lookup = {};

  this.listen = function(room, cb) {
    lookup[room] = cb;
  };

  this.send = function(room, cb) {
    if (lookup.hasOwnProperty(room))
      return socket.send(room + cb);
    throw 'Room not found: ' + room;
  };

  socket.onmessage = function(e) {
    var room = e.data[0], body = e.data.slice(1);
    if (lookup.hasOwnProperty(room)) return lookup[room](body);
    return console.log('Unknonwn Message', e.data);
  };

  var that = this;
  socket.onclose = function(e) {
    if (that.hasOwnProperty('onclose')) that.onclose(e);
  };
}
