using System;
using System.Collections.Generic;
using System.Linq;
using System.IO;

string filePath = "input/text.txt";

if (!File.Exists(filePath))
{
    Console.WriteLine("File not found.");
    return;
}

var words = File.ReadAllText(filePath)
                .ToLower()
                .Replace(" ", "\n")
                .Split(new char[] { '\r', '\n', '.', ',', '!', '?' }, StringSplitOptions.RemoveEmptyEntries);

var wordCounts = new Dictionary<string, int>();

foreach (var word in words)
{
    if (wordCounts.ContainsKey(word))
    {
        wordCounts[word]++;
    }
    else
    {
        wordCounts[word] = 1;
    }
}

var sortedWords = wordCounts.OrderBy(w => w.Key).OrderByDescending(w => w.Value);

foreach (var wordCount in sortedWords)
{
    Console.WriteLine($"{wordCount.Key}: {wordCount.Value}");
}