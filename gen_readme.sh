#!/bin/bash

# Get the output of the `tree` command.
out=$(tree -L 2 000-assignments)

# Replace all the special characters with spaces.
markdown=$(echo "$out" | tr -d '├│─└' | tr -s ' ')

# Initialize some variables.
moduleName=""

# Loop over each line of the markdown output.
for line in $markdown; do

    # Replace all the spaces with dashes.
    line=$(echo "$line" | tr ' ' '-')

    # If the line contains "000-assignments", skip it.
    if [[ $line == *"000-assignments"* ]]; then
        continue
    fi

    # If the line contains "directories", skip it.
    if [[ $line == *"directories"* ]]; then
        continue
    fi

    # If the line contains "intermediate-dsa" or "advanced-dsa", set the module name.
    if [[ $line == *"intermediate-dsa"* ]] || [[ $line == *"advanced-dsa"* ]]; then
        moduleName=$line
    else
        # Otherwise, add the line to the markdown output.
        markdown="$markdown\n\t- [${line:0:-1}](./000-assignments/$moduleName/$line)"
    fi

done

# Print the markdown output.
printf "$markdown"
