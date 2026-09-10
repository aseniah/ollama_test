using System;

if (Args.Length == 0)
{
    return;
}

if (int.TryParse(Args[0], out int n))
{
    if (n < 1)
    {
        return;
    }

    long a = 1;
    long b = 1;

    // Print the first Fibonacci number
    Console.WriteLine(a);

    // If N is 1 or more, the second '1' should also be printed
    if (n >= 1)
    {
        Console.WriteLine(b);
    }

    // Calculate subsequent numbers until the limit N is exceeded
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