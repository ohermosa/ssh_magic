#!/usr/bin/env python

import os
import json
import argparse
import subprocess

class bcolors:
    HEADER = '\033[95m'
    OKBLUE = '\033[94m'
    OKGREEN = '\033[92m'
    WARNING = '\033[93m'
    FAIL = '\033[91m'
    ENDC = '\033[0m'
    BOLD = '\033[1m'
    UNDERLINE = '\033[4m'


def read_json(configfile):
    with open(configfile, 'r') as f:
        return json.load(f)


def write_json(configfile, data):
    with open(configfile, 'w') as f:
        json.dump(data, f)


environments = read_json("environments.json")
release = read_json("release.json")


def parse_input(input):
    output = {}
    for item in input.split(","):
        if item in environments:
            output[item] = environments[item]
        else:
            print("{}'{}' is not an available environment{}".format(bcolors.FAIL,
                                                                    item,
                                                                    bcolors.ENDC))
    return output


def build_binary(envs={}):
    sshmagic_version = release["version"]
    go_version = release["go_version"]

    for k, v in envs.items():
        environment = k
        ip = v["ip"]
        key = v["key"]
        user = v["user"]
        command = "go build -o bin/{} -ldflags \"-X 'main.SSHMagicVersion={}' -X 'main.GoVersion={}' -X 'main.Environment={}' -X 'main.EnvironmentIP={}' -X 'main.User={}' -X 'main.Key={}'\""
        os.popen(command.format(environment,
                                sshmagic_version,
                                go_version,
                                environment,
                                ip,
                                user,
                                key))
        print("{}'{}' binary has been created{}".format(bcolors.OKBLUE, k, bcolors.ENDC))


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("-e",
                        "--environments",
                        type=str,
                        help="List of environments (comma separated), for binary creation")
    parser.add_argument("-l",
                        "--list",
                        help="Show available environments", action="store_true")
    args = parser.parse_args()
    args = vars(args)

    if args["environments"]:
        list_of_envs = parse_input(args["environments"])
        build_binary(list_of_envs)
    elif args["list"]:
        print("{}Available environments{}:".format(bcolors.WARNING, bcolors.ENDC))
        for item in environments:
            print("\t- {}".format(item))
    else:
        build_binary(environments)
