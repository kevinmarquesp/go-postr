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
from lorem import get_sentence

DUMMYUSERS_API = "https://jsonplaceholder.typicode.com/users"
DOTENV_FILE = ".env"

MIN_SENTENCES = 3
MAX_SENTENCES = 7
MAX_ARTICLES = 15


def insert_new_user(conn, curs, username, password, bio):
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
    username = user["username"]
    password = "dummypassword123"

    name = user["name"]
    company_name = user["company"]["name"]
    email = user["email"]

    bio = f"I'm {name}. Working for {company_name}\n{email}"

    return username, password, bio


def reset_db(conn, curs):
    curs.execute("DELETE FROM relationship")
    curs.execute("DELETE FROM article")
    curs.execute('DELETE FROM "user"')
    conn.commit()


def get_users_len(conn, curs):
    curs.execute('SELECT (COUNT(*)) AS row_count FROM "user"')

    return curs.fetchone()[0]


def get_users_ids(conn, curs):
    curs.execute('SELECT (id) FROM "user"')

    return [row[0] for row in curs.fetchall()]


def insert_relationship(conn, curs, follower_id, followed_ids):
    for followed_id in followed_ids:
        try:
            curs.execute("INSERT INTO relationship VALUES (%s, %s)",
                         (follower_id, followed_id))

        except errors.CheckViolation:
            print(f"[ERRO]: Invalid relation: {follower_id} to {followed_id}")

        except errors.UniqueViolation:
            print(f"[WARN]: {follower_id} to {followed_id} already inserted")

        conn.commit()


def insert_articles(conn, curs, user_id, articles):
    curr_time = datetime.now()

    for art in articles:
        try:
            curs.execute("INSERT INTO article (content, user_id, created_at,"
                         "updated_at) VALUES (%s, %s, %s, %s)",
                         (art, user_id, curr_time, curr_time))
            print(f"[INFO]: Insert an {len(art)} word article to {user_id}")

        except Exception as err:
            print(f"[ERRO]: Couldn't insert article to {user_id} ID, {err}")

        conn.commit()


def insert_data(conn, curs, users):
    reset_db(conn, curs)

    for user in users:
        username, password, bio = get_login_info(user)

        insert_new_user(conn, curs, username, password, bio)

    users_len = get_users_len(conn, curs)
    ids = get_users_ids(conn, curs)

    for follower_id in ids:
        followed_ids = r.choices(ids, k=r.randint(0, users_len))

        insert_relationship(conn, curs, follower_id, followed_ids)

    for user_id in ids:
        articles = [get_sentence(count=r.randint(MIN_SENTENCES, MAX_SENTENCES))
                    for _ in range(r.randint(0, MAX_ARTICLES))]

        insert_articles(conn, curs, user_id, articles)


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
                            specified; also, it will wipe out any data on the\
                            database before inserting the dummy users and the\
                            dummy users content (relationships and articles)")

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
