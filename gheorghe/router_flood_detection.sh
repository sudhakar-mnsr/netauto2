#!/usr/bin/perl
# security threats, large network flood-detection tool
# 
use Net::Pcap;
sub loop_exit { Net::Pcap::breakloop($pcap); }
sub fail() { print STDERR "$IP telnet FAILED!\n"; };
local $pcap;
my @exceptions = ('1.1.10.'); # Exceptions list
my $bgp_peer = '1.3.1.89'; # BGP Peer that needs to be cleared
my $dev = 'eth0'; # Monitored interface
my $threshold = 6500; # Maximum limit of pps
my $timeinterval = 10; # Time interval to read from the interface
my $nr_packets = 50000; # Max number of packets to read
my $max_loss = 3000 * $timeinterval; # Max number of lost packets
my $snap_length = 128 # How much to read from a packet
my $pktsnapshot = 20; # Number of captured packets
