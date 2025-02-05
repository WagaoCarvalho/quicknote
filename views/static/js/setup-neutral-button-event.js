// Função para configurar o botão neutro
export function setupNeutralButtonEvent() {
    // Selecionar o botão neutro
    const neutralButton = document.querySelector("button.neutral");
    if (neutralButton) {
        neutralButton.addEventListener("click", function () {
            // Redirecionar para "/"
            window.location.href = "/";
        });
    }
}