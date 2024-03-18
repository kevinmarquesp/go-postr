#!/usr/bin/env python3

import os
import re
from requests import get
from json import loads
from sys import argv
from argparse import ArgumentParser

DUMMYUSERS_API = "https://jsonplaceholder.org/users"
DOTENV_FILE = ".env"


def insert_user_credentials(user):
    pass


def main(args) -> None:
    r = get(DUMMYUSERS_API)
    _ = loads(r.text)  # users

    print(args)


def parse_args(user_args):
    parser = ArgumentParser(description="insert dummy users information in a\
                            postgres database connection, by default it will\
                            read the environment variables if any argument was\
                            specified")

    parser.add_argument("-u", "--username", type=str,
                        default=os.getenv("POSTGRES_USER"),
                        help="username credential info, to login")

    parser.add_argument("-p", "--password",
                        type=str, default=os.getenv("POSTGRES_PASSWORD"),
                        help="password credential info")

    parser.add_argument("-P", "--port", type=int,
                        default=os.getenv("POSTGRES_PORT"),
                        help="database connection port number")

    parser.add_argument("-H", "--host", type=str,
                        default=os.getenv("POSTGRES_HOST"),
                        help="hostname of the database server, use 'localhost'\
                        for development environment")

    parser.add_argument("-d", "--database", type=str,
                        default=os.getenv("POSTGRES_DB"),
                        help="default database name, where the application\
                        information is stored")

    return parser.parse_args(user_args)


def load_dotenv(dotenv_file):
    dotenv = []

    with open(dotenv_file, "r") as file:
        dotenv = re.findall(r"^\s*([^#\s][^=]+)\s*=\s*(.*)\s*$", file.read(),
                            flags=re.MULTILINE)

    for key, value in dotenv:
        os.environ[key] = value


if __name__ == "__main__":
    load_dotenv(DOTENV_FILE)

    args = parse_args(argv[1:])

    main(args)
