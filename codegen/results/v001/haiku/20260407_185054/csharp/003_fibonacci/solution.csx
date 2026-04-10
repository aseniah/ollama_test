if (Args.Count == 0)
{
    return;
}

if (!int.TryParse(Args[0], out int n) || n < 1)
{
    return;
}

long prev = 0;
long curr = 1;

while (curr <= n)
{
    Console.WriteLine(curr);
    long next = prev + curr;
    prev = curr;
    curr = next;
}
