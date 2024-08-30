# HEX DDD CQS EXAMPLE

## Description
Provide a project as example to how implement hexagonal architecture, DDD and CQS pattern.
Starting with a coupled code and refactoring on each git branchs.

## Start the project

1. Start the containers
    ```shell
    make docker-up
    ```

2. Run migrations to create database structure
    ````shell
    make run-migrations
    ````

3. See sever container logs
    ```shell
    make docker-logs
    ```

4. Run functional tests to be sure all is working fine.
    ````shell
    make test-funcional
    ````
   
## Branches

Change the repository branches to see changes in the code.

1. main --> first approach (legacy code)
2. domain_services -> refactor user package creating domain and testing services.
3. domain_testing -> adding test to domain entity
3. infra -> move infra code to infra package
4. app -> creating app layer and test http package (controllers)
