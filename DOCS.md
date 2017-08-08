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
    SSH_HOST: "foo.com"
    SSH_USERNAME: "root"
    SSH_KEY: ${SSH_PRIVATE_KEY}
  script:
    - echo Hello
    - echo World    

```

Example configuration in your `.gitlab-ci.yml` file for multiple hosts using user and password:

```diff
pipeline:
  image: bitlogicos/gitlab-ssh
  stage: deploy
  variables:
    SSH_HOST:
  +  - foo.com
  +  - bar.com
    SSH_USERNAME: "root"
    SSH_PASSWORD: 1234
    SSH_PORT: 22
  script:
    - echo hello
    - echo world
```


Example configuration for command timeout (unit: second), default value is 60 seconds:

```diff
pipeline:
  image: bitlogicos/gitlab-ssh
  stage: deploy
  variables:
    SSH_HOST: "foo.com"
    SSH_USERNAME: "root"
  + SSH_TIMEOUT: 10
  script:
    - echo hello
    - echo world
```

Example configuration for execute commands on a remote server using ｀SSHProxyCommand｀:

```diff
pipeline:
  image: bitlogicos/gitlab-ssh
  stage: deploy
  variables:
    SSH_HOST: "foo.com"
    SSH_USERNAME: "root"
    SSH_PASSWORD: 1234
+   PROXY_SSH_HOST: 10.130.33.145
+   PROXY_SSH_USER: ubuntu
+   PROXY_SSH_PORT: 22
+   PROXY_SSH_KEY: ${PROXY_KEY}
  script:
    - echo hello
    - echo world
```


# Parameter Reference

ssh_host
: target hostname or IP

ssh_port
: ssh port of target host

ssh_username
: account for target host user

ssh_password
: password for target host user

ssh_key
: plain text of user private key

ssh_key_path
: key path of user private key

ssh_timeout
: Timeout is the maximum amount of time for the TCP connection to establish.

ssh_command_timeout
: Command timeout is the maximum amount of time for the execute commands, default is 60 secs.

proxy_ssh_host
: proxy hostname or IP

proxy_ssh_port
: ssh port of proxy host

proxy_ssh_username
: account for proxy host user

proxy_ssh_password
: password for proxy host user

proxy_key
: plain text of proxy private key

proxy_key_path
: key path of proxy private key
