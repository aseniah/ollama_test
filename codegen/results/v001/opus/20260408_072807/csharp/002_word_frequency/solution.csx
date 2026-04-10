var text = File.ReadAllText("input/text.txt");
var words = text.Split(new char[]{' ','\t','\n','\r'}, StringSplitOptions.RemoveEmptyEntries);

var freq = new Dictionary<string, int>();
foreach (var raw in words)
{
    var word = new string(raw.ToLower().Where(char.IsLetter).ToArray());
    if (word.Length == 0) continue;
    freq[word] = freq.TryGetValue(word, out var c) ? c + 1 : 1;
}

foreach (var kvp in freq.OrderByDescending(x => x.Value).ThenBy(x => x.Key))
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}
