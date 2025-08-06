#!/usr/bin/perl
use v5.38;
use HTTP::Tiny;
unless (@ARGV) {
    die "Usage: $0 <base-url>\n";
}
my $base_url = $ARGV[0];

if ($base_url =~ m{.*youtube\.com/@+([^/]+)/.*}i) {
    # TODO: extract channel id from arg url
    # my $channel = $1;
    # my $feed = "https://www.youtube.com/feeds/videos.xml?channel_id=".$channel;
    # die "YouTube channel feed:\n$feed\n";
    #
    print "Type the channel id at the end of this url:\n";
    die  "https://www.youtube.com/feeds/videos.xml?channel_id=\n";
};

my @variations = ( 
    '/feed/',
    '/rss/',
    '/atom/',
    '/feed.xml',
    '/rss.xml',
    '/atom.xml',
    '/index.xml',
    '/blog/feed/',
    '/blog/rss/',
    '/blog/atom/',
    '/blog-feed.xml',
    '/feeds/posts/default/',
    '/?feed=rss2',
);

my $http = HTTP::Tiny->new;
foreach my $var (@variations) {
    my $url = $base_url.$var;
    my $response = $http->get($url);

    unless ($response->{success}) {
        next
    };

    die "Found RSS feed at: $url\n";
};

print "Nothing found\n";
