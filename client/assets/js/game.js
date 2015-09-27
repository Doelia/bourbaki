$(document).ready(function() {

    $('#btn-howto').click(function() {
        $('#guide')
            .modal('setting', 'transition', 'vertical flip')
            .modal('show');
    });

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

    //pause();

    $('.ttip').popup({
        inverted: true,
        position: 'top center',
        duration: 150,
    });

});
