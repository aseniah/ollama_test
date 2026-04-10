using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Text.Json;
using System.Text.Json.Nodes;

string filePath = "input/text.txt";
string content = File.ReadAllText(filePath);

var words = content
    .ToLower()
    .Split(new char[] { ' ', '\n', '\r', '\t', '\f' }, StringSplitOptions.RemoveEmptyEntries)
    .Select(word => new string(word.Where(c => char.IsLetter(c)).ToArray()))
    .Where(word => word.Length > 0);

var wordCounts = words
    .GroupBy(w => w)
    .ToDictionary(g => g.Key, g => g.Count());

var sorted = wordCounts
    .OrderByDescending(x => x.Value)
    .ThenBy(x => x.Key)
    .ToList();

foreach (var kvp in sorted)
{
    Console.WriteLine($"{kvp.Key}: {kvp.Value}");
}