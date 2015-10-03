/**
 * Notation "_" préfixée pour les attributs/méthodes privées
 *
 */

 var socket;

function init_socket() {

    socket = io();

    socket.on('chat', function(data) {
        console.log(data);
    });

    socket.on('CONNECTACCEPT', function(data) {
        console.log(data[0].X);
    });

    socket.emit("chat", "hey");


}
