// Objets
var board;
var players;

// Globales
var myNum; // Numéro du joueur client

// Provisoire (pour les tests sans serveur)
var playersDemo = [
    {"numPlayer":"1", "name": "Portrick", "score":"23", isActive: true},
    {"numPlayer":"2", "name": "Faewynn", "score":"178", isActive: true},
    {"numPlayer":"3", "name": "Pancake", "score":"87", isActive: false},
];

function loadLogin() {
    $.ajax({url: "/content/login.html", success: function(result) {
        $("body").html(result);
        $("body").addClass('login')
        init_login();
    }});
}

function loadGame() {
    $.ajax({url: "/content/game.html", success: function(result) {
        $("body").html(result);
        $("body").removeClass('login')
        $("body").addClass('game')
        init_game();
    }});
}

$(document).ready(function() {
    loadLogin();
});
