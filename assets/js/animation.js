let currentIndex = 0;

function moveSlide(direction) {
    const items = document.querySelectorAll('.carrouselItem');
    currentIndex = (currentIndex + direction + items.length) % items.length;
    document.querySelector('.carrouselCocktail').style.transform = `translateX(-${currentIndex * 100}%)`;
}