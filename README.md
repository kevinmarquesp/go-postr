# Go Postr

Small web application that I created from scratch in orther to learn more about
some web development concepts that is useful to have when dealing with other
frameworks that does most what I'm doing by hand here for me.

## How to Run

### Dependencies

Basic dependencies to quickly run the project without any trouble:
*   `git`
*   `docker`
*   `docker-compose`

Development dependencies:
*   `python3`

#### Install Snippets

Ubuntu/Debian:

```bash
sudo apt install git docker docker-compose  # basic
sudo apt install golang python3 python3-dev python3.10-venv postgresql libpq-dev  # development

# Minimal Docker setup.

sudo groupadd docker
sudo usermod -aG docker $USER
sudo service docker restart

su - $USER  # or logout and login again to your user account
```

##### Install Snippets

*   [Install Docker Engine on Debian](https://docs.docker.com/engine/install/debian/)

### Quick Final Result

Git clone this repository and move the `.env.example` file to `.env` - maybe
you would like to change some auth information on this file, but it's not
required just for a quick test:

```bash
git clone https://github.com/kevinmarquesp/go-postr
cd go-postr

cp .env.example .env
```

Then you can just build the application and start both the application and the
Postgres database (only this two, check the *Docker* session for more
information) with:

```bash
docker-compose --profile app up -d
```

> [!NOTE]
> If you don't want to use Docker to run, you can always setup your own
> Postgres server by yourself, just make sure to update the `.env` variable
> values with the settings to connect to your database.
>
> And to run the application, you can compile and run with just `go run .` and
> expect that it will read the `.env` file if the environment variables is not
> defined in your system.

Once everything was built and it's running, open your browser and check the
result at the **localhost:8080** address.

#### Insert Dummy Users

You'll need to have **Python** installed. Start a new Python environment to
install the dependencies locally and run the `helpers/insert_dummy.py` script
with:

```bash
python3 -m venv .venv
source .venv/bin/activate  # this line can vary depending on your shell
python3 -m pip install -r requirements.txt

python3 helpers/insert_dummy.py
```

This script will read the `.env` file to connect to the Postgres server too,
execute it at the root of the project as the code snippet above shows.

If the application is already running, along side with the database, it should
update the home page with the new users content in some seconds.

#### Clean

Once you tested this application, don't forget to clean your system from the
images created by the Docker Compose command with:

```bash
docker-compose --profile app down --rmi all
```

It's good to remove the Python environment too if you don't plan to use it
anymore:

```bash
rm -rf .venv
```

## Todos

Features:
* [ ]   Save a login sessino cokie on the user's browser.
    * [ ]   Change the homepage navbar is the user has an active session.

Development:
* [ ]   Create a `GetRecentArticles()` alternative that returns a Golang's
        structure type.
* [ ]   Remove the `TODO` comments in the code and migrate them to this
        document.
