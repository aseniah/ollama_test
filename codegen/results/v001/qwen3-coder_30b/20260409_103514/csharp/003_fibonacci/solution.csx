if (Args.Count == 0 || !int.TryParse(Args[0], out int N) || N < 1)
{
    return;
}

int a = 1;
int b = 1;

while (a <= N)
{
    Console.WriteLine(a);
    int temp = a + b;
    a = b;
    b = temp;
}