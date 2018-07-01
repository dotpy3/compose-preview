# compose-preview

**WIP**

Preview your `docker-compose` environments.

## Usage

Assuming you have a Go environment setup:

```
go get github.com/dotpy3/compose-preview
cd $folder_with_your_docker-compose_files
compose-preview
```

## Example

+ `docker-compose.versions.yml`:
    
    ```yaml
    version: '3'
    services:
      web:
        image: dotpy3/my-web-app:2.0
        env:
          APP_DEBUG: 'true'
    ```

+ `docker-compose.yml`:

    ```yaml
    version: '3'
    services:
      db:
        image: mongo:3.4
      web:
        extends:
          file: docker-compose.versions.yml
          service: web
        depends_on: db
        env:
          APP_MONGO_ADDRESS: db:27017
    ```

+ Output of `compose-preview`:
    
    ```yaml
    version: '3'
    services:
      web:
        image: dotpy3/my-web-app:2.0
        env:
          APP_DEBUG: 'true'
          APP_MONGO_ADDRESS: db:27017
        depends_on: db
    ```

## What do I need this for?

+ Straightforward previews when you have multiple `docker-compose` files depending on each other

+ The output is deterministic, and thus makes `diff` operations easier.
