# Command
docs link: https://openapi-generator.tech/

Run this in your terminal (replace out-open-api-example with your own OpenAPI yaml file)

```
# Exmaple
docker run --rm \
    -v $PWD:/local openapitools/openapi-generator-cli generate \
    -i /local/open-api-example.yaml \
    -g go \
    -o /local/out-open-api-example/
```

# Permissions
Since the file is being created by docker, it might have a different user id. You may need to run `sudo chown -R $USER:$USER out-open-api-example`
