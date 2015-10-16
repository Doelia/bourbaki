## Quel est le but du jeu?
Pour gagner une partie, il faut réussir à former **le plus de carrés possible**.
Le jeu se joue au tour par tour où chaque joueur pose une ligne sur la grille à son tour

## Comment s'organise un tour ?
Chaque joueur place une ligne de façon à essayer de former un carré, ou au moins d'éviter de permettre aux autres joueurs de former un carré. Quand un joueur ferme un carré, il peut rejouer et donc replacer une ligne.

## Comment sont attribués les points ?

Il existe plusieurs manières de gagner ou de perdre des points :
- À chaque fois que vous placez une ligne, **vous gagnez 1 point**. Si vous ne jouez pas après 10 secondes, une ligne est jouée à votre place et vous ne gagnerez pas de points
- Quand vous parvenez à fermer un carré, **vous gagnez 5 points** et jouez à nouveau
- Par contre si vous permettez au joueur suivant de former un carré, vous **perdez 1 point**.

Si le joueur parvient à faire des combos en faisant une suite de carrés, le joueur précédant perdra un point par carré réalisé.
Grâce à cette distribution des points, le joueur est encouragé à jouer même s'il ne peut pas faire de coups lui rapportant de points, puisque de toute façon une IA jouera à sa place.

## Comment sortir le jeu de la pause?
Il faut être minimum deux joueurs pour pouvoir jouer. Lorsque le nombre de joueurs actifs est inférieur à deux, le jeu passe en état de pause.

## A quoi sert le classement général ?
A la fin d’une partie, les points gagnés par chaque joueur s’accumulent. Ce total est visible dans le **classement général**.  On y trouve aussi qui a accumulé le plus de points, et en combien de parties. On voit également qui a fini premier le plus de fois !

##  Que se passe t-il si je me déconnecte ?
Pour éviter le blocage de la partie, une ligne sera automatiquement placée de façon aléatoire lorsque ce sera votre tour. Vous ne gagnerez pas de points, - sauf si vous avez énormément de chance ! -. Il vaut donc mieux finir une partie commencée, même si elle est mal partie, pour perdre le moins de points possible.
