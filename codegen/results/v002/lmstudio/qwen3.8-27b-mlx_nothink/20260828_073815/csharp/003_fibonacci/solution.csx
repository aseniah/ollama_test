using System;
using System.Numerics;

if (Args.Count > 0 && BigInteger.TryParse(Args[0], out var n) && n >= 1)
{
    BigInteger a = 1;
    BigInteger b = 1;

    while (a <= n)
    {
        Console.WriteLine(a);
        BigInteger next = a + b;
        a = b;
        b = next;
    }
}