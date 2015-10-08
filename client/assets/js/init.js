// Objets
var board;
var players;

// Globales
var myNum; // Numéro du joueur client

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
    //socket.emit("LOGIN", 'Cassoulet', 'mot de passe');
});
