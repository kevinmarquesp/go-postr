from sys import argv
from argparse import ArgumentParser
import os
import re
from icecream import ic
from psycopg2 import connect, errors

TESTDBS_WILDCARD = "test_%"


def delete_testdb(conn, curs):
    curs.execute("SELECT datname FROM pg_database WHERE datistemplate = false\
                 AND datname LIKE %s;", (TESTDBS_WILDCARD,))

    testdb_list = curs.fetchall()

    for testdb_res in testdb_list:
        testdb = testdb_res[0]

        # kills the db process if it is open and then drop it

        curs.execute(f"SELECT pg_terminate_backend(pg_stat_activity.pid) FROM\
                     pg_stat_activity WHERE pg_stat_activity.datname =\
                     '{testdb}' AND pid <> pg_backend_pid();")
        curs.execute(f"DROP DATABASE {testdb};")

        info = f"Dopped {testdb} table with success!"

        ic(info)


# the with keyword creates a transaction block, wich breakes the drop query

def main(args):
    try:
        conn = connect(host=args.host, port=args.port, user=args.username,
                       password=args.password, database=args.database)
        conn.autocommit = True
        curs = conn.cursor()

        delete_testdb(conn, curs)

    except Exception as error:
        ic(error)

    finally:
        curs.close()
        conn.close()


def parse_args(uargs):
    parser = ArgumentParser(description="deletes all test database based on\
                            ther name, will delete every database that starts\
                            with the 'test_' prefix")

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
                        default="postgres",
                        help="database that has access to the entire posgres\
                        system and other databas")

    return parser.parse_args(uargs)


def load_dotenv(dotenv_file):
    dotenv = []

    with open(dotenv_file, "r") as file:
        dotenv = re.findall(r"^\s*([^#\s][^=]+)\s*=\s*(.*)\s*$", file.read(),
                            flags=re.MULTILINE)

    for key, value in dotenv:
        os.environ[key] = value


DOTENV_FILE = ".env"

if __name__ == "__main__":
    load_dotenv(DOTENV_FILE)

    args = parse_args(argv[1:])

    main(args)
