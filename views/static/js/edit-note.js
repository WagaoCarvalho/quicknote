export function editNote() {
    // Seleciona todos os elementos de cor
    const colorElements = document.querySelectorAll(".color");

    // Adiciona um evento de clique a cada cor
    colorElements.forEach(colorElement => {
        colorElement.addEventListener("click", () => {
            // Remove a classe 'active' de todas as cores
            colorElements.forEach(el => el.classList.remove("active"));

            // Adiciona a classe 'active' ao item clicado
            colorElement.classList.add("active");

            // Atualiza o valor do input oculto
            const colorInput = document.querySelector("#color");
            if (colorInput) {
                colorInput.value = colorElement.dataset.color;
                console.log("Cor selecionada:", colorInput.value); // Debug
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
