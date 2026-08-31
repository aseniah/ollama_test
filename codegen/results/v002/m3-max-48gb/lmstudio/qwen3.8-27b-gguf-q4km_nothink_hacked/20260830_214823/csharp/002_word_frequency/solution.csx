using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;

string[] lines = File.ReadAllLines("input/text.txt");
Dictionary<string, int> wordCount = new Dictionary<string, int>();

foreach (string line in lines)
{
    string[] tokens = line.Split(new[] { ' ', ',', '-', ';', ':', '.', '!', '?', '\t' }, StringSplitOptions.RemoveEmptyEntries);
    foreach (string token in tokens)
    {
        // Strip all non-letter characters, keep only letters
        string word = new string(token.Where(c => char.IsLetter(c)).ToArray());
        if (string.IsNullOrEmpty(word))
            continue;
        word = word.ToLower();
        if (!wordCount.ContainsKey(word))
            wordCount[word] = 0;
        wordCount[word]++;
    }
}

var sorted = wordCount
    .OrderByDescending(kv => kv.Value)
    .ThenBy(kv => kv.Key)
    .ToList();

foreach (var kv in sorted)
{
    Console.WriteLine($"{kv.Key}: {kv.Value}");
}