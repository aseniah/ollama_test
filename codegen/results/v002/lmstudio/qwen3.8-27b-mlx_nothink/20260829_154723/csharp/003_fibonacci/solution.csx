using System;
using System.Collections.Generic;

if (Args.Count < 1)
{
    return;
}

if (!int.TryParse(Args[0], out int n))
{
    return;
}

if (n < 1)
{
    return;
}

long a = 1, b = 1;
while (a <= n)
{
    Console.WriteLine(a);
    long next = a + b;
    a = b;
    b = next;
}