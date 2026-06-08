package particles

import (
	"container/list"
)

// System définit un système de particules.
// Pour le moment il ne contient qu'une liste de particules, mais cela peut
// évoluer durant votre projet.
// Commentaire supplémentaire :
// Il contient également un gestionnaire de sons.
type System struct {
	Content *list.List
	Sound   *SoundManager
}
