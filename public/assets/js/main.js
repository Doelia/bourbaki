// Objets
var board;
var players;

// Globales
var myNum; // Numéro du joueur client
var myName = "undefined";

function init_all() {
    $('.btn-howto').click(function() {
        $('#guide')
            .modal('setting', 'transition', 'vertical flip')
            .modal('show');
    });

    $('.btn-ladder').click(function() {
        socket.emit("ASKLADDER", "");
    });
}

function loadLogin() {
    $.ajax({url: "/content/login.html", success: function(result) {
        $("#interface").html(result);
        $("#interface").addClass('login');
        init_all();
        init_login();
    }});
}

function loadGame() {
    $.ajax({url: "/content/game.html", success: function(result) {
        $("#interface").html(result);
        $("#interface").removeClass('login');
        $("#interface").addClass('game');
        init_all();
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
    // $('#endGame').modal('show');
});
