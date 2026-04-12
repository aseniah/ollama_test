using System;
using System.Text;

int N = int.Parse(Args[0]);

if (N < 1)
{
    return;
}

int a = 1;
int b = 1;

Console.WriteLine(a);
if (b <= N)
{
    Console.WriteLine(b);
}

while (true)
{
    int next = a + b;
    if (next > N)
        break;
    
    Console.WriteLine(next);
    a = b;
    b = next;
}