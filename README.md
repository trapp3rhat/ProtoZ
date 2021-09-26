# Protoz
Prototype Pollution Finder tool

# Dependencies

1,Google Chromium
2,Go-lang
3,unfurl

# How to build this tool

1, download this repo

2, run this command inside the ProtoZ folder 
 # go build -o ProtoZ
 
3, Place the file into the $Path directories
 
 
# Protoz Help 
-j - JS which needs to be run 

# Usage

Here place your target URLs into the url file then run this below command

cat urls | unfurl -u format "%s://%d%p" | awk '{print $1 "?__proto__[hacker]=1337"}' | ./ProtoZ -j 'window.hacker? "Vulnerable" : "Not Vulnerable"'

or simply run 

./ProtoZ.sh 

# XSS 

cat urls | ./ProtoZ -j "document.cookie"







