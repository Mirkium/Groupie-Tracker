const img = document.querySelectorAll('.IMG_ingredient'); // Récupérer toutes les images

const maxX = -2000; // Limite minimale
const newX = 2000; // Position de redémarrage
const vitesse = 2; // Vitesse du déplacement
const decalage = 200; // Décalage en millisecondes (2 secondes)

function animer() {
  img.forEach((image, index) => {
    let positionX = newX; // Position initiale

    setTimeout(() => {
      function bouger() {
        // Déplacer l'image
        positionX -= vitesse;
        image.style.transform = `translateX(${positionX}px)`;

        // Replacer l'image si elle dépasse la limite
        if (positionX < maxX) {
          positionX = newX; // Repositionner
        }

        // Répéter l'animation
        requestAnimationFrame(bouger);
      }

      bouger(); // Lancer l'animation pour cette image après le délai
    }, index * decalage); // Appliquer le décalage pour chaque image
  });
}

// Lancer l'animation
animer();
