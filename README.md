# ssh_magic

Small utility to organize access to the different environments

## Input

`ssh_magic` needs a file `environments.json` with a list of all environments like this:

```json
{
  "env1": {
    "user": "foo",
    "ip": "1.2.3.4",
    "key": "/root/foo/bar.key"
  },
  "env2": {
    "user": "bar",
    "ip": "5.6.7.8",
    "key": "/root/foo/bar2.key"
  },
  "env3": {
    "user": "john-doe",
    "ip": "9.10.11.12",
    "key": "/root/john/doe.key"
  }
}
```

## Build

For compile all binaries, execute:

```bash
./build.py
```

If you only need one/several of then, you have to execute:

```bash
./build.py -e env1,env3
```

For get a list of available environments:

```bash
./build.py -l
```
