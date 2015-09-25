## Client -> Serveur
- CO(user,pass) : paquet de connexion
- LINE(x,y,o) : demande d'ajout de la barre x,y,o

## Serveur -> Client
- DISPLAYLINE(x,y,o,c) : rajoute la barre x,y,o à la grille
- DISPLAYBOX(x,y,c) : rajoute le carre à la grille
- UPDATEPLAYERS(json) : mise à jour du tableau des scores de la partie en jeu
  - user
  - score
  - couleur(chiffre entre 1 et n)
  - actif?
- SETACTIVEPLAYER(user)
- STARTPAUSE()
- ENDPAUSE()
- EMPTYGRID() : remise à zéro de la grille
- GRID(json) : envoi de la grille, des barres, et des carrés pour un nouveau joueur
  - liste des barres
  - liste des carrés
