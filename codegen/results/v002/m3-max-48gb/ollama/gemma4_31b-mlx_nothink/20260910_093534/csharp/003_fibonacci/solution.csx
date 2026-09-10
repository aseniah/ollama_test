using System;

if (Args.Length == 0)
{
    return;
}

if (long.TryParse(Args[0], out long n))
{
    if (n < 1)
    {
        return;
    }

    long a = 1;
    long b = 1;

    // Print the first number of the sequence
    Console.WriteLine(a);

    // If N is at least 1, we already printed the first 1. 
    // Check if we can print the second 1.
    if (n >= 1)
    {
        // The sequence starts 1, 1...
        // Since the prompt specifies the sequence 1, 1, 2, 3, 5...
        // We handle the first two 1s specially.
        Console.WriteLine(b);
    }

    // Calculate subsequent numbers
    while (true)
    {
        long next = a + b;
        if (next > n)
        {
            break;
        }
        Console.WriteLine(next);
        a = b;
        b = next;
    }
}