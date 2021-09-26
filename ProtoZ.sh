
#!/bin/bash

cat urls | unfurl -u format "%s://%d%p" | awk '{print $1 "?__proto__[hacker]=1337"}' | ./ProtoZ -j 'window.hacker? "Vulnerable" : "Not Vulnerable"'
