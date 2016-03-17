
var FRAMES_PER_SECOND = 30;
var FRAME_INTERVAL = 1000/FRAMES_PER_SECOND; // JS usa ms como unidad, no ns

var canvas = document.getElementById("canvas");
// var canvas = $("#canvas");
var ctx = canvas.getContext("2d");
var part;

var radiusParticle = 1/50;
var lx = 1;
var ly = 1;

var Dt = 0.001; // Valor necesario para inicializar
var lastTime = null;

function animate(time) {
    if (lastTime != null) {
        Dt = (time - lastTime) / FRAME_INTERVAL;
    }
    lastTime = time;
    ctx.clearRect(0,0,canvas.width,canvas.height)
    part.move(Dt)
    drawPart(part);
    requestAnimationFrame(animate);
}

function Particle(pos, vel, i) {
    this.pos = pos
    this.vel = vel;
    this.move = moveParticle;
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
