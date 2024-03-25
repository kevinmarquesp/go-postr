"use strict";

class StatusBox {
	#bsColor;
	#bsIcon;
	#message;

	constructor(color, prefix, message) {
		this.#bsColor = color;
		this.#bsIcon = prefix;
		this.#message = message;
	}

	build() {
		const BS_DISPLAY = "d-block mt-1 mb-3 pe-3 py-1 border-start border-5"
		const BS_COLOR_STYLE = `border-${this.#bsColor} bg-${this.#bsColor}-subtle text-${this.#bsColor}-emphasis`

		return `<span class="${BS_DISPLAY} ${BS_COLOR_STYLE}">
			<i class="mx-3 bi ${this.#bsIcon}"></i>${this.#message}
		</span>`
	}

	render($target, overwrite = true) {
		if (overwrite)
			$target.innerHTML = this.build();
		else
			$target.innerHTML += this.build();
	}
}

const usernameAlreadyTakenError = new StatusBox("danger", "bi-x-circle-fill", "This username was already taken")
const passwordTooShortError = new StatusBox("danger", "bi-x-circle-fill", "You're password is too weak, just like you");
const passwordCouldBeBetterWarning = new StatusBox("warning", "bi-exclamation-triangle-fill", "C'mon, you could do better than this")
const passwordGoodEnoughInfo = new StatusBox("success", "bi-check-circle-fill", "Meh... Good enough...")

const $UsernameStatus = document.querySelector("#UsernameStatus");
const $PasswordStatus = document.querySelector("#PasswordStatus");

// usernameAlreadyTakenError.render($UsernameStatus);

// passwordTooShortError.render($PasswordStatus);
// passwordCouldBeBetterWarning.render($PasswordStatus);
// passwordGoodEnoughInfo.render($PasswordStatus);
