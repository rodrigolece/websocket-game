
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

function left(turn) {
    if (turn) {
        control.turn = 1;
    } else if (control.turn == 1) {
        control.turn = 0;
    }
    send();
}

function right(turn) {
    if (turn) {
        control.turn = -1;
    } else if (control.turn == -1) {
        control.turn = 0;
    }
    send();
}

function up(accel) {
    if (accel) {
        control.accel = 1;
    } else if (control.accel == 1) {
        control.accel = 0;
    }
    send();
}

function down(accel) {
    if (accel) {
        control.accel = -1;
    } else if (control.accel == -1) {
        control.accel = 0;
    }
    send();
}

// function kewUpEvent(direction) {
//     return JSON.stringify( {
//         event: "keyUp",
//         direction: direction
//     } )
// }
