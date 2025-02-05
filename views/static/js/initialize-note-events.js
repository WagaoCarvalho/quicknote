import { setupColorCardEvents } from './setup-color-card-events.js';
import { setupNeutralButtonEvent } from './setup-neutral-button-event.js';
import { setupRedirectToNoteView } from './redirect-to-note-view.js';
import { deleteNote } from './delete-note.js';
import { editNote } from './edit-note.js';

function initializeNoteEvents() {
    setupColorCardEvents();
    setupNeutralButtonEvent();
    setupRedirectToNoteView();
    editNote();

    // Deleta pelo batão dentro da nota
    document.querySelectorAll(".delete-button").forEach(button => {
        button.addEventListener("click", deleteNote);
    });

    // Deleta pelo link da nota na homepage
    document.querySelectorAll(".note a").forEach(link => {
        link.addEventListener("click", deleteNote); // Passa a referência da função
    });
    
}

document.addEventListener("DOMContentLoaded", initializeNoteEvents);
