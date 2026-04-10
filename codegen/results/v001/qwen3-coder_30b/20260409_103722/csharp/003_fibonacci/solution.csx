using System;
using System.Collections.Generic;

if (Args.Count == 0)
{
    return;
}

if (int.TryParse(Args[0], out int n) && n >= 1)
{
    int a = 1;
    int b = 1;
    
    Console.WriteLine(a);
    
    if (n >= 2)
    {
        Console.WriteLine(b);
        
        while (true)
        {
            int next = a + b;
            if (next > n)
                break;
                
            Console.WriteLine(next);
            a = b;
            b = next;
        }
    }
}