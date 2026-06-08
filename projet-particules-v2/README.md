**Projet Particle : un programme d'animation 2D écrit en golang.**


Ce projet implémente un système de particules de base, conforme aux attentes spécifiées dans le "Premier Travail".
A ce système de base, nous avons ajoutés des implémentations d'extensions proposées dans la partie "Deuxième travail" : 

L'extension 5.1 est matérialisée par la variable Gravity, paramétrable dans config.json, la valeur qu'on y associe est divisée par 360 (une valeur élevée afin que Gravity n'aie pas tout de suite un impact trop fort) avant d'être ajouté aux coordonées Y de chaque particule. Une gravité négative est possible, et fait chuter les particules vers le haut.

L'extension 5.2 est matérialisée par la variable DrawOblivionMargin, parametrable dans config.json, elle est désactivable en associant une valeur négative à la variable. DrawOblivionMargin donne le nombre de pixels qu'une particule doit entièrement franchir après être sortie entièrement de l'écran.

L'extension 5.3 est représentée par la variable DeathTime, paramétrable dans config.json, pour désactiver cette extension il suffit d'assigner 0 à la variable DeathTime. Chaque particule est generée avec un LifeTime égal à 0, et mettre par exemple 500 à DeathTime signifie que la particule disparaîtra au 500ème appel de la méthode Update() ayant pour receiver le System.

L'extension 5.4 est particulière, car nous l'avions en partie réalisée pour notre Premier Travail. 
Les variations de couleur sont configurables dans le fichier json par un booléen appelé EnableColors. 
Les variations de taille des particules sont désactivables en mettant la variable ScaleVariation à 0, et activable en mettant un nombre positif. 
Les variations d'opacité sont manipulables via le booléen OpacityVariation du json. 
Les variations de rotation sont parametrables par la variable ParticleRotationSpeed du json, y affecter 0 désactive la rotation.

L'extension 5.5 est ici directement liée à deux variables du fichier config.json, RandomSpawn et RotorSpeed. 
Si RandomSpawn est sur false, et que RotorSpeed est à 0, les particules ont une vitesse aléatoire et l'extension est considerée comme inactive. 
Si RotorSpeed est supérieur à 0, le mouvement de notre système se met en place.
Si Randomspawn est sur true, et que RotorSpeed est à 0, les particules se déplaçent aléatoirement mais elles ne varient plus de couleur.
Enfin, si RandomSpawn est sur true est que RotorSpeed dépasse 0, les particules se déplaçent aléatoirement, et leur couleur varie en fonction de la vitesse du rotor. 

L'extension 5.6 est liée a deux variables de config.json, RightClickSpawning et LeftClickMovesSpawner. La première, mise à true, génère au clic droit un nombre de particules égal à la valeur de InitNumParticles. La seconde mise à true permet de changer les variabels SpawnX et SpawnY (et n'est donc utile que si RandomSpawn vaut false).

L'extension 5.7 implémente un interface utilisateur, cette extension est switchable via la variable EnableInterface située dans config.json. Si cette dernière est activée, l'interface est caché au lancement du programme, et est activable en pressant la touche Espace. Les cases rouges sont les variables booléennes de config.json, cliquer dessus inverse leur valeur, et les cases bleues permettent de faire varier les valeurs de cinq variables de la config.  

L'extension 5.8 se matérialise par la variable MemoryThanksYou de config.json. La mettre à false désactive l'extension, et les particules "mortes" seront figées sur l'ecran si elles s'y trouvent encore à leur mort. Assigner true permet de gerer ces particules mortes via une seconde list.List, qui recupère ces dernières et se vide à chaque appel de la méthode (s *System) Update(). 


Nous avons aussi expérimentés une implémentation : 

Ajout d'une fonctionnalité sonore, les clics gauche et droit dans la fenêtre du jeu (extension 5.6) activent deux sons. Pour ce faire il a fallu completer la structure du System et y ajouter un champ Sound, pointant vers une structure en relation avec l'import des fonctionnalités audio ebiten. Cette extension est désactivable dans config.json en assignant false aux booléens EnableLeftClickSound et EnableRightClickSound. 



**Explication des modifications apportées aux fichiers à la racine du projet :** 


_1) draw.go_

Nous avons du ajouter deux fonctions : notInRange() qui renvoie un booléen utilisé par la deuxième fonction : needToKill(). Ces fonctions sont expliquées en commentaires, mais nous en avions besoin pour conditionner l'affichage par la fonction Draw() des particules à leur position dans l'ecran, et en prenant en compte la marge de la config. 

Dans la fonction Draw() elle-même, nous avons ajouté :

- if needToKill(p) && config.General.DrawOblivionMargin > 0 {
				continue
			} else {
Cet ajout renvoie à l'explication donnée plus haut, si l'extension d'optimisation de la mémoire 5.2 est activée, une particule qui entre dans les conditions de la fonction needToKill() n'a plus besoin d'être dessinée. Modification suggerée par les consignes du PDF SAE.

- options.GeoM.Translate((-float64(w) / 2), (-float64(h) / 2))
Cet ajout permettent de "recentrer" le centre de gravité de notre particule, et lui donne une rotation plus cohérente visuellement. 

- if particles.TheInterface.Visible {
		particles.DrawInterface(screen)
	} else if config.General.EnableInterface {
		ebitenutil.DebugPrintAt(screen, "Press SPACE to show the interface", 5, 15)
	}
Pour l'extension 5.7 nous avions besoin de modifier draw.go pour implementer l'interface, comme le préconisait le PDF SAE. Si l'extension est activée, un texte situé sous le compteur de frames nous indique la touche à presser pour afficher l'interface. Lorsqu'il devient visible, il est dessiné par la fonction.

_2) main.go_

Nous avons ajoutés les lignes : 
- audioContext := audio.NewContext(44100)
- sound := particles.NewSoundManager(audioContext)

Ces deux lignes servent à initialiser le système audio fourni par ebiten, il est requis de le faire dans la fonction main() afin d'éviter des crashs et des appels multiples (fortement déconseillé avec cette fonctionnalité).

La ligne : 
- g := game{system: particles.NewSystem(sound)}
A également été enrichie en conséquence.

_3) update.go_

Nous avons ajoutés les lignes : 
- if inpututil.IsKeyJustPressed(ebiten.KeySpace) && config.General.EnableInterface {
		particles.TheInterface.SwitchInterface()
	}
	if particles.TheInterface.Visible && config.General.EnableInterface {
		particles.TheInterface.UpdateInterface()
	}

Conformement à ce qui est précisé dans le PDF SAE, gerer l'extension 5.7 et l'interaction avec le clavier et la souris se fait au niveau de la méthode Update() agissant sur le *Game. 


**Paramètres de lancement optimaux (les valeurs non-spécifiées ont un impact moindre):**

-SpawnRate : une valeur comprise entre 0 et 1500.

-ScaleVariation : une valeur comprise entre 0 et 3, avec des pas de 0.5.

-DeathTime : intêret limité au delà de 1000.

Ces trois valeurs sont aussi modifiables dans l'interface dynamique. 









 