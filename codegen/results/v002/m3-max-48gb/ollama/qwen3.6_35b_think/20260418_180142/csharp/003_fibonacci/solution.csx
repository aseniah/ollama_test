using System;

if (Args.Count > 0)
{
    long n = long.Parse(Args[0]);
    if (n >= 1)
    {
        long a = 1;
        long b = 1;
        Console.WriteLine(a);
        while (b <= n)
        {
            Console.WriteLine(b);
            long temp = a + b;
            a = b;
            b = temp;
        }
    }
}