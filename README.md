# Protoz
Prototype Pollution Finder tool

# Dependencies

Google Chromium
Go-lang
unfurl

# How to build this tool

1, download this repo

2, run this command inside the protoz folder 
 # go build -o ProtoZ
 
# Usage

Here place your target URLs into the url file then run this below command

cat urls | unfurl -u format "%s://%d%p" | awk '{print $1 "?__proto__[hacker]=1337"}' | ./Protoz -j 'window.hacker? "Vulnerable" : "Not Vulnerable"'

# XSS 

cat urls | ./ProtoZ -j "document.cookie"



Protoz Help 

-j - JS which needs to be run 



