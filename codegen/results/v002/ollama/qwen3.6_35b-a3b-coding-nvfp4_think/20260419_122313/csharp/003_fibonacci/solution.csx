using System;

if (Args.Count > 0 && long.TryParse(Args[0], out long n))
{
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