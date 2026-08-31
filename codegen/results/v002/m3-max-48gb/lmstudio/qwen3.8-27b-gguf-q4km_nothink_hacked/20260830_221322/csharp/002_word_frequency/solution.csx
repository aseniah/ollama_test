using System;
using System.IO;
using System.Linq;
using System.Text;
using System.Collections.Generic;

string content = File.ReadAllText("input/text.txt");
var wordFreq = new Dictionary<string, int>();

foreach (var token in content.Split(new[] { ' ', '\t', '\n', '\r', ',', '.', '!', '?', ';', ':', '-', '(', ')', '[', ']', '{', '}', '"', '\'', '`', '~' }, StringSplitOptions.RemoveEmptyEntries))
{
    // Strip all punctuation, keep only letters
    var cleaned = new StringBuilder();
    foreach (char c in token)
    {
        if (char.IsLetter(c))
            cleaned.Append(char.ToLowerInvariant(c));
    }
    var word = cleaned.ToString();
    if (word.Length == 0) continue;
    
    if (wordFreq.ContainsKey(word))
        wordFreq[word]++;
    else
        wordFreq[word] = 1;
}

var sorted = wordFreq.OrderByDescending(kvp => kvp.Value)
                     .ThenBy(kvp => kvp.Key)
                     .ToList();

foreach (var kvp in sorted)
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}