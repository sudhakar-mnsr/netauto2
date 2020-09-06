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
my $snapshotto = 4; # Snapshot timeout
my $promisc = 1; # Put interface in promiscuous mode
my $to_ms = 0;
my $filter_str = ''; # Some tcpdump filter if we need any
my $optimize = 1; # Optimize filter
my $netmask =0; 
my $log_path = '/tmp/'; # Where we should save the tcpdump snapshots
my $IP = '127.0.0.1'; #bgpd IP address
my $err;
my %count;
my $filter = '';
local $dumper;
$pcap = Net::Pcap::open_live($dev, $snap_length, $promisc, $to_ms, \$err);
$pcap || die "Cant create packet decscriptor. Error was $err";
if ( Net::Pcap::compile($pcap, \$filter, $filter_str, $optimize, $netmask) == -1) {
$err = Net::Pcap::geterr($pcap);
die "Invalid filter: $filter_str (error was: $err)\n";
} else {
Net::Pcap::setfilter($pcap, $filter);
};
local $SIG{ALRM} = \&loop_exit;
my $start_time = time;
alaram $timeinterval;
my $exitcode = Net::Pcap::loop($pcap, $nr_packets, \&process_packet, '');
my $end_time = time;
$timeinterval = $end_time - $start_time;
alaram 0;
my %stats; Net::Pcap::stats($pcap, \%stats);
my $message = ''
my @add2bgp = ();
my $subject = 'FLOOD';
my $maxx = 0; my $loss_ip = '';
while (($key, $value) = each %count) {
$key2 = sprintf("%d.%d.%d.%d", ord(substr($key, 0, 1)),
                               ord(substr($key, 1, 1)),
                               ord(substr($key, 2, 1)),
                               ord(substr($key, 3, 1)));
$key = $key2;
my $nr = map { $key =~ /$_/x } @exceptions;
next if $nr > 0;
if ($value > $maxx) {
$maxx = $value;
$loss_ip = $key;
};

if (($value/$timeinterval) > $threshold) {
################# FLOOOOOOOOOOOOOOOOOOOOOD
$subject .= ' '.$key;
$message .= "\nFlood To ". $key . "\t(".int($value/ $timeinterval).' pps - Max. limit '.$threshold.')';
push @add2bgp, $key;
};
};
if ($message eq '') {
if (($stats{ps_drop}+$stats{ps_ifdrop}) > $max_loss){
## Flood to $loss_ip
$subject .= ' '.$loss_ip;
$message .= "\nFlood TO ". $loss_ip . "\t(".int($maxx/ $timeinterval).' pps - Max.limit'.$threshold.")\npackets loss by filter \t".$stats{ps_drop}.
"\npackets loss by the interface:\t".$stats{ps_ifdrop}..
"\nAllowed Max number of lost packets ".$max_loss;
push @add2bgp, $loss_ip;
};
};
if ($message ne '') {
use Net::Telnet::Cisco;
my $session = Net::Telnet::Cisco->new(Host => $IP, 
                                      Port => '2605',
                                      Timeout => 20, 
                                      Errmode => \&fail); 
$ok = $session->cmd(String => 'p[',
Prompt => '/bgpd[\$#>]/');
$ok = $session->cmd(String => 'enable',
Prompt => '/assword/',
Timeout => 20);
$ok = $session->cmd(String => 'p[',
Prompt => '/bgpd[\$#>]/',
Timeout => 20);
$ok = $session->cmd(String => "conf t",
Timeout => 20,
Prompt => '/bgpd\(config\)[\$#]/');
for ($i=0; $i<=$#add2dgp; $i++) {
$key = $add2bgp[$i];
$ok = $session->cmd(String => "access-list flood permit $key/32 exact-match",
Timeout => 20,
Prompt => '/bgpd\(config\)[\$#>]/');
$ok = $session->cmd(String => "router bgp 65000",
Timeout => 20,
Prompt => '/bgpd\(config\-router\)[\$#>]/');
$ok = $session->cmd(String => "network $key/32",
Timeout => 20,
Prompt => '/bgpd\(config\-router\)[\$#>]/');

$ok = $session->cmd(String => "exit",
Timeout => 20,
Prompt => '/bgpd[\$#>]/');
$ok = $session->cmd(String => "clear ip bgp $bgp_peer out",
Timeout => 20,
Prompt => '/bgpd[\$#>]/');
$ok = $session->cmd(String => "wr me",
Timeout => 20,
Prompt => '/bgpd[\$#>]/');

$ok = $session->cmd(String => "exit",
Timeout => 20);
$session->close();
};
