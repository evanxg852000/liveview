# Diffuser

# CLI

Structure based on https://github.com/golang-standards/project-layout

Initialise db, creates users & display credentials (api_key, api_secret)
```bash
diffuser init --data-path ./data --config path/to/config/file.toml
```

starts server process
```bash
diffuser run --config path/to/config/file.toml
```

relead config by creating object and destroying old ones 
```bash
diffuser reload --config path/to/config/file.toml
```

Actor Framework:
https://proto.actor/

https://graphviz.org/doc/info/lang.html


env_var support: https://www.elastic.co/guide/en/beats/winlogbeat/current/using-environ-vars.html


```yaml

groups = [ 'root', 'admin', 'api' ]


[[users]]
name = "jhone.doe"
email = "jhone@mail.com"
meta = { phone = "+0465352426"}
api_key = "only for api users (auto generated on init)"
api_secret = "only for api users"
groups = [ "main", "api"]

[[users]]
name = "jane.doe"
email = "jane@mail.com"
meta = { phone = "+0253444545"}
api_key = "key"
api_secret = "secret"
groups = [ "api"]
```

TODO:
ci
read_config
create_tasks got github
