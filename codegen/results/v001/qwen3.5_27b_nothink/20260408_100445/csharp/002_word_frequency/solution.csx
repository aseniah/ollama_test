using System;
using System.IO;
using System.Collections.Generic;
using System.Linq;
using System.Text;

var filePath = "input/text.txt";
if (!File.Exists(filePath))
{
    Console.Error.WriteLine($"File not found: {filePath}");
    Environment.Exit(1);
}

var text = File.ReadAllText(filePath);
var wordCounts = new Dictionary<string, int>(StringComparer.OrdinalIgnoreCase);

foreach (var word in text.Split())
{
    var cleanedWord = new string(word.Where(char.IsLetter).ToArray());
    if (!string.IsNullOrEmpty(cleanedWord))
    {
        var lowerWord = cleanedWord.ToLowerInvariant();
        if (wordCounts.ContainsKey(lowerWord))
        {
            wordCounts[lowerWord]++;
        }
        else
        {
            wordCounts[lowerWord] = 1;
        }
    }
}

var sortedWords = wordCounts.OrderByDescending(k => k.Value)
                            .ThenBy(k => k.Key)
                            .ToList();

foreach (var pair in sortedWords)
{
    Console.WriteLine($"{pair.Key}: {pair.Value}");
}