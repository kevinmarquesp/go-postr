const $template = document.createElement("div");

const BOOTSTRAP_STATUS_PLACEHOLDER = "BOOTSTRAP_STATUS_PLACEHOLDER";
const MESSAGE_PLACEHOLDER = "MESSAGE_PLACEHOLDER";

function writeStatusBox($statusBox, bootstrapStatus, message) {
	if ($template.innerHTML.length <= 0) {
		console.error("Could not parse the status box element because the $template is invalid")
		return;
	}

	const statusBoxNewContent = $template.innerHTML
		.replace(new RegExp(BOOTSTRAP_STATUS_PLACEHOLDER, "g"), bootstrapStatus)
		.replace(new RegExp(MESSAGE_PLACEHOLDER, "g"), message);

	$statusBox.innerHTML = statusBoxNewContent;
}

function handleFormFieldValidation() {
	const $usernameInputField = document.querySelector("#username");
	const $usernameStatusBox = document.querySelector("#username_status");

	// const $passwordInputField = document.querySelector("#password");
	// const $passwordStatusBox = document.querySelector("#password_status");

	$usernameInputField.onkeyup = () => {
		const username = $usernameInputField.value.trim();

		if (username.indexOf(" ") !== -1)
			writeStatusBox($usernameStatusBox, "warning", "Space characters is not allowed");

		else if (!/^[A-Za-z0-9-_]+$/.test(username) && username.length)
			writeStatusBox($usernameStatusBox, "warning", "The only special characters allowed is - and _");

		else if (username.length === 0)
			$usernameStatusBox.innerHTML = "";

		else
			$usernameInputField.dispatchEvent(new Event("username-server-validation"));
	};
}

htmx.ajax("POST", "/component/FieldValidationStatus", {
	target: $template,
	values: {
		"bootstrap-status": BOOTSTRAP_STATUS_PLACEHOLDER,
		"message": MESSAGE_PLACEHOLDER
	}
}).then(handleFormFieldValidation);
