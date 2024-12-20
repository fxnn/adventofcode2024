#!/bin/bash
set -e

leaderboard="$(./leaderboard-fetch.sh)"

awk -v leaderboard="${leaderboard}" \
  '
  BEGIN {
    isLeaderboard="false" ;
    isLeaderboardPrinted="false" ;
  }
  /^.*\(LEADERBOARD_BEGIN\)$/ {
    isLeaderboard="true" ;
    isLeaderboardPrinted="false" ;
    print $0 ;
    next ;
  }
  /^.*\(LEADERBOARD_END\)$/ {
    isLeaderboard="false" ;
    print $0 ;
    next ;
  }
  isLeaderboard=="true" && isLeaderboardPrinted=="false" {
    print "" ;
    print "> " leaderboard ;
    print "" ;
    isLeaderboardPrinted="true" ;
    next ;
  }
  isLeaderboard=="true" && isLeaderboardPrinted=="false" {
    next ;
  }
  isLeaderboard=="false" {
    print $0 ;
  }
  ' \
  README.md \
  > README.md.tmp

mv README.md.tmp README.md

