using System;
using System.Collections.Generic;

if (Args.Count > 0 && long.TryParse(Args[0], out long n))
{
    if (n < 1)
    {
        return;
    }

    long a = 1;
    long b = 1;

    while (a <= n)
    {
        Console.WriteLine(a);
        long temp = a + b;
        a = b;
        b = temp;
    }
}