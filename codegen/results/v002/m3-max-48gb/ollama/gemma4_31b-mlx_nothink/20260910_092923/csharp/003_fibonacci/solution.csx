using System;

if (Args.Length == 0)
{
    return;
}

if (!long.TryParse(Args[0], out long n) || n < 1)
{
    // If N is not a valid integer or N < 1, print nothing as per requirements.
    return;
}

long a = 1;
long b = 1;

// Print the first Fibonacci number
Console.WriteLine(a);

// If n is at least 1, we can potentially print the second '1'
if (n >= 1)
{
    // The problem specifies the sequence: 1, 1, 2, 3, 5...
    // Since we already printed the first 1, we handle the second 1.
    // However, we must ensure we don't print beyond N.
    // Because we already printed 'a', and the second number is 'b' (which is 1),
    // we check if b <= n.
    if (b <= n)
    {
        // We only print the second 1 if it's within the range.
        // To avoid an infinite loop or logic error, we'll use a while loop 
        // to calculate the next terms.
    }
}

// Since we already printed the first 1, we start the loop to handle 
// the second 1 and all subsequent numbers.
long current = 1; 
long next = 1;

// Special handling for the starting sequence 1, 1...
// We printed the first 1. Now we check if we can print the second 1 and beyond.
if (n >= 1)
{
    // Print the second 1
    Console.WriteLine(b);
    
    // Calculate subsequent numbers
    while (true)
    {
        long temp = a + b;
        if (temp > n) break;
        
        Console.WriteLine(temp);
        a = b;
        b = temp;
    }
}