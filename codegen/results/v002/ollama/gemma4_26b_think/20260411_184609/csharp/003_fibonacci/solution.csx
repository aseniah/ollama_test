using System;

// Check if there's at least one argument provided
if (Args.Count == 0)
{
    return;
}

// Parse the first argument as a long to accommodate large values of N
// and check if N is at least 1. If N < 1, the requirement says to print nothing.
if (long.TryParse(Args[0], out long n) && n >= 1)
{
    // We start with the first two numbers of the Fibonacci sequence provided: 1, 1.
    // We'll use 'a' to represent the current number in the sequence and 'b' for the next.
    long a = 1;
    long b = 1;

    // The loop continues as long as the current Fibonacci number 'a' does not exceed N.
    while (a <= n)
    {
        // Print the current Fibonacci number
        Console.WriteLine(a);

        // Before calculating the next Fibonacci number (a + b), check for potential long overflow.
        // If (long.MaxValue - a) < b, then the sum (a + b) would exceed the capacity of a long.
        if (long.MaxValue - a < b)
        {
            // If overflow is imminent, the next number will definitely be greater than N.
            // We break to avoid calculation errors or infinite loops.
            break;
        }

        // Calculate the next term in the sequence
        long next = a + b;

        // Update 'a' and 'b' for the next iteration.
        // We set 'a' to the current 'b' and 'b' to the newly calculated 'next'.
        // Note: On the first iteration (a=1, b=1), the second iteration will have a=1, b=2,
        // which correctly results in the sequence 1, 1, 2...
        a = b;
        b = next;
    }
}