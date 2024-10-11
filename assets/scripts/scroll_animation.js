function initializeAnimations() {
    // Create the observer
    const observer = new IntersectionObserver((entries) => {
        entries.forEach((entry) => {
            if (entry.isIntersecting) {
                entry.target.classList.add("show");
            } 
            // if uncommented fires everytime
            // else {
            //     entry.target.classList.remove("show");
            // }
        });
    });

    // Get all hidden elements and observe them
    const hiddenElements = document.querySelectorAll(".hidden, .hidden-swipe");
    hiddenElements.forEach((el) => {
        // Remove any existing show class when reinitializing
        el.classList.remove("show");
        observer.observe(el);
    });
}

// Run on initial page load
document.addEventListener("DOMContentLoaded", initializeAnimations);

// When using HTMX, run after content is loaded
document.addEventListener("htmx:afterSettle", initializeAnimations);

// For regular navigation events
window.addEventListener("popstate", initializeAnimations);
