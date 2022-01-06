#!/bin/sh

iptables -I INPUT -p tcp -m tcp --dport 8080 -j ACCEPT

iptables -I INPUT -p tcp -m tcp --dport 8081 -j ACCEPT

iptables -t nat -A PREROUTING -i ap0 -p tcp --dport 80 -j REDIRECT --to-port 8080

iptables -t nat -A PREROUTING -i ap0 -p tcp --dport 443 -j REDIRECT --to-port 8081