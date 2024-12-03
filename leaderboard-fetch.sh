#!/bin/bash
set -e
set -o pipefail

aoc_year=2024
export aoc_username="fxnn"

[ -z "${AOC_SESSION_COOKIE}" ] && { echo "ERROR: missing AOC_SESSION_COOKIE"; exit 1; }
[ -z "${AOC_LEADERBOARD_ID}" ] && { echo "ERROR: missing AOC_LEADERBOARD_ID"; exit 1; }

curl \
  -sS \
  -H "Cookie: ${AOC_SESSION_COOKIE}" \
  https://adventofcode.com/${aoc_year}/leaderboard/private/view/${AOC_LEADERBOARD_ID}.json \
  | jq -r \
  '.members 
  | values[] 
  | select(.name==$ENV.aoc_username) 
  | "user **" 
  + .name 
  + "** (" 
  + (.stars | tostring) 
  + " stars, last one at " 
  + (.last_star_ts | todateiso8601) 
  + ")"'

