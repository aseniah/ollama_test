using System;

// Check if an argument was provided and parse it to a long to prevent overflow
// during comparison with potentially large Fibonacci numbers.
if (Args.Count > 0 && long.TryParse(Args[0], out long n) && n >= 1)
{
    // We initialize the sequence as per the requirement: 1, 1, 2, 3, 5...
    // 'current' represents the Fibonacci number we are currently evaluating.
    // 'next' represents the subsequent Fibonacci number in the sequence.
    long current = 1;
    long next = 1;

    // The loop continues as long as the current Fibonacci number does not exceed N.
    while (current <= n)
    {
        // Print the current number in the sequence.
        Console.WriteLine(current);

        // Calculate the next term in the sequence: F(i) + F(i+1) = F(i+2).
        long temp = current + next;

        // Shift the window: the old 'next' becomes the new 'current',
        // and the 'temp' (sum) becomes the new 'next'.
        current = next;
        next = temp;
    }
}