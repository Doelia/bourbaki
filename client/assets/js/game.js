
function pause() {
    $('#pause')
        .modal({
            blurring: true,
        })
        .modal('setting', 'closable', false)
        .modal('setting', 'transition', 'horizontal flip')
        .modal('show');
}

function unpause() {
    $('#pause').modal('hide');
}

function isMyTurn() {
    return players.activePlayer == myNum;
}

function onRecvScore(json) {
    players.updatePlayers(json);
    players.updateActivePlayer();
}

function onRecvActivePlayer(numPlayer) {
    players.activePlayer = numPlayer;
    players.updateActivePlayer();
    players.startTimer();
}

function sendAddLine(x,y,o) {
    console.log("sendAddLine x="+x+", y="+y+", o="+o+"");
    socket.emit('PUTLINE', x, y, o);
}

function init_game() {

    board = new Board();
    board.createGrid();

    players = new Players();

    $('#btn-howto').click(function() {
        $('#guide')
            .modal('setting', 'transition', 'vertical flip')
            .modal('show');
    });

    $('.ttip').popup({
        inverted: true,
        position: 'top center',
        duration: 150,
    });

    $('.line.inactive').click(function() {
        if (isMyTurn()) {
            if ($(this).hasClass('inactive')) {
                var x = parseInt($(this).attr('x'));
                var y = parseInt($(this).attr('y'));
                var o = $(this).attr('o');
                sendAddLine(x,y,o==('v')?1:0);
            }
        }
    });

    console.log("send ready");
    socket.emit("READY", "OK");

}
