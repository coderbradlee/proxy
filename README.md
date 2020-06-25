# proxy
To build the server:

	go build .

To run server:

	export PORT=8080
    export ENDPOINT=https://did.iotex.one
    ./proxy
To build the Docker image:

	docker build -f ./Dockerfile . -t iotexproject/uni-resolver-proxy

To run the Docker image:

	docker run -p 8080:8080 iotexproject/uni-resolver-proxy
