export function hideSuccessMessage() {
    let successMessage = document.querySelector("p.success");
    if (successMessage) {
        setTimeout(() => successMessage.style.display = "none", 2000);
    }
}

