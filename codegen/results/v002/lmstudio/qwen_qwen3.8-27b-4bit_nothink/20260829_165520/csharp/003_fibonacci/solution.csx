using System;
using System.Collections.Generic;
using System.Linq;

if (Args.Count != 1)
{
    // If no argument or multiple arguments, print nothing or handle error?
    // The task says "Accept a single integer argument N"
    // If invalid input, we'll just exit or print nothing.
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

// Generate Fibonacci numbers: 1, 1, 2, 3, 5, ...
// Print all that don't exceed N

var fibs = new List<int>();
int a = 1, b = 1;
while (a <= N)
{
    fibs.Add(a);
    int next = a + b;
    a = b;
    b = next;
    // Handle overflow safety, though for typical int range this is fine
    if (b < 0) break; // overflow detection
}

foreach (var f in fibs)
{
    Console.WriteLine(f);
}