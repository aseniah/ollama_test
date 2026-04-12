using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;

string content = File.ReadAllLines("input/text.txt").Concat(new string[] { "" }).FirstOrDefault();
Regex regex = new Regex(@"\b[a-z]+\b");
Dictionary<string, int> wordCounts = new Dictionary<string, int>();

if (!string.IsNullOrEmpty(content))
{
    while (Regex.IsMatch(content, @"[^a-z ]"))
    {
        content = Regex.Replace(content, @"[^a-z ]", "");
    }
    foreach (Match match in regex.Matches(content))
    {
        string word = match.Value;
        if (!wordCounts.ContainsKey(word))
        {
            wordCounts[word] = 0;
        }
        wordCounts[word]++;
    }
}

List<(string word, int count)> sortedWords = wordCounts.ToDictionary().Select(x => (x.Key, x.Value)).OrderByDescending(x => x.count).ThenBy(x => x.word).ToList();

foreach (var item in sortedWords)
{
    Console.WriteLine($"{item.word}: {item.count}");
}