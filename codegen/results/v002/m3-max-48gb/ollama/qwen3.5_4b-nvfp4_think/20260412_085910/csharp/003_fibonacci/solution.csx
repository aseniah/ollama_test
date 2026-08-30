int n = int.Parse(System.Args[0]);
if (n < 1) { return; }
List<int> fibs = new List<int> { 1, 1 };
int a = 1, b = 2;
while (true) {
    if (b > n) break;
    fibs.Add(b);
    int next = a + b;
    a = b;
    b = next;
}
Console.WriteLine(string.Join("\n", fibs));