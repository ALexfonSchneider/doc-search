import { Action } from "@/lib/utils"
import { Document } from "../article/article.interfaces"

export interface SearchContentPaginationProps {
    page: number
    total_size: number
    size: number

    onSelect: Action<number>
}

export interface SearchResultsPaginate {
    articles: Document[]
    page: number,
    size: number,
    total_size: number
}
