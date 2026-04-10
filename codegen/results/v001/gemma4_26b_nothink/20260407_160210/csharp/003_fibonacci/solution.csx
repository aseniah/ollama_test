using System;

// Check if an argument was provided
if (Args.Count == 0)
{
    return;
}

// Attempt to parse the first argument as an integer
if (!int.TryParse(Args[0], out int n) || n < 1)
{
    // If N < 1, print nothing as per requirements
    return;
}

// Initialize the first two numbers of the sequence
long prev = 1;
long current = 1;

// Print the first number
Console.WriteLine(prev);

// If n is at least 1, the second '1' should be printed if it doesn't exceed n
if (current <= n)
{
    // To handle the sequence 1, 1, 2... specifically, we print the second 1
    // We check if n >= 1, which is already checked above.
    // However, we must ensure we don't print it if the logic requires strictly up to N.
    // Since the sequence is 1, 1, 2..., and 1 is <= n, we print it.
    Console.WriteLine(current);
}

// Calculate the rest of the sequence
while (true)
{
    long next = prev + current;

    // If the next Fibonacci number exceeds N, stop
    if (next > n)
    {
        break;
    }

    Console.WriteLine(next);

    // Update values for next iteration
    prev = current;
    current = next;
}