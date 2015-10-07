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
        $("#interface").html(result);
        $("#interface").addClass('login');
        init_login();
    }});
}

function loadGame() {
    $.ajax({url: "/content/game.html", success: function(result) {
        $("#interface").html(result);
        $("#interface").removeClass('login');
        $("#interface").addClass('game');
        init_game();
    }});
}

$(document).ready(function() {
    init_socket();
    loadLogin();

    socket.emit("LOGIN", 'Cassoulet', 'mot de passe');
});
