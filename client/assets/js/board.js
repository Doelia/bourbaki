$(document).ready(function() {

    var gridSpace = 56; // Espace entre 2 points

    var _placeE = function(e, x, y) {
        return e
        .css({
            top: x*gridSpace,
            left: y*gridSpace
        })
        .attr('x', x)
        .attr('y', y);
    }

    var _addDot = function(x, y) {
        $('#board').append(
            $('<div class="dot"></dot>')
        );

        _placeE($('#board .dot:last'), x, y);
    }

    var _createInactiveLine = function(x, y, o) {
        $('#board').append(
            $('<div class="line '+o+'"></dot>')
        );

        _placeE($('#board .line:last'), x, y)
            .addClass('inactive');
    }

    var _createInactiveSquare = function(x, y) {
        $('#board').append(
            $('<div class="square"></dot>')
        );

        _placeE($('#board .square:last'), x, y)
            .addClass('inactive');
    }

    var createGrid = function(size) {
        for (var i = 0; i <= size; i++) {
            for (var j = 0; j <= size; j++) {
                _addDot(i, j);
                if (i < size) {
                    _createInactiveLine(i, j, 'v');
                }
                if (j < size) {
                    _createInactiveLine(i, j, 'h');
                }
                if (i < size && j < size) {
                    _createInactiveSquare(i, j);
                }
            }
        }
    }

    var activeLine = function(x, y, o, c) {
        $("line."+o+"[x='"+x+"'][y='"+y+"']")
            .removeClass('inactive')
            .addClass('cbg')
            .attr('num', c);
    }

    createGrid(10);
    activeLine(5,3,'v',3);

    $('.line').click(function() {
        $(this).removeClass('inactive');
    });

    $('.square').click(function() {
        $(this).removeClass('inactive');
    });

});
