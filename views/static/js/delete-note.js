export function deleteNote(event) {
    event.stopPropagation(); // Impede que o clique na nota seja acionado

    const confirmation = window.confirm("Tem certeza que deseja deletar essa anotação?");
    if (!confirmation) {
        return;
    }

    const noteId = event.currentTarget.dataset.noteid; // Obtém o ID da nota do atributo data-noteid

    fetch("/note/" + noteId, {
        method: "DELETE",
        
    })
    .then(response => {
        if (response.ok) {
            window.location.href = "/"; // Redireciona após exclusão bem-sucedida
        } else {
            alert("Erro ao tentar excluir a anotação.");
        }
    })
    .catch(error => {
        console.error("Erro ao tentar excluir a anotação:", error);
        alert("Erro ao tentar excluir a anotação.");
    });
}
