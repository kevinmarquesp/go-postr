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
	serverError: new StatusBox(BS_DANGER, BS_ERROR_ICON, "Internal server error, sorry..."),
	hasSpaces: new StatusBox(BS_WARNING, BS_WARNNING_ICON, "Use - or _ instead of spaces"),
	invalidChars: new StatusBox(BS_WARNING, BS_WARNNING_ICON, "Invalid characters are not allowed"),
	validUsername: new StatusBox(BS_SUCCESS, BS_VALID_ICON, "This username is valid"),
};

const passwordStatus = {
	tooShort: new StatusBox(BS_DANGER, BS_ERROR_ICON, "You're password is too weak, just like you"),
	couldBeBetter: new StatusBox(BS_WARNING, BS_WARNNING_ICON, "C'mon, you could do better than this"),
	goodEnough: new StatusBox(BS_SUCCESS, BS_VALID_ICON, "Meh... Good enough..."),
};

let isPasswordValid = false;
let isUsernameValid = false;

const $username = document.querySelector("#username");
const $userStatusBox = document.querySelector("#UsernameStatus");
const VERIFICATION_DELAY = 870;

let usernameTimeout;

function validateUsername(username) {
	if (username.indexOf(" ") !== -1)
		return "spaced";

	else if (!/^[A-Za-z0-9-_]+$/.test(username) && username.length)
		return "invalid";

	else if (username.length === 0)
		return "empty";

	return "valid";
}

$username.onkeyup = () => {
	const username = $username.value;
	const status = validateUsername(username);

	isUsernameValid = false;

	clearTimeout(usernameTimeout);

	switch (status) {
		case "spaced":
			usernameStatus.hasSpaces.render($userStatusBox);
			break;

		case "invalid":
			usernameStatus.invalidChars.render($userStatusBox);
			break;

		case "valid":
			usernameTimeout = setTimeout(() => {
				const xhr = new XMLHttpRequest();

				xhr.onload = () => {
					if (xhr.status === 200) {
						isUsernameValid = true;

						usernameStatus.validUsername.render($userStatusBox);

					} else if (xhr.status === 400) {
						usernameStatus.alreadyTaken.render($userStatusBox);

					} else {
						usernameStatus.serverError.render($userStatusBox);
					}
				};

				xhr.open("POST", "/validate/username")
				xhr.send(username)

			}, VERIFICATION_DELAY);

			break;

		default:
			$userStatusBox.innerHTML = "";
	}
};

const $password = document.querySelector("#password");
const $passStatusBox = document.querySelector("#PasswordStatus");

function validatePassword(password) {
	const length = password.length;

	if (length == 0)
		return "empty";

	else if (length < 8)
		return "weak";

	else if (length >= 8 && length <= 12)
		return "good";

	return "strong";
}

$password.onkeyup = () => {
	const password = $password.value;
	const status = validatePassword(password);

	isPasswordValid = false;

	switch (status) {
		case "weak":
			passwordStatus.tooShort.render($passStatusBox);
			break;

		case "good":
			passwordStatus.couldBeBetter.render($passStatusBox);
			break;

		case "strong":
			isPasswordValid = true;

			passwordStatus.goodEnough.render($passStatusBox);
			break;

		default:
			$passStatusBox.innerHTML = "";
	}
};
