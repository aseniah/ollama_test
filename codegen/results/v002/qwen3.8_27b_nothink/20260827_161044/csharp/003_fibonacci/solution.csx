using System;
using System.Collections.Generic;

// Parse the command line argument
if (Args == null || Args.Count < 1)
{
    return;
}

if (!int.TryParse(Args[0], out int N))
{
    return;
}

if (N < 1)
{
    return;
}

// Generate Fibonacci numbers up to and including the largest that doesn't exceed N
List<long> fibs = new List<long>();
long a = 1, b = 1;

fibs.Add(a);

if (N >= 1)
{
    fibs.Add(b);
}

long c;
while (true)
{
    c = a + b;
    if (c > N)
    {
        break;
    }
    fibs.Add(c);
    a = b;
    b = c;
}

foreach (var f in fibs)
{
    Console.WriteLine(f);
}