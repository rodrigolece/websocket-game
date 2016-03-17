
// gas.js

var FRAMES_PER_SECOND = 30;
var FRAME_INTERVAL = 1000/FRAMES_PER_SECOND; // JS usa ms como unidad, no ns

var canvas = document.getElementById("canvas");
// var canvas = $("#canvas");
var ctx = canvas.getContext("2d");
var part;

var radiusParticle = 1/50;
var lx = 1;
var ly = 1;

var turnAngle = Math.PI/30;
var velFactor = 0.7;

var Dt = 0.001; // Valor necesario para inicializar
var lastTime = null;


// conn.js

var conn;
var ownId;
var connected = [];

var log = $("#log");
