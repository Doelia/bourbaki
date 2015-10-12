# TODO

## Projet
- Écrire les instructions de compilation dans le readme
- Revoir le système de package -> normal d'avoir le nom du projet partout ?
- Revoir la structure pour ne pas devoir compiler dans /server/

## Serveur
- Empêcher l'injection HTML dans les pseudos (+client?)
- Gérer les coups qui générent 2 carrés d'un coup
- Faire l'IA en ramdom

## Client
- Ajouter un logo
- Rédiger le "Comment jouer ?"
- Faire propre le "Mot de passe incorrect"

## Cas de triche
Ces cas doivent être bloqués par le serveur :
- J'envoie une ligne alors que ce n'est pas mon tour
- J'envoie une ligne qui n'existe pas sur la grille

## Avant MEP
- Remettre la popup de première connexion (socket.js -> on CONNECTACCEPT)
- Retirer les boutons de login colorés
- Retirer le changement de title
