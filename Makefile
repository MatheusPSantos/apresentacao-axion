run-microservices:
	cd microservices && ( \
    cd frontend && npx http-server --port 8080 & \
    cd frete && python3 app.py  & \
    cd carrinho && ./main & \
		cd pagamento && node index.js \
  )

run-event-driven:
	cd event-driven && ( \
		docker run --name zookeeper -p 2181:2181 -e ZOOKEEPER_CLIENT_PORT=2181 zookeeper & \
		Zookeeper_Server_IP=$(docker inspect zookeeper --format='{{ .NetworkSettings.IPAddress }}') & \
		docker run --name kafka -p 9092:9092 -e KAFKA_ZOOKEEPER_CONNECT=${Zookeeper_Server_IP}:2181 \
			-e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092 \
			-e KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1 \
			-e MAX_REQUEST_SIZE=5048576 \
			confluentinc/cp-kafka &\
		KAFKA_HOST=$(docker inspect kafka --format='{{.NetworkSettings.IPAddress}}') & \
	)
