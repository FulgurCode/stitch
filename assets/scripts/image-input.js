function handleFileSelect(event, sectionId) {
    const files = event.target.files;
    const section = document.getElementById(sectionId);
    const uploadButton = section.querySelector('.upload-button');

    if (sectionId === 'mainImageSection') {
        // Remove existing preview for main image
        const existingImg = section.querySelector('img');
        if (existingImg) {
            section.removeChild(existingImg);
        }
        if (files.length > 0) {
            const file = files[0];
            previewImage(file, section, uploadButton);
        }
    } else if (sectionId === 'moreImagesSection') {
        // Remove all existing previews
        const existingImgs = section.querySelectorAll('img');
        existingImgs.forEach(img => section.removeChild(img));

        const filesToAdd = Math.min(MAX_ADDITIONAL_IMAGES, files.length);
        for (let i = 0; i < filesToAdd; i++) {
            previewImage(files[i], section, uploadButton);
        }
        updateImageCounter(filesToAdd);
    }
}

function previewImage(file, section, uploadButton) {
    const reader = new FileReader();
    reader.onload = function(e) {
        const newImg = document.createElement('img');
        newImg.src = e.target.result;
        newImg.alt = "Uploaded image";
        section.insertBefore(newImg, uploadButton);
    }
    reader.readAsDataURL(file);
}

function updateImageCounter(count) {
    const counter = document.getElementById('imageCounter');
    counter.textContent = `${count} / ${MAX_ADDITIONAL_IMAGES} images`;
}

document.getElementById('fileInput_main').addEventListener('change', function(event) {
    handleFileSelect(event, 'mainImageSection');
});

document.getElementById('fileInput_more').addEventListener('change', function(event) {
    const files = event.target.files;
    if (files.length > MAX_ADDITIONAL_IMAGES) {
        alert(`You can only select up to ${MAX_ADDITIONAL_IMAGES} additional images.`);
        event.target.value = ''; // Clear the file input
        return;
    }
    handleFileSelect(event, 'moreImagesSection');
});