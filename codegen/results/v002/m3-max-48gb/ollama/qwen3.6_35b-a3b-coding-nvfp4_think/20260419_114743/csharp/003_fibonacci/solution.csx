using System;

int n = int.Parse(Args[0]);
if (n < 1) return;

int a = 1;
int b = 1;
while (a <= n)
{
    Console.WriteLine(a);
    int next = a + b;
    a = b;
    b = next;
}