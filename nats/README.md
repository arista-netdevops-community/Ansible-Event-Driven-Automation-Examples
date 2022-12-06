# This is a demo of the nats source plugin with EDA 

![Alt text](natsoverall.jpg?raw=true "Overall")

## Summary
This is a simple demo where we are going to down an interface.  As we down the interface we want to kick off a playbook just for testing purposes.  This is one of the primary reasons for EDA.  An event happens then we want to react to it.

So the idea is that there is a gNMIC process running.  It is streaming telemetry for a interface Ethernet1.  When Ethernet1 goes into a "DOWN" state.  The response-playbook is ran.

This is all possible given that gNMIC will take the data from gNMI and output it within the nats message bus.

## Explanation of file structure
```
├── gnmic.yml //gNMIC configuration file that references the gNMI path and output for nats
├── inventory.yml //Inventory for ansible simple localhost
├── nats-eda.py // nats EDA plugin created for this
├── nats-example.yml // the ansible-rulebook created for this demo with a condition
├── natsoverall.jpg 
├── README.md
├── response-playbook.yml // Playbook that is ran once the config has been met.
└── switchcfg.cfg // The switch configuration
```

## Demo needs

#### Run nats 
```
docker pull nats:latest
docker run -p 4222:4222 -dit nats:latest
```

#### To run the eda source plugin you will need nats python library
```
pip install nats-py
```

#### gnmic is necessary
Install [gNMIC]("https://gnmic.kmrd.dev/#installation") if it is not installed already

```
gnmic subscribe --config $PWD/gnmic.yml
```

#### Subscribe to the nats topic if necessary for debugging purposes 
```
./natssub or go run main.go 
```

Response

```
2022/12/06 08:59:58 Listening on [test]
2022/12/06 09:00:01 [#1] Received on [test]: '[{"name":"sub1","timestamp":1670273612118710138,"tags":{"interface_name":"Ethernet1","source":"172.20.20.3","subscription-name":"sub1"},"values":{"/interfaces/interface/state/admin-status":"UP"}}]'
```

#### Run the rule-book
```
ansible-rulebook --rulebook nats-example.yml -S $PWD -i inventory.yml --verbos --debug
```
Response

```
DEBUG:asyncio:Using selector: EpollSelector
DEBUG:ansible_rulebook.app:Loading rules from the file system nats-example.yml
INFO:ansible_rulebook.app:Starting sources
INFO:ansible_rulebook.app:Starting rules
INFO:ansible_rulebook.engine:run_ruleset
SLF4J: Failed to load class "org.slf4j.impl.StaticLoggerBinder".
SLF4J: Defaulting to no-operation (NOP) logger implementation
SLF4J: See http://www.slf4j.org/codes.html#StaticLoggerBinder for further details.
INFO:ansible_rulebook.engine:ruleset define: {"name": "Listen for events on nats", "hosts": ["all"], "sources": [{"EventSource": {"name": "nats-eda", "source_name": "nats-eda", "source_args": {"subject": "test", "host": "127.0.0.1", "port": 4222}, "source_filters": []}}], "rules": [{"Rule": {"name": "check down interface", "condition": {"AllCondition": [{"EqualsExpression": {"lhs": {"Event": "values._interfaces_interface_state_admin_status"}, "rhs": {"String": "DOWN"}}}]}, "action": {"Action": {"action": "run_playbook", "action_args": {"name": "response-playbook.yml"}}}, "enabled": true}}]}
INFO:ansible_rulebook.engine:load source
INFO:ansible_rulebook.engine:load source filters
INFO:ansible_rulebook.engine:Calling main in nats-eda
INFO:ansible_rulebook.engine:Waiting for event from Listen for events on nats
```
This will sit and listen at this point for all nats events. 

### On the switch
Log into the switch and simply down the interface.
```
ceos2>en
ceos2#conf t
ceos2(config)#int eth1
ceos2(config-if-Et1)#shutdown 
ceos2(config-if-Et1)#
```

The running subscribe debugger
```
2022/12/06 09:03:51 [#5] Received on [test]: '[{"name":"sub1","timestamp":1670335407832866095,"tags":{"interface_name":"Ethernet1","source":"172.20.20.3","subscription-name":"sub1"},"values":{"/interfaces/interface/state/admin-status":"DOWN"}}]'
```

The ansible rule-book response

```
INFO:ansible_rulebook.builtin:ruleset: Listen for events on nats, rule: check down interface
INFO:ansible_rulebook.builtin:Calling Ansible runner

PLAY [localhost] ***************************************************************

TASK [Gathering Facts] *********************************************************
ok: [localhost]

TASK [debug] *******************************************************************
ok: [localhost] => {
    "msg": "You have a down interface get a fix going!"
}

PLAY RECAP *********************************************************************
localhost                  : ok=2    changed=0    unreachable=0    failed=0    skipped=0    rescued=0 
```