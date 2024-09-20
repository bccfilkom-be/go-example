# minimum-setup-with-ansible

## Ansible

### Run playbook

```bash
ansible-playbook -i inventory.yaml --vault-password-file {password-file-path} playbooks/{playbook-name}.yaml
```

### Create Vault

```bash
ansible-vault encrypt_string --vault-password-file {password-file-path} 'string' --name 'password'
```

### Ping

```bash
ansible -i inventory.yaml --vault-password-file {password-file-path} vm-demo -m ping
```
