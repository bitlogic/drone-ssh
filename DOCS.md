## Examples usages

---
date: 2017-08-08T00:00:00+00:00
title: SSH
author: bitlogic
tags: [ publish, ssh ]
repo: bitlogic/gitlab-ssh
logo: term.svg
image: bitlogicos/gitlab-ssh
---

Use the SSH plugin to execute commands on a remote server. The below pipeline configuration demonstrates simple usage in a .gitlab-ci.yml, in this example the plugin key is provided via gitlab-ci secrets

```yaml

pipeline:
  image: bitlogicos/gitlab-ssh
  stage: deploy
  variables:
    PLUGIN_HOST: "foo.com"
    PLUGIN_USERNAME: "root"
    PLUGIN_KEY: ${SSH_PRIVATE_KEY}
  script:
    - echo Hello
    - echo World    

```

Example configuration in your `.gitlab-ci.yml` file for multiple hosts using user and password:

```diff
pipeline:
  image: bitlogicos/gitlab-ssh
  host:
+  - foo.com
+  - bar.com
  username: root
  password: 1234
  port: 22
  script:
    - echo hello
    - echo world
```


Example configuration for command timeout (unit: second), default value is 60 seconds:

```diff
pipeline:
  image: bitlogicos/gitlab-ssh
  host: foo.com
  username: root
  password: 1234
  port: 22
+ command_timeout: 10
  script:
    - echo hello
    - echo world
```

Example configuration for execute commands on a remote server using ｀SSHProxyCommand｀:

```diff
pipeline:
  image: bitlogicos/gitlab-ssh
  host: foo.com
  username: root
  port: 22
  key: ${DEPLOY_KEY}
  script:
    - echo hello
    - echo world
+ proxy_host: 10.130.33.145
+ proxy_user: ubuntu
+ proxy_port: 22
+ proxy_key: ${PROXY_KEY}
```


# Parameter Reference

host
: target hostname or IP

port
: ssh port of target host

username
: account for target host user

password
: password for target host user

key
: plain text of user private key

key_path
: key path of user private key

script
: execute commands on a remote server

timeout
: Timeout is the maximum amount of time for the TCP connection to establish.

command_timeout
: Command timeout is the maximum amount of time for the execute commands, default is 60 secs.

proxy_host
: proxy hostname or IP

proxy_port
: ssh port of proxy host

proxy_username
: account for proxy host user

proxy_password
: password for proxy host user

proxy_key
: plain text of proxy private key

proxy_key_path
: key path of proxy private key
