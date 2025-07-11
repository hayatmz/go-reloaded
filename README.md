# go-reloaded
## 🧠 Contexte
**Go Reloaded** est un outil de traitement de texte développé en Go, capable de corriger automatiquement un fichier texte en appliquant des **transformations précises** sur les mots : corrections grammaticales, conversions (binaire, hexadécimal), mise en forme (majuscules, minuscules, ponctuations, etc.).

> Ce premier projet post-admission chez Zone01, est un retour aux bases acquises lors de la [piscine go](https://github.com/hayatmz/piscine-golang), mais avec une vision plus claire et structurée. C'est aussi l'introduction à l'évaluation entre pairs via un système d'audit.

## 🎯 Objectifs pédagogiques
- Réutiliser et refactorirer du code existant.<br>
- Manipuler des fichiers en Go (```fs```, ```os```, ```ioutil```).<br>
- Gérer les chaînes de caractères, conversions numériques et majuscules/minuscules.<br>
- Construire un outil de traitement de texte simple mais complet.<br>
- S'habituer à écrire du code lisible.<br>
- Se préparer à l'audit et à l'autoévaluation.<br>

## 🔧 Fonctionnalités principales
Le programme modifie un fichier texte selon les instructions suivantes :<br>

- ```(hex)``` ➜ Convertit le mot précédent de l'hexadécimal au décimal.<br>
- ```(bin)``` ➜ Convertit le mot précédent du binaire au décimal.<br>

**Exemple :**<br>
```sample.txt : Simply add 42 (hex) and 10 (bin)```<br>```result.txt : Simply add 66 and 2```

- ```(up)```, ```(low)```, ```(cap, x)``` ➜ Modifient le mot précédent en MAJ, minuscule ou Capitalisé.<br>
- ```(up, x)```, ```(low, x)```, ```(cap, x)``` ➜ Modifient les x mots précédents en MAJ, minuscules ou Capitalisés.<br>

**Exemple :**<br>
```sample.txt : welcome to (up, 2) the brooklyn bridge (cap)```<br>
```result.txt : WELCOME TO the brooklyn Bridge```

- Corrige l'usage de 'a' en 'an' selon la première lettre du mot qui suit.<br>

**Exemple :**<br>
```sample.txt : this is a untold story```<br>
```result.txt : this is an untold story```

- Corrige la ponctualité (espaces mal placés, etc).<br>
- Reformate les expressions entre guillemets simples ```' '``` sans espace superflu.<br>

**Exemple :**<br>
```sample.txt : Well , this is awkward ...isn't it ? Let's fix it ;please.```<br>
```result.txt : Well, this is awkward... isn't it ? Let's fix it; please.```

## 🗂️ Fichiers

- ```main.go``` : Fichier principal, contenant toute la logique.<br>
- ```sample.txt``` : Fichier texte d'entrée à modifier.<br>
- ```result.txt``` : Fichier de sortie avec les modifications appliquées, crée par le programme.<br>

## Installation

Pour tester le projet localement :<br>

1. Assure toi d'avoir Go installé sur ta machine. Tu peux vérifier avec la commande :<br>
```
go version
```

Si besoin, [installe Golang](https://go.dev/doc/install).<br>

2. Clône le dépôt :

```
git clone https://github.com/hayatmz/go-reloaded.git
```
et assure toi d'être dedans.
```
cd go-reloaded
```

3. Pour lancer le programme :

```
go run main.go sample.txt result.txt
```

> [sample.txt](./sample.txt) peut-être modifié pour tester d'autres textes et [result.txt](./result.txt) sera automatiquement mis à jour en fonction de son contenu (il peut également être supprimé avant de lancer le programme).
