#!/bin/bash

pg_ctl -D orbitaldb_dev -o "-p 5433" start
pg_ctl -D sessiondb_dev -o "-p 5434" start
