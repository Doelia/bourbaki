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

function stopAll() {
    $('#break')
        .modal({
            blurring: true,
        })
        .modal('setting', 'closable', false)
        .modal('show');
    socket.close();
}

$(document).ready(function() {
    init_socket();
    loadLogin();
    // socket.emit('LOGIN', 'tata', 'tututu');
    // init_game();
});
