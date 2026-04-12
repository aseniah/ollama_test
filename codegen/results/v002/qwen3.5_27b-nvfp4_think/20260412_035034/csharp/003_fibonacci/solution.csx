using System;
using System.Collections.Generic;

if (Args.Count > 0 && int.TryParse(Args[0], out int n) && n >= 1)
{
    long a = 1;
    long b = 1;
    
    Console.WriteLine(a); // First 1
    
    while (b <= n)
    {
        Console.WriteLine(b);
        long next = a + b;
        a = b;
        b = next;
    }
}