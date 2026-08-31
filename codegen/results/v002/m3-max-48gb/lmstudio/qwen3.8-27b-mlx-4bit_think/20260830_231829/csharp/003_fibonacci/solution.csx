using System;

if (Args.Count < 1)
    return;

int N = int.Parse(Args[0]);

if (N < 1)
    return;

int a = 1, b = 1;
while (a <= N)
{
    Console.WriteLine(a);
    int temp = a;
    a = b;
    b = temp + b;
}