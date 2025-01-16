## Explication des fichiers CSV

### 1. `agency.csv` - Informations sur les agences

- **agency_id** : Identifiant unique de l'agence.
- **agency_name** : Nom officiel de l'agence.
- **agency_url** : URL du site web de l'agence.
- **agency_timezone** : Fuseau horaire dans lequel l'agence opère.
- **agency_lang** : Code de la langue principale utilisée par l'agence.
- **agency_phone** : Numéro de téléphone de contact de l'agence.
- **agency_fare_url** : URL où les informations tarifaires de l'agence sont disponibles.
- **agency_email** : Adresse email de contact de l'agence.
- **ticketing_deep_link_id** : Identifiant pour les liens profonds vers les services de billetterie.

### 2. `routes.csv` - Informations sur les lignes
- **route_id** : Identifiant unique de la ligne.
- **route_short_name** : Nom court de la ligne, généralement une version abrégée ou un numéro.
- **route_long_name** : Nom complet de la ligne.
- **route_type** : Type de service de transport (par exemple, bus, train).
- **route_color** : Couleur associée à la ligne pour l'identification.

### 3. `stops.csv` - Informations sur les arrêts
- **stop_id** : Identifiant unique de l'arrêt.
- **stop_name** : Nom de l'arrêt.
- **stop_lat** : Latitude de l'arrêt.
- **stop_lon** : Longitude de l'arrêt.
- **wheelchair_boarding** : Indique si l'arrêt est accessible aux fauteuils roulants.

### 4. `transfers.csv` - Informations sur les transferts entre arrêts
- **transfer_id** : Identifiant unique du transfert.
- **from_stop_id** : Identifiant de l'arrêt de départ.
- **to_stop_id** : Identifiant de l'arrêt d'arrivée.
- **min_transfer_time** : Temps minimum (en seconde) nécessaire pour effectuer le transfert entre les deux arrêts.

### 5. `trips.csv` - Informations sur les voyages

- **trip_id** : Identifiant unique du voyage.
- **from_stop_id** : Identifiant de l'arrêt de départ.
- **to_stop_id** : Identifiant de l'arrêt d'arrivée.
- **time** : Temps minimum (en seconde) nécessaire.
- **route_id** : Identifiant de la ligne associé au voyage.
- **bikes_allowed** : Indique si les vélos sont autorisés sur ce voyage.

> **Info:** Un voyage correspond à chaque sous-trajet (Aller/Retour) d'une ligne. Par exemple pour le T1, on a : IUT --> Croix-Luizet / Croix-Luizet --> Einstein / ... / ENS Lyon --> Debourg

> **Remarque:** On suppose qu'il n'y pas de temps d'attente en station

## Comment ajouter une ligne

Voici un guide pour savoir où et comment ajouter une ligne (par exemple le T1) :

1. **agency.csv** : Ce fichier contient les informations sur les agences de transport. Si vous voulez ajouter une nouvelle agence ou modifier les détails d'une agence existante, c’est ici.

2. **routes.csv** : Ce fichier liste les différentes lignes. Normalement, il y a déjà tout. Cherchez l'identifiant de votre ligne `route_id`.

3. **stops.csv** : Ce fichier contient les informations sur les arrêts de chaque ligne. Il faut ajouter manuellement les stations du T1 par exemple. Vous avez une liberté sur le choix de `stop_id`. Essayez de suivre la logique pré-existante.

> **Attention:** La station T1 Charpennes est différente de celle du métro A ou du T4. Chaque ligne possède sa propre station. Par contre, on aura une correspondance T1-T4 et T1-MA.

4. **transfers.csv** : Ce fichier détaille les possibilités de correspondance entre les arrêts avec d'autres lignes. Par exemple, à Charpennes on a une correspondance avec le métro A, B, T4 et certains bus. Vous avez une liberté sur le choix de `transfer_id`. Essayez de suivre la logique pré-existante.

5. **trips.csv** : Ce fichier décrit les voyage au sein d'une ligne. Un voyage correspond à chaque sous-trajet (Aller/Retour) d'une ligne. Par exemple pour le T1, on a : IUT --> Croix-Luizet / Croix-Luizet --> Einstein / ... / ENS Lyon --> Debourg. Vous avez une liberté sur le choix de `trip_id`. Essayez de suivre la logique pré-existante.