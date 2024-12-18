#!/bin/bash
set -e
set -o pipefail

aoc_year=2024
export aoc_username="fxnn"

[ -z "${AOC_SESSION_COOKIE}" ] && { echo "ERROR: missing AOC_SESSION_COOKIE" >&2; exit 1; }
[ -z "${AOC_LEADERBOARD_ID}" ] && { echo "ERROR: missing AOC_LEADERBOARD_ID" >&2; exit 1; }

curl \
  -sS \
  -H "Cookie: ${AOC_SESSION_COOKIE}" \
  https://adventofcode.com/${aoc_year}/leaderboard/private/view/${AOC_LEADERBOARD_ID}.json \
  | jq -r \
  '.members 
  | values[] 
  | select(.name==$ENV.aoc_username) 
  | "![" 
  + (.stars | tostring) 
  + " stars]("
  + "https://img.shields.io/badge/"
  + (.stars | tostring)
  + "-%E2%AD%90_stars-gold) "
  + "user **" 
  + .name 
  + "** (" 
  + "last one at " 
  + (.last_star_ts | todateiso8601) 
  + ")"'

