using System;

if (Args.Count > 0 && int.TryParse(Args[0], out int n))
{
    if (n >= 1)
    {
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
}