﻿#!/usr/bin/perl -w
use POSIX;
use Parallel::ForkManager;
#use Tie::File;
#use Fcntl 'O_RDONLY';

my $max_process = 5;
my $pm = new Parallel::ForkManager($max_process); 
my %info;#store the result generated by sub thread

my $slices = 20;
my $total_line = 100000;
my $block = $total_line/$slices;#split the big file into 20 slices
my $file="test.fastq";#100000 lines,each block with 10k lines
open IN,$file or die $!;
my @array = <IN>;
close IN;
#tie @array, 'Tie::File',$file , mode => O_RDONLY;

#run_on_finish is the key to share the data from the multiple threads.It is the callback run at the time of finished of the sub thread
#data structure(which was transfered from the sub thread)will be set back to the %info, so data sharing is done.
$pm -> run_on_finish(
	sub{
		my $tmp_info = pop @_;##get the last item transfered after the sub thread finished its work
		while(my ($key,$value) = each %$tmp_info){
			$info{$key} = $value;
		}
	}
);

foreach ( 1 .. $slices ) {
	$pm -> start and next;#create a new thread
	
	#do what you want to do here...	
	print ">slice$_ start at: ".strftime("%Y-%m-%d %H:%M:%S", localtime);
	my $data = readinfile($_);#the things we actually do.
	print "\t$_ stop at: ".strftime("%Y-%m-%d %H:%M:%S", localtime)."\n";
	#what you want to do finished here...
	
	#$data is an array ref, it including pid, exit code of the thread..and so on. the last item of $data is the data structure returned from the sub routine
	$pm -> finish(0,$data);
}
$pm -> wait_all_children;#synchonize all the sub threads
#untie @array;
print "\nAll finish at: ".strftime("%Y-%m-%d %H:%M:%S", localtime)."\n";;

foreach ( 1 .. $slices ){
	print "slice$_: ".$info{$_}."\n";	
}

sub readinfile{
	my $start = $_;#number $rank of the slices
	my ($total,$totalbase,$total_q20,$total_q30)=(0,0,0,0);
	my $x = $block * ($start-1);#read the fastq from line $x to line ($x+$block)
	my $k = $x;
	while ($k < ($x+$block)){#each loop read in one fastq read
		my ($name,$seq,$qname,$quality) = ($array[$k],$array[$k+1],$array[$k+2],$array[$k+3]);
		#$name= ;#read name
  	#$seq=;#read sequence
  	#$qname=;#read phred score name
  	#$quality=;#read phred score
  	$k = $k + 4;#four line means one fastq read
  	$seq=~s/\s+//g;#remove the blank space
    $quality=~s/\s+//g;
    my $seqlen=length($seq);
    if($seqlen ne length($quality)){#check the sequence and quality length
    	print "Warning: $name have different sequence and qualitylength\n";
    	next;#
    }
  	$total++;#count reads
    for($i=$seqlen-1;$i>=0;$i--){
    	$qual=substr($quality,$i,1);
    	$numqual=ord($qual)-33;#33+ based phred score
    	if($numqual>=20){
    		$total_q20++;#count q20 base
    	}
    	if($numqual>=30){
    		$total_q30++;#count q30 base
    	}
    	$totalbase++;#count bases
    }
	}
	print "\t$total;$totalbase;$total_q20;$total_q30\t";
	my %person_info;
	$person_info{$start} = join "\t",$total,$totalbase,$total_q20,$total_q30;
	return \%person_info;
}