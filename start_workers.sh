#!/bin/bash

# Number of worker processes to start
worker_count=$1
# Base directory
base_dir=~/Project/Go_POC
# Logs directory
logs_dir="$base_dir/.logs"
# Memory limit per worker (in MB)
mem_limit_mb=512

cd "$base_dir" || exit 1

# Create logs directory if it doesn't exist
mkdir -p "$logs_dir"

# Stop any running workers
pkill -f "./worker/main" || true
sleep 2

# Remove old logs
rm -f "$logs_dir/worker_"*.log

# Start workers
for ((i=1; i<=worker_count; i++)); do
    log_file="$logs_dir/worker_$i.log"
    echo "Starting worker $i -> $log_file"
    # GOMEMLIMIT sets an upper bound on the total memory the Go runtime (heap) can use for a process.
    GOMEMLIMIT="${mem_limit_mb}MiB" ./worker/main > "$log_file" 2>&1 &
    sleep 0.5
done

echo "All workers started. Logs are in: $logs_dir/"