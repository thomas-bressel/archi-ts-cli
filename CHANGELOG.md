# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]


## [1.6.3] - 2025-06-05
## Fix
    - Correction dans le fichier npm.go en placant une condition pour l'installation d'Express dans le package.json
    - Ajout de la confirmation si Express a été installé ou non dans le compte rendu de l'installation



## [1.6.2] - 2025-06-04
## Added
    - Execution de npm init une fois la configuration validé.
    - Ececution de l'installation de dependances.
    - Execution de l'installation de pnpm.
## Fix (Merci à Mahery Randrianirina pour le petit feedback ;)
    - Ajout de l'initialisation automatique du module Go (go mod init) si le fichier go.mod est absent.
    - Installation automatique des dépendances via go mod tidy lors du make build.
    - Corrige les erreurs rencontrées lors d’un premier make install sur un dépôt fraîchement cloné.
## Fix 
    - Suppression de npm init .... stupide vu que le package.json a déjà été généré !
    - Ajout d'un lien de don paypal pour ceux qui souhaite aider à financer ce projet.


## [1.5.2] - 2025-06-03
### Added
    - Generate tsconfig.json file if typescript is selected
    - Generate .gitignore file
    - Generate jest.config file with different configuration ts or js
    - Add a new prompt to install Express Library or not
### Update
    - Update the package.json template and logic to add dev dependencies in the right order
    - Add a make file to automatize installation and configuration



## [1.4.0] - 2025-06-02
### Added
    - Generate package.json file with the selected language type.
    - Generate .env file template.



## [1.3.0] - 2025-05-31
### Added
- Adding a prompt to choose the language type between typescript and javascript. 
- Adding a 3rd choice into the architecture prompt.
- Create folders structure depends the architecture type.



## [1.1.0] - 2025-05-30
### Added
- Adding a first prompt to the project to enter the name of the project with command `./archi create`



## [1.0.0] - 2025-05-30
### Added
- Starting the project development
- Setting up the initial project structure