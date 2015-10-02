/**
 * Notation "_" préfixée pour les attributs/méthodes privées
 *
 */

 var socket

function init_socket() {

    socket = io();

    socket.on('chat', function(data) {
        console.log(data);
    });

    socket.emit("chat", "hey");


}
