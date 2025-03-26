#!/bin/bash
set -eo pipefail

LOG_FILE="test_httping.log"
TARGET_URL="https://example.com"  # Now contains full URL
TEST_COUNT=3

# Initialize log file
echo "=== httping test suite $(date) ===" > $LOG_FILE
echo "Version: $(./httping -h 2>&1 | head -n 1)" >> $LOG_FILE

# Function to run tests
run_test() {
  local test_name="$1"
  local command="$2"
  local expect_fail="${3:-false}"
  
  echo -n "Running $test_name... " | tee -a $LOG_FILE
  echo -e "\n\n=== $test_name ===\n$command" >> $LOG_FILE
  
  if $expect_fail; then
    if eval "$command" >> $LOG_FILE 2>&1; then
      echo "FAIL (should have failed)" | tee -a $LOG_FILE
      return 1
    else
      echo "PASS" | tee -a $LOG_FILE
      return 0
    fi
  else
    if eval "$command" >> $LOG_FILE 2>&1; then
      echo "PASS" | tee -a $LOG_FILE
      return 0
    else
      echo "FAIL" | tee -a $LOG_FILE
      return 1
    fi
  fi
}

# Clean up on exit
cleanup() {
  echo -e "\nTests completed. Full output in $LOG_FILE"
  [ -n "$LISTENER_PID" ] && kill $LISTENER_PID 2>/dev/null
}
trap cleanup EXIT

# Test Cases
echo -e "\n=== Test Results ==="

run_test "Basic HTTP" "./httping -url http://${TARGET_URL#*://} -count $TEST_COUNT"
run_test "Basic HTTPS" "./httping -url $TARGET_URL -count $TEST_COUNT"
run_test "JSON Output" "./httping -url $TARGET_URL -json -count 1 | jq -e . >/dev/null"
run_test "Invalid Cert" "./httping -url https://expired.badssl.com -count 1" true
run_test "Invalid URL" "./httping -url 'not a url' -count 1" true
run_test "HEAD Requests" "./httping -url $TARGET_URL -httpverb HEAD -count 2"
run_test "Custom Host Header" "./httping -url http://127.0.0.1 -hostheader ${TARGET_URL#*://} -count 1" true

# Listener test using httping
echo -n "Running Listener Mode... " | tee -a $LOG_FILE
echo -e "\n\n=== Listener Test ===" >> $LOG_FILE
( ./httping -listen 8080 >> $LOG_FILE 2>&1 ) &
LISTENER_PID=$!
sleep 1  # Give it time to start

if ./httping -url http://localhost:8080 -count 1 >> $LOG_FILE 2>&1; then
  echo "PASS" | tee -a $LOG_FILE
else
  echo "FAIL" | tee -a $LOG_FILE
fi

