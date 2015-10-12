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
                    },
                    {
                        type   : 'regExp[/^[a-zA-Z0-9-_]{3,15}$/]',
                        prompt : 'Pseudo invalide. Trop long ? Caractères spéciaux ?'
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
            document.title = login;
            socket.emit("LOGIN", login, pass);
        }
    });

}
