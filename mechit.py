#!/usr/bin/env python

import argparse
import datetime
import json
import logging
import os
import sys
from dataclasses import asdict, dataclass

import requests

logger = logging.getLogger(__name__)


def debug_except_hook(type, value, tb):
    print("Converter Error {0}".format(type.__name__))
    print(str(type))
    import pdb
    import traceback

    traceback.print_exception(type, value, tb)
    pdb.post_mortem(tb)


host = os.getenv("MECHSERV_HOST", "http://localhost")
port = os.getenv("MECHSERV_PORT", "9999")
base_url = f"{host}:{port}"
endpoint = "api/v1/mechs"
create_url = f"{base_url}/{endpoint}"

time_string = "%Y-%m-%dT%H:%M:%SZ"


def as_release(year, month, day):
    d = datetime.datetime.now()
    return datetime.datetime(year, month, day, d.hour, d.minute, d.second).strftime(time_string)


@dataclass
class Mech:
    """Class for keeping track of an mechs in inventory."""

    name: str
    version: str
    release: str
    description: str

    def as_dict(self):
        return asdict(self)

    def as_json(self):
        return json.dumps(asdict(self))


default_mechs = [
    Mech(name="simonon", version="1.0.1", release=as_release(1979, 12, 4), description="Mech of Brixton"),
    Mech(name="strummer", version="1.0.2", release=as_release(1977, 4, 8), description="Mech of Somerset"),
    Mech(name="mustaine", version="1.0.3", release=as_release(1985, 6, 12), description="Mech of La Messa"),
    Mech(name="matheos", version="1.0.4", release=as_release(1984, 9, 1), description="Mech of Massachusetts"),
    Mech(name="vera", version="1.0.5", release=as_release(1983, 8, 9), description="Mech of Poughkeepsie"),
    Mech(name="smith", version="0.1.0", release=as_release(1981, 2, 2), description="Mech of Hackney"),
    Mech(name="willis", version="0.2.0", release=as_release(1980, 3, 14), description="Mech of Sheffield"),
]


def get_request(url, hdr=None):
    if hdr is None:
        hdr = {}
    ses = requests.Session()
    req = ses.get(url, headers=hdr)
    req.raise_for_status()
    return req.json()


def post_request(url, data=None, hdr=None):
    if hdr is None:
        hdr = {}
    ses = requests.Session()
    req = ses.post(url, data=data, headers=hdr)
    req.raise_for_status()
    return req.json()


def make_mech(data):
    return post_request(create_url, data, hdr={"Content-Type": "application/json"})


def init_mechs(info):
    results = []
    for mech in default_mechs:
        result = make_mech(data=mech.as_json())
        logger.info("Created : {id} Errors : {error}".format(**result))
        results.append(result)
    return results


def add_mech(info):
    new = Mech(
        name=info.get("name") or "foo",
        version=info.get("version") or "1.0.0",
        release=info.get("release") or as_release(2020, 9, 5),
        description=info.get("description") or "Mech of Foo",
    )
    result = make_mech(data=new.as_json())
    logger.info("Created : {id} Errors : {error}".format(**result))
    return result


class CmdLine(object):
    def __init__(self):

        parser = argparse.ArgumentParser(
            description="Mech Creation Tool",
            usage="""mechit <command> [<args>]

            mechit commands are:
                init          creates default Mechs
                create        creates a new Mech
            """,
        )

        parser.add_argument("command", help="Subcommand to run")
        # parse_args defaults to [1:] for args, but you need to
        # exclude the rest of the args too, or validation will fail
        args = parser.parse_args(sys.argv[1:2])
        if not hasattr(self, args.command):
            logger.error("Unrecognized command")
            parser.print_help()
            exit(1)
        # use dispatch pattern to invoke method with same name
        getattr(self, args.command)()

    def init(self):
        """
        Creates initial mechs
        """
        parser = argparse.ArgumentParser(description="Creates the initial mechs\n")
        parser.add_argument("--debug", dest="debug", action="store_true", default=False, help="Turn debug on")
        args = parser.parse_args(sys.argv[2:])
        debug = args.debug
        level = logging.INFO
        if debug:
            sys.excepthook = debug_except_hook
            level = logging.DEBUG
        logging.basicConfig(stream=sys.stderr, level=level, format="%(asctime)s %(name)s:[%(levelname)s] %(message)s")
        info = vars(args)
        return init_mechs(info)

    def create(self):
        """
        Creates new mechs
        """
        parser = argparse.ArgumentParser(description="Creates a new mechs\n")
        parser.add_argument("--name", dest="name", action="store", default="", help="Name of Mech")
        parser.add_argument(
            "--description", dest="description", action="store", default=None, help="Description of Mech"
        )
        parser.add_argument("--version", dest="version", action="store", default=None, help="Version for the Mech")
        parser.add_argument("--release", dest="release", action="store", default=None, help="Release for the Mech")
        parser.add_argument("--debug", dest="debug", action="store_true", default=False, help="Turn debug on")
        args = parser.parse_args(sys.argv[2:])
        debug = args.debug
        level = logging.INFO
        if debug:
            sys.excepthook = debug_except_hook
            level = logging.DEBUG
        logging.basicConfig(stream=sys.stderr, level=level, format="%(asctime)s %(name)s:[%(levelname)s] %(message)s")
        info = vars(args)
        return add_mech(info)


def main():
    CmdLine()


if __name__ == "__main__":
    sys.exit(CmdLine())
