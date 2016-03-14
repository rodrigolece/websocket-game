
function keyDownEvent(direction) {
    return JSON.stringify( {
        event: "keyDown",
        direction: direction
    } )
}

// function kewUpEvent(direction) {
//     return JSON.stringify( {
//         event: "keyUp",
//         direction: direction
//     } )
// }

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
        var j = keyDownEvent(direction)
        log.append($("<div/>").text(JSON.parse(j).direction))
        conn.send(j);
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
