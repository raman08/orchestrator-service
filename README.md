# orchestrator-service

-   This is an orchestrator service which would read any request it receives and forwards it to other orchestrator services and data services.
-   The flow of the program is:
    ```
    client ---RPC--> orchestrator_1(:9000) ---RPC--> orchestrator_2(:9001) ---RPC--> mock_data_service(:10000)
    ```

## Usage

> **Assumption:** You have already Setup to go and protoc. Also make sure that go is in your path.

-   Clone the Repository

    ```
    git clone https://github.com/raman08/orchestrator-service.git
    ```

-   Inside the Repository

    ```
    cd orchestrator-service
    ```

-   Start Mock data server, orchestrator 1 and orchestrator 2 in 3 different terminal.

    ```
    go run ./datamock/main.go
    go run ./logic/orchestrator_2.go
    go run ./logic/orchestrator_1.go
    ```

-   Run the client
    ```
    go run ./client/main.go
    ```

**Side Note:** You can add more names in _client/main.go_ (Line No: 30)

### System

-   I have user _Arch Linux_ :) as development enviornment.
-   **Go Version:** _go1.17.6 linux/amd64_
-   **GO path:** _/home/raman/.local/share/go_
-   **PWD:** _/home/raman/orchestrator-service_
