function lightTap() {
    if ("vibrate" in window.navigator) {
        window.navigator.vibrate(50);
    }
}

function initializeHaptics(){
    for (let button of document.getElementsByTagName("button")) {

        button.addEventListener("mousedown", (e) => {
            lightTap();
        }, { passive: false });  // Explicitly set passive to false to avoid interference
    }
    for (let a of document.getElementsByTagName("a")) {

        a.addEventListener("mousedown", (e) => {
            lightTap();
        }, { passive: false });  // Explicitly set passive to false to avoid interference
    }
    for (let input of document.getElementsByTagName("input")) {

        input.addEventListener("mousedown", (e) => {
            lightTap();
        }, { passive: false });  // Explicitly set passive to false to avoid interference
    }
}

// Run on initial page load
document.addEventListener("DOMContentLoaded", initializeHaptics);

// When using HTMX, run after content is loaded
document.addEventListener("htmx:afterSettle", initializeHaptics);

// For regular navigation events
window.addEventListener("popstate", initializeHaptics);
