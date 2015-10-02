/**
 * Notation "_" préfixée pour les attributs/méthodes privées
 *
 */

 var socket

function init_socket() {

    socket = io();

    socket.on('connection', function(socket) {
        console.log('connection OK');
    });

    socket.on('test', function(data) {
        console.log(data);
    });

    socket.on('chat', function(data) {
        console.log(data);
    });

    socket.emit("chat", "hey");


}
