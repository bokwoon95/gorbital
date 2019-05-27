#!/bin/bash

pg_ctl -D .orbitaldb_dev -o "-p 5433" stop
pg_ctl -D .sessiondb_dev -o "-p 5434" stop
