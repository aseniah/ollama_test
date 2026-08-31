using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Text.Json;
using System.Text.Json.Nodes;

if (Args.Count == 0)
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

List<long> fibs = new List<long>();
long a = 1;
long b = 1;

fibs.Add(a);
fibs.Add(b);

while (true)
{
    long next = a + b;
    if (next > n)
    {
        break;
    }
    fibs.Add(next);
    a = b;
    b = next;
}

foreach (long fib in fibs)
{
    Console.WriteLine(fib.ToString());
}