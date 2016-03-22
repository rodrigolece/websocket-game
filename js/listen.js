
addEventListener("keydown", function(evt) {
    var direction = null;
    // Left arrow
    if (evt.keyCode == 37) {
        direction = "left";
        left(true);
        evt.preventDefault();
    }
    // Right arrow
    if (evt.keyCode == 39) {
        direction = "right";
        right(true);
        evt.preventDefault();
    }
    // Up arrow
    if (evt.keyCode == 38) {
        direction = "up";
        up(true);
        evt.preventDefault();
    }
    // Down arrow
    if (evt.keyCode == 40) {
        direction = "down";
        down(true);
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

addEventListener("keyup", function(evt) {
    // Left arrow
    if (evt.keyCode == 37) {
        left(false);
        evt.preventDefault();
    }
    // Right arrow
    if (evt.keyCode == 39) {
        right(false);
        evt.preventDefault();
    }
    // Up arrow
    if (evt.keyCode == 38) {
        up(false);
        evt.preventDefault();
    }
    // Down arrow
    if (evt.keyCode == 40) {
        down(false);
        evt.preventDefault();
    }
});
