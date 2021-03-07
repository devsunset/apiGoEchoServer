kill -9 $(ps aux | grep 'go-build' | awk '{print $2}')
