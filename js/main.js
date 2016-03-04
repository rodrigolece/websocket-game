
var conn;
// var ownID;
// var connected = [];

var log = $("#log");


if (window["WebSocket"]) {
    // Dirección metida a mano abajo
    conn = new WebSocket("ws://127.0.0.1:8080/ws");
    conn.onclose = function(evt) {
        log.html("<div><b>Connection closed.</b></div>")
    }
    conn.onmessage = function(evt) {
        handleEvent(evt);
    }
    conn.onerror = function(evt) {
        log.html($("<div/>").text('Ws.error:', evt));
    };
} else {
    log.html("<div><b>Your browser does not support WebSockets.</b></div>")
}

function handleEvent(evt) {
    var wsEvent = JSON.parse(evt.data);
    log.html($("<div/>").wsEvent);
}

function broadcastEvent(data) {
    if (!conn) {
        return false;
    }
    var j = JSON.stringify({
        Action: "broadcast",
        Data: data
    });
    log.html($("<div/>").j);
//     conn.send(j);
//     return false
}

addEventListener("keydown", function(evt) {
    var direction = null;
    // Left arrow
    if (evt.keyCode == 37) {
        direction = "left";
        evt.preventDefault();
    }
    // Right arrow
    if (evt.keyCode == 39) {
        direction = "right";
        evt.preventDefault();
    }
    // Up arrow
    if (evt.keyCode == 38) {
        direction = "up";
        evt.preventDefault();
    }
    // Down arrow
    if (evt.keyCode == 40) {
        direction = "down";
        evt.preventDefault();
    }

    if (direction != null) {
        var data =  [
            {
                type: "turn",
                content: direction
            }
        ];
        broadcastEvent(data);

        // self guarda el índice de la partícula del cliente
        // gas.particles[self].changeVel(direction);
    }

    /*
    // d key --> add particle
    if (evt.keyCode == 68) {
        gas.addParticle();
    }
    // n key --> new gas
    if (evt.keyCode == 78) {
        gas.addParticle();
    }
    */

});
/* Tal vez esto no se necesita
addEventListener("keyup", function(evt) {
    // Left or right arrow
    if (evt.keyCode == 37 || evt.keyCode == 39) {
        var data =  [{ Type: "stopturn" }];
        broadcastEvent(data);
        evt.preventDefault();
    }
});*/

// $(function() { // equivalente a $(document).ready(function () {...})
