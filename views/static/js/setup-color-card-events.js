// Função para manipular as cores
export function setupColorCardEvents() {
    // Selecionar todas as cores
    const colors = document.querySelectorAll(".color");

    // Adicionar evento de clique a cada elemento com classe "color"
    colors.forEach(color => {
        color.addEventListener("click", function () {
            // Remover a classe "active" de todos os elementos
            colors.forEach(c => c.classList.remove("active"));

            // Adicionar a classe "active" ao elemento clicado
            this.classList.add("active");

            // Atualizar o valor do input hidden com o valor do atributo data-color
            const colorInput = document.getElementById("color");
            if (colorInput) {
                colorInput.value = this.getAttribute("data-color");
            }
        });
    });
}