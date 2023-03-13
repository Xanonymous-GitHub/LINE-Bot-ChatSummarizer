#!/usr/bin/env bash

docker build --pull -t xanonymous/chat-gpt-line-bot . && \
docker push xanonymous/chat-gpt-line-bot