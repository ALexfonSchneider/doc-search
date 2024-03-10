
export interface SuggestionsItemProps extends React.InputHTMLAttributes<HTMLInputElement> {
    suggestion: string
}

export interface SuggestionsListProps extends React.InputHTMLAttributes<HTMLInputElement> {
    suggestions: string[]
    active: boolean
}

