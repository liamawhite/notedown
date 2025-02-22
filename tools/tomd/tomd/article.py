from langchain_community.document_loaders import AsyncHtmlLoader
from langchain_community.document_transformers import MarkdownifyTransformer

def extract_article_remote(url: str) -> str:
    docs = AsyncHtmlLoader([url]).load()
    md = MarkdownifyTransformer()
    return md.transform_documents(docs)[0].page_content
