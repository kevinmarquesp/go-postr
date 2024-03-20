#!/usr/bin/env python3

import os
import re
import bcrypt

from argparse import ArgumentParser
from datetime import datetime
from json import loads
from psycopg2 import connect, errors
from requests import get
from sys import argv

DUMMYUSERS_API = "https://jsonplaceholder.org/users"
DOTENV_FILE = ".env"


def insert_new_user(args, username, password, bio):
    """
    Will insert the user (with the specified username, password and bio
    description) to the database. Every user is meant to be unique, on the
    cases that the current username is in conflict with another, it wont do
    nothing at all - Which means that you can run this function multiple times
    but still have the same ammount of inserted users on your database.

    :param argparse.Namespace args:
        This is the arguments given by the command line interface, it should
        have the correct information to connect to the Postgres server.
    :param str username:
        Username to be inserted, should be unique, or else it wont be inserted
        to the database (will not throw any errors).
    :param str password:
        Password for the user, be aware the it will use the bcrypt library to
        hash the password bites array before actually insert it to the table.
    :param str bio:
        Text description about the user, it isn't important at all, but it can
        be useful when developing the front-end page.
    """
    if args is None:
        return

    concatenated_password = username.encode("utf-8") + password.encode("utf-8")
    hashed_password = bcrypt.hashpw(concatenated_password, bcrypt.gensalt())
    curr_time = datetime.now()

    with connect(host=args.host, port=args.port, user=args.username,
                 password=args.password, database=args.database) as conn:
        curs = conn.cursor()

        try:
            curs.execute('INSERT INTO "user" (username, password, bio,\
                         created_at, updated_at) VALUES (%s, %s, %s, %s, %s)',
                         (username, hashed_password.decode("utf-8"), bio,
                          curr_time, curr_time))
            conn.commit()
            print(f"[INFO]: Inserted {username} to '{args.database}.user'")

        except errors.UniqueViolation:
            print(f"[ERRO]: {username} already inserted!")

        curs.close()


def main(args) -> None:
    r = get(DUMMYUSERS_API)
    users = loads(r.text)

    for user in users:
        username = user["login"]["username"]
        password = user["login"]["password"]
        bio = "I'm {} {}. Working for {}\n{}".format(user["firstname"],
                                                     user["lastname"],
                                                     user["company"]["name"],
                                                     user["email"])

        insert_new_user(args, username, password, bio)


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
