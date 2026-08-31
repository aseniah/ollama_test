using System;
using System.Collections.Generic;
using System.Linq;

if (Args.Count != 1)
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

var fibs = new List<int>();
int a = 1, b = 1;
fibs.Add(a);
if (a <= N)
{
    // print first 1
}

// We need to print all Fibonacci numbers up to and including the largest that does not exceed N.
// Sequence: 1, 1, 2, 3, 5, 8, ...
// The first two are both 1.

// Let's generate:
if (N >= 1)
{
    Console.WriteLine("1");
}

// Now generate subsequent Fibonacci numbers
int x = 1, y = 1;
while (true)
{
    int next = x + y;
    if (next > N)
    {
        break;
    }
    Console.WriteLine(next);
    x = y;
    y = next;
}