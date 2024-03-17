#!/usr/bin/env python3

from requests import get
from json import loads

DUMMYUSERS_API: str = "https://jsonplaceholder.org/users"


def insert_user_credentials(user):
    pass


def main() -> None:
    r = get(DUMMYUSERS_API)
    _ = loads(r.text)  # users

    print("\n[TODO]: Insert user credentials to database\n")


if __name__ == "__main__":
    main()
