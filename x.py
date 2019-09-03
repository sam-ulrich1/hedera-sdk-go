#!/usr/bin/env python3
import os, re, argparse, sys
from subprocess import Popen, check_call, PIPE, check_output, CalledProcessError
from shutil import copy2, copytree, rmtree


def sh(command, silent=False, cwd=None, shell=True, env=None):
    if env is not None:
        env = dict(**os.environ, **env)

    if silent:
        p = Popen(
            command, shell=shell, stdout=PIPE, stderr=PIPE, cwd=cwd, env=env)
        stdout, _ = p.communicate()

        return stdout

    else:
        check_call(command, shell=shell, cwd=cwd, env=env)


def update_submodules():
    print(" :: update hedera-sdk-c")

    # Ensure the submodule is initialized
    sh("git submodule update --init", silent=True)

    # Fetch upstream changes
    sh("git submodule foreach git fetch", silent=True)

    # Reset to upstream
    sh("git submodule foreach git reset --hard origin/HEAD", silent=True)

    # Update include/
    rmtree("include")
    copytree("vendor/hedera-sdk-c/include", "include")


def get_default_target():
    targets = sh("rustup target list", silent=True).decode()
    m = re.search(r"(.*?)\s*\(default\)", targets)
    t = m[1]

    if t == "x86_64-unknown-linux-gnu":
        t = "x86_64-unknown-linux-musl"
    elif t == "x86_64":
        t == "x86_64-apple-darwin"

    return t


def build(release=False):
    default_target = get_default_target()
    targets = [
        "x86_64-apple-darwin", "x86_64-pc-windows-gnu",
        "x86_64-unknown-linux-musl"
    ]

    prefix = {
        "x86_64-apple-darwin": "",
        "x86_64-pc-windows-gnu": "x86_64-w64-mingw32-",
        "x86_64-unknown-linux-musl": "x86_64-linux-musl-"
    }

    artifact = {
        "x86_64-apple-darwin": "libhedera.a",
        "x86_64-pc-windows-gnu": "hedera.lib",
        "x86_64-unknown-linux-musl": "libhedera.a"
    }

    if release:
        for target in targets:
            print(f" :: build hedera-sdk-c for {target}")

            if target != default_target:
                sh(["rustup", "target", "add", target],
                   shell=False,
                   silent=True,
                   cwd="vendor/hedera-sdk-c")

            if 'musl' in target:
                cc = 'musl-gcc'
            else:
                cc = f"{prefix[target]}gcc"

            profile = "--release" if release else ''
            sh(f"cargo build --target {target} {profile}",
               cwd="vendor/hedera-sdk-c",
               env={
                   "CC": cc,
                   "CARGO_TARGET_X86_64_UNKNOWN_LINUX_MUSL_LINKER": cc,
                   "CARGO_TARGET_X86_64_PC_WINDOWS_GNU_LINKER": cc,
               })

            if target.endswith("-apple-darwin"):
                sh(f"strip -Sx {artifact[target]}",
                   cwd=f"vendor/hedera-sdk-c/target/{target}/release", silent=True)

            elif target.endswith("-musl"):
                sh(f"strip --strip-unneeded -d -x {artifact[target]}",
                   cwd=f"vendor/hedera-sdk-c/target/{target}/release")

            else:
                sh(f"{prefix[target]}strip --strip-unneeded -d -x {artifact[target]}",
                   cwd=f"vendor/hedera-sdk-c/target/{target}/release")

            copy2(f"vendor/hedera-sdk-c/target/{target}/release/{artifact[target]}", f"libs/{target}/")

    else:
        target = default_target

        # For development; build only the _default_ target
        print(f" :: build hedera-sdk-c for {target}")
        sh(f"cargo build --target {target}", cwd="vendor/hedera-sdk-c")

        # Copy _default_ lib over
        copy2(f"vendor/hedera-sdk-c/target/{target}/debug/{artifact[target]}", f"libs/{target}/")


def commit():
    sha = sh("git rev-parse --short HEAD", cwd="vendor/hedera-sdk-c", silent=True).decode().strip()
    sh("git add ./vendor/hedera-sdk-c ./libs ./include")

    try:
        sh(f"git commit -m \"build libs/ and sync include/ from hedera-sdk-c#{sha}\"")
        sh("git push")

    except CalledProcessError:
        # Commit likely failed because there was nothing to commit
        pass


parser = argparse.ArgumentParser()
parser.add_argument(
    "-R", "--release", help="build in release mode", action="store_true")
parser.add_argument(
    "-C", "--commit", help="commit include/ and libs/", action="store_true")

argv = parser.parse_args(sys.argv[1:])

update_submodules()
build(release=argv.release)

if argv.commit and argv.release:
    commit()
