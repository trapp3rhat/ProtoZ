# Protoz
Prototype Pollution Finder tool

This tool is inspired from the NahamSec 2021 talk by @Tomnomnom(https://www.youtube.com/channel/UCyBZ1F8ZCJVKSIJPrLINFyA)

https://www.youtube.com/watch?v=Gv1nK6Wj8qM

# Dependencies

1,Google Chromium
2,Go-lang
3,unfurl

# How to build this tool

1, go get -u github.com/trapp3rhat/ProtoZ

2, Then download this repository

3, run ./ProtoZ.sh 

##*Note: Add path to the ProtoZ directory or GOPATH need to be added in the PATH*
 
What is inside this ProtoZ.sh ? 

cat urls | unfurl -u format "%s://%d%p" | awk '{print $1 "?__proto__[hacker]=1337"}' | ProtoZ -j 'window.hacker? "Vulnerable" : "Not Vulnerable"'
 
# ProtoZ Help 
-j - JS which needs to be run 

# XSS 

cat urls | ./ProtoZ -j "document.cookie"







