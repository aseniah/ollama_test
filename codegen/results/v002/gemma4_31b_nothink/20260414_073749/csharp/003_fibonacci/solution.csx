using System;

if (Args.Length == 0)
{
    return;
}

if (!long.TryParse(Args[0], out long n) || n < 1)
{
    // If N is not a valid integer or is less than 1, print nothing.
    return;
}

long a = 1;
long b = 1;

// Print the first number of the sequence
Console.WriteLine(a);

// If N is 1, we only print the first 1. 
// However, the Fibonacci sequence starts 1, 1...
// Since the requirement is "up to and including the largest that does not exceed N",
// if N=1, the largest Fibonacci number <= 1 is 1. 
// But the sequence is defined as 1, 1, 2... so we check if we can print the second 1.
if (n >= 1)
{
    // We already printed the first '1'. Now check for the second '1'.
    // The loop below handles sequences from the 3rd element onwards.
    // Let's handle the second '1' explicitly.
    Console.WriteLine(b);
}

// Calculate subsequent numbers
long next = a + b;
while (next <= n)
{
    Console.WriteLine(next);
    a = b;
    b = next;
    next = a + b;
    
    // Prevent overflow for very large N
    if (next < 0) break; 
}