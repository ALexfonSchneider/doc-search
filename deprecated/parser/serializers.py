from rest_framework import serializers

class AuthorSerializer(serializers.Serializer):
    name = serializers.CharField()
    affiliation = serializers.CharField()
    orcid = serializers.CharField()
    
    
class ArticleSerializer(serializers.Serializer):
    article_id = serializers.CharField()
    authors = AuthorSerializer(many=True)
    title = serializers.CharField()
    anatation = serializers.CharField()
    published = serializers.CharField()
    archive_id = serializers.CharField()
    preview_id = serializers.CharField()
    keywords = serializers.ListField()
    filename = serializers.CharField()
    content = serializers.CharField()
    udk = serializers.CharField()

class ArchiveSerializer(serializers.Serializer):
    archive_id = serializers.CharField()
    name = serializers.CharField()
    series = serializers.CharField()
    articles = ArticleSerializer(many=True)