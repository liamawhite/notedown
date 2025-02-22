import PyInstaller.__main__
from pathlib import Path

HERE = Path(__file__).parent.absolute()
path_to_main = str(HERE / "run.py")

def bundle():
    PyInstaller.__main__.run([
        path_to_main,
        # '--onefile',
        '--name', 'tomd',
        '--noconfirm',
        '--hidden-import', 'pydantic.deprecated.decorator',
        '--hidden-import', 'skops.io._sklearn',
        '--hidden-import', 'skops.io._quantile_forest',
        '--hidden-import', 'skops.io.old',
        '--hidden-import', 'skops.io.old._general_v0',
        '--hidden-import', 'skops.io.old._numpy_v0',
        '--hidden-import', 'skops.io.old._numpy_v1',
    ])
