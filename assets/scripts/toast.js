function showToast(message, type = 'default', duration = 3000) {
    const toastContainer = document.getElementById('toastContainer');
    const toast = document.createElement('div');
    toast.className = `toast ${type}`;

    if (type == "success") {
        toast.innerHTML = `<img src="/static/icons/info.svg"/> <span>${message}</span>`;
    } else if (type == "error") {
        toast.innerHTML = `<img src="/static/icons/warning.svg"/> <span>${message}</span>`;
    } else {
        toast.innerHTML = `<span>${message}</span>`;
    }

    toastContainer.appendChild(toast);
    toast.offsetHeight;
    toast.classList.add('show');

    setTimeout(() => {
        toast.classList.remove('show');
        setTimeout(() => {
            toastContainer.removeChild(toast);
        }, 300);
    }, duration);
}