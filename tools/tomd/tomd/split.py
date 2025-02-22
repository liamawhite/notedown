from wtpsplit import SaT

def split(text: str) -> str:
    sat = SaT("sat-12l")
    paragraphs = []
    for paragraph in sat.split(text, do_paragraph_segmentation=True):
        paragraphs.append(''.join(paragraph))
    return '\n\n'.join(paragraphs)
