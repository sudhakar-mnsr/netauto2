#!/usr/bin/perl
# security threats, large network flood-detection tool
# 
use Net::Pcap;
sub loop_exit { Net::Pcap::breakloop($pcap); }
sub fail() { print STDERR "$IP telnet FAILED!\n"; };
local $pcap;
my @exceptions = ('1.1.10.'); # Exceptions list
