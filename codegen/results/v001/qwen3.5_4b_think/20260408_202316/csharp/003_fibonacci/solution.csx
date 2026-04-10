using System;

int N = int.Parse(Args[0]);

if (N >= 1)
{
    int a = 1;
    int b = 1;
    
    while (a <= N)
    {
        Console.WriteLine(a);
        int temp = a + b;
        a = b;
        b = temp;
    }
}