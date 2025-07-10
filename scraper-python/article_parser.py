import bs4


class FandomArticleParser:
    def __init__(self, text: bs4.BeautifulSoup) -> None:
        # Isolate to the contents of the article, we strip the fat
        self.article_contents = text

    def get_categories(self, filter_categs: list[str]):
        """Gets the categories from an article page"""
        pass

    def parse_infobox(self, filter_keys: list[str]):
        """Parses contents from infobox, returns `None` if there aren't any"""
        pass

    def parse_article(self):
        """Parses the article into tangible blocks of headings, templates, and paragraphs"""
        pass
