var text = File.ReadAllText("input/text.txt");
var words = text.Split(new char[]{' ','\t','\n','\r'}, StringSplitOptions.RemoveEmptyEntries);

var freq = new Dictionary<string, int>();
foreach (var raw in words)
{
    var word = new string(raw.Where(char.IsLetter).ToArray()).ToLower();
    if (word.Length == 0) continue;
    freq[word] = freq.GetValueOrDefault(word, 0) + 1;
}

foreach (var kv in freq.OrderByDescending(x => x.Value).ThenBy(x => x.Key))
{
    Console.WriteLine($"{kv.Key}: {kv.Value}");
}
