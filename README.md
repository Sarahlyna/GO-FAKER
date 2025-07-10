# Go Faker From Scratch

Projet d'école : Générateur de fausses données **sans dépendances externes** en Go, réalisé **from scratch** pour comprendre la génération conditionnelle intelligente, la structuration backend/frontend, et l'architecture d'une application Go simple.

- Racha RAMOUL
- Sarah Lina SALAMANI
- Jay BURY

## Sommaire

- [Objectifs](#objectifs)
- [Fonctionnalités](#fonctionnalités)
- [Architecture](#architecture)
- [Installation](#installation)
- [Utilisation](#utilisation)
  - [Génération simple](#génération-simple)
  - [Génération par schéma](#génération-par-schéma)
  - [Téléchargement CSV/JSON](#téléchargement-csvjson)
- [Structure du projet](#structure-du-projet)
- [Tests](#tests)
- [Explications techniques](#explications-techniques)
- [Améliorations possibles](#améliorations-possibles)

## Objectifs

- Créer un **faker** permettant de générer :
  - Noms, emails, téléphones, adresses, âges, villes, codes postaux.
  - Données réalistes selon la langue (français ou anglais).
  - Codes postaux cohérents avec les villes.
  - Téléphones cohérents selon le pays.
- Réaliser le projet **sans dépendances tierces**, uniquement avec la **stdlib Go**.
- Comprendre l'architecture backend Go + frontend minimaliste.
- Fournir un outil pédagogique utilisable pour générer des datasets de test.

## Fonctionnalités

- Génération de données aléatoires configurables via requêtes HTTP.
- Support des locales (`fr`, `en`).
- Génération par schéma JSON pour des structures complexes.
- Téléchargement des données générées en JSON ou CSV.
- Visualisation immédiate via frontend simple.
- Utilisation via navigateur ou via `curl` pour automatisation.

## Architecture

**Backend Go :**

- `main.go` : point d'entrée du serveur HTTP, routes `/api/fake`, `/api/fake/schema`, `/api/fakers`.
- `faker/` : contient les fakers structurés en interfaces Go (NameFaker, EmailFaker, PhoneFaker, etc.).
- Aucune dépendance externe, tout est basé sur `net/http`, `encoding/json`, `math/rand`.

**Frontend :**

- `index.html`, `style.css` :
  - Interface de test.
  - Génération via schéma éditable JSON.
  - Boutons de téléchargement JSON/CSV.
  - Visualisation en tableau.

**Déploiement Docker :**

- `Dockerfile` et `docker-compose.yml` pour exécution rapide.
- Utilisation d'Air pour le live reload pendant le développement.

## Lancement

Lancer l'application :

- Docker compose up --build
