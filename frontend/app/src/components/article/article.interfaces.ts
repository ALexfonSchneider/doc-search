import { Action } from "@/lib/utils"

export interface ArticleAuthor {
    affiliation: string
    name: string
    orcid: string
}


export interface Article {
    article_id: string
    preview_id: string
    title: string
    anatation: string
    udk: string
    published: string
    keywords: string[]
    authors: ArticleAuthor[]
    link: string
}


export interface Archive {
    archive_id: string
    name: string
    series: string
    url: string
}

export interface DocHighlight {
    ["article.content"]: string[]
}

export interface WordCloudItem {
    value: string
    count: number
}

export interface Metrics {
    word_cloud: WordCloudItem[]
}


export interface Document {
    article: Article
    archive: Archive
    metrics: Metrics
    highlight: DocHighlight
}


export interface SearchedDocument {
    document: Document

    onBadgeClick: Action<string>
}