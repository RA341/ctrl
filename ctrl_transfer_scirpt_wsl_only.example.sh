#!/bin/bash

# Define variables
local_file=""
remote_user="username"
remote_server="remote ip"
remote_directory="insert remote dir"

# Transfer file using scp
scp $local_file $remote_user@$remote_server:$remote_directory
