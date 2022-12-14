# Ansible Event Driven Automation with Arista EOS

This is a repo for Ansible event-driven automation examples with Arista EOS and CloudVision.

## What is EDA?

Event-Driven Ansible is a new way to enhance and expand automation. It improves IT speed and agility, while enabling consistency and resilience. The Event-Driven Ansible technology was developed by Red Hat and is available as a developer preview. Community input is essential. Since we are building a solution to best meet your needs, we're providing an opportunity for you to advocate for those needs.

Event-Driven Ansible is designed for simplicity and flexibility. By writing an Ansible Rulebook (similar to Ansible Playbooks, but more oriented to "if-then" scenarios) and allowing Event-Driven Ansible to subscribe to an event listening source, your teams can more quickly and easily automate a variety of tasks across the organization. EDA is providing a way of codifying operational logic.

## Current Arista implementation Examples

- NATS

## Installation

### ubuntu

```shell
apt-get --assume-yes install build-essential maven openjdk-17-jdk python3-dev python3-pip
export JDK_HOME=/usr/lib/jvm/java-17-openjdk-amd64
export JAVA_HOME=$JDK_HOME
export PIP_NO_BINARY=jpy
export PATH=$PATH:~/.local/bin
pip3 install -U Jinja2
pip3 install ansible ansible-rulebook ansible-runner wheel
```

### Fedora

```shell
dnf --assumeyes install gcc java-17-openjdk maven python3-devel python3-pip
export JDK_HOME=/usr/lib/jvm/java-17-openjdk
export JAVA_HOME=$JDK_HOME
export PIP_NO_BINARY=jpy
pip3 install -U Jinja2
pip3 install ansible ansible-rulebook ansible-runner wheel
```

All other information such as conditions, rules and alternative installations features can be found within the [readthedocs of EDA]("https://ansible-rulebook.readthedocs.io/en/latest/index.html").
