## Approach:
1. Iterate through string, using `first` and `last` to keep track of result.
2. If digit, store in variables defined above and continue.
3. Check if substring starting from current position contains one of the "word numbers" and store its integer mapping in the variables
4. Return `<first><last>`