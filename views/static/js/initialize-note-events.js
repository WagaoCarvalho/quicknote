import { setupColorCardEvents } from './setup-color-card-events.js';
import { setupNeutralButtonEvent } from './setup-neutral-button-event.js';
import { setupRedirectToNoteView } from './redirect-to-note-view.js';
import { deleteNote } from './delete-note.js';
import { editNote } from './edit-note.js';
import { togglePasswordVisibility } from './toggle-password-visibility.js';
import { hideSuccessMessage } from './hide_success_message.js';

function initializeEvents() {
    setupColorCardEvents();
    setupNeutralButtonEvent();
    setupRedirectToNoteView();
    editNote();
    togglePasswordVisibility();
    hideSuccessMessage();
    
    document.querySelectorAll(".delete-button").forEach(button => {
        button.addEventListener("click", deleteNote);
    });
    
}

document.addEventListener("DOMContentLoaded", initializeEvents);
