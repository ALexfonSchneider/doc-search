import { Action } from "@/lib/utils"

export interface ComboboxKeywordsProps {
    selected_keywords: string[]

    onAddKeyword: Action<string>
    onRemoveKeyword: Action<string>
}

export interface KeywordsSuggestions {
    suggestions: string[]
}