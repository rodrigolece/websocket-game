
addEventListener("keydown", function(evt) {
    var direction = null;
    // Left arrow
    if (evt.keyCode == 37) {
        direction = "left";
        left();
        evt.preventDefault();
    }
    // Right arrow
    if (evt.keyCode == 39) {
        direction = "right";
        right();
        evt.preventDefault();
    }
    // Up arrow
    if (evt.keyCode == 38) {
        direction = "up";
        up();
        evt.preventDefault();
    }
    // Down arrow
    if (evt.keyCode == 40) {
        direction = "down";
        down();
        evt.preventDefault();
    }

    if (direction != null) {
        // conn.send(keyDownEvent(direction));
        part.changeVel(direction)
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
