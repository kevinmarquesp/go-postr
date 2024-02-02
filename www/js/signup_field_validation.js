const $template = document.createElement("div");

const FIELD_VALIDATION_COMPONENT_PATH = "/component/FieldValidationStatus";
const USERNAME_SERVER_VALIDATION_PATH = "/auth/validate/username";

const BOOTSTRAP_STATUS_PLACEHOLDER = "BOOTSTRAP_STATUS_PLACEHOLDER";
const MESSAGE_PLACEHOLDER = "MESSAGE_PLACEHOLDER";
const USERNAME_SERVER_VALIDATION_DELAY_MS = 1000;

let isPasswordValid = false;
let isUsernameValid = false;

function writeStatusBoxHtmlElement($statusBox, bootstrapStatus, message) {
	if ($template.innerHTML.length <= 0) {
		console.error("Could not parse the status box element because the $template is invalid")
		return;
	}

	const statusBoxNewContent = $template.innerHTML
		.replace(new RegExp(BOOTSTRAP_STATUS_PLACEHOLDER, "g"), bootstrapStatus)
		.replace(new RegExp(MESSAGE_PLACEHOLDER, "g"), message);

	$statusBox.innerHTML = statusBoxNewContent;
}

function getUsernameValidationStatus(username) {
	if (username.indexOf(" ") !== -1)
		return "has_space";
	else if (!/^[A-Za-z0-9-_]+$/.test(username) && username.length)
		return "invalid_character";
	else if (username.length === 0)
		return "is_empty";

	return "is_valid";
}

function handleUsernameFieldValidation() {
	const $usernameInputField = document.querySelector("#username");
	const $usernameStatusBox = document.querySelector("#username_status");

	let usernameTimeout;

	$usernameInputField.onkeyup = () => {
		const validationStatus = getUsernameValidationStatus($usernameInputField.value.trim());

		isUsernameValid = false;
		clearTimeout(usernameTimeout);
		
		switch (validationStatus) {
			case "has_space":
				writeStatusBoxHtmlElement($usernameStatusBox, "warning", "Space characters is not allowed");
				break;

			case "invalid_character":
				writeStatusBoxHtmlElement($usernameStatusBox, "warning", "The only special characters allowed is - and _");
				break;

			case "is_valid":  //an htmx:afterRequest event listener will check if the return status was valid or not
				usernameTimeout = setTimeout(() => {
					const UsernameServerValidation = new Event("username-server-validation");
					$usernameInputField.dispatchEvent(UsernameServerValidation);
				}, USERNAME_SERVER_VALIDATION_DELAY_MS);
				break;

			default:
				$usernameStatusBox.innerHTML = "";
		}
	};
}

function getPasswordValidationStatus(password) {
	const length = password.length;
		
	if (length == 0)
		return "empty";
	else if (length < 8)
		return "weak";
	else if (length >= 8 && length <= 12)
		return "good";

	return "is_valid";
}

function handlePasswordFieldValidation() {
	const $passwordInputField = document.querySelector("#password");
	const $passwordStatusBox = document.querySelector("#password_status");

	$passwordInputField.onkeyup = () => {
		const validationStatus = getPasswordValidationStatus($passwordInputField.value);
		isPasswordValid = false;

		switch (validationStatus) {
			case "weak":
				writeStatusBoxHtmlElement($passwordStatusBox, "danger", "Too weak, just like you...");
				break;

			case "good":
				writeStatusBoxHtmlElement($passwordStatusBox, "warning", "Comon, you could do better!");
				break;

			case "is_valid":
				isPasswordValid = true;
				writeStatusBoxHtmlElement($passwordStatusBox, "success", "Eh. Good enough.");
				break;

			default:
				$passwordStatusBox.innerHTML = "";
		}
	};
}

htmx.ajax("POST", FIELD_VALIDATION_COMPONENT_PATH, {
	target: $template,
	values: {
		"bootstrap-status": BOOTSTRAP_STATUS_PLACEHOLDER,
		"message": MESSAGE_PLACEHOLDER
	}
}).then(() => {
	handleUsernameFieldValidation();
	handlePasswordFieldValidation();
});

document.body.addEventListener("htmx:afterRequest", (event) => {
	if (event.detail.pathInfo.requestPath !== USERNAME_SERVER_VALIDATION_PATH)
		return;

	if (event.detail.xhr.status !== 200)
		console.log("invalid");
	else
		console.log("valid");
});
