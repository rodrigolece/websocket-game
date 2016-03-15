
var canvas = document.getElementById("canvas");
// var canvas = $("#canvas");
var ctx = canvas.getContext("2d");
var part;

var radiusParticle = 1/50;
var lx = 1;
var ly = 1;

var Dt = 0.5;
var lastTime = null;

function animate(time) {
    if (lastTime != null) {
        Dt = (time - lastTime) * 0.001;
    }
    lastTime = time;
    ctx.clearRect(0,0,canvas.width,canvas.height)
    part.move(Dt)
    drawPart(part);
    requestAnimationFrame(animate);
}

function Particle(pos, vel, i) {
    this.pos = {x: pos[0], y: pos[1]};
    this.vel = {vx: vel[0], vy: vel[1]};
    this.move = moveParticle;
    // this.index = i;
}
function drawPart(part) {
    ctx.beginPath();
    ctx.arc(canvas.width * part.pos.x, canvas.height * part.pos.y,
        canvas.width*radiusParticle, 0, 2*Math.PI);
    ctx.stroke();
}


function moveParticle(Dt) {
    futurex = this.pos.x + this.vel.vx * Dt;
    futurey = this.pos.y + this.vel.vy * Dt;

    if (futurex + radiusParticle > lx || futurex - radiusParticle < 0) {
        this.vel.vx *= -1;
    }
    if (futurey + radiusParticle > ly || futurey - radiusParticle < 0) {
        this.vel.vy *= -1;
    }
    this.pos = {
        x: this.pos.x + this.vel.vx * Dt,
        y: this.pos.y + this.vel.vy * Dt
    };
}
