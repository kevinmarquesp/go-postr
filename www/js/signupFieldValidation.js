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

const BS_DANGER = "danger";
const BS_WARNING = "warning";
const BS_SUCCESS = "success";

const BS_ERROR_ICON = "bi-x-circle-fill";
const BS_WARNNING_ICON = "bi-exclamation-triangle-fill";
const BS_VALID_ICON = "bi-check-circle-fill";

const usernameStatus = {
	alreadyTaken: new StatusBox(BS_DANGER, BS_ERROR_ICON, "This username was already taken"),
	validUsername: new StatusBox(BS_SUCCESS, BS_VALID_ICON, "This username is valid")
};

const passwordStatus = {
	tooShort: new StatusBox(BS_DANGER, BS_ERROR_ICON, "You're password is too weak, just like you"),
	couldBeBetter: new StatusBox(BS_WARNING, BS_WARNNING_ICON, "C'mon, you could do better than this"),
	goodEnough: new StatusBox(BS_SUCCESS, BS_VALID_ICON, "Meh... Good enough..."),
};

const $userStatusBox = document.querySelector("#UsernameStatus");
const $passStatusBox = document.querySelector("#PasswordStatus");

// usernameStatus.alreadyTaken.render($userStatusBox);
// usernameStatus.validUsername.render($userStatusBox);

// passwordStatus.tooShort.render($passStatusBox);
// passwordStatus.couldBeBetter.render($passStatusBox);
// passwordStatus.goodEnough.render($passStatusBox);
