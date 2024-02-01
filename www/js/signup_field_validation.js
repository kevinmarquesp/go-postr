const $usernameInputField = document.querySelector("#username");
const $usernameStatusBox = document.querySelector("#username_status");

const $passwordInputField = document.querySelector("#password");
const $passwordStatusBox = document.querySelector("#password_status");

$usernameInputField.onkeyup = () => {
	const username = $usernameInputField.value.trim();

	if (username.indexOf(" ") !== -1)
		$usernameStatusBox.innerText = "Space characters is not allowed";

	else if (!/^[A-Za-z0-9-_]+$/.test(username) && username.length)
		$usernameStatusBox.innerText = "The only special characters allowed is - and _";

	else if (username.length === 0)
		$usernameStatusBox.innerText = "";

	else
		$usernameInputField.dispatchEvent(new Event("username-server-validation"));
};
