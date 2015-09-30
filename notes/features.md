# Parties fonctionelles du serveur
- Login user + mot de passe
  - Gestion BD : user, mot de passe, score
  - Si le compte existe pas déjà, il est crée
  - Sinon, connexion
- Réception des actions des joueurs : poser une ligne
- Calcul du score partie en fonction des actions
- Test de validité du carré
- Mise en pause de la partie si moins de 2 joueurs
- Gestion des connexions/déconnexions
  - Lors de la déconnexion, le score du joueur visible dans le tableau "score partie" passe en grisé et le serveur joue à sa place
  - S'il se reconnecte, il repasse en coloration normale et récupère son score
- Gestion de la fin de la partie
  - Remise à zéro du tableau de scores
  - Enregistrement des scores en fin de partie "classement général"
  - Nettoyage de la grille

# Parties fonctionelles du client

- Connexion (user/mot de passe)
- Affichage dynamique :
    - De la grille (barres, carrés)
    - Des joueurs actifs/inactifs (pseudo, score)
    - Timer
    - Mode pause
- Affichage statique :
    - Classement général
- Actions :
    - Poser une ligne
