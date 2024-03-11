# taskrun

Simple cli that helps to create jobs from configurable commands for different stages.

## Installation

```bash
go install github.com/Drafteame/taskrun@latest
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
        # it will be omitted.
        #
        # Remote is the first to be loaded, so if there is a conflict with the vars, the remote value will be
        # overwritten.
        remote:
          source: secretsmanager
          key: secret-name
          
        # This vars will be loaded after the remote. The resolve order for var sources is env, ssm and static.
        # If there are dependencies between vars, these are resolved at the end by replacing the collected values
        # and resolving recursively all missing vars at that moment.
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

### Stages

Stages are a way to configure different behavior for different environments. For example, you can have a stage for 
development, another for testing, and another for production.

You should define at least one stage in the job file, so it can be executed as the default stage.

If you have more than one stage defined, you can execute a specific stage using the `-s` flag.

```bash
taskrun print job-name -s stage1
```

When you have defined more than one stage, if you don't specify the stage, the default stage will be taken as the first 
stage defined:

```yaml
jobs:
  stage1:
    - name: some-job
      command: echo
      args:
        - '"Hello, World! default"'

  stage2:
    - name: some-job
      command: echo
      args:
            - '"Hello, World!"'
```

```bash
taskrun print same-job

# This will execute the stage1
```

If you want to define a different stage as the default stage, you can use the `default_stage` configuration.

```yaml
default_stage: stage2

jobs:
  stage1:
    - name: some-job
      command: echo
      args:
        - '"Hello, World! default"'

  stage2:
    - name: some-job
      command: echo
      args:
        - '"Hello, World!"'
```

```bash
taskrun print same-job

# This will execute the stage2
```


## Templating

Jobs file supports variable replacing using notation similar to serverless framework.

```yaml
jobs:
  stage1:
    - name: some-job
      command: echo
      args:
        - '"${env:MY_ENV_VAR} ${sys:cwd}"'
      env:
        vars:
          MY_ENV_VAR:
            value: "Hello from"
```

### Supported replacing vars

| Notation            | Description                                                           |
|---------------------|-----------------------------------------------------------------------|
| `${env:<some-env>}` | Replace the value of the environment variable after source resolution |
| `${sys:cwd}`        | It sets the value of the current working directory as absolute path   |
| `${self:stage}`     | Sets the current selected stage and replace it on templates           |