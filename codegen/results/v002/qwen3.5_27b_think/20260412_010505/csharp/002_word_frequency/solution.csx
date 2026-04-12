#r "System"
using System;
using System.IO;
using System.Collections.Generic;
using System.Linq;
using System.Text;

string[] lines = File.ReadAllLines("input/text.txt");
string text = string.Join(" ", lines);

List<string> words = new List<string>();
StringBuilder currentWord = new StringBuilder();

foreach (char c in text)
{
    if (char.IsLetter(c))
    {
        currentWord.Append(char.ToLower(c));
    }
    else
    {
        if (currentWord.Length > 0)
        {
            words.Add(currentWord.ToString());
            currentWord.Clear();
        }
    }
}

if (currentWord.Length > 0)
{
    words.Add(currentWord.ToString());
}

Dictionary<string, int> wordCounts = new Dictionary<string, int>();
foreach (string word in words)
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

var sorted = wordCounts
    .OrderByDescending(kvp => kvp.Value)
    .ThenBy(kvp => kvp.Key);

foreach (var kvp in sorted)
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}