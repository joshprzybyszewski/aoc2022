
#!/bin/bash

# get today's date, if not provided
if [ -z "$DAY" ]; then
  # if the day wasn't provided...
  # if it's not 2024, exit
  YEAR=$(date +%Y)
  if [ "$YEAR" -ne 2024 ]; then
    echo "It's not 2024. Exiting."
    exit 1
  fi
  # if it's not December, exit
  MONTH=$(date +%m)
  if [ "$MONTH" -ne 12 ]; then
    echo "It's not December. Exiting."
    exit 1
  fi
  DAY=$(date +%d)
fi

# Make today's old input file empty: 
INPUT_FILE="util/inputfiles/day$DAY.txt"
> $INPUT_FILE
echo "Emptied $INPUT_FILE"

# delete today's old answer files: 
ANSWER_FILE1="answers/day$DAY-level1.txt"
ANSWER_FILE2="answers/day$DAY-level2.txt"
rm -f $ANSWER_FILE1 $ANSWER_FILE2
echo "Deleted $ANSWER_FILE1 and $ANSWER_FILE2"

# Fetch today's new input file by running go run main.go -day=X:
go run main.go -day=$DAY

echo "Fetched new input file for day $DAY"