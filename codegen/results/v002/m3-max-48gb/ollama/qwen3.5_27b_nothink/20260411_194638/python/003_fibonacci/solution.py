import sys

def main():
    if len(sys.argv) != 2:
        # If no argument or extra arguments are provided, exit without output
        # as per the spirit of "Accept a single integer argument"
        return

    try:
        n = int(sys.argv[1])
    except ValueError:
        # If the argument is not a valid integer, exit without output
        return

    if n < 1:
        return

    # Start the sequence: 1, 1, 2, 3, 5, ...
    a, b = 1, 1

    # Print the first number if it doesn't exceed n
    if a <= n:
        print(a)
    
    # If n is at least 1, the second '1' also qualifies
    # Note: The sequence is 1, 1, 2... so we need to handle the second 1 explicitly
    # or adjust the loop logic to generate pairs.
    # Simple approach: print first, then loop to generate subsequent.
    
    if b <= n:
        print(b)
    else:
        # If the second 1 is greater than n (which is impossible if n >= 1),
        # but logically if n=0 we already returned. 
        # If n=1, we print 1, then check b=1 <= 1 -> print 1.
        # The loop below generates 2, 3, 5...
        pass

    # Generate subsequent numbers
    while True:
        c = a + b
        if c > n:
            break
        print(c)
        a = b
        b = c

if __name__ == "__main__":
    main()