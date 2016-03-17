
var control = {
    accel: 0,
    turn: 0
};

var lastSent = 0.

function send() {
    if (control.accel != lastSent.accel || control.turn != lastSent.turn) {
        conn.send(JSON.stringify(control))
        lastSent = {
            accel: control.accel,
            turn: control.turn
        };
    }

}

function left() {
    control.turn = 1;
    send();
}

function right() {
    control.turn = -1;
    send();
}

function up() {
    control.accel = 1;
    send();
}

function down() {
    control.accel = -1;
    send();
}

// function kewUpEvent(direction) {
//     return JSON.stringify( {
//         event: "keyUp",
//         direction: direction
//     } )
// }
