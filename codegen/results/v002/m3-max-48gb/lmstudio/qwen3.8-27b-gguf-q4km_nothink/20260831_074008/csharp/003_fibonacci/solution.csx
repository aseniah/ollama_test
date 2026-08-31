using System;
using System.Collections.Generic;

if (Args.Count != 1)
{
    Console.Error.WriteLine("Usage: provide a single integer argument N");
    Environment.Exit(1);
}

if (!int.TryParse(Args[0], out int n))
{
    Console.Error.WriteLine("Invalid integer argument");
    Environment.Exit(1);
}

if (n < 1)
{
    return;
}

long a = 1, b = 1;
var result = new List<long>();

while (a <= n)
{
    result.Add(a);
    long next = a + b;
    a = b;
    b = next;
}

foreach (var num in result)
{
    Console.WriteLine(num);
}