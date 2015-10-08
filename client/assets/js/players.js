/**
 * Classe gèrant la liste des joueurs connectés (ou inacfis) à la parte en cours
 * Notation "_" préfixée pour les attributs/méthodes privées
 */

var Players = function() {

    this.activePlayer = 0;

    this._nbrPlayers = 0; // Nombre de joueurs dans la partie
    this._actualPercent = 100; // Pour réaffichage du timer au refresh des scores

    this._TIMETURN = 15; // Temps en seconde pour faire une action. Se coordoner avec le serveur.

    /**
     * Retourne le numéro du joueur précedent le joueur actif
     */
    this.getLastActivePlayer = function() {
        return (this.activePlayer-1 > 0) ? this.activePlayer-1 : this._nbrPlayers;
    };

    /**
     *  Met à jour totalement la liste des joueurs (scores, lables) à partir d'un json
     *  Executer ensuite updateActivePlayer() pour rééaficher les labels et timers
     */
    this.updatePlayers = function(json) {
        $('#playersList').html('');
        this._nbrPlayers = json.length;
        for (var i in json) {
            var player = json[i];

            $('#playersList')
                .append($('<div class="item player"></div>'));

            $('#playersList > div:last')
                .attr('num', player.NumPlayer)
                .append('<div class="right floated content">')
                .append('<div class="ui left floated cbg circular label" num="'+player.NumPlayer+'"></div>')
                .append('<div class="content"><div class="header">'+player.Name+'</div>'+player.Score+' points</div>');

            if (!player.IsActive) {
                $('#playersList > div:last .right')
                    .append('<div class="ui label gray">Inactif</div>');
            }
        }
    };


    /**
     * Met à jour les labels liés aux joueurs actifs (timers et "points en attente")
     * Gére aussi la classe "myTurn" du board
     * Peut être executée n'importe quand
     */
    this.updateActivePlayer = function() {
        $(".item.player .right .ui.label").not('.gray').remove();
        $(".item.player .right .ui.progress").remove();

        $(".item.player[num='"+this.getLastActivePlayer()+"'] .right")
            .append('<div class="ui label blue">Points en attente...</div>');

        $(".item.player[num='"+this.activePlayer+"'] .right")
            .append('<div class="ui progress green"><div class="bar"></div></div>');

        // Pour mise à jour instantannée du timer, évite le clignotement
        $('.progress').progress({
          percent: this._actualPercent,
        });

        if (this.activePlayer == myNum) {
            $('#board_container').addClass('myTurn');
        } else {
            $('#board_container').removeClass('myTurn');
        }
    };


    /**
     * Déclanche une tâche qui fait descendre le timer (non désigné)
     * A appeler quand il y a un changement de joueur actif
     *
     * Tâche réinitialisée automatiquement quand il y a un changement de joueur
     * grace au parametre _idTimer
     */

    this._idTimer = 0; // Pour stop la tâche timer quand on en lance un nouveau

    this.startTimer = function() {
        var that = this;
        this._idTimer++;
        var timeTurn = this._TIMETURN; // Secondes
        var nSteps = timeTurn*2;

        for (var i = 0; i <= nSteps; i++) {
            var timeStep = (timeTurn / nSteps) * i;
            var percent = 100 - (100 / nSteps * i);
            setTimeout(function(p, myIdTimer) {
                return function() {
                    if (that._idTimer == myIdTimer) {
                        that._actualPercent = p;
                        $('.progress').progress({
                          percent: p,
                          showActivity: false
                        });
                    }
                };
            }(percent, this._idTimer), timeStep*1000);
        }
    };

};
