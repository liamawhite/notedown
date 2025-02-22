import click, warnings, sys, mdformat, os
from pathlib import Path
from contextlib import redirect_stdout
from persistentkv.fs import FileSystemKV
os.environ['USER_AGENT'] = 'tomd' # set before importing langchain things otherwise we get a warning
from .youtube import extract_youtube_transcription
from .article import extract_article_remote
from .split import split

cache_dir = Path.home() / '.cache' / 'tomd'
warnings.filterwarnings("ignore")

@click.command()
@click.argument('url')
@click.option('--disable-cache', is_flag=True, default=False, help='force extraction even if cached result is present')
def extract(url: str, disable_cache: bool):
    """
    Extracts content as Markdown from a URL and prints it to stdout.

    Assumes text content by default but supports YouTube videos.
    """

    content = ""
    kv = FileSystemKV(str(cache_dir.resolve()))

    if not disable_cache:
        try:
            content = kv.get(url)
            click.echo('cache hit, source previously processed', err=True)
            click.echo(content)
            return
        except:
            click.echo('cache miss, source not previously processed', err=True)

    with redirect_stdout(sys.stderr):
        match url:
            case m if "youtube.com" in m or "youtu.be" in m:
                click.echo('inferred YouTube video', err=True)
                transcript = extract_youtube_transcription(url)
                content = split(transcript)

            case _:
                click.echo('no specific format detected, defaulting to text article', err=True)
                content = extract_article_remote(url)

        content = mdformat.text(content) 
        kv.set(url, content)

    with redirect_stdout(sys.stdout):
        click.echo(content)

if __name__ == '__main__':
    extract()
