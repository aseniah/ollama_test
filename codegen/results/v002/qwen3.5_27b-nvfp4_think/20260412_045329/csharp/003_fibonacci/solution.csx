using System;
using System.Collections.Generic;
using System.IO;

// Parse the command line argument
int n = 0;
if (Args.Count > 0 && int.TryParse(Args[0], out parsed))
{
    n = parsed;
}

// Print Fibonacci numbers up to N (inclusive) if N >= 1
if (n >= 1)
{
    long a = 1, b = 1;
    
    while (a <= n)
    {
        Console.WriteLine(a);
        long next = a + b;
        a = b;
        b = next;
    }
}

long parsed;