# taskrun

Simple cli that helps to create jobs from configurable commands

## Installation

```bash
go install github.com/Drafteame/taskrun/cmd/taskrun@latest
```

## Usage

```bash
taskrun -h
```

## Job file definition

```yaml
jobs:
  stage1:
    - name: some-job
      command: echo
      args:
        - '"Hello, World!"'
      env:
        # This remote loads the values from the specified source and key. If the source is not specified
        # it will be omitted
        remote:
          source: secretsmanager
          key: secret-name
          
        # This vars will be loaded after the remote, so if there is a conflict the remote value will be overwritten.
        vars:
          # This is a hardcoded value
          SOME_VAR:
            value: "some-value"
            
          # This is a value from SSM
          SOME_SSM_VAR:
            source: ssm
            key: /path/to/ssm/parameter

          # This is a value from environment variable
          SOME_ENV_VAR:
            source: env
            key: SOME_ENV_VAR_NAME
```

## Templating

WORK IN PROGRESS