### 1. Check for an existing virtual environment

```bash
find ~ -type f -name "pyvenv.cfg"
```

If you know what is the name in this case (.venv)
From the repository root:

```bash
ls -d .venv && .venv/bin/python --version
```

### 2. Create the virtual environment and install dependencies

Only create `.venv` if it does not already exist:

```bash
python3 -m venv .venv
```

Install or update the dependencies inside it:

```bash
.venv/bin/python -m pip install -r requirements.txt
```

### 3. Activate the environment


```bash
source .venv/bin/activate
```