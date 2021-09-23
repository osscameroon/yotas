#!/bin/bash

#file used only with docker
/start-postgres/dev/setup-postgres/setup-postgres.sh
make -C /start-postgres/ e2etest
