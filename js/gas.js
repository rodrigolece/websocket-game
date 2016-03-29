
function animate(time) {
    if (lastTime != null) {
        Dt = (time - lastTime) / FRAME_INTERVAL;
    }
    lastTime = time;
    ctx.clearRect(0,0,canvas.width,canvas.height)
    for (var id in gas) {
        // Importante usar [id] y no .id porque se tiene que evaluar
        var part = gas[id];
        part.move(Dt)
        drawPart(part);
    }
    requestAnimationFrame(animate);
}

function Particle(pos, vel, i) {
    this.pos = pos;
    this.vel = vel;
    this.move = moveParticle;
    this.changeVel = changeVel;
    // this.index = i;
}
function drawPart(part) {
    ctx.beginPath();
    ctx.arc(canvas.width * part.pos[0], canvas.height * part.pos[1],
        canvas.width*radiusParticle, 0, 2*Math.PI);
    ctx.stroke();
}


function moveParticle(Dt) {
    var x = this.pos[0]; var y = this.pos[1];
    futurex = x + this.vel[0] * Dt;
    futurey = y + this.vel[1] * Dt;

    if (futurex + radiusParticle > lx || futurex - radiusParticle < 0) {
        this.vel[0] *= -1;
    }
    if (futurey + radiusParticle > ly || futurey - radiusParticle < 0) {
        this.vel[1] *= -1;
    }
    this.pos = [x + this.vel[0] * Dt, y + this.vel[1] * Dt] ;
}

function changeVel(direction) {
    var angle = null;
    var factor = null;

    if (direction == "left") {
        // Cuidado! JS define los ángulos en sentido de las manecillas del reloj
        angle = -turnAngle;
    }
    if (direction == "right") {
        angle = turnAngle;
    }
    if (direction == "up") {
        factor = 1.02;
    }
    if (direction == "down") {
        factor = 0.98;
    }

    if (angle != null) {
        var c = Math.cos(angle);
        var s = Math.sin(angle);
        this.vel = [ c * this.vel[0] - s * this.vel[1], s * this.vel[0] + c * this.vel[1] ];
    } else if (factor != null) {
        this.vel = [ this.vel.vx * factor, this.vel.vy * factor ]
    }
}
