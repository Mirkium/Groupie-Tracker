let currentIndex = 0;
const carrousel = document.querySelector(".carrouselCocktail");

function moveSlide(direction) {
    const items = document.querySelectorAll('.Cocktail');
    currentIndex = (currentIndex + direction + items.length) % items.length;
    document.querySelector('.slide').style.transform = `translateX(-${currentIndex * 100}%)`;
}

