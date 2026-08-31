using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;

string text = File.ReadAllText("input/text.txt");

var words = text
    .ToLower()
    .Split(new[] { ' ', '\t', '\r', '\n' }, StringSplitOptions.RemoveEmptyEntries)
    .Select(w => new string(w.Where(char.IsLetter).ToArray()))
    .Where(w => w.Length > 0);

var counts = words
    .GroupBy(w => w)
    .Select(g => new { Word = g.Key, Count = g.Count() })
    .OrderByDescending(x => x.Count)
    .ThenBy(x => x.Word)
    .ToList();

foreach (var item in counts)
{
    Console.WriteLine($"{item.Word}: {item.Count}");
}