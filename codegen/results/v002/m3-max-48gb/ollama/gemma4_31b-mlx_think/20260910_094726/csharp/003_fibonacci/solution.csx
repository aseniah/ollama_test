using System;

if (Args.Length == 0)
{
    return;
}

if (long.TryParse(Args[0], out long n))
{
    if (n < 1)
    {
        // Print nothing as per requirements
        return;
    }

    long a = 1;
    long b = 1;

    while (a <= n)
    {
        Console.WriteLine(a);
        
        long next = a + b;
        a = b;
        b = next;
    }
}