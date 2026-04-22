using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;

string[] lines = File.ReadAllLines("input/text.txt");
List<string> words = new List<string>();

foreach (string line in lines)
{
    // Split by whitespace and punctuation
    // First, replace all non-letter characters with spaces
    string cleaned = "";
    foreach (char c in line)
    {
        if (char.IsLetter(c))
        {
            cleaned += c;
        }
        else
        {
            cleaned += ' ';
        }
    }

    string[] splitWords = cleaned.Split(new char[] { ' ', '\t', '\n', '\r' }, StringSplitOptions.RemoveEmptyEntries);
    foreach (string word in splitWords)
    {
        words.Add(word.ToLower());
    }
}

Dictionary<string, int> freq = new Dictionary<string, int>();
foreach (string word in words)
{
    if (freq.ContainsKey(word))
    {
        freq[word]++;
    }
    else
    {
        freq[word] = 1;
    }
}

var sorted = freq.OrderByDescending(kv => kv.Value)
                 .ThenBy(kv => kv.Key)
                 .ToList();

foreach (var kv in sorted)
{
    Console.WriteLine($"{kv.Key}: {kv.Value}");
}