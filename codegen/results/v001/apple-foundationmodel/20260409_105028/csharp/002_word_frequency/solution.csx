using System;
using System.IO;
using System.Text.Json;
using System.Text.Json.Nodes;

static void Main(string[] args)
{
    if (args.Length != 1)
    {
        Console.WriteLine("Usage: " + $"{Environment.GetCommandLine().Split(' ')[0]} input/text.txt");
        return;
    }

    var text = File.ReadAllText(args[0]);
    var words = text.Split(' ').Select(word =>
    {
        return word.ToLower().Replace("[^a-zA-Z]", "");
    }).ToArray();

    var wordCount = new Dictionary<string, int>();
    foreach (var word in words)
    {
        if (wordCount.ContainsKey(word))
        {
            wordCount[word]++;
        }
        else
        {
            wordCount[word] = 1;
        }
    }

    foreach (var (word, count) in wordCount.OrderByDescending(x => x.Value).ThenBy(x => x.Key))
    {
        Console.WriteLine($"{word}: {count}");
    }
}