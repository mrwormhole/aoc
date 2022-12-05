#!/usr/bin/perl
use warnings;
use strict;

my @stacks = (); # sample.txt STACK
push(@stacks, ["Z", "N"]);
push(@stacks, ["M", "C", "D"]);
push(@stacks, ["P"]);
#=pod
@stacks = (); # input.txt STACK
push(@stacks, ["R", "P", "C", "D", "B", "G"]);
push(@stacks, ["H", "V", "G"]);
push(@stacks, ["N", "S", "Q", "D", "J", "P", "M"]);
push(@stacks, ["P", "S", "L", "G", "D", "C", "N", "M"]);
push(@stacks, ["J", "B", "N", "C", "P", "F", "L", "S"]);
push(@stacks, ["Q", "B", "D", "Z", "V", "G", "T", "S"]);
push(@stacks, ["B", "Z", "M", "H", "F", "T", "Q"]);
push(@stacks, ["C", "M", "D", "B", "F"]);
push(@stacks, ["F", "C", "Q", "G"]);
#=cut

my ($filename, $keep_stack_order) = ('input.txt', 1);
open(FH, '<', $filename) or die $!;

while(<FH>){
  my @spl = split / /;
  my ($move, $from, $to) = ($spl[1], $spl[3], $spl[5]);
  #print("move $move from $from to $to \n");
    
  if ($keep_stack_order) {
    my @temp = ();
    for (1..$move) {
        my $elem = pop(@{$stacks[$from-1]});
        push(@temp, $elem);
    }
    foreach my $elem (reverse(@temp)) {
        #print("element: $elem\n");
        push(@{$stacks[$to-1]}, $elem);
    }
  } else {
    for (1..$move) {
      my $elem = pop(@{$stacks[$from-1]});
      #print("element: $elem\n");
      push(@{$stacks[$to-1]}, $elem);
    }
  }
}

foreach my $stack (@stacks) { print "@{$stack} \n" }
close(FH);