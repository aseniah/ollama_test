using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;

string[] lines = File.ReadAllLines("input/text.txt");
Dictionary<string, int> wordCounts = new Dictionary<string, int>();

foreach (string line in lines)
{
    string[] words = line.Split(new char[] { ' ', ',', '.', ';', ':', '!', '?', '\'', '"', '(', ')', '[', ']', '{', '}', '\t', '\n', '\r' }, StringSplitOptions.RemoveEmptyEntries);
    
    foreach (string word in words)
    {
        string cleaned = new string(word.Where(char.IsLetter).ToArray()).ToLower();
        if (string.IsNullOrEmpty(cleaned))
            continue;
        
        if (wordCounts.ContainsKey(cleaned))
            wordCounts[cleaned]++;
        else
            wordCounts[cleaned] = 1;
    }
}

var sortedWords = wordCounts
    .OrderByDescending(kvp => kvp.Value)
    .ThenBy(kvp => kvp.Key);

foreach (var kvp in sortedWords)
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}