const $template = document.createElement("div");

const FIELD_VALIDATION_COMPONENT_PATH = "/component/FieldValidationStatus";
const USERNAME_SERVER_VALIDATION_PATH = "/auth/validate/username";

const BOOTSTRAP_STATUS_PLACEHOLDER = "BOOTSTRAP_STATUS_PLACEHOLDER";
const MESSAGE_PLACEHOLDER = "MESSAGE_PLACEHOLDER";

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

function handleUsernameFieldValidation() {
	const $usernameInputField = document.querySelector("#username");
	const $usernameStatusBox = document.querySelector("#username_status");

	let usernameTimeout;

	$usernameInputField.onkeyup = () => {
		const username = $usernameInputField.value.trim();

		clearTimeout(usernameTimeout);

		if (username.indexOf(" ") !== -1)
			writeStatusBoxHtmlElement($usernameStatusBox,
				"warning", "Space characters is not allowed");
		else if (!/^[A-Za-z0-9-_]+$/.test(username) && username.length)
			writeStatusBoxHtmlElement($usernameStatusBox,
				"warning", "The only special characters allowed is - and _");
		else if (username.length === 0)
			$usernameStatusBox.innerHTML = "";
		else
			usernameTimeout = setTimeout(() => {
				$usernameInputField.dispatchEvent(new Event("username-server-validation"));
			}, 1000);
	};
}

function handlePasswordFieldValidation() {
	const $passwordInputField = document.querySelector("#password");
	const $passwordStatusBox = document.querySelector("#password_status");

	$passwordInputField.onkeyup = () => {
		const length = $passwordInputField.value.length;
			
		if (length == 0)
			$passwordStatusBox.innerHTML = "";
		else if (length < 8)
			writeStatusBoxHtmlElement($passwordStatusBox,
				"danger", "Too weak, just like you...");
		else if (length >= 8 && length <= 12)
			writeStatusBoxHtmlElement($passwordStatusBox,
				"warning", "Comon, you could do better!");
		else
			writeStatusBoxHtmlElement($passwordStatusBox,
				"success", "Eh. Good enough.");
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
