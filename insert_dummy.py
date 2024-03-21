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


def insert_new_user(conn, curs, username, password, bio):
    """
    Will insert the user (with the specified username, password and bio
    description) to the database. Every user is meant to be unique, on the
    cases that the current username is in conflict with another, it wont do
    nothing at all - Which means that you can run this function multiple times
    but still have the same ammount of inserted users on your database.

    :param psycopg2.exensions.connection conn:
        Database connection object, it will be used to actually change the
        dtabases's schema.
    :param psycopg2.exensions.cursor curs:
        Connection cursor, this will be use to get information from the
        database and cache the changes to the conn object commit them.
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

    try:
        curs.execute('INSERT INTO "user" (username, password, bio,\
                     created_at, updated_at) VALUES (%s, %s, %s, %s, %s)',
                     (username, hashed_password.decode("utf-8"), bio,
                      curr_time, curr_time))
        print(f"[INFO]: Inserted {username} to '{args.database}.user'")

    except errors.UniqueViolation:
        print(f"[WARN]: {username} already inserted!")

    conn.commit()


def get_login_info(user):
    username = user["login"]["username"]
    password = user["login"]["password"]

    firstname = user["firstname"]
    lastname = user["lastname"]
    company_name = user["company"]["name"]
    email = user["email"]

    bio = f"I'm {firstname} {lastname}. Working for {company_name}\n{email}"

    return username, password, bio


def insert_data(conn, curs, users):
    curs.execute("DELETE FROM relationship")
    curs.execute("DELETE FROM article")
    curs.execute('DELETE FROM "user"')
    conn.commit()

    for user in users:
        username, password, bio = get_login_info(user)

        insert_new_user(conn, curs, username, password, bio)


def main(args) -> None:
    req = get(DUMMYUSERS_API)
    users = loads(req.text)

    try:
        with connect(host=args.host, port=args.port, user=args.username,
                     password=args.password, database=args.database) as conn:
            with conn.cursor() as curs:
                insert_data(conn, curs, users)

    except Exception as err:
        print(f"\033[31m{err}\033[m")


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
