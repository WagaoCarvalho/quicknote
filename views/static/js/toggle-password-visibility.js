export function togglePasswordVisibility() {
    const checkbox = document.querySelector("input[type='checkbox']");
    const passwordInput = document.querySelector("#password");

    if (checkbox && passwordInput) {
        checkbox.addEventListener("click", function () {
            passwordInput.type = this.checked ? "text" : "password";
        });
    }
}
