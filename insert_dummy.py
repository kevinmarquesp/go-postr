#!/usr/bin/env python3

import os
import re
from requests import get
from json import loads

DUMMYUSERS_API: str = "https://jsonplaceholder.org/users"
DOTENV_FILE = ".env"


def insert_user_credentials(user):
    pass


def main() -> None:
    r = get(DUMMYUSERS_API)
    _ = loads(r.text)  # users

    print("\n[TODO]: Insert user credentials to database\n")


def load_dotenv(dotenv_file):
    dotenv = []

    with open(dotenv_file, "r") as file:
        dotenv = re.findall(r"^\s*([^#\s][^=]+)\s*=\s*(.*)\s*$", file.read(),
                            flags=re.MULTILINE)

    for key, value in dotenv:
        os.environ[key] = value


if __name__ == "__main__":
    load_dotenv(DOTENV_FILE)

    main()
