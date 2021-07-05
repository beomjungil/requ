<p align="center">
<img src="https://user-images.githubusercontent.com/24666383/124452986-e11af900-ddc1-11eb-9789-70348bdf89df.png" />
</p>

> **🚨 This project is under development. DO NOT USE.**

# requ
> Convenient cli for [IntelliJ HTTP client file](https://www.jetbrains.com/help/idea/exploring-http-syntax.html)

### Usage

```
Usage:
  requ [HTTP file path] [flags]

Flags:
  -h, --help   help for requ
```

#### Basic request

```bash
# Name: test.http
#
# GET https://google.com
#

$ requ ./test.http
```

#### With [variables](https://www.jetbrains.com/help/idea/exploring-http-syntax.html#c259614)

TODO
```bash
$ requ ./test.http --env ./http-client.env.json
```

#### With [environment variables](https://www.jetbrains.com/help/idea/exploring-http-syntax.html#environment-variables)

**TODO**
```bash
# Name: test.http
#
# GET {{host}}
#

$ requ ./test.http --variable {"host": "https://google.com"}
```

or

```bash
$ requ ./test.http
============
 Variables
============

host: **_**
```

#### With multiple requests in .http file

**TODO**
