/**
 * Notation "_" préfixée pour les attributs/méthodes privées
 *
 */

 var socket;

function init_socket() {

    socket = io();

    socket.on('CONNECTACCEPT', function(data) {
        var code = data[0]; // 0 incorrect, 1 login, 2 inscription
        var numPlayer = data[1];
        console.log("Recv CONNECTACCEPT. code="+code+", numPlayer="+numPlayer);

        if (code == 1 || code == 2) {
            myNum = numPlayer;
            if (code == 2) {
                //$('#register_done').modal('show');
            }
            loadGame();
        } else { // Erreur
            $('.ui.form').removeClass('loading');
            alert("[Message provisoire] Mot de passe incorrect");
        }
    });

    socket.on('DISPLAYLINE', function(data) {
        var line = data[0];
        console.log("Recv DISPLAYLINE. line="+line);
        board.activeLine(line.X, line.Y, line.O, line.N);
    });

    socket.on('DISPLAYSQUARE', function(data) {
        var square = data[0];
        console.log("Recv DISPLAYSQUARE. square="+square);
        board.activeSquare(square.X, square.Y, square.N);
    });

    socket.on('GRID', function(data) {
        board.enableLastLineColoration = false;
        var lines = data[0];
        var squares = data[1];
        for (var l in lines) {
            var line = lines[l];
            board.activeLine(line.X, line.Y, line.O, line.N);
        }
        for (var s in squares) {
            var square = square[s];
            board.activeSquare(square.X, square.Y, square.N);
        }
        board.enableLastLineColoration = true;
    });

    socket.on('UPDATEPLAYERS', function(data) {
        var json = data[0];
        console.log("Recv UPDATEPLAYERS : ");
        console.log(json);
        players.updatePlayers(json);
    });

    socket.on('SETACTIVEPLAYER', function(data) {
        var nPlayer = data[0];
        console.log("Recv: SETACTIVEPLAYER : "+nPlayer);
        onRecvActivePlayer(nPlayer);
    });

    socket.on('PAUSE', function() {
        console.log("Recv pause");
        pause();
    });

    socket.on('UNPAUSE', function() {
        console.log("Recv unpause");
        unpause();
    });

}
