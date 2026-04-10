if (Args.Count == 0 || !int.TryParse(Args[0], out int N) || N < 1)
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