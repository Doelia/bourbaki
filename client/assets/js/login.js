function init_login() {
    $('.ui.form')
    .form({
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

    $('.ui.form')
        .submit(function(event) {
            event.preventDefault();
            socket.emit("LOGIN", "Doe", "John");
        });
}
