export function editNote() {
    // Configura o evento de clique para alterar a cor da nota
    const colorElements = document.querySelectorAll(".color");
    colorElements.forEach(colorElement => {
        colorElement.addEventListener("click", () => {
            colorElements.forEach(el => el.classList.remove("active"));
            colorElement.classList.add("active");

            const colorInput = document.querySelector("#color");
            if (colorInput) {
                colorInput.value = colorElement.dataset.color;
            }
        });
    });

    // Configura o evento de clique para o botão "Editar"
    document.querySelectorAll(".edit-button").forEach(button => {
        button.addEventListener("click", event => {
            const noteId = event.currentTarget.getAttribute("data-noteid");
            if (noteId) {
                window.location.href = `/note/${noteId}/edit`;
            } else {
                console.error("Erro: ID da nota não encontrado.");
            }
        });
    });

    // Configura o evento de clique para o botão "Cancelar"
    const cancelButton = document.querySelector(".neutral");
    if (cancelButton) {
        cancelButton.addEventListener("click", () => {
            window.location.href = "/";
        });
    }
}
