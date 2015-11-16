FROM rabbitmq:3.5.5-management
RUN echo "[{rabbit, [{tcp_listeners, [{\"0.0.0.0\", 5672}]},{loopback_users, []}]}]." > /etc/rabbitmq/rabbitmq.config
RUN service rabbitmq-server start

