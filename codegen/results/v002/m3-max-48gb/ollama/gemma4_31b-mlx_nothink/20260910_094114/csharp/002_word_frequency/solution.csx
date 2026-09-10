using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Text.RegularExpressions;

string filePath = "input/text.txt";

if (!File.Exists(filePath))
{
    return;
}

string text = File.ReadAllText(filePath);

// Split by whitespace to get initial tokens
string[] tokens = text.Split(new[] { ' ', '\t', '\n', '\r' }, StringSplitOptions.RemoveEmptyEntries);

Dictionary<string, int> frequencies = new Dictionary<string, int>();

foreach (var token in tokens)
{
    // Convert to lowercase
    string lowercase = token.ToLower();
    
    // Strip all punctuation (keep only letters)
    StringBuilder sb = new StringBuilder();
    foreach (char c in lowercase)
    {
        if (char.IsLetter(c))
        {
            sb.Append(c);
        }
    }
    
    string word = sb.ToString();
    
    if (!string.IsNullOrEmpty(word))
    {
        if (frequencies.ContainsKey(word))
        {
            frequencies[word]++;
        }
        else
        {
            frequencies[word] = 1;
        }
    }
}

// Sort by count descending, then by word ascending
var sortedWords = frequencies
    .OrderByDescending(pair => pair.Value)
    .ThenBy(pair => pair.Key);

foreach (var pair in sortedWords)
{
    Console.WriteLine($"{pair.Key}: {pair.Value}");
}