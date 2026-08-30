using System;

// Check if an argument was provided
if (Args.Length == 0)
{
    return;
}

// Parse the input argument N as a long to handle larger Fibonacci numbers
if (!long.TryParse(Args[0], out long n))
{
    return;
}

// If N < 1, print nothing as per requirements
if (n < 1)
{
    return;
}

// Initialize the first two numbers of the Fibonacci sequence
long current = 1;
long next = 1;

// The loop continues as long as the current Fibonacci number does not exceed N
while (current <= n)
{
    Console.WriteLine(current);

    // Calculate the next number in the sequence
    long temp = current + next;
    current = next;
    next = temp;

    // Handle potential overflow for very large N (long.MaxValue)
    // If current becomes negative due to overflow, stop the loop
    if (current < 0)
    {
        break;
    }
}