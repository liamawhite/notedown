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

import tempfile
from langchain_community.document_loaders import YoutubeAudioLoader
from langchain_community.document_loaders.generic import GenericLoader
from langchain_community.document_loaders.parsers.audio import OpenAIWhisperParserLocal


def extract_youtube_transcription(url: str) -> str:
    with tempfile.TemporaryDirectory() as tmpdir:
        loader = GenericLoader(
            YoutubeAudioLoader([url], tmpdir),
            OpenAIWhisperParserLocal(lang_model="openai/whisper-large-v3-turbo"),
        )
        files = loader.load()
        return files[0].page_content
