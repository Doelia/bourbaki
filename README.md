# Projet Bourbaki

Version web du jeu des *petits carrés* implémenté en Golang.
Réalisé en collabaration avec [Marlène Guillemette](https://github.com/MarleneGuillemette)

Démonstration : http://bourbaki.doelia.fr/

## Principe du jeu

https://fr.wikipedia.org/wiki/La_Pipopipette

La Pipopipette ou « jeu des petits carrés » est un jeu de société au tour par tour.

![](http://jeuxstrategieter.free.fr/jeu_pipopipette/ex1.jpg)

Le jeu se joue généralement avec papier et crayon sur du papier quadrillé. À chaque tour, chaque joueur trace un petit trait suivant le quadrillage de la feuille. Le but du jeu est de former des carrés. Le gagnant est celui qui a fermé le plus de carrés.

## Fonctionnalités

- Connexion / Inscription rapide
- Partie unique à laquelle des joueurs peuvent se connecter à tout moment
- Gestion des connexions / déconnexions en pleine partie
- Temps limité pour jouer
- IA en cas de déconnexion d'un joueur ou de temps écoulé
- Calcul et affichage des scores en temps réél
- Classement général avec cumul des points gagnés

Détail complet des fonctionnalités :  https://github.com/doelia/go-bourbaki/blob/master/doc/features.md

## Démonstration en ligne

[![](https://raw.githubusercontent.com/Doelia/go-bourbaki/master/doc/screenshot.png)](http://bourbaki.doelia.fr:2000)

Une démo du jeu est en ligne sur http://bourbaki.doelia.fr/

## Installation et utilisation

### En local avec Golang

Récupération des sources :
```
go get github.com/doelia/go-bourbaki
```

Build et installation du binaire :
```
go install github.com/doelia/go-bourbaki
```

Lancement du serveur :
```
cd $GOPATH/bin
./go-bourbaki -port 2000
```
Ce qui ouvre un serveur web à l'adresse http://locahost:2000

### Avec docker

```
git clone https://github.com/Doelia/go-bourbaki
cd go-bourbaki
docker build . -t go-bourbaki
docker run -d -p 80:2000 go-bourbaki app
```

Ce qui ouvre un serveur web à l'adresse http://locahost/

## Technologies utilisées

- Serveur GO
    - Envoi des ressources via HTTP
    - Sauvegarde des données persitantes en [Bolt](https://github.com/boltdb/bolt) (système clé/valeur)
- Client Web
    - Framework [Semantic-ui](http://semantic-ui.com/)
    - Librairie [jQuery](https://jquery.com/)
    - Bibliothèque d'icônes [Font Awesome](http://fortawesome.github.io/Font-Awesome/icons/)
    - Pré-processeur  [Less](http://lesscss.org/)
- Communication via WebSockets avec la librairie [Socket.io](http://socket.io/)
    - Implémentation GO : https://github.com/googollee/go-socket.io

## Documentation technique
- Documentation Go du projet générée : http://bourbaki-doc.doelia.fr/pkg/go-bourbaki/
- [Protocole](https://github.com/doelia/go-bourbaki/blob/master/doc/protocol.md) (liste des paquets)
- [Règles du jeu](https://github.com/doelia/go-bourbaki/blob/master/doc/FAQ.md)
- [Ébauche diagramme UML](https://github.com/doelia/go-bourbaki/raw/master/doc/UMLBourbaki.pdf)


