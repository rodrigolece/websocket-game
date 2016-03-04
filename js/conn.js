// $(document).ready(function () {
var conn;
var ownId;
// var connected = [];

var log = $("#log");


if (window["WebSocket"]) {
    // Dirección metida a mano abajo
    conn = new WebSocket("ws://127.0.0.1:8080/ws");
    conn.onclose = function(evt) {
        log.append("<div><b>Connection closed.</b></div>")
    }
    conn.onmessage = function(evt) {
        handleEvent(evt);
    }
    conn.onerror = function(evt) {
        log.append($("<div/>").text('Ws.error:', evt));
    };
} else {
    log.html("<div><b>Your browser does not support WebSockets.</b></div>")
}

function handleEvent(evt) {
    var wsEvent = JSON.parse(evt.data);
    // log.append($("<div/>").text(wsEvent.event, wsEvent.id));

    if (wsEvent.event == "identity" && ownId == undefined) {
        ownId = wsEvent.id
        log.append($("<div/>").text("Succesfully registered id: " + ownId))
    }
    if (wsEvent.event == "createPlayer") {
        log.append($("<div/>").text("Created player: " + wsEvent.id));
    }
}

// });
