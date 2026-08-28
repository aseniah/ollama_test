using System;
using System.Collections.Generic;
using System.Linq;

int n = int.Parse(Args[0]);

if (n < 1)
{
    // Print nothing
}
else
{
    List<int> fibs = new List<int>();
    int a = 1, b = 1;
    fibs.Add(a);
    if (n >= 1)
    {
        fibs.Add(b);
    }
    
    while (true)
    {
        int c = a + b;
        if (c > n)
            break;
        fibs.Add(c);
        a = b;
        b = c;
    }
    
    // But we need to start with 1, 1, 2, 3, 5...
    // Let me reconsider: the sequence is 1, 1, 2, 3, 5, 8, 13, ...
    // So fib(1)=1, fib(2)=1, fib(3)=2, etc.
    // Print all that don't exceed N
    
    if (n >= 1)
    {
        Console.WriteLine(1);
    }
    
    if (n >= 1)
    {
        Console.WriteLine(1);
    }
    
    int prev = 1, curr = 1;
    while (true)
    {
        int next = prev + curr;
        if (next > n)
            break;
        Console.WriteLine(next);
        prev = curr;
        curr = next;
    }
}