// Função para redirecionar para a visualização da nota
export function redirectToNoteView(event) {
    const note = event.target.closest(".note");
    if (note) {
        const id = note.id;
        window.location.href = "note/" + id;
    }
}

// Configurar evento no contêiner das notas
export function setupRedirectToNoteView() {
    const container = document.querySelector(".notes-container");
    if (container) {
        container.addEventListener("click", redirectToNoteView);
    }
}
