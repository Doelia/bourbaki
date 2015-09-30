# Protocole

Décrit la liste des paquets et leurs parametres qui passent ente le client et le serveur.

## Envoyées par le client

### CONNEXION (user,pass)
Paquet envoyé à la connexion. Mot de passe envoyé en md5.

### PUTLINE(x,y,o)
Prévient le placement d'une ligne par le joueur.
- **x** : int, position X
- **y** : int, position Y
- **o** : string, orientation. 'v' pour vertical, 'h' pour horizontal

## Envoyées par le serveur

### CONNECT_ACCEPT(code, numPlayer)
Packet de retour acceptant ou non la demande de connexion du joueur.

- **code** : int, 0 si mot de passe incorrect, 1 si connexion OK, 2 si connexion OK + compte créé
- **numPlayer** : int, numéro du joueur dans la partie (entre 1 en n). 0 si le code vaut 0 (connexion refusée)

Tous les packets d'initialiastion du jeu suiveront ce packet

### DISPLAYLINE(x,y,o,n)
Ajoute la barre x,y,o à la grille

- **x** : int, position X
- **y** : int, position Y
- **o** : string, orientation. 'v' pour vertical, 'h' pour horizontal
- **n** : int, numéro identifiant du joueur (1 à N)

### DISPLAYBOX(x,y,n)
Ajoute le carré à la grille

- **x** : int, position X
- **y** : int, position Y
- **n** : int, numéro identifiant du joueur (1 à N)

### UPDATEPLAYERS(json)
Met à jour du tableau des scores de la partie.
Appelé autant de fois que necéssaire

Exemple :
```
[
    {"numPlayer":"1", "name": "Portrick", "score":"23", isActive: true},
    {"numPlayer":"2", "name": "Faewynn", "score":"178", isActive: true},
    {"numPlayer":"3", "name": "Pancake", "score":"87", isActive: false},
]
```

### SETACTIVEPLAYER(numPlayer)
Définit le joueur actif (celui qui est en train de joueur). Envoyé a chaque changement de joueur


### PAUSE()
Place le jeu en pause (si pas assez de joueurs)

### UNPAUSE()
Retire la pause

### EMPTYGRID()

Remet à zéro la grille
Utilisé pour les fins de partie

### GRID(json)

Envoi la grille (lignes, carrés) à la connexion d'un nouveau joueur.

Exemple :
```
{
    lines:
    [
        {"x": 3, "y": 4, "o": "v", "n": 1},
        {"x": 2, "y": 3, "o": "h", "n": 2},
    ],
    squares:
    [
        {"x": 3, "y": 4, "n": 1},
        {"x": 2, "y": 3, "n": 2},
    ],
}
```

## Pas encore traité

Traitements pas encore intégrés au protocole :

- Envoi des fichiers html et assets (Apache séparé ? http intégré au GO ?)
- Affichage du classement (websockets ou ajax ?)
