var isValid = false;

function init_login() {
    $('.ui.form').form({
        onSuccess: function() {
            isValid = true;
        },
        onFailure: function() {
            isValid = false;
        },
        fields: {
            user: {
                identifier  : 'user',
                rules: [
                    {
                        type   : 'empty',
                        prompt : 'Entrez votre pseudo'
                    }
                ]
            },
            password: {
                identifier  : 'password',
                rules: [
                    {
                        type   : 'empty',
                        prompt : 'Veillez entrer (ou choisir) un mot de passe'
                    },
                    {
                        type   : 'length[6]',
                        prompt : 'Votre mot de passe doit comporter au moins 6 caractères'
                    }
                ]
            }
        }
    });

    $('.ui.form').submit(function(event) {
        event.preventDefault();
        if (isValid) {
            $('.ui.form').addClass('loading');
            var login = $('.ui.form').form('get value', 'user');
            var pass = $('.ui.form').form('get value', 'password');
            console.log("Login with "+login+", "+pass);
            socket.emit("LOGIN", login, pass);
        }
    });
}
