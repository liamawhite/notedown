# -*- mode: python ; coding: utf-8 -*-


a = Analysis(
    ['/Users/liam/github.com/notedownorg/notedown/monorepo/tools/tomd/run.py'],
    pathex=[],
    binaries=[],
    datas=[],
    hiddenimports=['pydantic.deprecated.decorator', 'skops.io._sklearn', 'skops.io._quantile_forest', 'skops.io.old', 'skops.io.old._general_v0', 'skops.io.old._numpy_v0', 'skops.io.old._numpy_v1'],
    hookspath=[],
    hooksconfig={},
    runtime_hooks=[],
    excludes=[],
    noarchive=False,
    optimize=0,
)
pyz = PYZ(a.pure)

exe = EXE(
    pyz,
    a.scripts,
    [],
    exclude_binaries=True,
    name='tomd',
    debug=False,
    bootloader_ignore_signals=False,
    strip=False,
    upx=True,
    console=True,
    disable_windowed_traceback=False,
    argv_emulation=False,
    target_arch=None,
    codesign_identity=None,
    entitlements_file=None,
)
coll = COLLECT(
    exe,
    a.binaries,
    a.datas,
    strip=False,
    upx=True,
    upx_exclude=[],
    name='tomd',
)
