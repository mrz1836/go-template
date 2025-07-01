import subprocess
import tempfile
import re
from pathlib import Path

REQ_FILE = Path('.github/pip-requirements.txt')

def get_latest_version(pkg: str) -> str:
    result = subprocess.check_output(['pip', 'index', 'versions', pkg], text=True)
    first = result.splitlines()[0]
    m = re.search(r'\(([^\)]+)\)', first)
    if not m:
        raise RuntimeError(f'Cannot parse version for {pkg}')
    return m.group(1)

def get_hash(pkg: str, version: str) -> str:
    with tempfile.TemporaryDirectory() as tmp:
        subprocess.check_call(['pip', 'download', '--no-deps', f'{pkg}=={version}', '-d', tmp])
        files = list(Path(tmp).iterdir())
        if not files:
            raise RuntimeError(f'No files downloaded for {pkg}=={version}')
        file_path = files[0]
        hash_result = subprocess.check_output(['sha256sum', file_path], text=True)
        return hash_result.split()[0]

def main() -> None:
    content = REQ_FILE.read_text().splitlines()
    pkgs = []
    for line in content:
        if not line or line.startswith('#'):
            continue
        if '==' in line:
            pkg = line.split('==')[0].strip().replace('\\', '')
            pkgs.append(pkg)
    new_lines = []
    for pkg in pkgs:
        version = get_latest_version(pkg)
        sha = get_hash(pkg, version)
        new_lines.append(f'{pkg}=={version} \\')
        new_lines.append(f'    --hash=sha256:{sha}')
    REQ_FILE.write_text('\n'.join(new_lines) + '\n')

if __name__ == '__main__':
    main()
