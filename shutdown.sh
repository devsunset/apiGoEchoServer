# kill -9 $(ps aux | grep 'go-build' | awk '{print $2}')
kill -9 $(ps aux | grep './main' | awk '{print $2}')
