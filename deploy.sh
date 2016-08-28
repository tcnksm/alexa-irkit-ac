#!/bin/bash

# FIXME: apex --env-file doesn't work...
apex deploy \
     -s ACCESS_TOKEN=$ACCESS_TOKEN \
     -s IR_CLIENT_KEY=$IR_CLIENT_KEY \
     -s IR_DEVICE_ID=$IR_DEVICE_ID \
     ac
