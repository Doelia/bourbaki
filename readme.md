# Jeu de la pipopipette

## Principe du jeu

https://fr.wikipedia.org/wiki/La_Pipopipette

La Pipopipette ou « jeu des petits carrés » est un jeu de société au tour par tour.

![](http://jeuxstrategieter.free.fr/jeu_pipopipette/ex1.jpg)

Le jeu se joue généralement avec papier et crayon sur du papier quadrillé. À chaque tour, chaque joueur trace un petit trait suivant le quadrillage de la feuille. Le but du jeu est de former des carrés. Le gagnant est celui qui a fermé le plus de carrés.

## Fonctionnalités

- Connexion / Inscription rapide
- Une partie unique à laquelle des joueurs peuvent se connecter à tout moment
- Gestion des connexions / déconnexions
- Calcul et affichage des scores en temps réél
- Classement général avec cumul des points (ou parties gagnées ?)

Détail complet :  https://gitlab.info-ufr.univ-montp2.fr/HMIN302/go-bourbaki/blob/master/notes/features.md

## Technologies

- Serveur GO
    - Envoi des ressources via HTTP
    - Stockage en [Bolt](https://github.com/boltdb/bolt) (système clé/valeurs)
- Client Web
    - Framework [Semantic-ui](http://semantic-ui.com/)
    - Librairie [jQuery](https://jquery.com/)
    - Bibliothèque d'icônes [Font Awesome](http://fortawesome.github.io/Font-Awesome/icons/)
    - Pré-processeur  [Less](http://lesscss.org/)
- Communication via WebSockets avec la librairie [Socket.io](http://socket.io/)
    - Implémentation GO : https://github.com/googollee/go-socket.io

*Le maximum du développement sera déporté sur le serveur, idéalement le client se tâchera de faire uniquement de l'affichage.*
