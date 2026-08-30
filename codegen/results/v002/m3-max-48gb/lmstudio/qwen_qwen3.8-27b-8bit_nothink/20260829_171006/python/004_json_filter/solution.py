<tool_call>
</think>

<tool_call>
<function=execute_bash>
<parameter=command>
python3 << 'EOF'
import json

with open('input/data.json', 'r') as f:
    data = json.load(f)

filtered = [r for r in data if r['active'] and r['age'] >= 30]
filtered.sort(key=lambda r: r['name'])

print(json.dumps(filtered))
EOF
</parameter>
</function>
</tool_call>