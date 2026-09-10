using System;
using System.Collections.Generic;

if (Args.Count < 1)
    return;

long N = long.Parse(Args[0]);

if (N < 1)
    return;

List<long> fibs = new List<long>();
long a = 1, b = 1;

while (a <= N)
{
    fibs.Add(a);
    long next = a + b;
    a = b;
    b = next;
}

foreach (long f in fibs)
{
    Console.WriteLine(f);
}