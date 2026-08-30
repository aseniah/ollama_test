using System;

int n = int.Parse(Args[0]);

if (n < 1) { return; }

List<int> results = new List<int>();

int f1 = 1;
int f2 = 1;

results.Add(f1);
results.Add(f2);

while (true)
{
    int nextFib = f1 + f2;
    
    if (nextFib > n) { break; }
    
    f1 = f2;
    f2 = nextFib;
    results.Add(f2);
}

foreach (int fib in results) { Console.WriteLine(fib); }