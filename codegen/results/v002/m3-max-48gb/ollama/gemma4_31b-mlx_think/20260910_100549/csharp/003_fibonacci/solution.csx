using System;

if (Args.Length == 0)
{
    return;
}

if (long.TryParse(Args[0], out long n))
{
    if (n < 1)
    {
        // If N < 1, print nothing as per requirements
        return;
    }

    long a = 1;
    long b = 1;

    // Fibonacci sequence: 1, 1, 2, 3, 5, ...
    // We continue printing while the current number in the sequence is <= N
    while (a <= n)
    {
        Console.WriteLine(a);
        
        long next = a + b;
        a = b;
        b = next;
    }
}