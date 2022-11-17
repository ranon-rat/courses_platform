"use strict";
document.addEventListener('scroll', () => {
    const navbar = document.querySelector('navbar');
    if (window.scrollY !== 0) {
        navbar.style.boxShadow = '0 0 10px rgba(0, 0, 0, 0.5)';
        return;
    }
    navbar.style.boxShadow = 'none';
});
