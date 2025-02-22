import tempfile
from langchain_community.document_loaders import YoutubeAudioLoader
from langchain_community.document_loaders.generic import GenericLoader
from langchain_community.document_loaders.parsers.audio import OpenAIWhisperParserLocal

def extract_youtube_transcription(url: str) -> str:
    with tempfile.TemporaryDirectory() as tmpdir:
        loader = GenericLoader(
            YoutubeAudioLoader([url], tmpdir),
            OpenAIWhisperParserLocal(lang_model="openai/whisper-large-v3-turbo")
        )
        files = loader.load()
        return files[0].page_content
        
