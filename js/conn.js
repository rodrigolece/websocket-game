
// $(document).ready(function () {

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

    if (wsEvent.event == "identity" && ownId == undefined) {
        ownId = wsEvent.id
        log.append($("<div/>").text("Succesfully registered id: " + ownId))
    }
    if (wsEvent.event == "createPlayer") {
        var part = new Particle(wsEvent.pos, wsEvent.vel);
        gas[wsEvent.id] = part;
        numParticles++ ;

        log.append($("<div/>").text("Created player: " + wsEvent.id));
        requestAnimationFrame(animate);
    }
    // if (wsEvent.event == "destroyPlayer") {
    //     connected...
    // }
    if (wsEvent.event == "update") {
        var part = gas[wsEvent.id];
        part.pos = wsEvent.pos;
        part.vel = wsEvent.vel;
    }
}

// requestAnimationFrame(animate);

// });
