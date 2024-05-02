# Go Postr

Small web application that I created from scratch in orther to learn more about
some web development concepts that is useful to have when dealing with other
frameworks that does most what I'm doing by hand here for me.

## How to Run

### Quick Result

Git clone this repository and move the `.env.example` file to `.env` - maybe
you would like to change some auth information on this file, but it's not
required just for a quick test:

```bash
git clone https://github.com/kevinmarquesp/go-postr
cd go-postr

cp .env.example .env
```

Then you can just build the application and start both the application and the
Postgres database (only this two, check the *Docker* session for mor
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
