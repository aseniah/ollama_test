using System;

// Check if an argument was provided
if (Args.Count == 0)
{
    return;
}

// Attempt to parse the first argument as a long to handle large values of N
if (long.TryParse(Args[0], out long n) && n >= 1)
{
    // The Fibonacci sequence starts with 1, 1, 2, 3, 5...
    // We maintain 'current' as the number to be printed and 'next' as the subsequent number.
    long current = 1;
    long next = 1;

    // Iterate as long as the current Fibonacci number does not exceed N
    while (current <= n)
    {
        Console.WriteLine(current);
        
        // Calculate the next number in the sequence
        long temp = current + next;
        
        // Shift the window: current becomes the previous 'next', 
        // and 'next' becomes the newly calculated sum.
        current = next;
        next = temp;
    }
}