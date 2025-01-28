# TcTurtleProject

## Description

TcTurtleProject est une application web développée en Elm permettant de visualiser des dessins générés par des commandes de tracé basées sur le langage TcTurtle. Inspiré des Turtle Graphics, TcTurtle permet d'exprimer le chemin suivi par un crayon pour dessiner des formes complexes en utilisant des instructions simples.

## Fonctionnalités

- **Instructions de Base** :
  - `Forward x` : Avancer de x unités.
  - `Left x` : Tourner à gauche de x degrés.
  - `Right x` : Tourner à droite de x degrés.
- **Instructions Avancées** :
  - `Repeat x [ ... ]` : Répéter x fois la suite d'instructions spécifiée.

## Démarrer le Serveur de Développement

```
elm reactor
```

Ouvrez votre navigateur et naviguez vers http://localhost:8000/src/Main.elm