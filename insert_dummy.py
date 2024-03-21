#!/usr/bin/env python3

import os
import re
import bcrypt
import random as r

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
    """
    Simple wrapper to not polute the code with a bunch of variable definitions
    and values from a dictionary. Will only return the necessary information
    to construct and insert a new user to the database.

    :param dict[] user:
        This user is borrowed from a dummy users API, so this function expects
        that this user object is a JSON converted into a  dictionary.

    :returns tuple[str, str, str]:
        Returns the username, the password (raw) and the autogenerated user bio
        description, all of that inside a tuple!
    """
    username = user["login"]["username"]
    password = user["login"]["password"]

    firstname = user["firstname"]
    lastname = user["lastname"]
    company_name = user["company"]["name"]
    email = user["email"]

    bio = f"I'm {firstname} {lastname}. Working for {company_name}\n{email}"

    return username, password, bio


def reset_db(conn, curs):
    """
    Deletes all users, relationships and post articles from the database. The
    rest of the code will generate all that information again properly only if
    the database is empty.

    :param psycopg2.exensions.connection conn:
        PostgreSQL connection object, will be used to commit the SQL command
        actions/changes.
    :param psycopg2.exensions.cursor curs:
        Cursor to send the commands to the database, this parameter can be
        given from the conn object with conn.cursor().
    """
    curs.execute("DELETE FROM relationship")
    curs.execute("DELETE FROM article")
    curs.execute('DELETE FROM "user"')
    conn.commit()


def get_users_len(conn, curs):
    """
    Yet another wrapper. This will ask the database how much rows the users
    table has, then it will count and return that number for you.

    :param psycopg2.exensions.connection conn:
        PostgreSQL connection object, will be used to commit the SQL command
        actions/changes.
    :param psycopg2.exensions.cursor curs:
        Cursor to send the commands to the database, this parameter can be
        given from the conn object with conn.cursor().

    :returns int:
        The length of the users tables, how much rows it has.
    """
    curs.execute('SELECT (COUNT(*)) AS row_count FROM "user"')

    return curs.fetchone()[0]


def get_users_ids(conn, curs):
    """
    Get the id column of the table "user", it will also convert the the list
    of tuples into a simple list of integers.

    :param psycopg2.exensions.connection conn:
        PostgreSQL connection object, will be used to commit the SQL command
        actions/changes.
    :param psycopg2.exensions.cursor curs:
        Cursor to send the commands to the database, this parameter can be
        given from the conn object with conn.cursor().

    :returns list[int]:
        List of the user IDs, its important to know that it may not starts with
        0 or 1, even if the users table already has ben deleted!
    """
    curs.execute('SELECT (id) FROM "user"')

    return [row[0] for row in curs.fetchall()]


def insert_data(conn, curs, users):
    reset_db(conn, curs)

    for user in users:
        username, password, bio = get_login_info(user)

        insert_new_user(conn, curs, username, password, bio)

    users_len = get_users_len(conn, curs)
    ids = get_users_ids(conn, curs)

    for id in ids:
        followed_ids = r.choices(ids, k=r.randint(0, users_len))

        for followed_id in followed_ids:
            try:
                curs.execute("INSERT INTO relationship VALUES (%s, %s)",
                             (id, followed_id))

            except errors.CheckViolation:
                print(f"[ERRO]: Not valid relationship: {id} to {followed_id}")

            except errors.UniqueViolation:
                print(f"[WARN]: Duplicated {id} to {followed_id} relationship")

            conn.commit()


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
