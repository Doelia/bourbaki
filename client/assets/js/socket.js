/**
 * Notation "_" préfixée pour les attributs/méthodes privées
 *
 */

function init_socket() {

    var socket = io.connect('http://localhost:2000/socket.io/');

    socket.on('test', function(data) {
        console.log("oui :)")
    	console.log(data);
	});

    socket.emit("chat message", "hey")

}
