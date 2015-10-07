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
        code = 2;

        if (code == 1 || code == 2) {
            myNum = numPlayer;
            if (code == 2) {
                $('#register_done').modal('show');
            }
            loadGame();
        } else { // Erreur
            $('.ui.form').removeClass('loading');
            alert("[Message provisoire] Mot de passe incorect");
        }
    });

    socket.on('DISPLAYLINE', function(data) {
        var line = data[0];
        console.log("Recv DISPLAYLINE. line="+line);
        board.activeLine(line.X, line.Y, line.O, line.N);
    });

}
