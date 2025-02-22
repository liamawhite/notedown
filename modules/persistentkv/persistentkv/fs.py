# Copyright 2025 Notedown Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import os, base64

class FileSystemKV:
    def __init__(self, path: str):
        if not os.path.isabs(path):
            raise Exception(f"path must be absolute, got {path}")
        if not os.path.exists(path):
            os.makedirs(path)
        self.path = path

    def _encode_key(self, key: str) -> str:
        return base64.b32encode(key.encode()).decode()

    def _decode_key(self, key: str) -> str:
        return base64.b32decode(key.encode()).decode()

    def keys(self) -> list[str]:
        return [self._decode_key(f) for f in os.listdir(self.path)]

    def get(self, key: str) -> str:
        with open(os.path.join(self.path, self._encode_key(key)), 'r') as f:
            return f.read()

    def set(self, key: str, value: str):
        with open(os.path.join(self.path, self._encode_key(key)), 'w') as f:
            f.write(value)

    def delete(self, key: str):
        os.remove(os.path.join(self.path, self._encode_key(key)))

