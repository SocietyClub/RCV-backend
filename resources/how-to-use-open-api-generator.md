# Open API Generator Tutorial

## Command
Documentation: https://openapi-generator.tech/

Run this in your terminal (replace out-open-api-example with your own OpenAPI yaml file)

## Example

> Note: the command above assumes your UID is 1000. We pass the UID into the docker run command so it does not create files as root (Dockers default).

> :warning: DO NOT RUN THIS. This is only for the initial code generation when the first rcvSpec.yaml was created!!

```
docker run --rm \
    -v $PWD:/local \
    -u ${UID}:${UID} \
    openapitools/openapi-generator-cli generate \
        -i /local/resources/rcvSpec.yaml \
        -g go-server \
        -o /local/api/
```

## Permissions
Since the file is being created by docker, it might have a different user id. You may need to run `sudo chown -R $USER:$USER out-open-api-example`
