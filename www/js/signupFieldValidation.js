"use strict";

const BS_DANGER = "danger";
const BS_WARNING = "warning";
const BS_SUCCESS = "success";

const BS_ERROR_ICON = "bi-x-circle-fill";
const BS_WARNNING_ICON = "bi-exclamation-triangle-fill";
const BS_VALID_ICON = "bi-check-circle-fill";

const VERIFICATION_DELAY = 870;

const $retype = document.querySelector("#retype");
const $password = document.querySelector("#password");
const $username = document.querySelector("#username");
const $submit = document.querySelector("#submit");

const $userStatusBox = document.querySelector("#UsernameStatus");
const $passStatusBox = document.querySelector("#PasswordStatus");
const $retypeStatusBox = document.querySelector("#RetypeStatus");

class StatusBox {
	#$status;
	#bsColor;
	#bsIcon;
	#message;

	constructor($status, color, prefix, message) {
		this.#$status = $status;
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

	render($target = this.#$status, overwrite = true) {
		if (overwrite)
			$target.innerHTML = this.build();
		else
			$target.innerHTML += this.build();
	}
}

const usernameStatus = {
	alreadyTaken: new StatusBox($userStatusBox, BS_DANGER, BS_ERROR_ICON, "This username was already taken"),
	serverError: new StatusBox($userStatusBox, BS_DANGER, BS_ERROR_ICON, "Internal server error, sorry..."),
	hasSpaces: new StatusBox($userStatusBox, BS_WARNING, BS_WARNNING_ICON, "Use - or _ instead of spaces"),
	invalidChars: new StatusBox($userStatusBox, BS_WARNING, BS_WARNNING_ICON, "Invalid characters are not allowed"),
	validUsername: new StatusBox($userStatusBox, BS_SUCCESS, BS_VALID_ICON, "This username is valid"),
};

const passwordStatus = {
	tooShort: new StatusBox($passStatusBox, BS_DANGER, BS_ERROR_ICON, "You're password is too weak, just like you"),
	couldBeBetter: new StatusBox($passStatusBox, BS_WARNING, BS_WARNNING_ICON, "C'mon, you could do better than this"),
	goodEnough: new StatusBox($passStatusBox, BS_SUCCESS, BS_VALID_ICON, "Meh... Good enough..."),
};

const retypeStatus = {
	match: new StatusBox($retypeStatusBox, BS_SUCCESS, BS_VALID_ICON, "You don't seem to have alzeimer, good"),
	missmatch: new StatusBox($retypeStatusBox, BS_DANGER, BS_ERROR_ICON, "Can you type like a human for once?"),
}

let isUsernameValid = false;
let isPasswordValid = false;
let isRetypeValid = false;

let usernameBuffer;
let usernameTimeout;

function updateSubmitButtonState() {
	if (isUsernameValid && isPasswordValid && isRetypeValid)
		$submit.removeAttribute("disabled", "");
	else
		$submit.setAttribute("disabled", "");
}

function validateUsername(username) {
	if (username.indexOf(" ") !== -1)
		return "spaced";

	else if (!/^[A-Za-z0-9-_]+$/.test(username) && username.length)
		return "invalid";

	else if (username.length === 0)
		return "empty";

	return "valid";
}

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

function validateRetype(retype, password) {
	const passwordStatus = validatePassword(password);

	if (retype.length === 0 || password.length === 0)
		return "empty"

	else if (retype !== password || passwordStatus !== "strong")
		return "missmatch";

	return "match";

}

$username.onkeyup = () => {
	const username = $username.value;
	const status = validateUsername(username);

	isUsernameValid = false;

	clearTimeout(usernameTimeout);

	if (usernameBuffer !== username)
		updateSubmitButtonState();

	switch (status) {
		case "spaced":
			usernameStatus.hasSpaces.render();
			break;

		case "invalid":
			usernameStatus.invalidChars.render();
			break;

		case "valid":
			usernameTimeout = setTimeout(() => {
				const xhr = new XMLHttpRequest();

				xhr.onload = () => {
					if (xhr.status === 200) {
						isUsernameValid = true;
						usernameBuffer = username;

						usernameStatus.validUsername.render();
						updateSubmitButtonState();

					} else if (xhr.status === 400) {
						usernameStatus.alreadyTaken.render();

					} else {
						usernameStatus.serverError.render();
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

$password.onkeyup = () => {
	const password = $password.value;
	const status = validatePassword(password);

	isPasswordValid = false;

	updateSubmitButtonState();
	$retype.onkeyup();

	switch (status) {
		case "weak":
			passwordStatus.tooShort.render();
			break;

		case "good":
			passwordStatus.couldBeBetter.render();
			break;

		case "strong":
			isPasswordValid = true;

			passwordStatus.goodEnough.render();
			updateSubmitButtonState();
			break;

		default:
			$passStatusBox.innerHTML = "";
	}
};

$retype.onkeyup = () => {
	const retype = $retype.value;
	const password = $password.value;
	const status = validateRetype(password, retype);

	isRetypeValid = false;

	updateSubmitButtonState();

	switch (status) {
		case "missmatch":
			retypeStatus.missmatch.render();
			break;

		case "match":
			isRetypeValid = true;

			retypeStatus.match.render();
			updateSubmitButtonState();
			break;

		default:
			$retypeStatusBox.innerHTML = "";
	}
};
