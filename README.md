# HW1 Ansible
## Plays
1. Copy Binary
```
$ ansible-playbook hellosjsu_deploy_ansibleplaybook.yml --tags "deploy"
```

2. Start Hello World Server 
```
$ ansible-playbook hellosjsu_deploy_ansibleplaybook.yml --tags "run"
```

3. Stop Hello World Service
```
$ ansible-playbook hellosjsu_deploy_ansibleplaybook.yml --tags "stop"
```

4. Remove Binary
```
$ ansible-playbook hellosjsu_deploy_ansibleplaybook.yml --tags "remove"
```
