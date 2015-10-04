/**
 * Notation "_" préfixée pour les attributs/méthodes privées
 *
 */

 var socket;

function init_socket() {

    socket = io();

    socket.on('CONNECTACCEPT', function(data) {
        var code = data[0];
        var numPlayer = data[1];
        myNum = numPlayer;
        console.log("Recv CONNECTACCEPT. code="+code+", numPlayer="+numPlayer);
        loadGame();
    });

    socket.on('DISPLAYLINE', function(data) {
        var line = data[0];
        console.log("Recv DISPLAYLINE. line="+line);
        board.activeLine(line.X, line.Y, line.O, line.N);
    });

}
