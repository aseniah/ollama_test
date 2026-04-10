#!/usr/bin/env bash
set -euo pipefail

PYPROJECT="pyproject.toml"

# --- Homebrew ---
if ! command -v brew &>/dev/null; then
    echo "error: Homebrew is not installed. Visit https://brew.sh to install it." >&2
    exit 1
fi

# --- Tools ---
for tool in ruff pyright; do
    if ! command -v "$tool" &>/dev/null; then
        echo "Installing $tool..."
        brew install "$tool"
    fi
done

# --- pyproject.toml ---
if [[ ! -f "$PYPROJECT" ]]; then
    echo "Creating $PYPROJECT..."
    cat > "$PYPROJECT" <<'EOF'
[tool.pyright]
pythonVersion = "3.11"
typeCheckingMode = "strict"

[tool.ruff.lint]
select = ["E", "F", "W", "ANN"]
EOF
fi

# --- Run checks ---
PASS=0

echo ""
echo "=== ruff ==="
if ruff check .; then
    echo "ruff: no issues"
else
    PASS=1
fi

echo ""
echo "=== pyright ==="
if pyright; then
    echo "pyright: no issues"
else
    PASS=1
fi

echo ""
exit $PASS
