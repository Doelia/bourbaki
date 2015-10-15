/**
 * Notation "_" préfixée pour les attributs/méthodes privées
 *
 */

var socket;
var stopRecv = false;

function init_socket() {

    socket = io();

    socket.on('CONNECTACCEPT', function(data) {
        stopRecv = false;
        var code = data[0]; // 0 incorrect, 1 login, 2 inscription
        var numPlayer = data[1];
        console.log("Recv CONNECTACCEPT. code="+code+", numPlayer="+numPlayer);

        if (code == 1 || code == 2) {
            myNum = numPlayer;
            if (code == 2) {
                //$('#register_done').modal('show'); // Désactivé pour le dev
            }
            loadGame();
        } else { // Erreur
            $('.ui.form').removeClass('loading');
            if (code == -1) {
                alert("[Message provisoire] Nom de compte invalide. Trop long ? Caractères spéciaux ?");
            } else {
                alert("[Message provisoire] Mot de passe incorrect");
            }
        }
    });

    socket.on('DISPLAYLINE', function(data) {
        if (stopRecv) return;
        var line = data[0];
        console.log("Recv DISPLAYLINE. line="+line);
        board.activeLine(line.X, line.Y, line.O, line.N);
    });

    socket.on('DISPLAYSQUARE', function(data) {
        if (stopRecv) return;
        var square = data[0];
        console.log("Recv DISPLAYSQUARE. square="+square);
        board.activeSquare(square.X, square.Y, square.N);
    });

    socket.on('GRID', function(data) {
        if (stopRecv) return;
        board.enableLastLineColoration = false;
        var lines = data[0];
        var squares = data[1];
        for (var l in lines) {
            var line = lines[l];
            board.activeLine(line.X, line.Y, line.O, line.N);
        }
        for (var s in squares) {
            var square = squares[s];
            board.activeSquare(square.X, square.Y, square.N);
        }
        board.enableLastLineColoration = true;
    });

    socket.on('UPDATEPLAYERS', function(data) {
        if (stopRecv) return;
        var json = data[0];
        console.log("Recv UPDATEPLAYERS : ");
        console.log(json);
        players.updatePlayers(json);
        players.updateActivePlayer();
    });

    socket.on('SETACTIVEPLAYER', function(data) {
        if (stopRecv) return;
        var nPlayer = data[0];
        console.log("Recv: SETACTIVEPLAYER : "+nPlayer);
        onRecvActivePlayer(nPlayer);
    });

    socket.on('ENDGAME', function(data) {

        stopRecv = true;

        var json = data[0];

        $('#endGame tbody').html('');

        console.log(json);
        for (var i in json) {
            var p = json[i];

            var classement = p.Classement;
            var numPlayer = p.NumPlayer;
            var name = p.Name;
            var score = p.Score;

            $('#endGame tbody').append('<tr></tr>');

            var tr = $('#endGame tbody tr:last');

            tr.append('<td>'+classement+'.</td>');
            tr.append('<h4 class="ui image header">'+
                '<div class="ui left floated cbg circular label" num="'+numPlayer+'"></div>'+
                '<div class="content">'+name+'</div>'+
                '</h4>');
            tr.append('<td>'+score+' points</td>');

            if (numPlayer == myNum) {
                // TODO mettre en avant soi même
                $('#myScore').html(score+' points');
                $('#myClassement').html(classement + (classement == 1 ? 'er' : 'ème'));
            }
        }

        $('#replay').click(function(event) {
            socket.emit('GOAGAIN', "");
        });

        $('#endGame')
            .modal('setting', 'closable', false)
            .modal('show');

    });

    socket.on('PAUSE', function() {
        if (stopRecv) return;
        console.log("Recv pause");
        pause();
    });

    socket.on('UNPAUSE', function() {
        if (stopRecv) return;
        console.log("Recv unpause");
        unpause();
    });

    socket.on('LADDER', function(data) {
        var json = data[0];
        console.log('Recv ladder. Json:');
        console.log(json);

        $('#ladder tbody').html('');

        for (var i in json) {
            var p = json[i];

            var classement = p.Classement;
            var name = p.Name;
            var nbrGames = p.NbrGames;
            var nbrWins = p.NbrWins;
            var score = p.Score;

            $('#ladder tbody').append('<tr></tr>');

            var tr = $('#ladder tbody tr:last');

            tr.append('<td># '+classement+'</td>');
            tr.append('<td><strong>'+name+'</strong></td>');
            tr.append('<td>'+nbrGames+' partie jouées</td>');
            tr.append('<td>'+nbrWins+' partie gagnées</td>');
            tr.append('<td><strong>'+score+' points</strong></td>');

            if (name == myName) {
                // TODO mettre en avant soi même
            }
        }

        $('#ladder')
            .modal('show');
    });

	socket.on('disconnect', function() {
		stopAll();
	});

}
