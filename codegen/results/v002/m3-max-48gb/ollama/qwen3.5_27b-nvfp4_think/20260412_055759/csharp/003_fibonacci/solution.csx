using System;

if (Args.Count > 0 && int.TryParse(Args[0], out int n))
{
    if (n >= 1)
    {
        int a = 1;
        int b = 1;
        
        while (a <= n)
        {
            Console.WriteLine(a);
            int next = a + b;
            a = b;
            b = next;
        }
    }
}